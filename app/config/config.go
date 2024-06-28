package config

type Config struct {
	DB DB
}

type DB struct {
	DB_Driver string
	DB_Host   string
	DB_Port   string
	DB_User   string
	DB_Pass   string
	DB_Name   string
}
