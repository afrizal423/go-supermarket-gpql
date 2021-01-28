package customer

import "github.com/afrizal423/go-supermarket-gpql/config"

type Customers struct {
	CustomerID    string `gorm:"primary_key;auto_increment;column:customer_id" json:"customer_id"`
	CustomerNames string `gorm:"column:customer_name" json:"customer_name"`
	Segment       string `gorm:"column:segment" json:"segment"`
	Age           string `gorm:"column:age" json:"age"`
	Country       string `gorm:"column:country" json:"country"`
	City          string `gorm:"column:city" json:"city"`
	State         string `gorm:"column:state" json:"state"`
	PostalCode    string `gorm:"column:postal_code" json:"postal_code"`
	Region        string `gorm:"column:region" json:"region"`
}

// GetAllData ...
func GetAllData(limit *int, offset *int) ([]*Customers, error) {
	var err error
	hasil := []*Customers{}
	err = config.Db.Raw("select * from customer limit ? offset ?", limit, offset).
		Scan(&hasil).Error
	if err != nil {
		return []*Customers{}, err
	}
	return hasil, err
}

func GetDataByID(id string) ([]*Customers, error) {
	var err error
	hasil := []*Customers{}
	err = config.Db.Raw("select * from customer where customer_id LIKE ?", "%"+id+"%").
		Scan(&hasil).Error
	if err != nil {
		return []*Customers{}, err
	}
	return hasil, err
}
