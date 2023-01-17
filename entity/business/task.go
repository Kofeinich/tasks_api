package business

import "time"

type Task struct {
	Id   int       `json:"-"`
	Text string    `json:"-"`
	Tags []string  `json:"-"`
	Due  time.Time `json:"-"`
}
