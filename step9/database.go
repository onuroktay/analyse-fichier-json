package main

type DATABASE struct {
	// List of Methods to be implemented in the db struct (couchbase, elasticsearch, ...)
	accesser interface {
		SaveItem(*ITEM) error
	}
}
