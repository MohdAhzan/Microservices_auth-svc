package services

import (
	"context"
	"net/http"

	dbconnection "github.com/MohdAhzan/auth-svc/pkg/dbConnection"
	"github.com/MohdAhzan/auth-svc/pkg/helper"
	"github.com/MohdAhzan/auth-svc/pkg/models"
	"github.com/MohdAhzan/auth-svc/pkg/pb"
)


type AuthService struct{
  
  Repo dbconnection.Repository
  Jwt helper.JwtWrapper    
    pb.UnimplementedAuthServiceServer
}
	// Register(context.Context, *RegisterRequest) (*RegisterResponse, error)

func (s *AuthService)Register(ctx context.Context , req *pb.RegisterRequest)(*pb.RegisterResponse, error){
       
 var user models.Users

if result := s.Repo.DB.Where(&models.Users{Email: req.Email}).First(&user); result.Error == nil {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "User with this email already exists",
		}, nil
	}

  user.Name=req.Name
  user.Email=req.Email
	
  hashedPass,err:= helper.HashPassword(req.Password)

	if err!=nil {
		return &pb.RegisterResponse{
			Status: http.StatusNotFound,
			Error:  "error hashing password",
		}, nil
	}
  user.Password=hashedPass

  err=s.Repo.DB.Create(&user).Error
  if err!=nil{
		return &pb.RegisterResponse{
			Status: http.StatusInternalServerError,
			Error:  "Error creating User",
		}, nil
  }
    return &pb.RegisterResponse{
    Status: http.StatusCreated,
  },nil

}

func (s *AuthService)Login(ctx context.Context,req *pb.LoginRequest)(*pb.LoginResponse,error){

    var user models.Users

	if result := s.Repo.DB.Where(&models.Users{Email: req.Email}).First(&user); result.Error != nil {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	match := helper.CheckPasswordHash(req.Password, user.Password)

	if !match {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	jwtToken,err  := s.Jwt.GenerateToken(user)
  if err!=nil{
		return &pb.LoginResponse{
			Status: http.StatusInternalServerError,
			Error:  "error generating token",
		}, nil

  }
    
  return &pb.LoginResponse{
    Status: http.StatusOK,
    JwtToken: jwtToken,
  },nil
}


func (s *AuthService)JwtValidate(ctx context.Context,req *pb.JwtRequest)(*pb.JwtResponse,error){
 
   claims, err := s.Jwt.ValidateToken(req.JwtToken)

	if err != nil {
		return &pb.JwtResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	var user models.Users

	if result := s.Repo.DB.Where(&models.Users{Email: claims.Email}).First(&user); result.Error != nil {
		return &pb.JwtResponse{

			Status: http.StatusNotFound,
			Error:  "user not found",
		}, nil
	}

	return &pb.JwtResponse{
		Status: http.StatusOK,
		UserId: user.Id,
      Role: claims.Role,
    
	}, nil 

}

func (s *AuthService) AdminLogin(ctx context.Context, req *pb.AdminLoginRequest) (*pb.AdminLoginResponse, error) {
	var admin models.Admin

	if result := s.Repo.DB.Where(&models.Admin{Email: req.Email}).First(&admin); result.Error != nil {
		return &pb.AdminLoginResponse{
			Status: http.StatusNotFound,
			Error:  "admin not found 1",
		}, nil
	}

	match := helper.CheckPasswordHash(req.Password, admin.Password)

	if !match {
		return &pb.AdminLoginResponse{
			Status: http.StatusInternalServerError,
			Error:  "Error admin password",
		}, nil
	}

	adminToken, err:= s.Jwt.GenerateAdminToken(admin)
  if err!=nil{
		return &pb.AdminLoginResponse{
			Status: http.StatusInternalServerError,
			Error:  "error generating adminToken",
		}, nil
	}

	return &pb.AdminLoginResponse{
		Status: http.StatusOK,
		JwtToken: adminToken,
	}, nil
}

