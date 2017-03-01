package controllers

import (
	"encoding/binary"
	"encoding/json"

	"github.com/boltdb/bolt"
)

type LocalDataStore struct {
	DB *bolt.DB
}

func InitDB(name string) (*LocalDataStore, error) {
	ndb, err := bolt.Open(name, 0600, nil)
	ds := LocalDataStore{}
	if err != nil {
		return &ds, err
	}

	ds.DB = ndb

	return &ds, nil
}

func (ds *LocalDataStore) FindAndUpdate(bname string, key interface{}, value interface{}) error {
	ds.DB.Begin(true)
	err := ds.DB.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bname))

		id, _ := b.NextSequence()
		key.ID = int(id)

		buf, err := json.Marshal(key)

		if err != nil {
			return err
		}

		return b.Put(itob(key.ID), buf)
	})
	return nil
}

func (ds *LocalDataStore) FindAndReturn(param string) (interface{}, error) {

}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
