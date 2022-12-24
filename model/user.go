package model

type User struct {
	Id      uint             `xorm:"pk autoincr"`                                // 主键 自增
	Name    string           `xorm:"varchar(120) comment(姓名)"`                   // 指定类型 注释
	Gender  string           `xorm:"default('男') comment(性别)"`                   // 默认值 注释
	Age     uint             `xorm:"int 'user_age' default(18) comment(年龄)"`     // 指定类型 指定列名 默认值 注释
	Address string           `xorm:"notnull comment(地址)"`                        // 不为空 注释
	Account string           `xorm:"varchar(120) index unique(acc) comment(账号)"` // 指定类型 设置为索引 唯一索引 注释
	Email   string           `xorm:"notnull unique index(ema) comment(邮箱)"`      // 不为空 唯一 设置为索引并添加索引名 注释
	Ver     int              `xorm:"version comment(版本)"`                        // 在insert时默认为1且每次更新自动加1() 注释
	Other   string           `xorm:"-"`                                          // 不进行字段映射(这个字段不写入数据库)
	Comm    `xorm:"extends"` // 此结构体的所有成员也映射到数据库中，extends可加载无限级
}
type Comm struct {
	Create int64 `xorm:"created 'ct'"` // 在Insert时自动赋值为当前时间 指定列名
	Update int64 `xorm:"updated 'ut'"` // 在Insert或Update时自动赋值为当前时间 指定列名
	Delete int64 `xorm:"deleted 'dt'"` // 在Delete时设置为当前时间并且软删除 指定列名
}

// 为User结构体对应数据库表指定表名
func (User) TableName() string {
	return "xorm_user"
}

type Work struct {
	Id     int `xorm:"pk autoincr"`
	UserId int
	Works  string
}

// 为User结构体对应数据库表指定表名
func (Work) TableName() string {
	return "work"
}

type UserWork struct {
	TempUser User `xorm:"extends"`
	TempWork Work `xorm:"extends"`
}
