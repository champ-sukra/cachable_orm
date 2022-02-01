package main

type Token struct {
	Token     string `gorm:"primary_key"`
	ProfileId int64
}
