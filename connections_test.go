package etcd

import (
	"os"
	"testing"
)

var etcd_points = os.Getenv("ETCD_ADDR")
var etcd_prefix = os.Getenv("ETCD_PREFIX")
var etcd_service = os.Getenv("ETCD_SVC")

func TestConnections(t *testing.T) {

	etcdcli := NewEtcdCli([]string{etcd_points})

	etcdcli.SetPrefix(etcd_prefix)
	etcdcli.SetService(etcd_service)
	etcdcli.SetEtcdType("users")

	etcdcli.MakeWatchRoot()

	_, err := etcdcli.OpenEtcd()
	if err != nil {
		t.Error(err)
	}
}
