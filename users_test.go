package etcd

import (
	"log"
	"testing"
)

func TestCreateOrUpdateOneUser(t *testing.T) {
	etcdcli := NewEtcdCli([]string{etcd_points})

	etcdcli.SetPrefix(etcd_prefix)
	etcdcli.SetService(etcd_service)
	etcdcli.SetEtcdType("users")

	etcdcli.MakeWatchRoot()

	cli, err := etcdcli.OpenEtcd()
	if err != nil {
		t.Error(err)
	}

	newusr, err := NewUser("dev", "dev", 0, "dev")
	if err != nil {
		t.Error(err)
	}

	err = newusr.CreateOrUpdateOneUser(etcdcli, cli)
	if err != nil {
		t.Error(err)
	}

	newusr.SetMaxConnections(999)

	err = newusr.CreateOrUpdateOneUser(etcdcli, cli)
	if err != nil {
		t.Error(err)
	}

}

func TestQueryAllUsers(t *testing.T) {
	etcdcli := NewEtcdCli([]string{etcd_points})

	etcdcli.SetPrefix(etcd_prefix)
	etcdcli.SetService(etcd_service)
	etcdcli.SetEtcdType("users")

	etcdcli.MakeWatchRoot()

	cli, err := etcdcli.OpenEtcd()
	if err != nil {
		t.Error(err)
	}

	usrs, err := QueryAllUsers(etcdcli, cli)
	if err != nil {
		t.Error(err)
	}

	log.Println(usrs)
}

func TestDeleteOneUser(t *testing.T) {
	etcdcli := NewEtcdCli([]string{etcd_points})

	etcdcli.SetPrefix(etcd_prefix)
	etcdcli.SetService(etcd_service)
	etcdcli.SetEtcdType("users")

	etcdcli.MakeWatchRoot()

	cli, err := etcdcli.OpenEtcd()
	if err != nil {
		t.Error(err)
	}

	newusr, err := NewUser("dev", "dev", 0, "dev")
	if err != nil {
		t.Error(err)
	}

	err = newusr.DeleteOneUser(etcdcli, cli)
	if err != nil {
		t.Error(err)
	}

}
