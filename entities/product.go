package entities

type Product struct {
	ID         int64     `gorm:"column:id;primaryKey"`
	UUID       string    `gorm:"column:displayname;size:200"`
	Code       string    `gorm:"column:code;size:200"`
	FullName   string    `gorm:"column:displayname;size:200"`
	CategoryID int64     `gorm:"column:category_id;omitempty"`
	Category   *Category `gorm:"references:category_id;foreignKey:id"`
	Base
}

func (*Product) TableName() string {
	return "products"
}
