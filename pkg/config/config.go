package config

import "github.com/spf13/viper"


type Config struct{

	DBHost       string `mapstructure:"DBHost"`
	DBPort       string `mapstructure:"DBPort"`
	DBUser       string `mapstructure:"DBUser"`
	DBPassword   string `mapstructure:"DBPassword"`
	DBName       string `mapstructure:"DBName"`
}


func LoadConfig()( cfg Config ,err error){
  
    
    viper.AddConfigPath(".")
    viper.SetConfigName(".auth.env")
    viper.SetConfigType("env")

    viper.AutomaticEnv()

  err = viper.ReadInConfig()
    if err != nil {
        return cfg,err
    }

    err = viper.Unmarshal(&cfg)
    if err!=nil{
    return cfg,err
  }
    
    return cfg ,nil

}
