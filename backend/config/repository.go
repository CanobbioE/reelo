package config

import "time"

var (
	dbDriver            = "mysql"
	user                = "reeloUser"
	password            = "password"
	host                = "localhost:3306"
	dbName              = "reelo"
	bkpDir              = "bkp"
	maxConnections      = 5
	maxIdleConnections  = 5
	maxConnTries        = 10
	connectionsLifetime = (time.Minute * 5)
	instanceEsists      = false
)
