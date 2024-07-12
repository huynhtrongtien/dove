package entities

type Product struct {
	ID         int64     `gorm:"column:id;primaryKey" redis:"id"`
	UUID       string    `gorm:"column:uuid;omitempty" redis:"uuid"`
	Code       string    `gorm:"column:code;size:200" redis:"code"`
	FullName   string    `gorm:"column:fullname;size:200" redis:"fullname"`
	CategoryID int64     `gorm:"column:category_id;omitempty" redis:"category_id"`
	Category   *Category `gorm:"references:category_id;foreignKey:id" redis:"-"`
	Base
}

func (*Product) TableName() string {
	return "products"
}
