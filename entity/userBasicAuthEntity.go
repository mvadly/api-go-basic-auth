package entity

type UserBasicAuth struct {
	ID       int    `gorm:"column:id;type:int(11) autoIncrement;primaryKey;"`
	Username string `gorm:"column:username;type:varchar(20);not null:"`
	Password string `gorm:"column:password;type:varchar(100);not null;"`
}

func (UserBasicAuth) TableName() string {
	return "user_basic_auth"
}
