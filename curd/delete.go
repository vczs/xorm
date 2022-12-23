package curd

import (
	"fmt"
	"xorm/model"
	"xorm/vlog"

	"xorm.io/xorm"
)

func D(engine *xorm.Engine) {
	del(engine)
}

func del(engine *xorm.Engine) {
	// 通过主键id删除数据(软删除)
	affected, _ := engine.ID(ids[1]).Delete(new(model.User))
	if affected < 1 {
		fmt.Printf("id=%d数据删除失败!\n", ids[1])
	} else {
		fmt.Printf("id=%d数据删除成功!\n", ids[1])
	}

	// 删除全部数据(软删除)
	num, err := engine.Where("1=1").Delete(new(model.User))
	// engine.Delete(new(model.User))
	if err != nil {
		vlog.Vlog("删除全部数据失败", err)
	} else {
		fmt.Printf("删除全部数据成功,共删除%d条数据!\n", num)
	}

	// 硬删除
	num, err = engine.ID(ids[4]).Unscoped().Delete(new(model.User))
	if num < 1 {
		vlog.Vlog("硬删除数据失败", err)
	} else {
		fmt.Println("硬删除数据成功!")
	}
}
