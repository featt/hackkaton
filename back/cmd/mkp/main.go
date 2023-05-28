package main

import (
	"backend/configs"
	"backend/pkg/api/handler"
	"backend/pkg/api/middleware"
	"backend/pkg/repository"
	"backend/pkg/service"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net/http"
)

func main() {
	// Используем конфигурацию из файла config
	config := configs.GetConfig()

	// Инициализируем подключение к БД
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		config.DB.User, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.DBName, config.DB.SSLMode)

	dbpool, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalf("Unable to connection to database: %v\n", err)
	}
	defer dbpool.Close()

	// Получаем секретный ключ JWT и продолжительность токена из конфигурации
	jwtSecretKey := []byte(config.JWTSecretKey)
	tokenDuration := config.TokenDuration

	// Инициализируем репозитории
	technicalRepo := repository.NewTechnicalRepository(dbpool)
	userRepo := repository.NewUserRepository(dbpool)
	internshipRepo := repository.NewInternshipRepository(dbpool)
	candidateRepo := repository.NewCandidateRepository(dbpool)

	// Инициализируем сервисы
	technicalService := service.NewTechnicalService(technicalRepo)
	sessionService := service.NewSessionService(string(jwtSecretKey)) // Преобразование []byte в строку
	// Создаем экземпляр UserService с поддержкой JWT токенов
	userService := service.NewUserService(userRepo, string(jwtSecretKey), tokenDuration) // Преобразование []byte в строку
	internshipService := service.NewInternshipService(internshipRepo)
	loyaltyService := service.NewLoyaltyService(userRepo) // Инициализируем сервис для системы лояльности
	candidateService := service.NewCandidateService(candidateRepo)

	// Инициализируем хендлеры
	technicalHandler := handler.NewTechnicalHandler(technicalService)
	userHandler := handler.NewUserHandler(userService, sessionService)

	internshipHandler := handler.NewInternshipHandler(internshipService)
	loyaltyHandler := handler.NewLoyaltyHandler(loyaltyService) // Инициализируем хендлер для системы лояльности

	candidateHandler := handler.NewCandidateHandler(candidateService)

	// Инициализируем роутинг
	r := mux.NewRouter()

	// Добавляем маршруты для проверки подключения к базе данных и вывода всех таблиц
	r.HandleFunc("/check-connection", technicalHandler.CheckDatabaseConnection).Methods("GET")
	r.HandleFunc("/tables", technicalHandler.GetAllTables).Methods("GET")

	// Добавляем маршруты для пользователя
	r.HandleFunc("/api/users", userHandler.RegisterUser).Methods("POST")
	r.HandleFunc("/api/users/login", userHandler.LoginUser).Methods("POST")
	r.HandleFunc("/api/users/all", userHandler.GetAllUsers).Methods("GET")
	r.HandleFunc("/api/users/me", userHandler.GetCurrentUser).Methods("GET")
	r.HandleFunc("/api/users/{id}", userHandler.UpdateUser).Methods("PUT")
	r.HandleFunc("/api/users/{id}", userHandler.GetUserByID).Methods("GET")
	r.HandleFunc("/api/users/{id}/role", userHandler.UpdateUserRole).Methods("PUT")
	r.HandleFunc("/api/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	// Добавляем маршруты для стажировок
	r.HandleFunc("/api/internships", internshipHandler.CreateInternship).Methods("POST")
	r.HandleFunc("/api/internships/all", internshipHandler.GetAllInternships).Methods("GET")
	r.HandleFunc("/api/internships/{id}", internshipHandler.GetInternship).Methods("GET")
	r.HandleFunc("/api/internships/{id}", internshipHandler.UpdateInternship).Methods("PUT")
	r.HandleFunc("/api/internships/{id}", internshipHandler.DeleteInternship).Methods("DELETE")

	// Регистрация эндпоинтов для системы лояльности
	r.HandleFunc("/api/users/{id}/loyalty", loyaltyHandler.GetUserLoyalty).Methods("GET")
	r.HandleFunc("/api/users/{id}/loyalty", loyaltyHandler.UpdateUserLoyalty).Methods("PUT")

	// Маршрут для автоматической проверки анекты кандидата
	r.HandleFunc("/api/candidate", candidateHandler.ProcessCandidateForm).Methods("POST")

	// маршруты для импорта и экспорта табелей
	r.HandleFunc("/api/save-excel", internshipHandler.SaveExcelHandler).Methods("POST")
	r.HandleFunc("/api/download-excel", internshipHandler.DownloadExcelHandler).Methods("GET")

	corsMiddleware := middleware.Cors()

	http.ListenAndServe(":8080", corsMiddleware(r))
}
