package etcd

import (
	"log"
	"testing"
)

func TestCreateOrUpdateOneRHG(t *testing.T) {
	etcdcli := NewEtcdCli([]string{etcd_points})

	etcdcli.SetPrefix(etcd_prefix)
	etcdcli.SetService(etcd_service)
	etcdcli.SetEtcdType("users")

	etcdcli.MakeWatchRoot()

	cli, err := etcdcli.OpenEtcd()
	if err != nil {
		t.Error(err)
	}

	newrhg, err := NewRHG(1, 2)
	if err != nil {
		t.Error(err)
	}

	err = newrhg.CreateOrUpdateOneRHG(etcdcli, cli)
	if err != nil {
		t.Error(err)
	}

	newrhg.SetWriterHostGroup(1)
	newrhg.SetReaderHostGroup(3)

	err = newrhg.CreateOrUpdateOneRHG(etcdcli, cli)
	if err != nil {
		t.Error(err)
	}

}

func TestQueryAllRHG(t *testing.T) {
	etcdcli := NewEtcdCli([]string{etcd_points})

	etcdcli.SetPrefix(etcd_prefix)
	etcdcli.SetService(etcd_service)
	etcdcli.SetEtcdType("users")

	etcdcli.MakeWatchRoot()

	cli, err := etcdcli.OpenEtcd()
	if err != nil {
		t.Error(err)
	}

	allrhgs, err := QueryAllRHG(etcdcli, cli)
	if err != nil {
		t.Error(err)
	}

	log.Println(allrhgs)
}

func TestDeleteOneRHG(t *testing.T) {
	etcdcli := NewEtcdCli([]string{etcd_points})

	etcdcli.SetPrefix(etcd_prefix)
	etcdcli.SetService(etcd_service)
	etcdcli.SetEtcdType("users")

	etcdcli.MakeWatchRoot()

	cli, err := etcdcli.OpenEtcd()
	if err != nil {
		t.Error(err)
	}

	newrhg, err := NewRHG(1, 3)
	if err != nil {
		t.Error(err)
	}

	err = newrhg.DeleteOneRHG(etcdcli, cli)
	if err != nil {
		t.Error(err)
	}

}
