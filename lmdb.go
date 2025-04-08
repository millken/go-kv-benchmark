package main

import (
	"os"

	"github.com/PowerDNS/lmdb-go/lmdb"
)

type lmdbEngine struct {
	env  *lmdb.Env
	db   lmdb.DBI
	path string
}

var bucketlmdb = "bucketlmdb"

func newLmdb(path string) (kvEngine, error) {
	if err := os.MkdirAll(path, 0755); err != nil {
		return nil, err
	}
	env, err := lmdb.NewEnv()
	if err != nil {
		return nil, err
	}
	env.Sync(false)
	err = env.SetMapSize((1000 << 20)) // 1GB
	if err != nil {
		return nil, err
	}
	err = env.SetMaxDBs(11)
	if err != nil {
		return nil, err
	}
	err = env.SetMaxReaders(100)
	if err != nil {
		return nil, err
	}
	err = env.Open(path, 0, 0644)
	if err != nil {
		return nil, err
	}
	var dbi lmdb.DBI
	err = env.Update(func(txn *lmdb.Txn) (err error) {
		dbi, err = txn.CreateDBI(bucketlmdb)
		return err
	})
	if err != nil {
		return nil, err
	}
	return &lmdbEngine{db: dbi, env: env, path: path}, err
}

func (db *lmdbEngine) Put(key []byte, value []byte) error {
	return db.env.Update(func(txn *lmdb.Txn) (err error) {
		return txn.Put(db.db, key, value, 0)
	})
}

func (db *lmdbEngine) Get(key []byte) ([]byte, error) {
	var val []byte
	err := db.env.View(func(txn *lmdb.Txn) (err error) {
		val, err = txn.Get(db.db, key)
		if err != nil {
			return err
		}
		return nil
	})
	return val, err
}

func (db *lmdbEngine) Close() error {
	return db.env.Close()
}

func (db *lmdbEngine) FileSize() (int64, error) {
	return dirSize(db.path)
}
