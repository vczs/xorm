package main

import (
	"fmt"
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

	// 创建表
	err = engine.CreateTables(&model.User{})
	if err != nil {
		vlog.Vlog("数据库表创建失败: ", err)
		return
	}
	fmt.Println("数据库表创建成功！")
}
