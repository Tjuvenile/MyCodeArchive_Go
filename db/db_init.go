package db

import (
	"MyCodeArchive_Go/logging"
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

type DbRepo struct {
	DbConn *gorm.DB
}

const (
	UserName        = "storage"      //用户名
	EncryptPassword = "c3RvckAxMjM=" //加密密码
	DbName          = "storage"      //数据库名
	maxRetries      = 3              //最大重试次数

	// mariadb连接信息
	MariadbIP   = "gal-writer-svc.ccos-mariadb.svc" //Mariadb IP
	MariadbPort = "3306"                            //Mariadb 端口
	// mysql连接信息
	MysqlIP   = "ccos-mysql.ccos-mysql.svc" // Mysql IP
	MysqlPort = "3331"                      // Mysql Port
)

var DBPassword string
var DbConnect DbRepo

// ProductName 产品形态
var ProductName string

type SessionConfig struct {
	SessionTime float64 //session最大持续时间，单位：秒
}

func (d *DbRepo) GetSession(conf *SessionConfig) (*DbRepo, func()) {
	var tx *gorm.DB
	cancelCtx, cancel := context.WithCancel(context.Background())
	if *d == (DbRepo{}) {
		logging.Log.Infof("database connection does not exist, start connecting...")
		if ret := ConnectDB(); ret != 0 {
			logging.Log.Errorf("database connection failed.")
			return &DbRepo{
				DbConn: nil,
			}, cancel
		}
		tx = DbConnect.DbConn.Session(&gorm.Session{Context: cancelCtx})

	} else {
		tx = d.DbConn.Session(&gorm.Session{Context: cancelCtx})
	}
	return &DbRepo{
		DbConn: tx,
	}, cancel
}

func ConnectDB() int {
	ip := MysqlIP
	port := MysqlPort
	// 产品形态是CeaStor的时候采用mariadb
	if strings.EqualFold(ProductName, "CeaStor") {
		ip = MariadbIP
		port = MariadbPort
	}

	path := strings.Join([]string{UserName, ":", DBPassword, "@(", ip, ":", port, ")/", DbName, "?charset=utf8&parseTime=True&loc=Local&timeout=10s&readTimeout=10s&writeTimeout=10s"}, "")
	var db *gorm.DB
	var err error

	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(mysql.Open(path), &gorm.Config{})
		if err == nil {
			break
		} else {
			time.Sleep(time.Second * 3)
			logging.Log.Infof("Database connection failed! Have Retried %d times", i)
		}
	}
	if err != nil {
		logging.Log.Errorf("Database connection failed! err: %v", err.Error())
	}
	sqlDB, err := db.DB()
	if err != nil {
		logging.Log.Errorf("Database connection failed! err: %v", err.Error())
	}
	sqlDB.SetMaxOpenConns(32)           // SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxIdleConns(32)           // SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetConnMaxLifetime(time.Hour) // SetConnMaxLifetime 设置了连接可复用的最大时间。

	//db.SingularTable(false) //如果使用gorm来帮忙创建表时，这里填写false的话gorm会给表添加s后缀，填写true则不会
	//db.LogMode(true)        //打印sql语句

	DbConnect.DbConn = db
	return 0
}

func GetConnect() (*DbRepo, error) {
	if DbConnect == (DbRepo{}) {
		logging.Log.Infof("database connection does not exist, start connecting...")
		if ret := ConnectDB(); ret != 0 {
			logging.Log.Errorf("database connection failed.")
			return &DbRepo{
				DbConn: nil,
			}, fmt.Errorf("database connection failed.\n")
		}
	}
	return &DbConnect, nil
}

// CreateTableAuto 也会判断表是否存在.但AutoMigrate会在表已经存在的情况下进行智能迁移。
// 如果没有这个字段，将会创建这个字段，并且自动设置模型中定义的外键和索引。
// 如果已经有这个字段了，也会进行智能迁移。会比对数据类型，大小，精度，是非为null，唯一性，默认值，注释，然后确保数据库中的列和gorm模型中定义的列保持一致。
func (d *DbRepo) CreateTableAuto(table interface{}) int {
	var err error
	for i := 0; i < maxRetries; i++ {
		if err = d.DbConn.Migrator().AutoMigrate(table); err == nil {
			break
		}
	}
	if err != nil {
		logging.Log.Infof("Table creation failed, %v", err)
		return 1
	}
	logging.Log.Infof("Table creation succeeded.")
	return 0
}

// CreateTable 0正常，1异常。 hastable用来判断表是否存在，当表不存在时才创建表，如果表存在，报错。
func (d *DbRepo) CreateTable(table interface{}) int {
	for i := 0; i < maxRetries; i++ {
		if !d.DbConn.Migrator().HasTable(table) {
			if err := d.DbConn.Migrator().CreateTable(table); err != nil {
				logging.Log.Errorf("CreateTable: Table creation failed, %v", err)
				return 1
			}
		} else {
			logging.Log.Info("CreateTable: Table already exists.")
			return 0
		}
	}
	return 0
}
