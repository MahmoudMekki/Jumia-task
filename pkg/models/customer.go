package model

const (
	CustomersTableName = "customer"
)

// CutomerInfo /*CustomerInfo model for the customer table*/
type CustomerInfo struct {
	ID    int32  `gorm:"column:id"`
	Name  string `gorm:"column:name"`
	Phone string `gorm:"column:phone"`
}

/*returns TableName for Gorm*/
func (c CustomerInfo) TableName() string {
	return CustomersTableName
}
