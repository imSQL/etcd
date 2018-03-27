package etcd

import (
	"context"
	"encoding/base64"
	"encoding/json"

	"github.com/coreos/etcd/clientv3"
	"github.com/juju/errors"
)

type (
	EtcdUsers struct {
		Username              string   `db:"username" json:"username"`
		Password              string   `db:"password" json:"password"`
		Active                uint64   `db:"active" json:"active"`
		UseSsl                uint64   `db:"use_ssl" json:"use_ssl"`
		DefaultHostgroup      uint64   `db:"default_hostgroup" json:"default_hostgroup"`
		DefaultSchema         string   `db:"default_schema" json:"default_schema"`
		SchemaLocked          uint64   `db:"schema_locked" json:"schema_locked"`
		TransactionPersistent uint64   `db:"transaction_persistent" json:"transaction_persistent"`
		FastForward           uint64   `db:"fast_forward" json:"fast_forward"`
		Backend               uint64   `db:"backend" json:"backend"`
		Frontend              uint64   `db:"frontend" json:"frontend"`
		MaxConnections        uint64   `db:"max_connections" json:"max_connections"`
		Privileges            []string `db:"privileges" json:"privileges"`
	}
)

func NewUser(username string, password string, default_hostgroup uint64, default_schema string) (*EtcdUsers, error) {

	newuser := new(EtcdUsers)

	newuser.Username = username
	newuser.Password = password
	newuser.DefaultHostgroup = default_hostgroup
	newuser.DefaultSchema = default_schema

	newuser.Active = 1
	newuser.UseSsl = 0
	newuser.SchemaLocked = 0
	newuser.TransactionPersistent = 0
	newuser.FastForward = 0
	newuser.Backend = 1
	newuser.Frontend = 1
	newuser.MaxConnections = 10000
	newuser.Privileges = []string{"ALL PRIVILEGES"}

	return newuser, nil
}

// set fast_forward
func (users *EtcdUsers) SetFastForward(fast_forward uint64) {
	if fast_forward >= 1 {
		users.FastForward = 1
	} else {
		users.FastForward = 0
	}
}

// set max_connections
func (users *EtcdUsers) SetMaxConnections(max_connections uint64) {
	switch {
	case max_connections >= 10000:
		users.MaxConnections = 10000
	case max_connections <= 1:
		users.MaxConnections = 1
	default:
		users.MaxConnections = max_connections
	}
}

// set backend
func (users *EtcdUsers) SetBackend(backend uint64) {
	if backend >= 1 {
		users.Backend = 1
	} else {
		users.Backend = 0
	}
}

// set fronted
func (users *EtcdUsers) SetFrontend(frontend uint64) {
	if frontend >= 1 {
		users.Frontend = 1
	} else {
		users.Frontend = 0
	}
}

// set user active/disactive
func (users *EtcdUsers) SetUserActive(active uint64) {
	if active >= 1 {
		users.Active = 1
	} else {
		users.Active = 0
	}
}

// Set users UseSSL
func (users *EtcdUsers) SetUseSSL(use_ssl uint64) {
	if use_ssl >= 1 {
		users.UseSsl = 1
	} else {
		users.UseSsl = 0
	}
}

// set users SchemaLocked
func (users *EtcdUsers) SetSchemaLocked(schema_locked uint64) {
	if schema_locked >= 1 {
		users.SchemaLocked = 1
	} else {
		users.SchemaLocked = 0
	}
}

// set users transaction_persistent
func (users *EtcdUsers) SetTransactionPersistent(transaction_persistent uint64) {
	if transaction_persistent >= 1 {
		users.TransactionPersistent = 1
	} else {
		users.TransactionPersistent = 0
	}
}

// set privileges
func (users *EtcdUsers) AddPrivileges(privileges ...string) {
	if len(privileges) != 0 {
		if users.Privileges[0] == "ALL PRIVILEGES" {
			users.Privileges = []string{}
			users.Privileges = append(users.Privileges, privileges...)
		} else {
			users.Privileges = append(users.Privileges, privileges...)
		}
	}
}

func (usr *EtcdUsers) CreateOrUpdateOneUser(etcdcli *EtcdCli, cli *clientv3.Client) error {

	key := []byte(usr.Username)
	value, err := json.Marshal(usr)
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

func (usr *EtcdUsers) DeleteOneUser(etcdcli *EtcdCli, cli *clientv3.Client) error {

	key := []byte(usr.Username)
	//value, err := json.Marshal(newuser)
	//if err != nil {
	//	return errors.Trace(err)
	//}

	encodeKey := base64.StdEncoding.EncodeToString(key)
	//encodeValue := base64.StdEncoding.EncodeToString(value)

	ctx, cancel := context.WithTimeout(context.Background(), etcdcli.RequestTimeout)

	//create user at etcd
	response, err := cli.Delete(ctx, etcdcli.Root+"/"+encodeKey, clientv3.WithPrefix())
	cancel()
	if err != nil {
		return errors.Trace(err)
	}

	orig_key, err := base64.StdEncoding.DecodeString(encodeKey)
	if err != nil {
		return errors.Trace(err)
	}

	if response.Deleted == 0 {
		return errors.NotFoundf(string(orig_key))
	}

	return nil

}

func QueryAllUsers(etcdcli *EtcdCli, cli *clientv3.Client) ([]EtcdUsers, error) {
	// new users handler
	var allusers []EtcdUsers

	ctx, cancel := context.WithTimeout(context.Background(), etcdcli.RequestTimeout)
	resp, err := cli.Get(ctx, etcdcli.Root, clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))
	cancel()
	if err != nil {
		return nil, errors.Trace(err)
	}
	for _, ev := range resp.Kvs {
		var tmpusr EtcdUsers
		value, _ := base64.StdEncoding.DecodeString(string(ev.Value))
		//decode user information.
		if err := json.Unmarshal(value, &tmpusr); err != nil {
			return nil, errors.Trace(err)
		}
		//merge user informations.
		allusers = append(allusers, tmpusr)

	}
	return allusers, nil
}
