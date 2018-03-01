package etcd

import (
	"log"
	"testing"
)

func TestUpdateOneConfig(t *testing.T) {
	etcdcli := NewEtcdCli([]string{etcd_points})

	etcdcli.SetPrefix(etcd_prefix)
	etcdcli.SetService(etcd_service)
	etcdcli.SetEtcdType("variables")

	etcdcli.MakeWatchRoot()

	cli, err := etcdcli.OpenEtcd()
	if err != nil {
		t.Error(err)
	}

	vars := new(EtcdVariables)
	vars.VariablesName = "mysql-wait_timeout"
	vars.Value = "9797"

	err = CreateOrUpdateOneConfig(etcdcli, cli, vars)
	if err != nil {
		t.Error(err)
	}

	res, err := QueryAllConfigs(etcdcli, cli)
	if err != nil {
		t.Error(err)
	}
	log.Println(res)

	vars.Value = "19898"
	err = CreateOrUpdateOneConfig(etcdcli, cli, vars)
	if err != nil {
		t.Error(err)
	}

	res, err = QueryAllConfigs(etcdcli, cli)
	if err != nil {
		t.Error(err)
	}
	log.Println(res)
}
