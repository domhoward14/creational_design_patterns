package db

import (
	"sync"
)

var (
	clientInstance *client
	mutex          sync.Mutex
	once           sync.Once
)

type QueryResult string

type Singleton struct {
	dbc *client
}

func (d *Singleton) Query(q string) QueryResult {
	return d.dbc.query(q)
}

func GetSingleton() Singleton {
	once.Do(func() {
		clientInstance = &client{}
	})
	return Singleton{dbc: clientInstance}
}
