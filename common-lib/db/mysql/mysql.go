package mysql

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/FlashFeiFei/my-gin/common-lib/db"
)

//db连接信息
//不同的key，连接不同的数据库，主从
//不需要加锁,因为只是在main入口函数导入而已
var db_map map[string]*gorm.DB

func init() {
	if db_map == nil {
		db_map = make(map[string]*gorm.DB)
	}
}

func InitDBConnect(db_info ...db.DBInfo) {
	if len(db_info) == 0 {
		return
	}

	for _, item := range db_info {
		if _, ok := db_map[item.Key]; !ok {
			connect_db, err := gorm.Open("mysql", item.Dsn)
			if err != nil {
				panic(err)
			}
			db_map[item.Key] = connect_db
		}
	}
}

//获取db
func GetDBConnect(key string) *gorm.DB {
	if connect_db, ok := db_map[key]; ok {
		return connect_db
	}
	panic(errors.New("找不到连接: " + key))
}
