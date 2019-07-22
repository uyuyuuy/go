package config

type Config struct {
	Common          *Common
	OldDobiDatabase *Database
	NewDobiDatabase *Database

}

type Common struct {
	InitDatabase int

}

type Database struct {
	DriverName string
	DataSourceName string
}




