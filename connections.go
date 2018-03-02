package etcd

import (
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/juju/errors"
)

type (
	EtcdCli struct {
		//etcd connection informations
		DialTimeout    time.Duration
		RequestTimeout time.Duration
		EndPoints      []string

		//watch path.
		// sucs as :  /database/parauser
		// database is prefix
		// parauser is service
		Prefix  string
		Service string
		Root    string
		Type    string
		Key     string
		Value   string

		//cli
		cli *clientv3.Client

		//proxysql connection informations.
		ProxySQLAddr  string
		ProxySQLPort  uint64
		ProxySQLAdmin string
		ProxySQLPass  string

		//error
		Err error
	}
)

// new etcd client
// return a new etcdcli.
func NewEtcdCli(endpoints []string) *EtcdCli {

	var etcdcli EtcdCli

	// set default timeout value
	etcdcli.DialTimeout = 5 * time.Second
	etcdcli.RequestTimeout = 3 * time.Second

	// set endporints
	etcdcli.EndPoints = endpoints

	// etcd watch path.
	// usage:  /<root path>/<service name>/<proxysql object>/<base64 key> <base64 value>
	// example:  /database/user_center/users/cmVsYXRpb25fc2VydmljZQ== eyJpZCI6MSwidXNlcm5hbWUiOiJ0aWFubGVpIiwiYWdlIjozMywiYWRkcmVzcyI6ImJlaWppbmcifQ==
	etcdcli.Prefix = "database"
	etcdcli.Service = "users"
	etcdcli.Type = ""
	etcdcli.Key = ""
	etcdcli.Value = ""

	return &etcdcli
}

// set dialtimeout
func (cli *EtcdCli) SetDilTimeout(num uint64) {

	cli.DialTimeout = time.Duration(num) * time.Second
}

// set request timeout
func (cli *EtcdCli) SetRequestTimeout(num uint64) {

	cli.RequestTimeout = time.Duration(num) * time.Second
}

// set root path
// default is database
func (cli *EtcdCli) SetPrefix(prefix string) {
	cli.Prefix = prefix
}

// set service name
func (cli *EtcdCli) SetService(service string) {
	cli.Service = service
}

// set type name
func (cli *EtcdCli) SetEtcdType(etcd_type string) {
	cli.Type = etcd_type
}

// set etcd key
func (cli *EtcdCli) SetEtcdKey(key string) {
	cli.Key = key
}

// set etcd value
func (cli *EtcdCli) SetEtcdValue(val string) {
	cli.Value = val
}

// set proxysql dbi
func (cli *EtcdCli) SetProxyAddr(proxy_addr string) {
	cli.ProxySQLAddr = proxy_addr
}

func (cli *EtcdCli) SetProxyPort(proxy_port uint64) {
	cli.ProxySQLPort = proxy_port
}

func (cli *EtcdCli) SetProxyAdmin(admin string) {
	cli.ProxySQLAdmin = admin
}

func (cli *EtcdCli) SetProxyPass(pass string) {
	cli.ProxySQLPass = pass
}

// prefix+service
func (cli *EtcdCli) MakeWatchRoot() {
	if len(cli.Type) == 0 {
		cli.Root = fmt.Sprintf("/%s/%s", cli.Prefix, cli.Service)
	} else {
		cli.Root = fmt.Sprintf("/%s/%s/%s", cli.Prefix, cli.Service, cli.Type)
	}
}

// open etcd connection.
func (cli *EtcdCli) OpenEtcd() (*clientv3.Client, error) {

	var ecli *clientv3.Client

	ecli, err := clientv3.New(clientv3.Config{
		Endpoints:   cli.EndPoints,
		DialTimeout: cli.DialTimeout,
	})
	if err != nil {
		return nil, errors.Trace(err)
	}

	return ecli, nil
}

// close etcd connection.
func (cli *EtcdCli) CloseEtcd(ecli *clientv3.Client) error {
	err := ecli.Close()
	return errors.Trace(err)
}
