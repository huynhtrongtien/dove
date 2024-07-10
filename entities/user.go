package entities

type Role int32

const (
	Admin   Role = 1
	EndUser Role = 2
)

func (s Role) Value() int32 {
	return int32(s)
}

type User struct {
	ID           int64  `gorm:"column:id;primaryKey"`
	UUID         string `gorm:"column:displayname;size:200"`
	DisplayName  string `gorm:"column:displayname;size:200"`
	Username     string `gorm:"column:username;size:200"`
	PasswordHash string `gorm:"column:password_hash;size:200"`
	Role         Role   `gorm:"column:role"`
}

func (*User) TableName() string {
	return "users"
}
