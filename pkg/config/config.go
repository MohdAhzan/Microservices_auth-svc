package config

import "github.com/spf13/viper"


type Config struct{
  Port       string `mapstructure:"PORT"`
  DBHost       string `mapstructure:"DBHost"`
  DBPort       string `mapstructure:"DBPort"`
  DBUser       string `mapstructure:"DBUser"`
  DBPassword   string `mapstructure:"DBPassword"`
  DBName       string `mapstructure:"DBName"`
  JWTSecretKey string `mapstructure:"JWT_SECRET_KEY"`
  ADMINPassword string  `mapstructure:"ADMIN_PASSWORD"`
  ADMINName string  `mapstructure:"ADMIN_PASSWORD"`
  ADMINEmail string  `mapstructure:"ADMIN_EMAIL"`
}


func LoadConfig()( cfg Config ,err error){

  viper.AddConfigPath(".")
  viper.SetConfigName(".auth.env")
  viper.SetConfigType("env")
  viper.AutomaticEnv()

  err = viper.ReadInConfig()
  if err != nil {
    return 
  }

  err = viper.Unmarshal(&cfg)
  if err!=nil{
    return 
  }

  return 

}

