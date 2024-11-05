package dbconnection

import (
	"database/sql"
	"fmt"

	"github.com/MohdAhzan/auth-svc/pkg/config"
	"github.com/MohdAhzan/auth-svc/pkg/models"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct{
  
  DB *gorm.DB

} 


func DBconnect(cfg config.Config)(Repository,error){
 
  connString:=fmt.Sprintf("host=%s user=%s password=%s ",cfg.DBHost,cfg.DBUser,cfg.DBPassword)

  db,err:=sql.Open("postgres",connString) 
  if err!=nil{

    fmt.Println("Errdfklsdjfkljdslkor checking if database exists")
    return Repository{},err
  }
  rows,err:=db.Query("SELECT 1 FROM pg_database WHERE datname = $1",cfg.DBName)
  if err!=nil{
    fmt.Println("Error checking if database exists")
    return Repository{},err
  } 

  if rows.Next() {
        rows.Close()
        // fmt.Println(cfg.DBName+" already exists...")
  }else{
    _,err:=db.Exec("CREATE DATABASE "+cfg.DBName)
    if err!=nil{
    fmt.Println("Error creating"+cfg.DBName)
      return Repository{},err
    }
        fmt.Println(cfg.DBName+" created")

  }

    
  dsn:=fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",cfg.DBHost,cfg.DBUser,cfg.DBPassword,cfg.DBName,cfg.DBPort)
 
  DB,err:=gorm.Open(postgres.Open(dsn),&gorm.Config{})

    if err!=nil{
    return Repository{},err
  }
  err=DB.AutoMigrate(&models.Users{})
  if err!= nil{
    return Repository{},err
  }
  err= DB.AutoMigrate(&models.Admin{})
  if err!= nil{
    return Repository{},err
  }
  
  return Repository{DB: DB},nil
}

