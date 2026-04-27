package authusecases

import (
	"blogapp/src/core/domain/entity"
	"blogapp/src/core/domain/ports/repository"
	"context"

	"golang.org/x/crypto/bcrypt"
)



type RegisterUseCase struct {
	userRepo repository.UserRepository
	profileRepo repository.ProfileRepository
}


func NewRegisterUseCase(userRepo repository.UserRepository, profileRepo repository.ProfileRepository) *RegisterUseCase{
	return &RegisterUseCase{
		userRepo: userRepo,
		profileRepo: profileRepo,
	}
}


func (uc *RegisterUseCase) Invoke(ctx context.Context ,username, email,password string) (*entity.UserEntity,error){
	hash,err :=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil {
		return nil ,err
	}
	user := entity.NewUserEntity(username,email,string(hash))

	user , err = uc.userRepo.Create(ctx,user)
	if err != nil {
		return nil,err
	}
	profile := entity.NewProfileEntity(user.ID,"","","")
	profile,err = uc.profileRepo.Create(ctx,profile)
	if err != nil {
		return nil,err
	}

	return user,nil




}