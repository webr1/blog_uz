package profileusecases

import (
	"blogapp/src/core/domain/entity"
	"blogapp/src/core/domain/ports/repository"
	"context"
)


type ProfileUseCase struct{
	profileRepo repository.ProfileRepository

}

func NewProfileUseCase(profileRepo repository.ProfileRepository) *ProfileUseCase{
	return &ProfileUseCase{
		profileRepo: profileRepo,

	}
}

func (uc *ProfileUseCase) Invoke(ctx context.Context,userID uint,fullname,bio,avatar string) (*entity.ProfileEntity,error){
	profile, err := uc.profileRepo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	profile.FullName = fullname
	profile.Bio = bio
	profile.Avatar = avatar
	return uc.profileRepo.Update(ctx, profile)
}