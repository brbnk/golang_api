package config

import (
	"flag"
	"os"
)

type Configurations struct {
	Srv *serverConfiguration
	Db  *databaseConfiguration
}

type serverConfiguration struct {
	APIHost string
	APIPort string
}

type databaseConfiguration struct {
	User     string
	Password string
	Name     string
}

func Get() *Configurations {
	return &Configurations{
		Srv: getServerCfg(),
		Db:  getDatabaseCfg(),
	}
}

func getServerCfg() *serverConfiguration {
	config := &serverConfiguration{}

	flag.StringVar(&config.APIHost, "apihost", os.Getenv("API_HOST"), "API Host Domain")
	flag.StringVar(&config.APIPort, "apiport", os.Getenv("API_PORT"), "API Port")

	flag.Parse()

	return config
}

func getDatabaseCfg() *databaseConfiguration {
	config := &databaseConfiguration{}

	flag.StringVar(&config.User, "user", os.Getenv("MYSQL_USER"), "User DB")
	flag.StringVar(&config.Password, "password", os.Getenv("MYSQL_PASSWORD"), "Password DB")
	flag.StringVar(&config.Name, "database", os.Getenv("MYSQL_DATABASE"), "Database name")

	flag.Parse()

	return config
}
