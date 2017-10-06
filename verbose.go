package main

import "log"

// no comment -_-
func verbose(args ...interface{}) {
	if *VERBOSE {
		log.Println(args...)
	}
}
