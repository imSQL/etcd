package users

import (
	"context"
	"encoding/base64"
	"encoding/json"

	"github.com/coreos/etcd/clientv3"
	"github.com/imSQL/etcd/connections"
	"github.com/imSQL/imSQL"
	"github.com/imSQL/proxysql"
	"github.com/juju/errors"
)

type (
	Users struct {
		// proxysql users
		ProxysqlUsers proxysql.Users
		// mysql users
		imSQLUsers imSQL.Users
	}
)

func (usr *Users) CreateOrUpdateOneUser(etcdcli *connections.EtcdCli, usr *Users) error {

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
	_, err = etcdcli.cli.Put(ctx, etcdcli.Root+"/"+encodeKey, encodeValue)
	cancel()
	if err != nil {
		return errors.Trace(err)
	}

}

func (usr *Users) DeleteOneUser(etcdcli *connections.EtcdCli, usr *Users) error {
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
	//encodeValue := base64.StdEncoding.EncodeToString(value)

	ctx, cancel := context.WithTimeout(context.Background(), etcdcli.RequestTimeout)

	//create user at etcd
	_, err = etcdcli.cli.Delete(ctx, etcdcli.Root+"/"+encodeKey, clientv3.WithPrefix())
	cancel()
	if err != nil {
		return errors.Trace(err)
	}

}

func (usr *Users) QueryAllUsers() error {

}
