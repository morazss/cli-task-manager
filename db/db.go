package db

import (
	"time"

	"github.com/boltdb/bolt"
)

func db() {
	db, err := bolt.Open("tasks.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
