package model

type User struct {
	ID      uint
	Account string
	Name    string
	Gender  string
	Age     int
	Address string
}

// 为User结构体对应数据库表设置表名
func (User) TableName() string {
	return "xorm_user"
}
