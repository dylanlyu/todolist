package configs

type Configs struct {
	PostgreSQL struct {
		Host     string
		Port     int
		User     string
		DBName   string
		Password string
		SSLMode  string
	}
}

func GetConfig() *Configs {
	var config Configs

	config.PostgreSQL.Host = "localhost"
	config.PostgreSQL.Port = 5432
	config.PostgreSQL.User = "dylan"
	config.PostgreSQL.Password = "123456"
	config.PostgreSQL.DBName = "todo"
	config.PostgreSQL.SSLMode = "disable"

	return &config
}
