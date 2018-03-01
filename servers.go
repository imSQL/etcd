package etcd

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"strconv"

	"github.com/coreos/etcd/clientv3"
	"github.com/juju/errors"
)

type (
	EtcdBackend struct {
		HostGroupId       uint64 `db:"hostgroup_id,omitempty" json:"hostgroup_id"`
		HostName          string `db:"hostname" json:"hostname"`
		Port              uint64 `db:"port" json:"port"`
		Status            string `db:"status" json:"status"`
		Weight            uint64 `db:"weight" json:"weight"`
		Compression       uint64 `db:"compression" json:"compression"`
		MaxConnections    uint64 `db:"max_connections" json:"max_connections"`
		MaxReplicationLag uint64 `db:"max_replication_lag" json:"max_replication_lag"`
		UseSsl            uint64 `db:"use_ssl" json:"use_ssl"`
		MaxLatencyMs      uint64 `db:"max_latency_ms" json:"max_latency_ms"`
		Comment           string `db:"comment" json:"comment"`
	}
)

// init a new servers.
func NewServer(hostgroup_id uint64, hostname string, port uint64) (*EtcdBackend, error) {
	newsrv := new(EtcdBackend)

	newsrv.HostGroupId = hostgroup_id
	newsrv.HostName = hostname
	newsrv.Port = port

	newsrv.Status = "ONLINE"
	newsrv.Weight = 1000
	newsrv.Compression = 0
	newsrv.MaxConnections = 10000
	newsrv.MaxReplicationLag = 0
	newsrv.UseSsl = 0
	newsrv.MaxLatencyMs = 0
	newsrv.Comment = ""

	return newsrv, nil
}

// set servers status
func (srvs *EtcdBackend) SetServerStatus(status string) {
	switch status {
	case "ONLINE":
		srvs.Status = "ONLINE"
	case "SHUNNED":
		srvs.Status = "SHUNNED"
	case "OFFLINE_SOFT":
		srvs.Status = "OFFLINE_SOFT"
	case "OFFLINE_HARD":
		srvs.Status = "OFFLINE_HARD"
	default:
		srvs.Status = "ONLINE"
	}
}

// set servers weight
func (srvs *EtcdBackend) SetServerWeight(weight uint64) {
	srvs.Weight = weight
}

// set servers compression
func (srvs *EtcdBackend) SetServerCompression(compression uint64) {
	srvs.Compression = compression
}

// set servers max_connections
func (srvs *EtcdBackend) SetServerMaxConnection(max_connections uint64) {
	if max_connections >= 10000 {
		srvs.MaxConnections = 10000
	} else {
		srvs.MaxConnections = max_connections
	}
}

// set servers max_replication_lag
func (srvs *EtcdBackend) SetServerMaxReplicationLag(max_replication_lag uint64) {
	if max_replication_lag > 126144000 {
		srvs.MaxReplicationLag = 1261440000
	} else {
		srvs.MaxReplicationLag = max_replication_lag
	}
}

// set servers use_ssl
func (srvs *EtcdBackend) SetServerUseSSL(use_ssl uint64) {
	if use_ssl >= 1 {
		srvs.UseSsl = 1
	} else {
		srvs.UseSsl = 0
	}
}

// set servers max_latency_ms
func (srvs *EtcdBackend) SetServerMaxLatencyMs(max_latency_ms uint64) {
	srvs.MaxLatencyMs = max_latency_ms
}

// set servers comment
func (srvs *EtcdBackend) SetServersComment(comment string) {
	srvs.Comment = comment
}

// create or update a backend informations.
func (srvs *EtcdBackend) CreateOrUpdateOneBackend(etcdcli *EtcdCli, cli *clientv3.Client) error {

	key := []byte(strconv.Itoa(int(srvs.HostGroupId)) + "|" + srvs.HostName + "|" + strconv.Itoa(int(srvs.Port)))
	value, err := json.Marshal(srvs)
	if err != nil {
		return errors.Trace(err)
	}

	encodeKey := base64.StdEncoding.EncodeToString(key)
	encodeValue := base64.StdEncoding.EncodeToString(value)

	ctx, cancel := context.WithTimeout(context.Background(), etcdcli.RequestTimeout)

	//create user at etcd
	_, err = cli.Put(ctx, etcdcli.Root+"/"+encodeKey, encodeValue)
	cancel()
	if err != nil {
		return errors.Trace(err)
	}

	return nil
}

// delete a backend.
func (srvs *EtcdBackend) DeleteOneBackend(etcdcli *EtcdCli, cli *clientv3.Client) error {

	key := []byte(strconv.Itoa(int(srvs.HostGroupId)) + "|" + srvs.HostName + "|" + strconv.Itoa(int(srvs.Port)))
	//value, err := json.Marshal(newuser)
	//if err != nil {
	//	return errors.Trace(err)
	//}

	encodeKey := base64.StdEncoding.EncodeToString(key)
	//encodeValue := base64.StdEncoding.EncodeToString(value)

	ctx, cancel := context.WithTimeout(context.Background(), etcdcli.RequestTimeout)

	//create user at etcd
	_, err := cli.Delete(ctx, etcdcli.Root+"/"+encodeKey, clientv3.WithPrefix())
	cancel()
	if err != nil {
		return errors.Trace(err)
	}

	return nil
}

// query all backens informations.
func (srvs *EtcdBackend) QueryAllBackends(etcdcli *EtcdCli, cli *clientv3.Client) ([]EtcdBackend, error) {

	var allbkd []EtcdBackend

	ctx, cancel := context.WithTimeout(context.Background(), etcdcli.RequestTimeout)
	resp, err := cli.Get(ctx, etcdcli.Root, clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))
	cancel()
	if err != nil {
		return nil, errors.Trace(err)
	}
	for _, ev := range resp.Kvs {
		var tmpbkd EtcdBackend
		value, _ := base64.StdEncoding.DecodeString(string(ev.Value))
		//decode user information.
		if err := json.Unmarshal(value, &tmpbkd); err != nil {
			return nil, errors.Trace(err)
		}
		//merge user informations.
		allbkd = append(allbkd, tmpbkd)

	}
	return allbkd, nil
}
