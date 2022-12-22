package curd

import (
	"fmt"
	"xorm/model"
	"xorm/vlog"

	"xorm.io/xorm"
)

func R(engine *xorm.Engine) {
	// 查询一条数据
	getData(engine)

	// 查询所有数据
	findData(engine)
}

// 查询一条数据 Get方法
func getData(engine *xorm.Engine) {
	// 根据id获取单条数据
	user1 := new(model.User)
	has, _ := engine.ID(102).Get(user1)
	if !has {
		fmt.Printf("id等于%d的数据不存在!\n", 102)
		return
	}
	fmt.Printf("查询到id等于%d的数据:%v\n", 102, *user1)

	// 根据结构体中已有的非空数据来获得单条数据
	user2 := model.User{Account: "vczs002"}
	has, _ = engine.Get(&user2)
	if !has {
		fmt.Printf("account等于%s的数据不存在!\n", "vczs002")
		return
	}
	fmt.Printf("查询到account等于%s的数据:%v\n", "vczs002", user2)

	// 根据Where来获得单条数据
	user3 := model.User{}
	has, _ = engine.Where("account=?", "vczs003").Get(&user3)
	if !has {
		fmt.Printf("account等于%s的数据不存在!\n", "vczs003")
		return
	}
	fmt.Printf("查询到account等于%s的数据:%v\n", "vczs003", user3)
}

// 查询所有数据 Find方法
func findData(engine *xorm.Engine) {
	// Slice接收数据
	sUsers := make([]model.User, 0)
	err := engine.Find(&sUsers)
	if err != nil {
		vlog.Vlog("Slice接收数据 查询失败:", err)
	}
	fmt.Println(sUsers)

	// Map接收数据
	// map的key必须为int类型,key为数据主键id(对于复合主键无法使用这种方式)
	mUsers := make(map[int64]*model.User)
	err = engine.Find(&mUsers)
	if err != nil {
		vlog.Vlog("Map接收数据 查询失败:", err)
	}
	// 因为mUsers的value为User的指针 所以要遍历mUsers获取每个value指向的值
	for k, v := range mUsers {
		fmt.Println(k, *v)
	}

	// Where获取符合要求的数据
	wUsers := make([]model.User, 0)
	err = engine.Where("user_age > ? or address = ?", 16, "CN").Limit(10, 0).Find(&wUsers)
	if err != nil {
		vlog.Vlog("Where获取符合要求的数据 查询失败:", err)
	}
	fmt.Println(wUsers)

	// 查询单个字段
	var strs []string
	err = engine.Table(&model.User{}).Cols("account").Find(&strs)
	if err != nil {
		vlog.Vlog("查询单个字段 查询失败:", err)
	}
	fmt.Println(strs)
}
