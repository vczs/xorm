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
	user := model.User{Id: 10, Account: "vczs001", Name: "vcz01", Gender: "女", Age: 12, Address: "US", Email: "vczs001@gmail.com"}
	num, err := engine.Insert(&user)
	if num < 1 {
		vlog.Vlog("数据添加失败", err)
	} else {
		fmt.Printf("用户[%s]添加成功\n", user.Account)
	}
}

// 插入多条数据
func insertMany(engine *xorm.Engine) {
	user1 := model.User{Id: 11, Account: "vczs002", Name: "vcz02", Gender: "男", Age: 16, Address: "US", Email: "vczs002@gmail.com"}
	user2 := model.User{Id: 12, Account: "vczs003", Name: "vcz03", Gender: "女", Age: 17, Address: "CN", Email: "vczs003@gmail.com"}
	num, err := engine.Insert(&user1, &user2)
	if num < 2 {
		vlog.Vlog("数据添加失败", err)
	} else {
		fmt.Printf("用户[%s]添加成功\n用户[%s]添加成功\n共%d条数据全部添加成功!\n", user1.Account, user2.Account, num)
	}
}

// 插入一组数据
func insertArray(engine *xorm.Engine) {
	var users []*model.User
	user1 := model.User{Account: "vczs004", Name: "vcz04", Gender: "男", Age: 26, Address: "US", Email: "vczs004@gmail.com"}
	user2 := model.User{Account: "vczs005", Name: "vcz05", Gender: "女", Age: 25, Address: "CN", Email: "vczs005@gmail.com"}
	user3 := model.User{Account: "vczs006", Name: "vcz06", Gender: "未知", Age: 33, Address: "US", Email: "vczs006@gmail.com"}
	users = append(users, &user1, &user2, &user3)
	num, err := engine.Insert(&users)
	if num < 3 {
		vlog.Vlog("数据添加失败", err)
	} else {
		fmt.Printf("用户[%s]添加成功\n用户[%s]添加成功\n用户[%s]添加成功\n共%d条数据全部添加成功!\n", user1.Account, user2.Account, user3.Account, num)
	}
}
