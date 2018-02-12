package model

import (
	"time"
)

var levelName = map[int]string{
	0:      "VERB",
	1 << 0: "VERB",
	1 << 1: "DEBUG",
	1 << 2: "INFO",
	1 << 3: "WARN",
	1 << 4: "ERR",
	1 << 5: "CRIT",
}

// Log the log mode
type Log struct {
	Level   int
	Message string
	File    string
	Line    int
	Time    time.Time
}

func (l *Log) String() string {
	return l.Time.Format("[2006-01-02 15:04:05]") + " [" + levelName[l.Level] + "] " + l.Message
}
