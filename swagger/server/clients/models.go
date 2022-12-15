package clients

type Customer struct {
	customerId uint   `gorm:"bigserial;unique:id_UNIQUE"`
	FirstName  string `gorm:"size:45;not null"`
	LastName   string `gorm:"size:45"`
	PhoneNo    uint   `gorm:"unique:phoneNo_UNIQUE"`
	Password   string `gorm:"size:10"`
	Email      string `gorm:"size:60;primarykey"`
}

func (m *Customer) TableName() string {
	return "customer"
}
