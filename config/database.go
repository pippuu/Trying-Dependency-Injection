package config

type DatabaseConfig struct {
	DATABASE_HOST     string
	DATABASE_PORT     string
	DATABASE_NAME     string
	DATABASE_USERNAME string
	DATABASE_PASSWORD string
}

func MakeDBConfig() DatabaseConfig {
	Host := GetENV("DATABASE_HOST", "127.0.0.1")
	Port := GetENV("DATABASE_PORT", "127.0.0.1")
	Name := GetENV("DATABASE_NAME", "DEFAULT")
	Username := GetENV("DATABASE_USERNAME", "DEFAULT")
	Password := GetENV("DATABASE_PASSWORD", "DEFAULT")
	return DatabaseConfig{
		DATABASE_HOST:     Host,
		DATABASE_PORT:     Port,
		DATABASE_NAME:     Name,
		DATABASE_USERNAME: Username,
		DATABASE_PASSWORD: Password,
	}
}
