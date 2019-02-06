package models

import (
	"time"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_"github.com/Go-SQL-Driver/MySQL"
)

// models 层 init 函数连接数据库，并注册模型
func init(){
	var dbhost string
	var dbport string
	var dbuser string
	var dbpassword string
	var dbname string

	// 从配置文件中获取数据库相关配置
	dbhost = beego.AppConfig.String("dbhost")
	dbport = beego.AppConfig.String("dbport")
	dbuser = beego.AppConfig.String("dbuser")
	dbpassword = beego.AppConfig.String("dbpassword")
	dbname = beego.AppConfig.String("dbname")

	// 注册 MySQL 驱动
	orm.RegisterDriver("mysql",orm.DRMySQL)

	// 拼接连接信息
	conn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"

	// 注册数据库连接
	orm.RegisterDataBase("default","mysql",conn)

	// 注册对象模型
	orm.RegisterModel(new(UserInfo))
	orm.RunSyncdb("default",false,true)
}

// 用户结构体
type UserInfo struct {
	Id        int       // 用户 Id
	UserName  string    // 用户名
	UserPwd   string    // 密码
	Remark    string    // 备注
	AddDate   time.Time // 添加日期
	ModifDate time.Time // 修改日期
	DelFlag   int       // 删除标记 软删除
}
