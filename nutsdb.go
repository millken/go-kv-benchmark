package main

import (
	"github.com/xujiajun/nutsdb"
)

type nutsdbEngine struct {
	db   *nutsdb.DB
	path string
}

var bucket001 = "bucket001"

func newNutsDB(path string) (kvEngine, error) {

	db, err := nutsdb.Open(
		nutsdb.DefaultOptions,
		nutsdb.WithSyncEnable(false),
		nutsdb.WithRWMode(nutsdb.MMap),
		nutsdb.WithDir(path),
	)
	if err != nil {
		return nil, err
	}
	return &nutsdbEngine{db: db, path: path}, err
}

func (db *nutsdbEngine) Put(key []byte, value []byte) error {

	if err := db.db.Update(
		func(tx *nutsdb.Tx) error {
			err := tx.Put(bucket001, key, value, nutsdb.Persistent)
			return err
		}); err != nil {
		return err
	}
	return nil
}

func (db *nutsdbEngine) Get(key []byte) ([]byte, error) {
	var e *nutsdb.Entry
	var err error
	if err := db.db.View(
		func(tx *nutsdb.Tx) error {
			if e, err = tx.Get(bucket001, key); err != nil {
				return err
			}
			return nil
		}); err != nil {
		return nil, err
	}
	return e.Value, nil
}

func (db *nutsdbEngine) Close() error {
	return db.db.Close()
}

func (db *nutsdbEngine) FileSize() (int64, error) {
	return dirSize(db.path)
}
