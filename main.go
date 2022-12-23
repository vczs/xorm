package main

import (
	"fmt"
	"xorm/curd"
	"xorm/model"
	"xorm/mysql"
	"xorm/vlog"
)

const (
	dbName = "xorm_db"
)

func main() {
	// 连接数据库
	engine, err := mysql.Conn(dbName)
	if err != nil {
		vlog.Vlog("数据库连接失败: ", err)
		return
	}
	fmt.Println("数据库连接成功！")

	// 创建表(将结构体和数据库表自动同步)
	err = engine.Sync(new(model.User))
	if err != nil {
		vlog.Vlog("数据库表创建失败: ", err)
		return
	}
	fmt.Println("数据库表创建成功！")

	// 是否显示SQL语句
	engine.ShowSQL(false)

	// 增
	curd.C(engine)

	// 查
	curd.R(engine)

	// 改
	curd.U(engine)

	// 删
	curd.D(engine)

	// 通用
	curd.Comm(engine)
}
