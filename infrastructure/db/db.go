package db

import (
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/lhps/desafio-01/domain/model"
	_ "github.com/lib/pq"
	_ "gorm.io/driver/sqlite"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	err := godotenv.Load(basepath + "/../../.env")

	if err != nil {
		log.Fatalf("Error loading .env files")
	}
}

func ConnectDB() *gorm.DB {
	var (
		dsn string
		db  *gorm.DB
		err error
	)

	dsn = os.Getenv("dsn")
	db, err = gorm.Open(os.Getenv("dbType"), dsn)

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		panic(err)
	}

	if os.Getenv("debug") == "true" {
		db.LogMode(true)
	}

	if os.Getenv("AutoMigrateDb") == "true" {
		db.AutoMigrate(&model.Product{})
	}

	return db
}
