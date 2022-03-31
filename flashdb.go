package main

import (
	"github.com/arriqaaq/flashdb"
)

type flashdbEngine struct {
	db   *flashdb.FlashDB
	path string
}

func newFlashDB(path string) (kvEngine, error) {
	opts := flashdb.DefaultConfig()
	opts.EvictionInterval = 0
	opts.Path = path
	opts.NoSync = false

	db, err := flashdb.New(opts)
	if err != nil {
		return nil, err
	}
	return &flashdbEngine{db: db, path: path}, err
}

func (db *flashdbEngine) Put(key []byte, value []byte) error {
	if err := db.db.Update(
		func(tx *flashdb.Tx) error {
			return tx.Set(string(key), string(value))
		}); err != nil {
		return err
	}
	return nil
}

func (db *flashdbEngine) Get(key []byte) ([]byte, error) {
	var val string
	var err error
	if err := db.db.View(
		func(tx *flashdb.Tx) error {
			if val, err = tx.Get(string(key)); err != nil {
				return err
			}
			return nil
		}); err != nil {
		return nil, err
	}
	return []byte(val), nil
}

func (db *flashdbEngine) Close() error {
	return db.db.Close()
}

func (db *flashdbEngine) FileSize() (int64, error) {
	return dirSize(db.path)
}
