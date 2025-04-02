package main

import (
	"fmt"
	"log"
	"net/http"

	exampleRepo "starter-go/internal/repository/example"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 1. Connect to MySQL
	dsn := "user:password@tcp(127.0.0.1:3306)/yourdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}

	// 2. Auto-migrate DB schema (for dev only)
	_ = db.AutoMigrate(&exampleRepo.ExampleModel{}) // comment out in prod

	// 3. Setup dependencies
	repo := userRepo.NewGormRepository(db)
	usecase := userUsecase.NewService(repo)
	handler := userHandler.NewHandler(usecase)

	// 4. Setup routes
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetUser(w, r)
		case http.MethodPost:
			handler.CreateUser(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// 5. Start server
	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
