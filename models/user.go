package models

type User struct {
	ID       uint
	Name     string `gorm:"type:TEXT(50) NOT NULL; check:name <> ''"`
	Email    string `gorm:"type:TEXT(50) NOT NULL; check:email <> ''"`
	Password []byte `gorm:"type:TEXT(50) NOT NULL; check:BLOB NOT NULL; check:password <> ''"`
}

type PasswordParam struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}
