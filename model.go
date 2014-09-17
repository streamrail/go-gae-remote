package main

import (
	"time"
)

type MyEntityKind struct {
	Category string
	Action   string
	Label    string
	Date     time.Time
}
