package main

import (
	"github.com/millken/godb"
)

type godbEngine struct {
	db   *godb.DB
	path string
}

func newGoDB(path string) (kvEngine, error) {
	db, err := godb.Open(path, godb.WithFsync(false))
	if err != nil {
		return nil, err
	}
	return &godbEngine{db: db, path: path}, err
}

func (db *godbEngine) Put(key []byte, value []byte) error {
	return db.db.Put(key, value)
}

func (db *godbEngine) Get(key []byte) ([]byte, error) {
	return db.db.Get(key)
}

func (db *godbEngine) Close() error {
	return db.db.Close()
}

func (db *godbEngine) FileSize() (int64, error) {
	return dirSize(db.path)
}
