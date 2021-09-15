package repo

import (
	"gorm.io/gorm"
	model "jumia-task/pkg/models"
)

type CustomersRepo interface {
	GetCustomers() (customers []model.CustomerInfo, errr error)
	GetCustomersByCountry(pagination model.Pagination) (customers []model.CustomerInfo, err error)
	GetCustomersPagination(pagination model.Pagination) (customers []model.CustomerInfo, err error)
}

func NewCustomersRepo(db *gorm.DB) CustomersRepo {
	return &customersImp{DBEngine: db}
}

type customersImp struct {
	DBEngine *gorm.DB
}

func (c customersImp) GetCustomers() (customers []model.CustomerInfo, err error) {
	err = c.DBEngine.Table(model.CustomersTableName).Find(&customers).Error
	return customers, err
}

func (c customersImp) GetCustomersByCountry(pagination model.Pagination) (customers []model.CustomerInfo, err error) {
	err = c.DBEngine.Raw("select * from customer where phone like '%%?%%'", pagination.FilterBy).
		Offset(int((pagination.Page - 1) * pagination.Limit)).
		Limit(int(pagination.Limit)).Find(&customers).Error
	return customers, err
}
func (c customersImp) GetCustomersPagination(pagination model.Pagination) (customers []model.CustomerInfo, err error) {
	err = c.DBEngine.Table(model.CustomersTableName).Find(&customers).
		Offset(int((pagination.Page - 1) * pagination.Limit)).
		Limit(int(pagination.Limit)).Find(&customers).Error
	return customers, err
}
