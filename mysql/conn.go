package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

const (
	host     = "127.0.0.1"
	port     = 3306
	account  = "root"
	password = "123456"
	charset  = "utf8mb4"
)

func Conn(dbName string) (*xorm.Engine, error) {
	connstr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", account, password, host, port, dbName, charset)
	engine, err := xorm.NewEngine("mysql", connstr)
	if err != nil {
		return nil, err
	}
	return engine, nil
}
