package main

import (
	authHandler "finback/internal/auth/handler"
	authRepo "finback/internal/auth/repository"
	authUsecase "finback/internal/auth/usecase"
	contentHandler "finback/internal/content/handler"
	contentRepo "finback/internal/content/repository"
	contentUsecase "finback/internal/content/usecase"
	"finback/internal/platform/config"
	"finback/internal/platform/logger"
	"finback/internal/platform/store"
	"net/http"
)

func main() {
	cfg := config.Load()
	appStore := store.New()

	userRepo := authRepo.NewUserRepository(appStore)
	contentRepository := contentRepo.New(appStore)

	authService := authUsecase.NewService(userRepo)
	contentService := contentUsecase.NewService(contentRepository, userRepo)

	authHTTP := authHandler.NewHTTPHandler(authService)
	contentHTTP := contentHandler.NewHTTPHandler(contentService)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})

	mux.HandleFunc("POST /auth/register", authHTTP.Register)
	mux.HandleFunc("POST /auth/login", authHTTP.Login)
	mux.HandleFunc("GET /auth/me", authHTTP.Me)

	mux.HandleFunc("GET /topics", contentHTTP.ListTopics)
	mux.HandleFunc("GET /topics/", contentHTTP.GetTopic)
	mux.HandleFunc("GET /lessons/", contentHTTP.GetLesson)

	mux.HandleFunc("POST /assessments/submit", contentHTTP.SubmitAssessment)
	mux.HandleFunc("GET /recommendations", contentHTTP.GetRecommendations)
	mux.HandleFunc("POST /progress/lesson/complete", contentHTTP.CompleteLesson)
	mux.HandleFunc("GET /progress", contentHTTP.GetProgress)

	logger.Info("mock finance backend started on http://localhost:%s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, withCORS(mux)); err != nil {
		panic(err)
	}
}

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
