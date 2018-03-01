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
	EtcdRHG struct {
		WriterHostgroup uint64 `json:"writer_hostgroup" db:"writer_hostgroup"`
		ReaderHostgroup uint64 `json:"reader_hostgroup" db:"reader_hostgroup"`
		Comment         string `json:"comment" db:"comment"`
	}
)

// new replication hostgroup instance
func NewRHG(whg uint64, rhg uint64) (*EtcdRHG, error) {

	newrhg := new(EtcdRHG)

	newrhg.WriterHostgroup = whg
	newrhg.ReaderHostgroup = rhg
	newrhg.Comment = ""

	// return new replication hostgroup instance.
	return newrhg, nil
}

// set writer hostgroup
func (rhg *EtcdRHG) SetWriterHostGroup(writer uint64) {
	rhg.WriterHostgroup = writer
}

// set reader hostgroup
func (rhg *EtcdRHG) SetReaderHostGroup(reader uint64) {
	rhg.ReaderHostgroup = reader
}

// set comment
func (rhg *EtcdRHG) SetComment(comment string) {
	rhg.Comment = comment
}

// create or update a rhg
func (rhg *EtcdRHG) CreateOrUpdateOneRHG(etcdcli *EtcdCli, cli *clientv3.Client) error {

	rw := fmt.Sprintf("%d|%d", rhg.WriterHostgroup, rhg.ReaderHostgroup)
	key := []byte(rw)

	value, err := json.Marshal(rhg)
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

// delete a rhg
func (rhg *EtcdRHG) DeleteOneRHG(etcdcli *EtcdCli, cli *clientv3.Client) error {

	rw := fmt.Sprintf("%d|%d", rhg.WriterHostgroup, rhg.ReaderHostgroup)
	key := []byte(rw)

	// base64
	encodeKey := base64.StdEncoding.EncodeToString(key)

	ctx, cancel := context.WithTimeout(context.Background(), etcdcli.RequestTimeout)
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

// query all replications hostgroups
func QueryAllRHG(etcdcli *EtcdCli, cli *clientv3.Client) ([]EtcdRHG, error) {
	var allrhgs []EtcdRHG

	ctx, cancel := context.WithTimeout(context.Background(), etcdcli.RequestTimeout)
	resp, err := cli.Get(ctx, etcdcli.Root, clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))
	cancel()
	if err != nil {
		return nil, errors.Trace(err)
	}
	for _, ev := range resp.Kvs {
		var tmprhg EtcdRHG
		value, _ := base64.StdEncoding.DecodeString(string(ev.Value))
		//decode user information.
		if err := json.Unmarshal(value, &tmprhg); err != nil {
			return nil, errors.Trace(err)
		}
		//merge user informations.
		allrhgs = append(allrhgs, tmprhg)

	}
	return allrhgs, nil
}
