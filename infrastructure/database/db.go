package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"internet_service/config"
	"internet_service/internal/domain"
	"log"
	"sync"
)

func dsn() string {
	conf := config.GetConfToml()
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		conf.Databases.Postgres.Host,
		conf.Databases.Postgres.User,
		conf.Databases.Postgres.Password,
		conf.Databases.Postgres.Name,
		conf.Databases.Postgres.Port,
		conf.Databases.Postgres.TimeZone,
	)
}

func dsnTest() string {
	conf := config.GetConfToml()
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		conf.Databases.Test.Host,
		conf.Databases.Test.User,
		conf.Databases.Test.Password,
		conf.Databases.Test.Name,
		conf.Databases.Test.Port,
		conf.Databases.Test.TimeZone,
	)
}

var once = &sync.Once{}

var handle *gorm.DB

func migrate() {
	err := handle.AutoMigrate(&domain.Provider{}, &domain.NetPackage{})
	if err != nil {
		log.Fatal(err)
	}
}

func initialize() {
	var err error
	handle, err = gorm.Open(postgres.Open(dsn()), &gorm.Config{})
	if err != nil {
		// Todo do something
		log.Fatal(err)
	}
	migrate()
}
func initializeTest() {
	var err error
	handle, err = gorm.Open(postgres.Open(dsnTest()), &gorm.Config{})
	if err != nil {
		// Todo do something
		log.Fatal(err)
	}
	migrate()
}

func Connect() *gorm.DB {
	once.Do(initialize)
	return handle
}

func ConnectToTest() *gorm.DB {
	once.Do(initializeTest)
	return handle
}
