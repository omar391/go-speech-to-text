package models

type User struct {
	ID       uint
	Name     string  `gorm:"type:TEXT(50) NOT NULL; check:name <> ''"`
	Email    string  `gorm:"type:TEXT(50) NOT NULL; check:email <> ''"`
	Password string  `gorm:"type:TEXT NOT NULL; check:password <> ''"`
}
