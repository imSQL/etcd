package users

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/coreos/etcd/clientv3"
	"github.com/imSQL/etcd/connections"
	"github.com/imSQL/imSQL"
	"github.com/imSQL/proxysql"
	"github.com/juju/errors"
)

type (
	EtcdUsers struct {
		// proxysql users
		ProxysqlUsers proxysql.Users
		// mysql users
		imSQLUsers imSQL.Users
	}
)

func (usr *EtcdUsers) CreateOrUpdateOneUser(etcdcli *connections.EtcdCli, cli *clientv3.Client) error {

	// new users handler
	newuser, err := proxysql.NewUser(usr.ProxysqlUsers.Username, usr.ProxysqlUsers.Password, usr.ProxysqlUsers.DefaultHostgroup, usr.ProxysqlUsers.DefaultSchema)
	if err != nil {
		return errors.Trace(err)
	}

	newuser.SetUserActive(usr.ProxysqlUsers.Active)
	newuser.SetFastForward(usr.ProxysqlUsers.FastForward)
	newuser.SetBackend(usr.ProxysqlUsers.Backend)
	newuser.SetFrontend(usr.ProxysqlUsers.Frontend)
	newuser.SetMaxConnections(usr.ProxysqlUsers.MaxConnections)
	newuser.SetSchemaLocked(usr.ProxysqlUsers.SchemaLocked)
	newuser.SetTransactionPersistent(usr.ProxysqlUsers.TransactionPersistent)
	newuser.SetUseSSL(usr.ProxysqlUsers.UseSsl)

	key := []byte(newuser.Username)
	value, err := json.Marshal(newuser)
	if err != nil {
		return errors.Trace(err)
	}

	encodeKey := base64.StdEncoding.EncodeToString(key)
	encodeValue := base64.StdEncoding.EncodeToString(value)

	ctx, cancel := context.WithTimeout(context.Background(), etcdcli.RequestTimeout)

	//create user at etcd
	_, err = cli.Put(ctx, etcdcli.Root+"/"+encodeKey, encodeValue)
	cancel()
	if err != nil {
		return errors.Trace(err)
	}

	return nil
}

func (usr *EtcdUsers) DeleteOneUser(etcdcli *connections.EtcdCli, cli *clientv3.Client) error {
	// new users handler
	newuser, err := proxysql.NewUser(usr.ProxysqlUsers.Username, usr.ProxysqlUsers.Password, usr.ProxysqlUsers.DefaultHostgroup, usr.ProxysqlUsers.DefaultSchema)
	if err != nil {
		return errors.Trace(err)
	}

	newuser.SetUserActive(usr.ProxysqlUsers.Active)
	newuser.SetFastForward(usr.ProxysqlUsers.FastForward)
	newuser.SetBackend(usr.ProxysqlUsers.Backend)
	newuser.SetFrontend(usr.ProxysqlUsers.Frontend)
	newuser.SetMaxConnections(usr.ProxysqlUsers.MaxConnections)
	newuser.SetSchemaLocked(usr.ProxysqlUsers.SchemaLocked)
	newuser.SetTransactionPersistent(usr.ProxysqlUsers.TransactionPersistent)
	newuser.SetUseSSL(usr.ProxysqlUsers.UseSsl)

	key := []byte(newuser.Username)
	//value, err := json.Marshal(newuser)
	//if err != nil {
	//	return errors.Trace(err)
	//}

	encodeKey := base64.StdEncoding.EncodeToString(key)
	//encodeValue := base64.StdEncoding.EncodeToString(value)

	ctx, cancel := context.WithTimeout(context.Background(), etcdcli.RequestTimeout)

	//create user at etcd
	_, err = cli.Delete(ctx, etcdcli.Root+"/"+encodeKey, clientv3.WithPrefix())
	cancel()
	if err != nil {
		return errors.Trace(err)
	}

	return nil

}

func QueryAllUsers(etcdcli *connections.EtcdCli, cli *clientv3.Client) ([]EtcdUsers, error) {
	// new users handler
	var allusers []EtcdUsers

	ctx, cancel := context.WithTimeout(context.Background(), etcdcli.RequestTimeout)
	resp, err := cli.Get(ctx, etcdcli.Root, clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))
	cancel()
	if err != nil {
		return nil, errors.Trace(err)
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
	return allusers, nil
}
