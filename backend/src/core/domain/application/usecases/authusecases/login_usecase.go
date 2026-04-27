package authusecases

import (
	"blogapp/src/core/domain/entity"
	"blogapp/src/core/domain/ports/repository"
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"        
	"golang.org/x/crypto/bcrypt"
)

type LoginUseCase struct {
	userRepo  repository.UserRepository
	jwtSecret []byte  
}


func NewLoginUseCase(userRepo repository.UserRepository, jwtSecret []byte) *LoginUseCase {
	return &LoginUseCase{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

func (uc *LoginUseCase) Invoke(ctx context.Context, email, password string) (string, *entity.UserEntity, error) {
	user, err := uc.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return "", nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", nil, errors.New("invalid password")
	}


	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(), 
	})

	tokenString, err := token.SignedString(uc.jwtSecret)
	if err != nil {
		return "", nil, err
	}

	return tokenString, user, nil
}