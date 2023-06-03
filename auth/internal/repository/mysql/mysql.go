package mysql

import (
	"minilib/auth/internal/repository"
	"minilib/auth/pkg/model"
	"minilib/pkg/auth"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	db *gorm.DB
}

func New(db *gorm.DB) repository.AuthRepository {
	return &AuthRepositoryImpl{
		db: db,
	}
}

func (repo *AuthRepositoryImpl) Register(userInput model.UserInput) (model.User, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)

	if err != nil {
		return model.User{}, err
	}

	var user model.User = model.User{
		Email:    userInput.Email,
		Password: string(password),
	}

	result := repo.db.Create(&user)

	if err := result.Error; err != nil {
		return model.User{}, err
	}

	err = result.Last(&user).Error

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (repo *AuthRepositoryImpl) Login(userInput model.UserInput) (string, error) {
	var user model.User

	err := repo.db.First(&user, "email = ?", userInput.Email).Error

	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password))

	if err != nil {
		return "", err
	}

	jwtConfig := auth.NewDefaultConfig()

	token, err := jwtConfig.GenerateToken(int(user.ID))

	if err != nil {
		return "", err
	}

	return token, nil
}
