package server

import (
    "encoding/json"
    "log"
    "net/http"
    "io"
    "bytes"

    "backend/internal/models"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
    r := chi.NewRouter()
    r.Use(middleware.Logger)

    r.Post("/api/results", s.ResultsHandler)
    r.Get("/api/results", s.GetResultsHandler)
    r.Get("/health", s.healthHandler)

    return r
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
    resp := make(map[string]string)
    resp["message"] = "Hello World"
    
    log.Printf("request received: %v", resp)
    jsonResp, err := json.Marshal(resp)
    if err != nil {
        log.Fatalf("error handling JSON marshal. Err: %v", err)
    }

    _, _ = w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
    jsonResp, _ := json.Marshal(s.db.Health())
    _, _ = w.Write(jsonResp)
}

func (s *Server) ResultsHandler(w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close()

    body, _ := io.ReadAll(r.Body)
    r.Body = io.NopCloser(bytes.NewBuffer(body))

    var req models.ResultRequest

    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        log.Fatalf("error decoding request body. Err: %v", err)
        http.Error(w, "bad request", http.StatusBadRequest)
        return
    }
    err = s.db.SaveResult(req)
    if err != nil {
        log.Println("in api")
        log.Println("error saving result to db. Err: %v", err)
        http.Error(w, "internal server error", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    return
}

func (s *Server) GetResultsHandler(w http.ResponseWriter, r *http.Request) {
    var resp *models.ResultResponse
    resp, err := s.db.GetResults()
    if err != nil {
        log.Println("error retrieving results from db. Err: %v", err)
        http.Error(w, "internal server error", http.StatusInternalServerError)
        return
    }

    jsonData, err := json.Marshal(resp)
    if err != nil {
        log.Printf("error marshalling response. Err: %v", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    _, err = w.Write(jsonData)
    if err!= nil {
        log.Printf("error writing response. Err: %v", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    return 
}
