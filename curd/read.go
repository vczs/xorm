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

	// 查询多条数据
	findData(engine)

	// 条件查询
	condData(engine)

	// 联表查询
	joinData(engine)

	// 其他查询
	otherData(engine)
}

// 查询一条数据 Get方法
func getData(engine *xorm.Engine) {
	// 根据id获取单条数据
	user1 := new(model.User)
	has, _ := engine.ID(ids[0]).Get(user1)
	if !has {
		fmt.Printf("id等于%d的数据不存在!\n", ids[0])
	} else {
		fmt.Printf("id等于%d的数据:%v\n", ids[0], *user1)
	}

	// 根据结构体中已有的非空数据来获得单条数据
	user2 := model.User{Account: accounts[0]}
	has, _ = engine.Get(&user2)
	if !has {
		fmt.Printf("account等于%s的数据不存在!\n", accounts[0])
	} else {
		fmt.Printf("account等于%s的数据:%v\n", accounts[0], user2)
	}

	// 根据Where来获得单条数据
	user3 := model.User{}
	has, _ = engine.Where("account=?", accounts[1]).Get(&user3)
	if !has {
		fmt.Printf("account等于%s的数据不存在!\n", accounts[1])
	} else {
		fmt.Printf("account等于%s的数据:%v\n", accounts[1], user3)
	}
}

// 查询多条数据 Find方法
func findData(engine *xorm.Engine) {
	// Slice接收数据
	sUsers := make([]model.User, 0)
	err := engine.Find(&sUsers)
	if err != nil {
		vlog.Vlog("Slice接收数据查询失败:", err)
	} else {
		fmt.Println("Slice接收数据:", sUsers)
	}

	// Map接收数据
	// map的key必须为int类型,key为数据主键id(对于复合主键无法使用这种方式)
	mUsers := make(map[int64]*model.User)
	err = engine.Find(&mUsers)
	if err != nil {
		vlog.Vlog("Map接收数据查询失败:", err)
	} else {
		// 因为mUsers的value为User的指针 所以要遍历mUsers获取每个value指向的值
		fmt.Print("Map接收数据:")
		for k, v := range mUsers {
			fmt.Println(k, *v)
		}
	}

	// Where获取符合要求的数据
	wUsers := make([]model.User, 0)
	err = engine.Where("user_age > ? or address = ?", 16, "CN").Limit(10, 0).Find(&wUsers)
	if err != nil {
		vlog.Vlog("Where获取符合要求的数据 查询失败:", err)
	} else {
		fmt.Println("Where获取符合要求的数据:", wUsers)
	}
}

// 条件查询
func condData(engine *xorm.Engine) {
	// Where And 组合条件
	users1 := make([]model.User, 0)
	err := engine.Where("user_age>?", 16).And("address=?", "CN").Find(&users1)
	if err != nil {
		vlog.Vlog("Where And 组合条件查询失败", err)
	} else {
		fmt.Printf("Where And 组合条件查询的数据:%v\n", users1)
	}

	// 指定字段名正序排序（可以组合）
	users2 := make([]model.User, 0)
	err = engine.Asc("name").And("address=?", "US").Find(&users2)
	if err != nil {
		vlog.Vlog("指定字段名正序排序查询失败", err)
	} else {
		fmt.Printf("指定字段名正序排序查询的数据:%v\n", users2)
	}

	// 指定字段名逆序排序（可以组合）
	users3 := make([]model.User, 0)
	err = engine.Desc("name").Asc("user_age").Find(&users3)
	if err != nil {
		vlog.Vlog("指定字段名逆序排序查询失败", err)
	} else {
		fmt.Printf("指定字段名逆序排序查询的数据:%v\n", users3)
	}

	// 查询指定字段
	user1 := new(model.User)
	has, _ := engine.Cols("name", "account", "email").Get(user1)
	if !has {
		fmt.Println("查询指定字段的数据不存在!")
	} else {
		fmt.Printf("查询指定字段的数据:%v\n", *user1)
	}

	// 查询并去重指定字段
	users4 := make([]model.User, 0)
	engine.Distinct("name", "gender").Find(&users4)
	if err != nil {
		vlog.Vlog("查询并去重指定字段查询失败", err)
	} else {
		fmt.Printf("查询并去重指定字段查询的数据:%v\n", users4)
	}

	// 分页查询
	users5 := make([]model.User, 0)
	// Limit(number, start) 第一个参数为条数，第二个参数表示开始位置（如果不传则为0）
	// Top(int): Limit(number, 0)
	engine.Limit(3, 0).Find(&users5)
	if err != nil {
		vlog.Vlog("分页查询查询失败", err)
	} else {
		fmt.Printf("分页查询查询的数据:%v\n", users5)
	}

	// 统计数量
	total, err := engine.Where("user_age >?", 17).Count(new(model.User))
	// total, err := engine.Count(&model.User{Account: accounts[3]})
	if err != nil {
		vlog.Vlog("统计数量 查询失败:", err)
	} else {
		fmt.Println("统计数量user_age大于17岁的数据有:", total)
	}

	// 查询数据是否存在 (与Get()方法相比 Exist()只判断数据是否存在不用将查询结果解析到结构体中 所以Exist的执行效率要比Get更高)
	has, _ = engine.Where("account = ?", accounts[2]).Exist(&model.User{})
	//  engine.Exist(&model.User{Account: accounts[2]})
	//  engine.SQL("select * from xorm_user where account = ?", accounts[2]).Exist()
	// engine.Table(&model.User{}).Where("account = ?", accounts[2]).Exist()
	if !has {
		fmt.Println("account=" + accounts[2] + "的数据不存在!")
	} else {
		fmt.Println("account=" + accounts[2] + "的数据存在!")
	}
}

// 联表查询
func joinData(engine *xorm.Engine) {
	// 创建表(用于联表查询的第二张表)
	err := engine.Sync(new(model.Work))
	if err != nil {
		vlog.Vlog("数据库表创建失败: ", err)
		return
	}

	// 给表添加数据
	var works []*model.Work
	work1 := model.Work{UserId: ids[0], Works: fmt.Sprintf("用户%d在works表中的数据", ids[0])}
	work2 := model.Work{UserId: ids[1], Works: fmt.Sprintf("用户%d在works表中的数据", ids[1])}
	work3 := model.Work{UserId: ids[2], Works: fmt.Sprintf("用户%d在works表中的数据", ids[2])}
	works = append(works, &work1, &work2, &work3)
	num, err := engine.Insert(&works)
	if num < 3 {
		vlog.Vlog("数据添加失败", err)
		return
	}

	// 执行联表查询
	userWork := make([]model.UserWork, 0)
	engine.SQL("SELECT a.*,b.`works` FROM `xorm_user` AS a RIGHT JOIN `work` AS b on a.id  = b.`user_id`").Find(&userWork)
	fmt.Println("联表查询数据:", userWork)
}

// 其他查询
func otherData(engine *xorm.Engine) {
	// 查询单个字段
	var strs []string
	err := engine.Table(&model.User{}).Cols("account").Find(&strs)
	if err != nil {
		vlog.Vlog("查询单个字段 查询失败:", err)
	} else {
		fmt.Println("查询单个字段", strs)
	}

	// Iterate方法提供逐条执行查询到的记录的方法，他所能使用的条件和Find方法完全相同
	fmt.Println("Iterate方法查询数据:")
	engine.Where("name >? or address=?", 17, "CN").Iterate(new(model.User), func(i int, bean interface{}) error {
		user := bean.(*model.User)
		fmt.Printf("第%d条数据:%v\n", i, *user)
		return nil
	})

	// Rows方法和Iterate方法类似，提供逐条执行查询到的记录的方法，不过Rows更加灵活好用。
	rows, err := engine.Rows(&model.User{Address: "US"})
	if err != nil {
		vlog.Vlog("Rows方法查询失败", err)
	}
	defer rows.Close()
	bean := new(model.User)
	fmt.Println("rows方法查询数据:")
	for rows.Next() {
		rows.Scan(bean)
		fmt.Println(*bean)
	}

	// SumInt求某个字段的和,返回int64 (Sum返回float64)
	total, err := engine.Where("id >?", 10).SumInt(new(model.User), "user_age")
	if err != nil {
		vlog.Vlog("SumInt查询失败", err)
	}
	fmt.Printf("umInt求某个字段的和:user_age的和: %d\n", total)

	// SumsInt求某几个字段的和,返回int64的Slice (Sums返回float64的Slice)
	totals, err := engine.Where("id < ?", 103).SumsInt(new(model.User), "user_age", "id")
	if err != nil {
		vlog.Vlog("SumsInt查询失败", err)
	} else {
		fmt.Println("SumsInt求某几个字段的和:")
		for k, v := range totals {
			fmt.Printf("第%d个字段和:%d\n", k+1, v)
		}
	}

	// Query查询 (QueryInterface返回值为[]map[string]interface{}  QueryString返回值为[]map[string]string)
	res, err := engine.Query("select * from " + model.User{}.TableName())
	if len(res) < 1 {
		vlog.Vlog("Query查询失败", err)
	} else {
		fmt.Println("Query查询:", res)
	}

	// 获取软删除的数据
	user1 := model.User{}
	has, _ := engine.ID(ids[2]).Unscoped().Get(&user1)
	if !has {
		fmt.Printf("未查询到id=%d的数据\n", ids[2])
	} else {
		fmt.Printf("查询到id=%d的数据%v\n", ids[2], user1)
	}

	// 执行指定的Sql语句，并把结果映射到结构体。（当选择内容或者条件比较复杂时，可以直接使用Sql）
	users1 := make([]model.User, 0)
	err = engine.SQL("select * from xorm_user where user_age > 18").Find(&users1)
	if err != nil {
		vlog.Vlog("执行指定的Sql语句查询失败", err)
	} else {
		fmt.Println("执行指定的Sql语句查询:", users1)
	}
}
