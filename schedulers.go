package etcd

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/coreos/etcd/clientv3"
	"github.com/juju/errors"
)

type (
	EtcdSchld struct {
		Id         int64  `json:"id" db:"id"`
		Active     int64  `json:"active" db:"active"`
		IntervalMs int64  `json:"interval_ms" db:"interval_ms"`
		FileName   string `json:"filename" db:"filename"`
		Arg1       string `json:"arg1" db:"arg1"`
		Arg2       string `json:"arg2" db:"arg2"`
		Arg3       string `json:"arg3" db:"arg3"`
		Arg4       string `json:"arg4" db:"arg4"`
		Arg5       string `json:"arg5" db:"arg5"`
		Comment    string `json:"comment" db:"comment"`
	}
)

//new NewSch
func NewSch(filename string, interval_ms int64) (*EtcdSchld, error) {

	sched := new(EtcdSchld)

	sched.FileName = filename
	switch {
	case interval_ms < 100:
		sched.IntervalMs = 100
	case interval_ms > 100000000:
		sched.IntervalMs = 100000000
	default:
		sched.IntervalMs = interval_ms
	}

	sched.Active = 0
	sched.Arg1 = "NULL"
	sched.Arg2 = "NULL"
	sched.Arg3 = "NULL"
	sched.Arg4 = "NULL"
	sched.Arg5 = "NULL"

	return sched, nil

}

// Set Scheduler id
func (sched *EtcdSchld) SetSchedulerId(id int64) {
	sched.Id = id
}

// Set Scheduler Active status
func (sched *EtcdSchld) SetSchedulerActive(active int64) {
	if active >= 1 {
		sched.Active = 1
	} else {
		sched.Active = 0
	}
}

// Set Scheduler all Args
func (sched *EtcdSchld) SetSchedulerArg1(arg1 string) {
	sched.Arg1 = arg1
}

func (sched *EtcdSchld) SetSchedulerArg2(arg2 string) {
	sched.Arg2 = arg2
}

func (sched *EtcdSchld) SetSchedulerArg3(arg3 string) {
	sched.Arg3 = arg3
}

func (sched *EtcdSchld) SetSchedulerArg4(arg4 string) {
	sched.Arg4 = arg4
}

func (sched *EtcdSchld) SetSchedulerArg5(arg5 string) {
	sched.Arg5 = arg5
}

// Set scheduler interval_ms
func (sched *EtcdSchld) SetSchedulerIntervalMs(interval_ms int64) {
	switch {
	case interval_ms < 100:
		sched.IntervalMs = 100
	case interval_ms > 100000000:
		sched.IntervalMs = 100000000
	default:
		sched.IntervalMs = interval_ms
	}
}

// create or update a scheduler
func (schld *EtcdSchld) CreateOrUpdateOneSchld(etcdcli *EtcdCli, cli *clientv3.Client) error {

	schld_id := fmt.Sprintf("%d", schld.Id)
	key := []byte(schld_id)

	value, err := json.Marshal(schld)
	if err != nil {
		return errors.Trace(err)
	}

	// base64
	encodeKey := base64.StdEncoding.EncodeToString(key)
	encodeValue := base64.StdEncoding.EncodeToString(value)

	ctx, cancel := context.WithTimeout(context.Background(), etcdcli.RequestTimeout)
	//create user
	_, err = cli.Put(ctx, etcdcli.Root+"/"+encodeKey, encodeValue)
	cancel()
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

// create or update a scheduler
func (schld *EtcdSchld) DeleteOneSchld(etcdcli *EtcdCli, cli *clientv3.Client) error {
	schld_id := fmt.Sprintf("%d", schld.Id)
	key := []byte(schld_id)

	// base64
	encodeKey := base64.StdEncoding.EncodeToString(key)

	ctx, cancel := context.WithTimeout(context.Background(), etcdcli.RequestTimeout)
	_, err := cli.Delete(ctx, etcdcli.Root+"/"+encodeKey, clientv3.WithPrefix())
	cancel()
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

// create or update a scheduler
func QueryAllSchlds(etcdcli *EtcdCli, cli *clientv3.Client) ([]EtcdSchld, error) {
	var allschlds []EtcdSchld

	ctx, cancel := context.WithTimeout(context.Background(), etcdcli.RequestTimeout)
	resp, err := cli.Get(ctx, etcdcli.Root, clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))
	cancel()
	if err != nil {
		return nil, errors.Trace(err)
	}
	for _, ev := range resp.Kvs {
		var tmpschld EtcdSchld
		value, _ := base64.StdEncoding.DecodeString(string(ev.Value))
		//decode user information.
		if err := json.Unmarshal(value, &tmpschld); err != nil {
			return nil, errors.Trace(err)
		}
		//merge user informations.
		allschlds = append(allschlds, tmpschld)

	}
	return allschlds, nil
}
