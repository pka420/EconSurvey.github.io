package database

import (
	"backend/internal/models"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

type Service interface {
	Health() map[string]string
    SaveResult(req models.ResultRequest) (error)
    GetResults() (*models.ResultResponse, error)
}

type service struct {
	db *sql.DB
}

var (
	dburl = os.Getenv("DB_URL")
)

func New() Service {
    path, _ := os.Getwd()
    db, err := sql.Open("sqlite3", path+"/"+dburl)
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}
	s := &service{db: db}
	return s
}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.db.PingContext(ctx)
	if err != nil {
		log.Println(fmt.Sprintf("db down: %v", err))
	}

	return map[string]string{
		"message": "It's healthy",
	}
}

func (s* service) SaveResult(req models.ResultRequest) (error){
    log.Println(req.Economic, req.Diplomatic, req.Civil, req.Societal)
    ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel()

    _, err := s.db.ExecContext(ctx, "INSERT into results (Economic, Diplomatic, Civil, Societal) VALUES($1, $2, $3, $4)", req.Economic, req.Diplomatic, req.Civil, req.Societal)
    if err != nil {
        log.Println("in db")
        log.Println("error saving result to db. Err: %v", err)
        return err
    }

    return nil
}

func(s* service) GetResults() (*models.ResultResponse, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel()

    rows, err := s.db.QueryContext(ctx, "SELECT * FROM results")
    if err != nil {
        return nil, err
    }

    var query models.ResultResponse

    for rows.Next() {
        var result models.ResultRequest
        err := rows.Scan(&result.Economic, &result.Diplomatic, &result.Civil, &result.Societal)
        if err != nil {
            return nil, err
        }
        query.Results = append(query.Results, result)
    }

    return &query, nil
}
