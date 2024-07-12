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
	ID           int64  `gorm:"column:id;primaryKey" redis:"id"`
	UUID         string `gorm:"column:displayname;size:200" redis:"uuid"`
	DisplayName  string `gorm:"column:displayname;size:200" redis:"displayname"`
	Username     string `gorm:"column:username;size:200" redis:"-"`
	PasswordHash string `gorm:"column:password_hash;size:200" redis:"-"`
	Role         Role   `gorm:"column:role" redis:"-"`
}

func (*User) TableName() string {
	return "users"
}
