package curd

import (
	"fmt"
	"xorm/model"

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
	user := model.User{Account: "vczs001", Name: "vcz01", Gender: "女", Age: 12, Address: "US", Email: "vczs001@gmail.com"}
	_, err := engine.Insert(&user)
	if err != nil {
		fmt.Println("数据添加失败:", err)
		return
	}
	fmt.Printf("用户[%s]添加成功 id=%d\n", user.Account, user.Id)
}

// 插入多条数据
func insertMany(engine *xorm.Engine) {
	user1 := model.User{Id: 102, Account: "vczs002", Name: "vcz02", Gender: "男", Age: 16, Address: "US", Email: "vczs002@gmail.com"}
	user2 := model.User{Id: 103, Account: "vczs003", Name: "vcz03", Gender: "女", Age: 17, Address: "CN", Email: "vczs003@gmail.com"}
	_, err := engine.Insert(&user1, &user2)
	if err != nil {
		fmt.Println("数据添加失败:", err)
		return
	}
	fmt.Printf("用户[%s]添加成功 id=%d\n", user1.Account, user1.Id)
	fmt.Printf("用户[%s]添加成功 id=%d\n", user2.Account, user2.Id)
}

// 插入一组数据
func insertArray(engine *xorm.Engine) {
	var users []*model.User
	user1 := model.User{Account: "vczs004", Name: "vcz04", Gender: "男", Age: 26, Address: "US", Email: "vczs004@gmail.com"}
	user2 := model.User{Account: "vczs005", Name: "vcz05", Gender: "女", Age: 25, Address: "CN", Email: "vczs005@gmail.com"}
	users = append(users, &user1, &user2)
	num, err := engine.Insert(&users)
	if err != nil {
		fmt.Println("数据添加失败:", err)
		return
	}
	fmt.Println("成功添加", num, "条数据")
}
