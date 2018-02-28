package etcd

import (
	"testing"

	"github.com/imSQL/etcd/connections"
	"github.com/imSQL/etcd/users"
)

func TestCreateOrUpdateOneUsers(t *testing.T) {
	etcdcli := connections.NewEtcdCli([]string{etcd_points})

	etcdcli.SetPrefix(etcd_prefix)
	etcdcli.SetService(etcd_service)
	etcdcli.SetEtcdType("users")

	etcdcli.MakeWatchRoot()

	cli, err := etcdcli.OpenEtcd()
	if err != nil {
		t.Error(err)
	}

	_, err = users.QueryAllUsers(etcdcli, cli)
	if err != nil {
		t.Error(err)
	}
}
