package config

type Config struct {
	Common          *CommonConfig
	OldDobiDatabase *MysqlConfig
	NewDobiDatabase *NewDobiDatabaseConfig
	RedisDatabase	*RedisConfig

}

type CommonConfig struct {
	InitDatabase int

}

type MysqlConfig struct {
	DriverName string
	DataSourceName string
}

type RedisConfig struct {
	Addr	string
	Password	string
	DB	int
}

type NewDobiDatabaseConfig struct {
	Core	MysqlConfig
	Trade	MysqlConfig
}




