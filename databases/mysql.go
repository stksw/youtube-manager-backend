package databases

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() (db *gorm.DB, err error) {
	err = godotenv.Load()
	if err != nil {
		logrus.Fatal("error loading .env file")
	}

	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	options := "?charset=utf8mb4&parseTime=True&Loc=Local"
	mysql := mysql.Open(user + ":" + password + "@tcp(db:3306)/" + dbName + options)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	db, err = gorm.Open(mysql, &gorm.Config{Logger: newLogger})
	if err != nil {
		logrus.Fatal(err)
	}

	return db, err
}
