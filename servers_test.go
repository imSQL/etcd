package etcd

import (
	"log"
	"testing"
)

func TestCreateOrUpdateOneBackend(t *testing.T) {
	etcdcli := NewEtcdCli([]string{etcd_points})

	etcdcli.SetPrefix(etcd_prefix)
	etcdcli.SetService(etcd_service)
	etcdcli.SetEtcdType("servers")

	etcdcli.MakeWatchRoot()

	cli, err := etcdcli.OpenEtcd()
	if err != nil {
		t.Error(err)
	}

	newbkd, err := NewServer(0, "127.0.0.1", 6032)
	if err != nil {
		t.Error(err)
	}

	err = newbkd.CreateOrUpdateOneBackend(etcdcli, cli)
	if err != nil {
		t.Error(err)
	}

	newbkd.SetServerMaxConnection(999)

	err = newbkd.CreateOrUpdateOneBackend(etcdcli, cli)
	if err != nil {
		t.Error(err)
	}

}

func TestQueryAllBackends(t *testing.T) {
	etcdcli := NewEtcdCli([]string{etcd_points})

	etcdcli.SetPrefix(etcd_prefix)
	etcdcli.SetService(etcd_service)
	etcdcli.SetEtcdType("servers")

	etcdcli.MakeWatchRoot()

	cli, err := etcdcli.OpenEtcd()
	if err != nil {
		t.Error(err)
	}

	bkds, err := QueryAllBackends(etcdcli, cli)
	if err != nil {
		t.Error(err)
	}

	log.Println(bkds)
}

func TestDeleteOneBackend(t *testing.T) {
	etcdcli := NewEtcdCli([]string{etcd_points})

	etcdcli.SetPrefix(etcd_prefix)
	etcdcli.SetService(etcd_service)
	etcdcli.SetEtcdType("servers")

	etcdcli.MakeWatchRoot()

	cli, err := etcdcli.OpenEtcd()
	if err != nil {
		t.Error(err)
	}

	newbkd, err := NewServer(0, "127.0.0.1", 6032)
	if err != nil {
		t.Error(err)
	}

	err = newbkd.DeleteOneBackend(etcdcli, cli)
	if err != nil {
		t.Error(err)
	}

}
