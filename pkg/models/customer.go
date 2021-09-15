package model

const (
	CustomersTableName = "customer"
)

type CutomerInfo struct {
	ID    int32  `gorm:"column:id"`
	Name  string `gorm:"column:name"`
	Phone string `gorm:"column:phone"`
}

func (c CutomerInfo) TableName() string {
	return CustomersTableName
}
