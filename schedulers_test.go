package etcd

import (
	"log"
	"testing"
)

func TestCreateOrUpdateOneSchld(t *testing.T) {
	etcdcli := NewEtcdCli([]string{etcd_points})

	etcdcli.SetPrefix(etcd_prefix)
	etcdcli.SetService(etcd_service)
	etcdcli.SetEtcdType("schedulers")

	etcdcli.MakeWatchRoot()

	cli, err := etcdcli.OpenEtcd()
	if err != nil {
		t.Error(err)
	}

	newschld, err := NewSch("/bin/ls", 1000)
	if err != nil {
		t.Error(err)
	}

	newschld.SetSchedulerId(999)

	err = newschld.CreateOrUpdateOneSchld(etcdcli, cli)
	if err != nil {
		t.Error(err)
	}

	newschld.SetSchedulerId(999)
	newschld.SetSchedulerArg1("-l")

	err = newschld.CreateOrUpdateOneSchld(etcdcli, cli)
	if err != nil {
		t.Error(err)
	}

}

func TestQueryAllSchlds(t *testing.T) {
	etcdcli := NewEtcdCli([]string{etcd_points})

	etcdcli.SetPrefix(etcd_prefix)
	etcdcli.SetService(etcd_service)
	etcdcli.SetEtcdType("schedulers")

	etcdcli.MakeWatchRoot()

	cli, err := etcdcli.OpenEtcd()
	if err != nil {
		t.Error(err)
	}

	allschlds, err := QueryAllSchlds(etcdcli, cli)
	if err != nil {
		t.Error(err)
	}

	log.Println(allschlds)
}

func TestDeleteOneSchlds(t *testing.T) {
	etcdcli := NewEtcdCli([]string{etcd_points})

	etcdcli.SetPrefix(etcd_prefix)
	etcdcli.SetService(etcd_service)
	etcdcli.SetEtcdType("schedulers")

	etcdcli.MakeWatchRoot()

	cli, err := etcdcli.OpenEtcd()
	if err != nil {
		t.Error(err)
	}

	newschld, err := NewSch("/bin/ls", 1000)
	if err != nil {
		t.Error(err)
	}

	newschld.SetSchedulerId(999)

	newschld.DeleteOneSchld(etcdcli, cli)
	if err != nil {
		t.Error(err)
	}

}
