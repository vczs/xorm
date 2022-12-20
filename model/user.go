package model

type User struct {
	ID      uint
	Account string
	Name    string
	Gender  string
	Age     uint
	Address string
	Email   string
}

// 为User结构体对应数据库表指定表名
func (User) TableName() string {
	return "xorm_user"
}
