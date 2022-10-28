package db

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DBHost string
	DBName string
	DBUser string
	DBPass string
}

func LoadConfig(path string) (dbConfig *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("database")
	viper.SetConfigType("yml")

	err = viper.ReadInConfig()

	if err != nil {
		return nil, err
	}

	host := viper.GetString("development.host")
	port := viper.GetString("development.port")
	return &Config{
		DBHost: fmt.Sprintf("tcp(%s:%s)", host, port),
		DBName: viper.GetString("development.database"),
		DBUser: viper.GetString("development.username"),
		DBPass: viper.GetString("development.password"),
	}, nil
}

func (dbConfig *Config) getConnectionString() string {
	return fmt.Sprintf(
		"%s:%s@%s/%s?charset=utf8&parseTime=true",
		dbConfig.DBUser,
		dbConfig.DBPass,
		dbConfig.DBHost,
		dbConfig.DBName,
	)
}

func InitDB(config *Config) *gorm.DB {
	url := config.getConnectionString()
	DB, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return DB
}
