package db

import (
	"MyCodeArchive_Go/utils/fault"
	"MyCodeArchive_Go/utils/logging"
	"MyCodeArchive_Go/utils/math_"
	"database/sql"
	"fmt"
	"time"
)

const BatchSizeMax = 100

type ExampleDb struct {
	Id         uint
	Name       string
	Phone      string
	CreateTime time.Time `gorm:"autoCreateTime"` // 创建的时候自动填充时间戳
	UpdateTime time.Time `gorm:"autoUpdateTime"` // 更新的时候自动填充时间戳
}

// QueryByNameMut Name不唯一的情况
func (e *ExampleDb) QueryByNameMut(name string) ([]ExampleDb, *fault.Fault) {
	dbCon, sessionClose := DbConnect.GetSession(&SessionConfig{})
	if dbCon == nil || dbCon.DbConn == nil {
		logging.Log.Errorf("failed to get session from database")
		return []ExampleDb{}, fault.ConnectDB
	}
	defer sessionClose()

	var ret []ExampleDb
	if err := dbCon.DbConn.Model(&ExampleDb{}).Where("BINARY name = ?", name).Find(&ret).Error; err != nil {
		return []ExampleDb{}, fault.Err("failed to query example by name", err, fault.QueryRecord)
	}

	if ret == nil {
		return []ExampleDb{}, nil
	}
	return ret, nil
}

// QueryByNameUni Name唯一的情况
func (e *ExampleDb) QueryByNameUni(name string) *fault.Fault {
	dbCon, sessionClose := DbConnect.GetSession(&SessionConfig{})
	if dbCon == nil || dbCon.DbConn == nil {
		logging.Log.Errorf("failed to get session from database")
		return fault.ConnectDB
	}
	defer sessionClose()

	out := dbCon.DbConn.Model(&ExampleDb{}).Where("BINARY name = ?", name).Find(e)
	if out.Error != nil {
		return fault.Err(fmt.Sprintf("fail to query export %s", name), out.Error, fault.QueryRecord)
	}
	if out.RowsAffected == 0 {
		return fault.NotExist("name", name)
	}
	return nil
}

func (e *ExampleDb) QueryFirst(name string) *fault.Fault {
	dbCon, sessionClose := DbConnect.GetSession(&SessionConfig{})
	if dbCon == nil || dbCon.DbConn == nil {
		logging.Log.Errorf("failed to get session from database")
		return fault.ConnectDB
	}
	defer sessionClose()

	// first如果没查到，会直接报错
	if err := dbCon.DbConn.Model(&ExampleDb{}).Where("BINARY name = ?", name).First(e).Error; err != nil {
		return fault.Err("failed to query example by name", err, fault.QueryRecord)
	}

	return nil
}

func (e *ExampleDb) QueryCount(value *int64) *fault.Fault {
	dbCon, sessionClose := DbConnect.GetSession(&SessionConfig{})
	if dbCon == nil || dbCon.DbConn == nil {
		logging.Log.Errorf("fail to get db session")
		return fault.ConnectDB
	}
	defer sessionClose()

	if err := dbCon.DbConn.Model(&ExampleDb{}).Count(value).Error; err != nil {
		return fault.Err("Failed to query nfs export count", err, fault.QueryRecord)
	}

	return nil
}

func (e *ExampleDb) QueryCountByName(value *int64, name string) *fault.Fault {
	dbCon, sessionClose := DbConnect.GetSession(&SessionConfig{})
	if dbCon == nil || dbCon.DbConn == nil {
		logging.Log.Errorf("fail to get db session")
		return fault.ConnectDB
	}
	defer sessionClose()

	if err := dbCon.DbConn.Model(&ExampleDb{}).Where("BINARY name = ?", name).Count(value).Error; err != nil {
		return fault.Err("Failed to query name count", err, fault.QueryRecord)
	}

	return nil
}

// QueryFirstName 获取到第一行的Name列的值
func (e *ExampleDb) QueryFirstName() *fault.Fault {
	dbCon, sessionClose := DbConnect.GetSession(&SessionConfig{})
	if dbCon == nil || dbCon.DbConn == nil {
		logging.Log.Errorf("fail to get db session")
		return fault.ConnectDB
	}
	defer sessionClose()

	dbVal := struct {
		Name string
	}{}
	// 必须要用结构体才能获取到这个值
	ret := dbCon.DbConn.Model(&ExampleDb{}).Select("name").First(&dbVal)
	if ret.Error != nil {
		return fault.Err("fail to query csc cluster config first record", ret.Error, fault.QueryRecord)
	}

	if ret.RowsAffected == 0 {
		return fault.QueryColEmpty("name")
	}

	if dbVal.Name != "张三" {
		return fault.Wrap("The current protocol is not NFS", "当前不是NFS协议")
	}
	return nil
}

func (e *ExampleDb) QueryByIds(ids []string) (list []ExampleDb, err *fault.Fault) {
	dbCon, sessionClose := DbConnect.GetSession(&SessionConfig{})
	if dbCon == nil || dbCon.DbConn == nil {
		logging.Log.Errorf("fail to get db session")
		return []ExampleDb{}, fault.ConnectDB
	}
	defer sessionClose()

	ret := dbCon.DbConn.Model(&ExampleDb{}).Find(&list, ids)
	if ret.Error != nil {
		return []ExampleDb{}, fault.Err("fail to query list", ret.Error, fault.QueryRecord)
	}
	if list == nil {
		return []ExampleDb{}, nil
	}
	return list, nil
}

func (e *ExampleDb) Create() *fault.Fault {
	dbCon, sessionClose := DbConnect.GetSession(&SessionConfig{})
	if dbCon == nil || dbCon.DbConn == nil {
		logging.Log.Errorf("fail to get db session")
		return fault.ConnectDB
	}
	defer sessionClose()

	ret := dbCon.DbConn.Create(e)
	if ret.Error != nil {
		return fault.Err(fmt.Sprintf("failed to insert %s into database", e.Name), ret.Error, fault.InsertRecord)
	}
	return nil
}

// BatchCreate 批量创建
func (e *ExampleDb) BatchCreate(clients []ExampleDb) *fault.Fault {
	dbCon, sessionClose := DbConnect.GetSession(&SessionConfig{})
	if dbCon == nil || dbCon.DbConn == nil {
		logging.Log.Errorf("failed to get db session")
		return fault.ConnectDB
	}
	defer sessionClose()

	// 批量创建记录，会分成BatchSize大小按批创建
	err := dbCon.DbConn.Model(&ExampleDb{}).CreateInBatches(clients, BatchSizeMax).Error
	if err != nil {
		return fault.Err("failed to perform a batch insertion of clients into the database", err, fault.InsertRecord)
	}
	return nil
}

func (e *ExampleDb) Delete() *fault.Fault {
	dbCon, sessionClose := DbConnect.GetSession(&SessionConfig{})
	if dbCon == nil || dbCon.DbConn == nil {
		logging.Log.Errorf("fail to get db session")
		return fault.ConnectDB
	}
	defer sessionClose()

	err := dbCon.DbConn.Where("BINARY name = ?", e.Name).Delete(&ExampleDb{}).Error
	if err != nil {
		return fault.Err(fmt.Sprintf("fail to delete policy %s", e.Name), err, fault.DeleteRecord)
	}
	return nil
}

func (e *ExampleDb) DeleteByName(name string) *fault.Fault {
	dbCon, sessionClose := DbConnect.GetSession(&SessionConfig{})
	if dbCon == nil || dbCon.DbConn == nil {
		logging.Log.Errorf("fail to get db session")
		return fault.ConnectDB
	}
	defer sessionClose()

	err := dbCon.DbConn.Where("BINARY name = ?", name).Delete(&ExampleDb{}).Error
	if err != nil {
		return fault.Err("fail to delete record from NfsExportPolicies", err, fault.DeleteRecord)
	}
	return nil
}

// DeleteAllDataByTx 通过事务完成这件事
func (e *ExampleDb) DeleteAllDataByTx() *fault.Fault {
	dbCon, sessionClose := DbConnect.GetSession(&SessionConfig{})
	if dbCon == nil || dbCon.DbConn == nil {
		logging.Log.Errorf("fail to get db session")
		return fault.ConnectDB
	}
	defer sessionClose()

	tx := dbCon.DbConn.Begin()
	defer func() {
		if r := recover(); r != nil {
			logging.Log.Errorf("roll back db operations")
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return fault.ConnectDB
	}

	err := tx.Delete(&ExampleDb{}).Error
	if err != nil {
		tx.Rollback()
		return fault.Err("fail to delete record", err, fault.DeleteRecord)
	}

	if err = tx.Commit().Error; err != nil {
		return fault.Err("fail to delete record from nfs", err, fault.DeleteRecord)
	}
	return nil
}

func (e *ExampleDb) List(filterBy, filterValue, order, sortBy string, pageSize, pageNumber int) (list []ExampleDb, total int64, err *fault.Fault) {
	dbCon, sessionClose := DbConnect.GetSession(&SessionConfig{})
	if dbCon == nil || dbCon.DbConn == nil {
		logging.Log.Errorf("fail to get db session")
		return []ExampleDb{}, -1, fault.ConnectDB
	}
	defer sessionClose()

	queryBuilder := dbCon.DbConn.Model(&ExampleDb{})
	if filterBy != "" && filterValue != "" {
		queryBuilder = queryBuilder.Where(fmt.Sprintf("%s = ?", filterBy), filterValue).Count(&total)
	} else {
		queryBuilder = queryBuilder.Count(&total)
	}

	if order != "" && sortBy != "" {
		queryBuilder = queryBuilder.Order(fmt.Sprintf("%s %s", sortBy, order))
	}

	if pageSize != -1 && pageNumber != -1 {
		queryBuilder = queryBuilder.Offset(math_.CalculateOffset(pageSize, pageNumber)).Limit(pageSize)
	}

	if ret := queryBuilder.Find(&list); ret.Error != nil {
		return []ExampleDb{}, -1, fault.Err("fail to query list", ret.Error, fault.QueryRecord)
	}

	if list == nil {
		return []ExampleDb{}, 0, nil
	}
	return list, total, nil
}

// Update update格式：map[string]interface{}{"name" : "12345"}
func (e *ExampleDb) Update(update map[string]interface{}, name string) *fault.Fault {
	dbCon, sessionClose := DbConnect.GetSession(&SessionConfig{})
	if dbCon == nil || dbCon.DbConn == nil {
		logging.Log.Errorf("failed to get db session")
		return fault.ConnectDB
	}
	defer sessionClose()

	ret := dbCon.DbConn.Model(&ExampleDb{}).Where("BINARY name = ?", name).Updates(update)
	if ret.Error != nil {
		return fault.Err("fail to update policy", ret.Error, fault.UpdateRecord)
	}
	return nil
}

func (e *ExampleDb) IsExist(name string) (bool, *fault.Fault) {
	dbCon, sessionClose := DbConnect.GetSession(&SessionConfig{})
	if dbCon == nil || dbCon.DbConn == nil {
		logging.Log.Errorf("fail to get db session")
		return false, fault.ConnectDB
	}
	defer sessionClose()

	var exports []ExampleDb
	ret := dbCon.DbConn.Model(&ExampleDb{}).Where("BINARY name = ?", name).Find(e)
	if ret.Error != nil {
		return false, fault.Err(fmt.Sprintf("fail to query %s", name), ret.Error, fault.QueryRecord)
	}
	if ret.RowsAffected == 1 {
		logging.Log.Infof("found example %s, %d, %+v", name, ret.RowsAffected, exports)
		return true, nil
	} else {
		logging.Log.Infof("not found example %s, affected rows %d", name, ret.RowsAffected)
		return false, nil
	}
}

func (e *ExampleDb) IsExistedByName(name string) (bool, *fault.Fault) {
	dbCon, sessionClose := DbConnect.GetSession(&SessionConfig{})
	if dbCon == nil || dbCon.DbConn == nil {
		logging.Log.Errorf("fail to get db session")
		return false, fault.ConnectDB
	}
	defer sessionClose()

	logging.Log.Infof("Start checking if this client has been created. name: %s", name)
	out := dbCon.DbConn.Where("BINARY name = ?", name).Find(e)
	if out.Error != nil {
		logging.Log.Errorf("fail to query %s , %v",
			name, out.Error)
		return false, fault.QueryRecord
	}
	if out.RowsAffected == 0 {
		logging.Log.Infof("not found %s", name)
		return false, nil
	} else {
		logging.Log.Errorf("found client %s, affected rows %d", name, out.RowsAffected)
		return true, nil
	}
}

func ListRpcIp() ([]string, *fault.Fault) {
	dbCon, sessionClose := DbConnect.GetSession(&SessionConfig{})
	if dbCon == nil || dbCon.DbConn == nil {
		logging.Log.Errorf("fail to get db session")
		return []string{}, fault.ConnectDB
	}
	defer sessionClose()

	return []string{"10.255.0.1"}, nil
}

// recordStats 关于defer中rollback的用法。 TODO 如果发生panic，还能调用rollback函数吗？
func recordStats(db *sql.DB, userID, productID int64) (err error) {
	// 开启事务
	// 操作views和product_viewers两张表
	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	// 更新products表
	if _, err = tx.Exec("UPDATE products SET views = views + 1"); err != nil {
		return
	}
	// product_viewers表中插入一条数据
	if _, err = tx.Exec(
		"INSERT INTO product_viewers (user_id, product_id) VALUES (?, ?)",
		userID, productID); err != nil {
		return
	}
	return
}
