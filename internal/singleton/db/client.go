package db

import "fmt"

type client struct{}

func (d *client) query(q string) QueryResult {
	return QueryResult(fmt.Sprint("Results for: ", q))
}
