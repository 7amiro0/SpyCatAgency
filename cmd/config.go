package main

import "os"

type config struct {
	host, port string
}

func newConfig() config {
	return config{
		host: os.Getenv("HOST"),
		port: os.Getenv("PORT"),
	}
}