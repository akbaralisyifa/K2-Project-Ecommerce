package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type setting struct {
	User 		string
	Host 		string
	Password 	string
	Port 		string
	DBName 		string
	JWTSecrat	string
}

	func InportSetting() setting {
		var result setting

		err := godotenv.Load(".env");

		if err != nil {
			return	 setting{}
		};

		result.User 	= os.Getenv("DB_USER");
		result.Host		= os.Getenv("DB_HOST");
		result.Port		= os.Getenv("DB_PORT");
		result.DBName 	= os.Getenv("DB_NAME");
		result.Password = os.Getenv("DB_PASSWORD")

		return result;
	}

	func ConnectDB(s *setting)(*gorm.DB, error){
		var connStr = fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s", s.Host, s.User, s.Password, s.Port, s.DBName)
		db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
				TablePrefix: "public.",
			},
		})

		if err != nil {
			log.Fatal("Error config database ", err.Error())
			return nil, err
		}

		return db, nil;
	}
