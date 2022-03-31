package main

import (
	"git.mills.io/prologic/bitcask"
)

type bitcaskEngine struct {
	db   *bitcask.Bitcask
	path string
}

func newBitcask(path string) (kvEngine, error) {
	db, err := bitcask.Open(path)
	if err != nil {
		return nil, err
	}
	return &bitcaskEngine{db: db, path: path}, err
}

func (db *bitcaskEngine) Put(key []byte, value []byte) error {
	return db.db.Put(key, value)
}

func (db *bitcaskEngine) Get(key []byte) ([]byte, error) {
	val, err := db.db.Get(key)
	return val, err
}

func (db *bitcaskEngine) Close() error {
	return db.db.Close()
}

func (db *bitcaskEngine) FileSize() (int64, error) {
	return dirSize(db.path)
}
