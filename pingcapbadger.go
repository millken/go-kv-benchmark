package main

import (
	"fmt"

	"github.com/pingcap/badger"
)

func newPingcapBadger(path string) (kvEngine, error) {
	opts := badger.DefaultOptions
	fmt.Println(path)
	opts.Dir = path
	opts.ValueDir = path
	opts.SyncWrites = false
	db, err := badger.Open(opts)
	return &pingcapBadgerEngine{db: db, path: path}, err
}

type pingcapBadgerEngine struct {
	path string
	db   *badger.DB
}

func (db *pingcapBadgerEngine) Put(key []byte, value []byte) error {
	return db.db.Update(func(tx *badger.Txn) error {
		return tx.Set(key, value)
	})
}

func (db *pingcapBadgerEngine) Get(key []byte) ([]byte, error) {
	var val []byte
	err := db.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		v, err := item.ValueCopy(nil)
		if err != nil {
			return err
		}
		val = v
		return nil
	})
	return val, err
}

func (db *pingcapBadgerEngine) Close() error {
	return db.db.Close()
}

func (db *pingcapBadgerEngine) FileSize() (int64, error) {
	return dirSize(db.path)
}
