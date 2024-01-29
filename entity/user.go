package entity

import "time"

type User struct {
	Id          uint      `gorm:"primaryKey;autoIncrement"`
	Email       string    `gorm:"not null;unique"`
	Password    string    `gorm:"not null"`
	Name        string    `gorm:"not null"`
	Address     string    `gorm:"not null"`
	DateOfBirth time.Time `gorm:"not null"`
}
