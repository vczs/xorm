package curd

import (
	"fmt"
	"xorm/model"

	"xorm.io/xorm"
)

func U(engine *xorm.Engine) {
	// 通过主键id更新数据
	user1 := new(model.User)
	// 因为user的Ver字段加了乐观锁,Update的内容必须包含version原来的值,所以要先查询要更新数据的ver值
	engine.ID(ids[1]).Get(user1)
	affected, _ := engine.ID(ids[1]).Update(&model.User{Name: "ids01", Age: 30, Ver: user1.Ver})
	if affected < 1 {
		fmt.Printf("id=%d数据更新失败!\n", ids[1])
	} else {
		fmt.Printf("id=%d数据更新成功!\n", ids[1])
	}

	// 指定更新的字段 未指定的将不更新 (指定了的即使为0或空字符串也会更新)
	user2 := new(model.User)
	engine.ID(ids[2]).Get(user2)
	affected, _ = engine.ID(ids[2]).Cols("name", "user_age").Update(&model.User{Name: "cols01", Ver: user2.Ver})
	if affected < 1 {
		fmt.Printf("id=%d数据更新失败!\n", ids[2])
	} else {
		fmt.Printf("id=%d数据更新成功!\n", ids[2])
	}

	// 指定必须更新的字段 (其它字段根据值的情况自动判断)
	user3 := new(model.User)
	engine.ID(ids[3]).Get(user3)
	affected, _ = engine.ID(ids[3]).MustCols("name").Update(&model.User{Name: "must01", Age: 38, Account: "must001", Address: "AB", Ver: user3.Ver})
	if affected < 1 {
		fmt.Printf("id=%d数据更新失败!\n", ids[3])
	} else {
		fmt.Printf("id=%d数据更新成功!\n", ids[3])
	}

	// 更新所有的字段 未赋值字段为对应类型的零值
	user4 := new(model.User)
	engine.ID(ids[4]).Get(user4)
	affected, _ = engine.ID(ids[4]).AllCols().Update(&model.User{Name: "all01", Age: 38, Account: "all001", Ver: user4.Ver})
	if affected < 1 {
		fmt.Printf("id=%d数据更新失败!\n", ids[4])
	} else {
		fmt.Printf("id=%d数据更新成功!\n", ids[4])
	}

	// 传入map进行更新 （需要额外指定更新到哪个表，因为通过map是无法自动检测更新哪个表）
	user5 := new(model.User)
	engine.ID(ids[5]).Get(user5)
	affected, _ = engine.Table(&model.User{}).ID(ids[5]).Update(map[string]interface{}{"name": "map01", "user_age": 10, "account": "map001"})
	if affected < 1 {
		fmt.Printf("id=%d数据更新失败!\n", ids[5])
	} else {
		fmt.Printf("id=%d数据更新成功!\n", ids[5])
	}
}
