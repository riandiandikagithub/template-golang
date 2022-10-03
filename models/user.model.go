package models

import "time"

// CREATE TABLE core.tbl_user (
// 	id int8 NOT NULL DEFAULT nextval('core.newtable_id_seq'::regclass),
// 	firstname varchar NULL,
// 	lastname varchar NULL,
// 	email varchar NULL,
// 	username varchar NULL,
// 	"password" varchar NULL,
// 	createdat timestamp NOT NULL DEFAULT now(),
// 	photo varchar NULL
// );

type User struct {
	Firstname string    `json:"firstname" gorm:"column:firstname"`
	Lastname  string    `json:"lastname" gorm:"column:lastname"`
	Email     string    `json:"email" gorm:"column:email"`
	Username  string    `json:"username" gorm:"column:username"`
	Password  string    `json:"password" gorm:"column:password"`
	Photo     string    `json:"photo" gorm:"column:photo"`
	Createdat time.Time `json:"createdat" gorm:"column:createdat"`
}
