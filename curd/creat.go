package curd

import (
	"fmt"
	"xorm/model"
	"xorm/vlog"

	"xorm.io/xorm"
)

func C(engine *xorm.Engine) {
	// 添加一条数据
	insertOone(engine)

	// 插入多条数据
	insertMany(engine)

	// 插入一组数据
	insertArray(engine)
}

// 添加一条数据
func insertOone(engine *xorm.Engine) {
	user := model.User{Id: uint(ids[0]), Account: accounts[0], Name: "vcz" + accounts[0], Gender: "女", Age: 12, Address: "US", Email: accounts[0] + "@gmail.com"}
	num, err := engine.Insert(&user)
	if num < 1 {
		vlog.Vlog("数据添加失败", err)
	} else {
		fmt.Printf("用户[%s]添加成功\n共%d条数据添加成功!\n", user.Account, num)
	}
}

// 插入多条数据
func insertMany(engine *xorm.Engine) {
	user1 := model.User{Id: uint(ids[1]), Account: accounts[1], Name: "vcz" + accounts[1], Gender: "男", Age: 16, Address: "US", Email: accounts[1] + "@gmail.com"}
	user2 := model.User{Id: uint(ids[2]), Account: accounts[2], Name: "vcz" + accounts[2], Gender: "女", Age: 17, Address: "CN", Email: accounts[2] + "@gmail.com"}
	num, err := engine.Insert(&user1, &user2)
	if num < 2 {
		vlog.Vlog("数据添加失败", err)
	} else {
		fmt.Printf("用户[%s]添加成功\n用户[%s]添加成功\n共%d条数据添加成功!\n", user1.Account, user2.Account, num)
	}
}

// 插入一组数据
func insertArray(engine *xorm.Engine) {
	var users []*model.User
	user1 := model.User{Account: accounts[3], Name: "vcz" + accounts[3], Gender: "男", Age: 26, Address: "US", Email: accounts[3] + "@gmail.com"}
	user2 := model.User{Account: accounts[4], Name: "vcz" + accounts[4], Gender: "女", Age: 25, Address: "CN", Email: accounts[4] + "@gmail.com"}
	user3 := model.User{Account: accounts[5], Name: "vcz" + accounts[5], Gender: "男", Age: 33, Address: "US", Email: accounts[5] + "@gmail.com"}
	users = append(users, &user1, &user2, &user3)
	num, err := engine.Insert(&users)
	if num < 3 {
		vlog.Vlog("数据添加失败", err)
	} else {
		fmt.Printf("用户[%s]添加成功\n用户[%s]添加成功\n用户[%s]添加成功\n共%d条数据添加成功!\n", user1.Account, user2.Account, user3.Account, num)
	}
}
