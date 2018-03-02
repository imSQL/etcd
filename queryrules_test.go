package etcd

import (
	"log"
	"testing"
)

func TestCreateOrUpdateOneQr(t *testing.T) {
	etcdcli := NewEtcdCli([]string{etcd_points})

	etcdcli.SetPrefix(etcd_prefix)
	etcdcli.SetService(etcd_service)
	etcdcli.SetEtcdType("queryrules")

	etcdcli.MakeWatchRoot()

	cli, err := etcdcli.OpenEtcd()
	if err != nil {
		t.Error(err)
	}

	newqr, err := NewQr("dev")
	if err != nil {
		t.Error(err)
	}
	newqr.SetQrRuleid(999)
	newqr.SetQrLog("0")
	newqr.SetQrCacheTTL("100")
	newqr.SetQrReconnect("0")
	newqr.SetQrRetries("5")

	err = newqr.CreateOrUpdateOneQr(etcdcli, cli)
	if err != nil {
		t.Error(err)
	}

	newqr.SetQrRuleid(999)
	newqr.SetQrLog("0")
	newqr.SetQrDigest("^SELECT")
	newqr.SetQrCacheTTL("100")
	newqr.SetQrReconnect("0")
	newqr.SetQrRetries("5")

	err = newqr.CreateOrUpdateOneQr(etcdcli, cli)
	if err != nil {
		t.Error(err)
	}

}

func TestQueryAllQrs(t *testing.T) {
	etcdcli := NewEtcdCli([]string{etcd_points})

	etcdcli.SetPrefix(etcd_prefix)
	etcdcli.SetService(etcd_service)
	etcdcli.SetEtcdType("queryrules")

	etcdcli.MakeWatchRoot()

	cli, err := etcdcli.OpenEtcd()
	if err != nil {
		t.Error(err)
	}

	qrs, err := QueryAllQrs(etcdcli, cli)
	if err != nil {
		t.Error(err)
	}

	log.Println(qrs)
}

func TestDeleteOneQr(t *testing.T) {
	etcdcli := NewEtcdCli([]string{etcd_points})

	etcdcli.SetPrefix(etcd_prefix)
	etcdcli.SetService(etcd_service)
	etcdcli.SetEtcdType("queryrules")

	etcdcli.MakeWatchRoot()

	cli, err := etcdcli.OpenEtcd()
	if err != nil {
		t.Error(err)
	}

	newqr, err := NewQr("dev")
	if err != nil {
		t.Error(err)
	}

	newqr.SetQrRuleid(999)
	newqr.SetQrLog("0")
	newqr.SetQrCacheTTL("100")
	newqr.SetQrReconnect("0")
	newqr.SetQrRetries("5")

	err = newqr.DeleteOneQr(etcdcli, cli)
	if err != nil {
		t.Error(err)
	}

}
