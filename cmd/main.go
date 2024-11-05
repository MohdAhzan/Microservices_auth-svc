package main

import (
	"log"
	"net"

	"github.com/MohdAhzan/auth-svc/pkg/config"
	dbconnection "github.com/MohdAhzan/auth-svc/pkg/dbConnection"
	"github.com/MohdAhzan/auth-svc/pkg/helper"
	"github.com/MohdAhzan/auth-svc/pkg/pb"
	"github.com/MohdAhzan/auth-svc/pkg/services"
	"google.golang.org/grpc"
)


func main(){

 cfg, err:=config.LoadConfig()
  if err!=nil{
    log.Fatal("error loading authCOnfig",err)
  }
  dbrepo,err:= dbconnection.DBconnect(cfg)
  
  if err!=nil{
    log.Fatal("error loading authCOnfig",err)
  }

  listner,err:=net.Listen("tcp",cfg.Port)
  if err!=nil{
    log.Fatal("error listening to port",cfg.DBPort,err)
  }

  
  auth_svc:= &services.AuthService{
   Repo: dbrepo, 
   Jwt: helper.JwtWrapper{
      SecretKey: cfg.JWTSecretKey,
      Issuer: "auth-svc",
      ExpiryHours: 24*30,
    },
  }

  grpcServer:=grpc.NewServer()
  pb.RegisterAuthServiceServer(grpcServer,auth_svc)
  
  
  log.Println("auth_service running on port ",cfg.Port)
  err=grpcServer.Serve(listner)
  if err!=nil{
    log.Fatal(err)

  }


}
