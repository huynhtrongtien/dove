package entities

type Category struct {
	ID       int64      `gorm:"column:id;primaryKey"`
	UUID     string     `gorm:"column:uuid;size:200"`
	Code     string     `gorm:"column:code;size:200"`
	FullName string     `gorm:"column:fullname;size:200"`
	Products []*Product `gorm:"references:id;foreignKey:category_id"`
	Base
}

func (*Category) TableName() string {
	return "categories"
}
