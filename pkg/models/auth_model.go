package models

type Users struct {

  Id int64 `json:"id" gorm:"primaryKey"` 
  Name string `json:"name"`
  Email string `json:"email"`
  Password string `json:"password"`

}

type UsersLogin struct{
  Email string 
  Password string 

}

type Admin struct{

  Id int64 `json:"id" gorm:"primaryKey"` 
  Name string `json:"name"`
  Email string `json:"email"`
  Password string `json:"password"`

}
