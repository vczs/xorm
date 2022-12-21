package curd

import (
	"fmt"
	"xorm/model"

	"xorm.io/xorm"
)

func C(engine *xorm.Engine) {
	// 添加一条数据
	_, err := engine.Insert(&model.User{Account: "vczs001", Name: "vcz01", Gender: "女", Age: 12, Address: "US", Email: "vczs001@gmail.com"})
	if err != nil {
		fmt.Println("数据添加失败:", err)
	}
	_, err = engine.Insert(&model.User{Account: "vczs002", Name: "vcz02", Address: "CN", Email: "vczs002@gmail.com"})
	if err != nil {
		fmt.Println("数据添加失败:", err)
	}
}
