package main

import (
	"fmt"
	"log"

	"github.com/MohdAhzan/auth-svc/pkg/config"
)


func main(){

 cfg, err:=config.LoadConfig()
  if err!=nil{
    log.Fatal("error loading authCOnfig",err)
  }
  fmt.Println("cfg",cfg)



}
