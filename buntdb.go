package main

import (
	"github.com/tidwall/buntdb"
)

type buntdbEngine struct {
	db   *buntdb.DB
	path string
}

func newBuntDB(path string) (kvEngine, error) {

	db, err := buntdb.Open(path)
	if err != nil {
		return nil, err
	}
	return &buntdbEngine{db: db, path: path}, err
}

func (db *buntdbEngine) Put(key []byte, value []byte) error {
	if err := db.db.Update(
		func(tx *buntdb.Tx) error {
			_, _, err := tx.Set(string(key), string(value), nil)
			return err
		}); err != nil {
		return err
	}
	return nil
}

func (db *buntdbEngine) Get(key []byte) ([]byte, error) {
	var val string
	var err error
	if err := db.db.View(
		func(tx *buntdb.Tx) error {
			if val, err = tx.Get(string(key)); err != nil {
				return err
			}
			return nil
		}); err != nil {
		return nil, err
	}
	return []byte(val), nil
}

func (db *buntdbEngine) Close() error {
	return db.db.Close()
}

func (db *buntdbEngine) FileSize() (int64, error) {
	return dirSize(db.path)
}
