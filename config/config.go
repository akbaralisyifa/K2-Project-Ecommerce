package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type setting struct {
	User 		string
	Host 		string
	Password 	string
	Port 		string
	DBName 		string
	JWTSecrat	string
	CldKey		string
}

	func ImportSetting() setting {
		var result setting

		err := godotenv.Load(".env");

		if err != nil {
			return	 setting{}
		};

		result.User 	= os.Getenv("DB_USER");
		result.Host		= os.Getenv("DB_HOST");
		result.Port		= os.Getenv("DB_PORT");
		result.DBName 	= os.Getenv("DB_NAME");
		result.Password = os.Getenv("DB_PASSWORD");
		result.JWTSecrat= os.Getenv("JWT_SECRATE");
		result.CldKey   = os.Getenv("CLOUDINARY_KEY");

		return result;
	}

	func ConnectDB(s *setting)(*gorm.DB, error){
		var connStr = fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s", s.Host, s.User, s.Password, s.Port, s.DBName)
		db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

		if err != nil {
			log.Fatal("Error config database ", err.Error())
			return nil, err
		}

		return db, nil;
	}
