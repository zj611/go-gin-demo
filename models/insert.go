package models

type Customer struct {
	ID         int    `gorm:"primary_key" json:"id"`
	User       string `gorm:"column:user;type:varchar(200)" json:"user"`
	CreatedOn  int    `json:"created_on"`
	ModifiedOn int    `json:"modified_on"`
	DeletedOn  int    `json:"deleted_on"`
}

func InsertCustomer(name string) error {
	customer := Customer{
		User: name,
	}
	if err := db.Create(&customer).Error; err != nil {
		return err
	}

	return nil
}
