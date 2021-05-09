package db

import(
    "log"
    "fmt"
    "strconv"
    "flowpipe-server/config"

    "github.com/jinzhu/gorm"
     _ "github.com/jinzhu/gorm/dialects/mysql"
)

var WfDb *gorm.DB

func init() {
    // 启动数据库连接
    InitDB()
}

func InitDB() {
    var err error

    c := config.Cfg()
    // TODO: not only mysql
    params := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
        c.Db.Username,
        c.Db.Password,
        c.Db.Host,
        c.Db.Port,
        c.Db.Database,
    )
    WfDb, err = gorm.Open("mysql", params)
    if err != nil {
        log.Fatalf("数据库连接失败 err: %v", err)
    }

    // 启用Logger，显示详细日志
    // 默认情况下会打印发生的错误
    mode, _ := strconv.ParseBool("true")
    WfDb.LogMode(mode)
    // TODO:let table prefix config
    gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
        return defaultTableName
    }

    // 全局禁用表名复数
    // WfDb.SingularTable(true)
    // DB.DB().SetMaxIdleConns(10)
    // DB.DB().SetMaxOpenConns(100)
    log.Println("database init on port ", c.Db.Host)
}
