package mysql

import (
	"fmt"
	"log"
	"starter-go/internal/pkg/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	conf := config.Database()
	// user:password@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.GetUser(),
		conf.GetPassword(),
		conf.GetHost(),
		conf.GetPort(),
		conf.GetName(),
	)

	var db *gorm.DB
	var err error

	// retry database connection at startup
	for i := 0; i < 3; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("Failed to connect to database (attempt %d/3): %v. Retrying in 5 seconds...", i+1, err)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		log.Fatalf("Failed to connect to database after 3 attempts: %v", err)
	}

	return db
}
