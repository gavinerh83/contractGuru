package storage

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	//Client represents the database connection
	Client     *sql.DB
	userID     = "mysql_users_userid"
	password   = "mysql_users_password"
	department = "mysql_users_host"
	firstname  = "mysql_users_firstname"
	lastname   = "mysql_users_lastname"
	email      = "mysql_users_email"
)

const (
	dbUser     = "root"
	dbPassword = "password"
	hostname   = "127.0.0.1:3306"
	dbName     = "contractManagement"
)
