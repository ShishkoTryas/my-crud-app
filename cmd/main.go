package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"my-crud-app/internal/config"
	db3 "my-crud-app/internal/repository/db"
	"my-crud-app/internal/service"
	"my-crud-app/internal/transport/rest"
	db2 "my-crud-app/pkg/db"
	"my-crud-app/pkg/hash"
	"net/http"
	"os"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

// @title           Books Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.email  ya.maksimka228@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey UserAuth
// @in header
// @name Authorization
func main() {
	cfg := config.New()

	log.Printf("config: %+v\n", cfg)

	db, err := db2.NewPostgresDB(db2.ConnectionInfo{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		Username: cfg.DB.Username,
		DBName:   cfg.DB.DBName,
		SSLMode:  cfg.DB.SSLMode,
		Password: cfg.DB.Password,
	})
	if err != nil {
		log.Fatal("Cannot create db", err)
	}
	defer db.Close()

	hasher := hash.New(cfg.Phrase.Salt)

	booksRepo := db3.NewBooks(db)
	userRepo := db3.New(db)
	booksService := service.NewBooks(booksRepo)
	usersService := service.New(userRepo, hasher, []byte(cfg.Phrase.Secret))

	handler := rest.NewHandler(booksService, usersService)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler.CreateRouter(),
	}

	log.Info("SERVER STARTED AT")

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
