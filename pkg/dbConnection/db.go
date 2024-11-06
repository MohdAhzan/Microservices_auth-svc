package dbconnection

import (
	"database/sql"
	"fmt"

	"github.com/MohdAhzan/auth-svc/pkg/config"
	"github.com/MohdAhzan/auth-svc/pkg/models"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
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
  
  err=CheckAndCreateAdmin(cfg,DB)
  if err!=nil{
    return Repository{},err
  }
  return Repository{DB: DB},nil
}


func CheckAndCreateAdmin(cfg config.Config, db *gorm.DB)error {
	var count int64
	db.Model(&models.Admin{}).Count(&count)
	if count == 0 {
		password := cfg.ADMINPassword
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		admin := models.Admin{
			Id:       1,
			Name:     cfg.ADMINName,
			Email:    cfg.ADMINEmail,
			Password: string(hashedPassword),
		}

		db.Create(&admin)
	}

    return nil
}

