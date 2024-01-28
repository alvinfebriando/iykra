package entity

import "time"

type User struct {
	Id          uint      `gorm:"primaryKey;autoIncrement"`
	Name        string    `gorm:"not null"`
	Address     string    `gorm:"not null"`
	DateOfBirth time.Time `gorm:"not null"`
}
