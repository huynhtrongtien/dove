package entities

type Category struct {
	ID       int64      `gorm:"column:id;primaryKey" redis:"id"`
	UUID     string     `gorm:"column:uuid;omitempty" redis:"uuid"`
	Code     string     `gorm:"column:code;size:200" redis:"code"`
	FullName string     `gorm:"column:fullname;size:200" redis:"fullname"`
	Products []*Product `gorm:"references:id;foreignKey:category_id"`
	Base
}

func (*Category) TableName() string {
	return "categories"
}
