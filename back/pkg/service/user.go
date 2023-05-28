// service/user.go
package service

import (
	"backend/pkg/model"
	"backend/pkg/repository"
	"context"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type UserService struct {
	repo          *repository.UserRepository
	jwtSecretKey  []byte
	tokenDuration time.Duration
}

func NewUserService(repo *repository.UserRepository, jwtSecretKey string, tokenDuration time.Duration) *UserService {
	return &UserService{
		repo:          repo,
		jwtSecretKey:  []byte(jwtSecretKey),
		tokenDuration: tokenDuration,
	}
}

func (s *UserService) CreateUser(ctx context.Context, user model.User) (int64, error) {
	return s.repo.CreateUser(ctx, user)
}

func (s *UserService) GetUserByEmailAndPassword(ctx context.Context, email string, password string) (model.User, error) {
	return s.repo.GetUserByEmailAndPassword(ctx, email, password)
}

func (s *UserService) GetUserByID(ctx context.Context, id int64) (model.User, error) {
	return s.repo.GetUserByID(ctx, id)
}

func (s *UserService) UpdateUser(ctx context.Context, id int64, user model.User) error {
	return s.repo.UpdateUser(ctx, id, user)
}

func (s *UserService) DeleteUser(ctx context.Context, id int64) error {
	return s.repo.DeleteUser(ctx, id)
}

func (s *UserService) UpdateUserRole(ctx context.Context, id int64, role string) error {
	return s.repo.UpdateUserRole(ctx, id, role)
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]model.User, error) {
	return s.repo.GetAllUsers(ctx)
}


func (s *UserService) RegisterUser(ctx context.Context, regReq model.RegistrationRequest) (int64, error) {
	user := model.User{
		Email:    regReq.Email,
		Password: regReq.Password, // TODO: hash and salt the password
		Name:     regReq.Name,
		LastName: regReq.LastName,
		Role:     model.Candidate, // assign a default role
	}
	return s.repo.CreateUser(ctx, user)
}


func (s *UserService) LoginUser(ctx context.Context, loginReq model.LoginRequest) (string, error) {
	user, err := s.repo.GetUserByEmailAndPassword(ctx, loginReq.Email, loginReq.Password)
	if err != nil {
		return "", err
	}

	// Генерируем новый JWT токен
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(s.tokenDuration).Unix()

	// Подписываем токен с использованием секретного ключа
	tokenString, err := token.SignedString(s.jwtSecretKey)
	if err != nil {
		return "", err
	}

	return "Bearer " + tokenString, nil
}