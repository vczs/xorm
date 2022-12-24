package curd

import (
	"fmt"
	"xorm/model"
	"xorm/vlog"

	"xorm.io/xorm"
)

var ids = []int{40, 41, 42, 43, 44, 45, 46}
var accounts = []string{"118540", "118541", "118542", "118543", "118544", "118545", "118546"}

func Comm(engine *xorm.Engine) {
	// 通用功能
	myComm(engine)

	// 事务
	MyTransaction(engine)
}

// 通用功能
func myComm(engine *xorm.Engine) {
	// 添加、查询、更新时排除指定字段  （添加时排除的字段给默认值）
	num, err := engine.Omit("user_age", "gender").Insert(&model.User{Account: accounts[6], Name: "vcz" + accounts[6], Address: "US", Email: accounts[6] + "@gmail.com"})
	if num < 1 {
		vlog.Vlog("添加、查询、更新时排除指定字段失败", err)
	} else {
		fmt.Println("添加、查询、更新时排除指定字段执行成功！")
	}

	// 执行SQL
	sql := "update xorm_user set name=? where id=?"
	res, err := engine.Exec(sql, "exe-sql", ids[5])
	if num, _ := res.RowsAffected(); num > 1 {
		vlog.Vlog("执行SQL失败", err)
	} else {
		fmt.Println("执行SQL成功!")
	}
}

// 事务
func MyTransaction(engine *xorm.Engine) {
	// 创建事务
	session := engine.NewSession()
	defer session.Close()

	// 开启事务
	if err := session.Begin(); err != nil {
		vlog.Vlog("开启事务失败", err)
		return
	}

	// 捕获该函数异常
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("事务执行失败!")
			session.Rollback() // 如果函数发生panic就回滚
		} else {
			fmt.Println("事务执行成功!")
			session.Commit() // 如果函数没有发生panic就提交事务
		}
	}()

	// 开始业务逻辑
	user1 := model.User{Account: "shiwu001", Name: "shiwu01", Gender: "男", Age: 38, Address: "US", Email: "shiwu001@gmail.com"}
	num, err := session.Insert(&user1)
	if num < 1 {
		panic(err)
	}
	fmt.Printf("事务操作 添加[%s]用户 id=%d\n", user1.Account, user1.Id)

	user2 := new(model.User)
	session.Table(&model.User{}).ID(user1.Id).Get(user2)
	affected, _ := session.ID(user1.Id).Update(&model.User{Name: "shiwuupdate", Age: 30, Ver: user2.Ver})
	if affected < 1 {
		panic(err)
	}
	fmt.Printf("事务操作 更新[%s]用户name为%s\n", user2.Account, "shiwuupdate")

	affected, _ = session.ID(user1.Id).Delete(new(model.User))
	if affected < 1 {
		panic(err)
	}
	fmt.Printf("事务操作 删除id为%d的用户\n", user1.Id)
}
