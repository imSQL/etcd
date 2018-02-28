package etcd

import (
	"os"
	"testing"

	"github.com/imSQL/etcd/connections"
)

var etcd_points = os.Getenv("ETCD_ADDR")
var etcd_prefix = os.Getenv("ETCD_PREFIX")
var etcd_service = os.Getenv("ETCD_SVC")

func TestConnection(t *testing.T) {

	etcdcli := connections.NewEtcdCli([]string{etcd_points})

	etcdcli.SetPrefix(etcd_prefix)
	etcdcli.SetService(etcd_service)
	etcdcli.SetEtcdType("users")

	etcdcli.MakeWatchRoot()

	_, err := etcdcli.OpenEtcd()
	if err != nil {
		t.Error(err)
	}
}
