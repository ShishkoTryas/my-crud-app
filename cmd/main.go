package main

import (
	log "github.com/sirupsen/logrus"
	"my-crud-app/internal/config"
	db3 "my-crud-app/internal/repository/db"
	"my-crud-app/internal/service"
	"my-crud-app/internal/transport/rest"
	db2 "my-crud-app/pkg/db"
	"net/http"
	"os"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
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

// @securityDefinitions.basic  BasicAuth
func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

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

	booksRepo := db3.NewBooks(db)
	booksService := service.NewBooks(booksRepo)
	handler := rest.NewHandler(booksService)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler.CreateRouter(),
	}

	log.Info("SERVER STARTED AT")

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
