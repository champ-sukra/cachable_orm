package main

import "time"

type Customer struct {
	Id             uint32    `json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email"`
	FacebookId     string    `json:"facebook_id"`
	Language       string    `json:"language"`
	CreatedDate    time.Time `json:"created_date,omitempty"`
	UpdatedDate    time.Time `json:"updated_date,omitempty"`
	ProfilePicture string    `json:"profile_picture"`
	MobileNumber   string    `json:"mobile_number"`
	BirthDay       string    `json:"birth_day"`
	BirthMonth     string    `json:"birth_month"`
	BirthYear      string    `json:"birth_year"`
	GenderID       string    `json:"gender_id"`
	PrefixID       string    `json:"prefix_id"`
	BigCard        string    `json:"bigcard"`
	IdCard         string    `json:"id_card"`
}
