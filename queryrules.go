package etcd

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/coreos/etcd/clientv3"
	"github.com/juju/errors"
)

type (
	EtcdQr struct {
		Rule_id               uint64 `db:"rule_id" json:"rule_id"`
		Active                uint64 `db:"active" json:"active"`
		Username              string `db:"username" json:"username"`
		Schemaname            string `db:"schemaname" json:"schemaname"`
		FlagIN                uint64 `db:"flagIN" json:"flagIN"`
		Client_addr           string `db:"client_addr" json:"client_addr"`
		Proxy_addr            string `db:"proxy_addr" json:"proxy_addr"`
		Proxy_port            string `db:"proxy_port" json:"proxy_port"`
		Digest                string `db:"digest" json:"digest"`
		Match_digest          string `db:"match_digest" json:"match_digest"`
		Match_pattern         string `db:"match_pattern" json:"match_pattern"`
		Negate_match_pattern  uint64 `db:"negate_match_pattern" json:"negate_match_pattern"`
		FlagOUT               string `db:"flagOUT" json:"flagOUT"`
		Replace_pattern       string `db:"replace_pattern" json:"replace_pattern"`
		Destination_hostgroup string `db:"destination_hostgroup" json:"destination_hostgroup"`
		Cache_ttl             string `db:"cache_ttl" json:"cache_ttl"`
		Reconnect             string `db:"reconnect" json:"reconnect"`
		Timeout               string `db:"timeout" json:"timeout"`
		Retries               string `db:"retries" json:"retries"`
		Delay                 string `db:"delay" json:"delay"`
		Mirror_flagOUT        string `db:"mirror_flagOUT" json:"mirror_flagOUT"`
		Mirror_hostgroup      string `db:"mirror_hostgroup" json:"mirror_hostgroup"`
		Error_msg             string `db:"error_msg" json:"error_msg"`
		Log                   string `db:"log" json:"log"`
		Apply                 uint64 `db:"apply" json:"apply"`
		Comment               string `db:"comment" json:"comment"`
	}
)

// new mysql query rules
func NewQr(username string) (*EtcdQr, error) {
	newqr := new(EtcdQr)

	if username == "" {
		return nil, errors.BadRequestf(username)
	}
	if strings.Index(username, "\"") == -1 {
		newqr.Username = fmt.Sprintf("\"%s\"", username)
	} else {
		newqr.Username = username
	}

	newqr.Destination_hostgroup = "NULL"
	newqr.Schemaname = "NULL"
	newqr.FlagIN = 0
	newqr.Client_addr = "NULL"
	newqr.Proxy_addr = "NULL"
	newqr.Proxy_port = "NULL"
	newqr.Digest = "NULL"
	newqr.Match_digest = "NULL"
	newqr.Match_pattern = "NULL"
	newqr.Negate_match_pattern = 0
	newqr.FlagOUT = "NULL"
	newqr.Replace_pattern = "NULL"
	newqr.Cache_ttl = "NULL"
	newqr.Reconnect = "NULL"
	newqr.Timeout = "NULL"
	newqr.Retries = "NULL"
	newqr.Delay = "NULL"
	newqr.Mirror_flagOUT = "NULL"
	newqr.Mirror_hostgroup = "NULL"
	newqr.Error_msg = "NULL"
	newqr.Log = "NULL"
	newqr.Apply = 0
	newqr.Active = 0
	newqr.Comment = "NULL"

	return newqr, nil
}

// set qr rule_id
func (qr *EtcdQr) SetQrRuleid(rule_id uint64) {
	qr.Rule_id = rule_id
}

// set qr active
func (qr *EtcdQr) SetQrActive(active uint64) {
	switch active {
	case 0:
		qr.Active = 0
	case 1:
		qr.Active = 1
	default:
		qr.Active = 1
	}
}

// set qr apply
func (qr *EtcdQr) SetQrApply(apply uint64) {
	switch apply {
	case 0:
		qr.Apply = 0
	case 1:
		qr.Apply = 1
	default:
		qr.Apply = 1
	}
}

// set qr schemaname
func (qr *EtcdQr) SetQrSchemaname(schema_name string) {
	if schema_name == "" || len(schema_name) == 0 {
		qr.Schemaname = "NULL"
	} else {
		if strings.Index(schema_name, "\"") == -1 {
			qr.Schemaname = fmt.Sprintf("\"%s\"", schema_name)
		} else {
			qr.Schemaname = schema_name
		}
	}
}

// set qr flagIN
func (qr *EtcdQr) SetQrFlagIN(flag_in uint64) {
	qr.FlagIN = flag_in
}

// set qr client_addr
func (qr *EtcdQr) SetQrClientAddr(client_addr string) {
	if client_addr == "" || len(client_addr) == 0 {
		qr.Client_addr = "NULL"
	} else {
		if strings.Index(client_addr, "\"") == -1 {
			qr.Client_addr = fmt.Sprintf("\"%s\"", client_addr)
		} else {
			qr.Client_addr = client_addr
		}
	}
}

// set qr proxy_addr
func (qr *EtcdQr) SetQrProxyAddr(proxy_addr string) {
	if proxy_addr == "" || len(proxy_addr) == 0 {
		qr.Proxy_addr = "NULL"
	} else {
		if strings.Index(proxy_addr, "\"") == -1 {
			qr.Proxy_addr = fmt.Sprintf("\"%s\"", proxy_addr)
		} else {
			qr.Proxy_addr = proxy_addr
		}
	}
}

// set qr proxy_port
func (qr *EtcdQr) SetProxyPort(proxy_port string) {
	if proxy_port == "" || len(proxy_port) == 0 {
		qr.Proxy_port = "NULL"
	} else {
		if strings.Index(proxy_port, "\"") == -1 {
			qr.Proxy_port = fmt.Sprintf("\"%s\"", proxy_port)
		} else {
			qr.Proxy_port = proxy_port
		}
	}
}

// set qr digest
func (qr *EtcdQr) SetQrDigest(digest string) {
	if digest == "" || len(digest) == 0 {
		qr.Digest = "NULL"
	} else {
		if strings.Index(digest, "\"") == -1 {
			qr.Digest = fmt.Sprintf("\"%s\"", digest)
		} else {
			qr.Digest = digest
		}
	}
}

// set qr match_digest
func (qr *EtcdQr) SetQrMatchDigest(match_digest string) {
	if match_digest == "" || len(match_digest) == 0 {
		qr.Match_digest = "NULL"
	} else {
		if strings.Index(match_digest, "\"") == -1 {
			qr.Match_digest = fmt.Sprintf("\"%s\"", match_digest)
		} else {
			qr.Match_digest = match_digest
		}
	}
}

// set qr match_pattern
func (qr *EtcdQr) SetQrMatchPattern(match_pattern string) {
	if match_pattern == "" || len(match_pattern) == 0 {
		qr.Match_pattern = "NULL"
	} else {
		if strings.Index(match_pattern, "\"") == -1 {
			qr.Match_pattern = fmt.Sprintf("\"%s\"", match_pattern)
		} else {
			qr.Match_pattern = match_pattern
		}
	}
}

// set qr mnegate_match_pattern
func (qr *EtcdQr) SetQrNegateMatchPattern(negate_match_pattern uint64) {
	switch negate_match_pattern {
	case 0:
		qr.Negate_match_pattern = 0
	case 1:
		qr.Negate_match_pattern = 1
	default:
		qr.Negate_match_pattern = 0
	}
}

// set qr flagout
func (qr *EtcdQr) SetQrFlagOut(flag_out string) {
	if flag_out == "" || len(flag_out) == 0 {
		qr.FlagOUT = "NULL"
	} else {
		if strings.Index(flag_out, "\"") == -1 {
			qr.FlagOUT = fmt.Sprintf("\"%s\"", flag_out)
		} else {
			qr.FlagOUT = flag_out
		}
	}
}

// set qr replace_pattern
func (qr *EtcdQr) SetQrReplacePattern(replace_pattern string) {
	if replace_pattern == "" || len(replace_pattern) == 0 {
		qr.Replace_pattern = "NULL"
	} else {
		if strings.Index(replace_pattern, "\"") == -1 {
			qr.Replace_pattern = fmt.Sprintf("\"%s\"", replace_pattern)
		} else {
			qr.Replace_pattern = replace_pattern
		}
	}
}

// set qr destination_hostgroup
func (qr *EtcdQr) SetQrDestHostGroup(destination_hostgroup string) {
	if destination_hostgroup == "" || len(destination_hostgroup) == 0 {
		qr.Destination_hostgroup = "NULL"
	} else {
		if strings.Index(destination_hostgroup, "\"") == -1 {
			qr.Destination_hostgroup = fmt.Sprintf("\"%s\"", destination_hostgroup)
		} else {
			qr.Destination_hostgroup = destination_hostgroup
		}
	}
}

// set qr cache_ttl
func (qr *EtcdQr) SetQrCacheTTL(cache_ttl string) {
	if cache_ttl == "" || len(cache_ttl) == 0 {
		qr.Cache_ttl = "NULL"
	} else {
		if strings.Index(cache_ttl, "\"") == -1 {
			qr.Cache_ttl = fmt.Sprintf("\"%s\"", cache_ttl)
		} else {
			qr.Cache_ttl = cache_ttl
		}
	}
}

// set qr reconnect
func (qr *EtcdQr) SetQrReconnect(reconnect string) {
	if reconnect == "" || len(reconnect) == 0 {
		qr.Reconnect = "NULL"
	} else {
		if strings.Index(reconnect, "\"") == -1 {
			qr.Reconnect = fmt.Sprintf("\"%s\"", reconnect)
		} else {
			qr.Reconnect = reconnect
		}
	}
}

// set qr timeout
func (qr *EtcdQr) SetQrTimeOut(timeout string) {
	if timeout == "" || len(timeout) == 0 {
		qr.Timeout = "NULL"
	} else {
		if strings.Index(timeout, "\"") == -1 {
			qr.Timeout = fmt.Sprintf("\"%s\"", timeout)
		} else {
			qr.Timeout = timeout
		}
	}
}

// set qr retries
func (qr *EtcdQr) SetQrRetries(retries string) {
	if retries == "" || len(retries) == 0 {
		qr.Retries = "NULL"
	} else {
		if strings.Index(retries, "\"") == -1 {
			qr.Retries = fmt.Sprintf("\"%s\"", retries)
		} else {
			qr.Retries = retries
		}
	}
}

// set qr delay
func (qr *EtcdQr) SetQrDelay(delay string) {
	if delay == "" || len(delay) == 0 {
		qr.Delay = "NULL"
	} else {
		if strings.Index(delay, "\"") == -1 {
			qr.Delay = fmt.Sprintf("\"%s\"", delay)
		} else {
			qr.Delay = delay
		}
	}
}

// set qr mirror_flagout
func (qr *EtcdQr) SetQrMirrorFlagOUT(mirror_flagout string) {
	if mirror_flagout == "" || len(mirror_flagout) == 0 {
		qr.Mirror_flagOUT = "NULL"
	} else {
		if strings.Index(mirror_flagout, "\"") == -1 {
			qr.Mirror_flagOUT = fmt.Sprintf("\"%s\"", mirror_flagout)
		} else {
			qr.Mirror_flagOUT = mirror_flagout
		}
	}
}

// set qr mirror_hostgroup
func (qr *EtcdQr) SetQrMirrorHostgroup(mirror_hostgroup string) {
	if mirror_hostgroup == "" || len(mirror_hostgroup) == 0 {
		qr.Mirror_hostgroup = "NULL"
	} else {
		if strings.Index(mirror_hostgroup, "\"") == -1 {
			qr.Mirror_hostgroup = fmt.Sprintf("\"%s\"", mirror_hostgroup)
		} else {
			qr.Mirror_hostgroup = mirror_hostgroup
		}
	}
}

// set qr error_msg
func (qr *EtcdQr) SetQrErrorMsg(error_msg string) {
	if error_msg == "" || len(error_msg) == 0 {
		qr.Error_msg = "NULL"
	} else {
		if strings.Index(error_msg, "\"") == -1 {
			qr.Error_msg = fmt.Sprintf("\"%s\"", error_msg)
		} else {
			qr.Error_msg = error_msg
		}
	}
}

// set qr log
func (qr *EtcdQr) SetQrLog(log string) {
	if log == "" || len(log) == 0 {
		qr.Log = "NULL"
	} else {
		if strings.Index(log, "\"") == -1 {
			qr.Log = fmt.Sprintf("\"%s\"", log)
		} else {
			qr.Log = log
		}
	}
}

// create or update a new queryrules
func (qr *EtcdQr) CreateOrUpdateOneQr(etcdcli *EtcdCli, cli *clientv3.Client) error {

	rule_id := fmt.Sprintf("%d", qr.Rule_id)

	key := []byte(rule_id)

	value, err := json.Marshal(qr)
	if err != nil {
		return errors.Trace(err)
	}

	// base64
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

// delete a queryrules
func (qr *EtcdQr) DeleteOneQr(etcdcli *EtcdCli, cli *clientv3.Client) error {
	rule_id := fmt.Sprintf("%d", qr.Rule_id)

	key := []byte(rule_id)

	// base64
	encodeKey := base64.StdEncoding.EncodeToString(key)

	ctx, cancel := context.WithTimeout(context.Background(), etcdcli.RequestTimeout)
	//create user at etcd
	response, err := cli.Delete(ctx, etcdcli.Root+"/"+encodeKey, clientv3.WithPrefix())
	cancel()
	if err != nil {
		return errors.Trace(err)
	}

	orig_key, err := base64.StdEncoding.DecodeString(encodeKey)
	if err != nil {
		return errors.Trace(err)
	}

	if response.Deleted == 0 {
		return errors.NotFoundf(string(orig_key))
	}

	return nil
}

// query all queryrules
func QueryAllQrs(etcdcli *EtcdCli, cli *clientv3.Client) ([]EtcdQr, error) {
	var allqrs []EtcdQr

	ctx, cancel := context.WithTimeout(context.Background(), etcdcli.RequestTimeout)
	resp, err := cli.Get(ctx, etcdcli.Root, clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))
	cancel()
	if err != nil {
		return nil, errors.Trace(err)
	}
	for _, ev := range resp.Kvs {
		var tmpqr EtcdQr
		value, _ := base64.StdEncoding.DecodeString(string(ev.Value))
		//decode user information.
		if err := json.Unmarshal(value, &tmpqr); err != nil {
			return nil, errors.Trace(err)
		}
		//merge user informations.
		allqrs = append(allqrs, tmpqr)

	}
	return allqrs, nil
}
