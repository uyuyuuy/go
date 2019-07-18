package config

type Config struct {
	Common *Common
	Database *Database

}

type Common struct {
	InitDatabase int

}

type Database struct {
	DriverName string
	DataSourceName string
}




