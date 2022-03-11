package main

type PersonalInformation struct {
	Id int64 `json:"id,omitempty" gorm:"primaryKey;column:id"`
	Name string `json:"name,omitempty" gorm:"column:name"`
	Sex string `json:"sex,omitempty" gorm:"column:sex"`
	Tall float32 `json:"tall,omitempty" gorm:"column:tall"`
	Weight float32 `json:"weight,omitempty" gorm:"column:weight"`
	Age int64 `json:"age,omitempty" gorm:"column:age"`
}

func (*PersonalInformation) TableName() string {
	return "personal_info"
}
