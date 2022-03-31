package main

import (
	"github.com/millken/archivedb"
)

type archivedbEngine struct {
	db   *archivedb.DB
	path string
}

func newArchiveDB(path string) (kvEngine, error) {
	db, err := archivedb.Open(path)
	if err != nil {
		return nil, err
	}
	return &archivedbEngine{db: db, path: path}, err
}

func (db *archivedbEngine) Put(key []byte, value []byte) error {
	return db.db.Put(key, value)
}

func (db *archivedbEngine) Get(key []byte) ([]byte, error) {
	return db.db.Get(key)
}

func (db *archivedbEngine) Close() error {
	return db.db.Close()
}

func (db *archivedbEngine) FileSize() (int64, error) {
	return dirSize(db.path)
}
