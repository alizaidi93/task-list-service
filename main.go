package main

import (
	"database/sql"
	"io"
	"log"
	"net/http"
	"os"
	"task-list-service/ayzee/database/repository"
	"task-list-service/ayzee/handlers"
	"task-list-service/ayzee/service"
	"time"

	logrus "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	var log *logrus.Logger = initialiseLogger()
	var db = initialiseDatabase()

	//Wire in the dependenciess

	//Pass the db to the repo
	repository := repository.ProvideRepository(db, log)
	//Pass the repo to the service
	service := service.ProvideService(repository, log)
	//Pass the service to the handler
	handler := handlers.ProvideHandlers(service, log)
	//handler then has access when providing paths

	log.Info("Starting Server")
	startServer(handler)
}

func initialiseLogger() *logrus.Logger {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Unable to Open / Create logging file.")
	}
	multiWriter := io.MultiWriter(os.Stdout, file)

	var log = logrus.New()
	log.SetOutput(multiWriter)
	log.SetLevel(logrus.InfoLevel)
	// log.SetFormatter(&logrus.JSONFormatter{})
	return log
}

func initialiseDatabase() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:<password>@localhost/postgres?sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect", err)
	}
	return db
}

func startServer(handler *handlers.Handlers) {
	router := mux.NewRouter()
	router.HandleFunc("/task/{uuid}", handler.GetTask).Methods("GET")
	router.HandleFunc("/task/user/{user}", handler.GetTasksByUser).Methods("GET")
	router.HandleFunc("/task", handler.CreateTask).Methods("POST")
	router.HandleFunc("/task/update", handler.UpdateTask).Methods("POST")

	//TODO : Add Middleware handler for logging
	//TODO : Add Middleware handler for Authentication

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:9100",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	srv.ListenAndServe()
}
