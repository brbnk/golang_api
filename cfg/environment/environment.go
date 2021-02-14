package environment

import (
	"flag"
	"fmt"
	"os"
)

type Configurations struct {
	APIHost    string
	APIPort    string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	DBHost     string
	migrate    string
}

func Get() *Configurations {
	config := &Configurations{}

	flag.StringVar(&config.APIHost, "apihost", os.Getenv("API_HOST"), "API Host Domain")
	flag.StringVar(&config.APIPort, "apiport", os.Getenv("API_PORT"), "API Port")
	flag.StringVar(&config.DBUser, "user", os.Getenv("MYSQL_USER"), "User DB")
	flag.StringVar(&config.DBPassword, "password", os.Getenv("MYSQL_PASSWORD"), "Password DB")
	flag.StringVar(&config.DBName, "database", os.Getenv("MYSQL_DATABASE"), "Database name")
	flag.StringVar(&config.DBPort, "port", os.Getenv("MYSQL_PORT"), "Database port")
	flag.StringVar(&config.DBHost, "host", os.Getenv("MYSQL_HOST"), "Database host")
	flag.StringVar(&config.migrate, "migrate", "up", "Specify if we should be migrating DB 'up' or 'down'")

	flag.Parse()

	return config
}

func (c *Configurations) GetDBConnStr() string {
	return c.getDBConnStr(c.DBHost, c.DBName)
}

func (c *Configurations) getDBConnStr(dbhost, dbname string) string {
	return fmt.Sprintf(
		"%s:%s@tcp(mysqlsrv)/%s",
		c.DBUser,
		c.DBPassword,
		c.DBName,
	)
}

func (c *Configurations) GetMigration() string {
	return c.migrate
}
