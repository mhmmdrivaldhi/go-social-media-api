package service

import (
	"github.com/mhmmdrivaldhi/go-social-media-api/model/dto"
	"github.com/mhmmdrivaldhi/go-social-media-api/model/entity"
	"github.com/mhmmdrivaldhi/go-social-media-api/repository"
	"github.com/mhmmdrivaldhi/go-social-media-api/utils"
)

type AuthService interface{
	Register(req *dto.RegisterRequest) error
	Login(req *dto.LoginRequest) (*dto.LoginResponse, error)
}

type authService struct {
	authRepo repository.AuthRepository
	jwtToken utils.JwtToken
}

func NewAuthService(authRepo repository.AuthRepository) AuthService {
	return &authService{authRepo: authRepo}
}

func (as *authService) Register(req *dto.RegisterRequest) error {
	emailExists := as.authRepo.IsEmailExists(req.Email)
	if emailExists {
		return &utils.BadRequestError{Message: "Email already exists"}
	}
	
	if req.Password != req.PasswordConfirmation {
		return &utils.BadRequestError{Message: "Password and password confirmation do not match"}
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return &utils.InternalServerError{Message: err.Error()}
	}

	user := entity.User{
		Name:     req.Name,
		Email:   req.Email,
		Password: hashedPassword,
		Gender: req.Gender,
	}

	err = as.authRepo.Register(&user)
	if err != nil {
		return &utils.InternalServerError{Message: err.Error()}
	}

	return nil
}

func (as *authService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := as.authRepo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, &utils.NotFoundError{Message: "Email address or password is wrong"}
	}

	err = utils.VerificationPassword(user.Password, req.Password)
	if err != nil {
		return nil, &utils.UnauthorizedError{Message: "Email address or Password is Wrong"}
	}

	token, err := as.jwtToken.GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, &utils.InternalServerError{Message: "Failed to generate token"}
	}

	return &dto.LoginResponse{
		Token: token,
		User: dto.UserResponse{
			ID: user.ID,
			Name: user.Name,
			Email: user.Email,
			Gender: user.Gender,
		},
	}, nil
}
