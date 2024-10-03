package main

import "time"

type Job struct {
	ID        int
	name      string
	desc      string
	createdAt time.Time
}
