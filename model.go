package main

import (
	"time"
)

type SampleKind struct {
	Category string
	Action   string
	Label    string
	Country  string
	Date     time.Time
}
