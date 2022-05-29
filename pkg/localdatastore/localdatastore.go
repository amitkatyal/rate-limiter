package localdatastore

import "github.com/amit.katyal/rate-limiter/datastore"

type LocalDataStore struct {

}

func NewLocalDataStore() datastore.DataStore {
	return &LocalDataStore{}
}

func (store *LocalDataStore) Check() bool  {
	return true
}