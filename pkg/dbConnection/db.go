package dbconnection

import (
	"fmt"

	"github.com/MohdAhzan/auth-svc/models"
	"github.com/MohdAhzan/auth-svc/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

 


func DBconnect(cfg config.Config)(*gorm.DB,error){

  dsn:=fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",cfg.DBHost,cfg.DBUser,cfg.DBPassword,cfg.DBName,cfg.DBPort)
 
  DB,err:=gorm.Open(postgres.Open(dsn),&gorm.Config{})

    if err!=nil{
    return nil,err
  }
  err=DB.AutoMigrate(&models.Users{})
  if err!= nil{
    return nil,err
  }
  err= DB.AutoMigrate(&models.Admin{})
  if err!= nil{
    return nil,err
  }
  
  return DB,nil
}

