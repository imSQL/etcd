package etcd

import (
	"context"
	"encoding/base64"
	"encoding/json"

	"github.com/coreos/etcd/clientv3"
	"github.com/juju/errors"
)

type (
	EtcdVariables struct {
		VariablesName string `db:"Variable_name" json:"variable_name"`
		Value         string `db:"Value" json:"variable_value"`
	}
)

func UpdateOneConfig(etcdcli *EtcdCli, cli *clientv3.Client, vars *EtcdVariables) error {

	key := []byte(vars.VariablesName)
	value := []byte(vars.Value)

	// base64
	encodeKey := base64.StdEncoding.EncodeToString(key)
	encodeValue := base64.StdEncoding.EncodeToString(value)

	ctx, cancel := context.WithTimeout(context.Background(), etcdcli.RequestTimeout)
	//create user
	_, err := cli.Put(ctx, etcdcli.Root+"/"+encodeKey, encodeValue)
	cancel()
	if err != nil {
		return errors.Trace(err)
	}

	return nil
}

func QueryAllConfigs(etcdcli *EtcdCli, cli *clientv3.Client) ([]EtcdVariables, error) {

	var allvars []EtcdVariables

	ctx, cancel := context.WithTimeout(context.Background(), etcdcli.RequestTimeout)
	resp, err := cli.Get(ctx, etcdcli.Root, clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))
	cancel()
	if err != nil {
		return nil, errors.Trace(err)
	}
	for _, ev := range resp.Kvs {
		var tmpvar EtcdVariables
		value, _ := base64.StdEncoding.DecodeString(string(ev.Value))
		//decode user information.
		if err := json.Unmarshal(value, &tmpvar); err != nil {
			return nil, errors.Trace(err)
		}
		//merge user informations.
		allvars = append(allvars, tmpvar)

	}
	return allvars, nil
}
