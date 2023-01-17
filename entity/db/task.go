package db

import "time"

type Task struct {
	Id   int       `db:"id"`
	Text string    `db:"text"`
	Tags []string  `db:"tags"`
	Due  time.Time `db:"due"`
}
