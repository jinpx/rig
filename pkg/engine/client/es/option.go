package es

import "log"

type Config struct {
	Name      string
	Host      string // A list of Elasticsearch nodes to use.
	Addresses []string
	Username  string // Username for HTTP Basic Authentication.
	Password  string // Password for HTTP Basic Authentication.
	logger    *log.Logger
}
