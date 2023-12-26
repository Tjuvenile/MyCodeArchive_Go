package dcs

import (
	"MyCodeArchive_Go/utils/fault"
	"MyCodeArchive_Go/utils/logging"
	"MyCodeArchive_Go/utils/math_"
	"MyCodeArchive_Go/utils/tool/db"
	"errors"
	"fmt"
	"time"
)

func CreateTable() {
	dbCon := db.DbConnect.GetConnect()
	if dbCon.DbConn == nil {
		logging.Log.Error("fail to get db connect")
		return
	}

	if dbCon.CreateTableAuto(&DcsRelations{}) != 0 {
		logging.Log.Error("failed to create table DcsRelations")
		return
	}

	if dbCon.CreateTableAuto(&DcsStrategies{}) != 0 {
		logging.Log.Error("failed to create table DcsStrategies")
		return
	}
	logging.Log.Infof("end to create DCS table")
}

type DcsRelations struct {
	UUID             string    `json:"UUID" gorm:"primary_key"`
	Name             string    `json:"Name"`
	MasterPool       string    `json:"MasterPool"`
	SlavePool        string    `json:"SlavePool"`
	MasterResourceId string    `json:"MasterResourceId"` // 主端资源
	SlaveResourceId  string    `json:"SlaveResourceId"`  // 从端资源
	ResourceType     string    `json:"ResourceType"`     // 资源类型：块/文件/对象
	StrategyIds      string    `json:"StrategyIds"`      // 策略id列表
	LastSyncTime     int64     `json:"LastSyncTime"`     // 上次同步时间
	LastSyncSnap     string    `json:"LastSyncSnap"`     // 上次同步快照
	State            string    `json:"State"`            // 中间状态
	RunningState     string    `json:"RunningState"`     // 运行状态
	HealthState      string    `json:"HealthState"`      // 健康状态
	DataState        string    `json:"DataState"`        // 数据状态
	Role             string    `json:"Role"`             // 记录本地角色（主/从）
	IsConfigSync     bool      `json:"IsConfigSync"`     // 是否开启资源配置同步
	CreateTime       time.Time `json:"CreateTime" gorm:"autoCreateTime"`
	UpdateTime       time.Time `json:"UpdateTime" gorm:"autoUpdateTime"`
}

type DcsStrategies struct {
	UUID         string    `json:"UUID" gorm:"primary_key"`
	Name         string    `json:"Name"`
	TimePoint    string    `json:"TimePoint"`    // 定时同步时间
	Interval     string    `json:"Interval"`     // 同步间隔
	StrategyType string    `json:"StrategyType"` // 定时策略类型
	Description  string    `json:"Description"`
	CreateTime   time.Time `json:"CreateTime" gorm:"autoCreateTime"`
}

func DcsRelationsName() string {
	return "dcs_relations"
}

func DcsStrategiesName() string {
	return "dcs_strategies"
}

func (re *DcsRelations) Create() *fault.Fault {
	dbCon := db.DbConnect.GetConnect()
	if dbCon.DbConn == nil {
		logging.Log.Error("fail to get db connect")
		return fault.ConnectDB
	}

	ret := dbCon.DbConn.Model(&DcsRelations{}).Create(re)
	if ret.Error != nil {
		return fault.Err(fmt.Sprintf("failed to create %s:%s into database", re.Name, re.UUID), ret.Error, fault.InsertRecord)
	}
	return nil
}

func (re *DcsRelations) Update(update map[string]interface{}) *fault.Fault {
	dbCon := db.DbConnect.GetConnect()
	if dbCon.DbConn == nil {
		logging.Log.Error("fail to get db connect")
		return fault.ConnectDB
	}

	var err error
	if len(re.UUID) != 0 {
		err = dbCon.DbConn.Model(&DcsRelations{}).Where("uuid = ?", re.UUID).Updates(update).Error
	} else if len(re.Name) != 0 {
		err = dbCon.DbConn.Model(&DcsRelations{}).Where("BINARY name = ?", re.Name).Updates(update).Error
	} else {
		err = errors.New("id and name is empty")
	}
	if err != nil {
		return fault.Err(fmt.Sprintf("fail to update relation %s:%s", re.Name, re.UUID), err, fault.UpdateRecord)
	}
	return nil
}

func (re *DcsRelations) Delete() *fault.Fault {
	dbCon := db.DbConnect.GetConnect()
	if dbCon.DbConn == nil {
		logging.Log.Error("fail to get db connect")
		return fault.ConnectDB
	}

	var err error
	if len(re.UUID) != 0 {
		err = dbCon.DbConn.Where("uuid = ?", re.UUID).Delete(&DcsRelations{}).Error
	} else if len(re.Name) != 0 {
		err = dbCon.DbConn.Where("BINARY name = ?", re.Name).Delete(&DcsRelations{}).Error
	} else {
		err = errors.New("id and name is empty")
	}
	if err != nil {
		return fault.Err(fmt.Sprintf("fail to delete relation %s:%s", re.Name, re.UUID), err, fault.DeleteRecord)
	}
	return nil
}

func (re *DcsRelations) QueryById() *fault.Fault {
	dbCon := db.DbConnect.GetConnect()
	if dbCon.DbConn == nil {
		logging.Log.Error("fail to get db connect")
		return fault.ConnectDB
	}

	out := dbCon.DbConn.Model(&DcsRelations{}).Where("uuid = ?", re.UUID).Find(re)
	if out.Error != nil {
		return fault.Err(fmt.Sprintf("fail to query %s", re.UUID), out.Error, fault.QueryRecord)
	}
	if out.RowsAffected == 0 {
		return fault.NotExist("relation uuid", re.UUID)
	}
	return nil
}

func (re *DcsRelations) QueryByName() *fault.Fault {
	dbCon := db.DbConnect.GetConnect()
	if dbCon.DbConn == nil {
		logging.Log.Error("fail to get db connect")
		return fault.ConnectDB
	}

	out := dbCon.DbConn.Model(&DcsRelations{}).Where("BINARY name = ?", re.Name).Find(re)
	if out.Error != nil {
		return fault.Err(fmt.Sprintf("fail to query %s", re.Name), out.Error, fault.QueryRecord)
	}
	if out.RowsAffected == 0 {
		return fault.NotExist("relation name", re.Name)
	}
	return nil
}

func (re *DcsRelations) List(filterBy, filterValue, order, sortBy string, pageSize, pageNumber int) (list []DcsRelations, total int64, err *fault.Fault) {
	dbCon := db.DbConnect.GetConnect()
	if dbCon.DbConn == nil {
		logging.Log.Error("fail to get db connect")
		return []DcsRelations{}, 0, fault.ConnectDB
	}

	queryBuilder := dbCon.DbConn.Model(&DcsRelations{})
	if filterBy != "" && filterValue != "" {
		if filterBy == "Name" {
			queryBuilder = queryBuilder.Where(fmt.Sprintf("BINARY %s = ?", filterBy), filterValue).Count(&total)
		} else {
			queryBuilder = queryBuilder.Where(fmt.Sprintf("%s = ?", filterBy), filterValue).Count(&total)
		}
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
		return []DcsRelations{}, -1, fault.Err("fail to query list", ret.Error, fault.QueryRecord)
	}

	if list == nil {
		return []DcsRelations{}, 0, nil
	}
	return list, total, nil
}

func (st *DcsStrategies) QueryByIds(ids []string) (list []DcsStrategies, err *fault.Fault) {
	dbCon := db.DbConnect.GetConnect()
	if dbCon.DbConn == nil {
		logging.Log.Error("fail to get db connect")
		return []DcsStrategies{}, fault.ConnectDB
	}

	ret := dbCon.DbConn.Model(&DcsStrategies{}).Find(&list, ids)
	if ret.Error != nil {
		return []DcsStrategies{}, fault.Err("fail to query list", ret.Error, fault.QueryRecord)
	}
	if list == nil {
		return []DcsStrategies{}, nil
	}
	return list, nil
}

func (st *DcsStrategies) Create() *fault.Fault {
	dbCon := db.DbConnect.GetConnect()
	if dbCon.DbConn == nil {
		return fault.ConnectDB
	}

	ret := dbCon.DbConn.Model(&DcsStrategies{}).Create(st)
	if ret.Error != nil {
		return fault.Err(fmt.Sprintf("failed to create %s:%s into database", st.Name, st.UUID), ret.Error, fault.InsertRecord)
	}
	return nil
}

func (st *DcsStrategies) Delete() *fault.Fault {
	dbCon := db.DbConnect.GetConnect()
	if dbCon.DbConn == nil {
		return fault.ConnectDB
	}

	var err error
	if len(st.UUID) != 0 {
		err = dbCon.DbConn.Where("uuid = ?", st.UUID).Delete(&DcsStrategies{}).Error
	} else if len(st.Name) != 0 {
		err = dbCon.DbConn.Where("BINARY name = ?", st.Name).Delete(&DcsStrategies{}).Error
	} else {
		err = errors.New("id and name is empty")
	}
	if err != nil {
		return fault.Err(fmt.Sprintf("fail to delete relation %s:%s", st.Name, st.UUID), err, fault.DeleteRecord)
	}
	return nil
}

func (st *DcsStrategies) List(filterBy, filterValue, order, sortBy string, pageSize, pageNumber int) (list []DcsStrategies, total int64, err *fault.Fault) {
	dbCon := db.DbConnect.GetConnect()
	if dbCon.DbConn == nil {
		return []DcsStrategies{}, 0, fault.ConnectDB
	}

	queryBuilder := dbCon.DbConn.Model(&DcsStrategies{})
	if filterBy != "" && filterValue != "" {
		if filterBy == "Name" {
			queryBuilder = queryBuilder.Where(fmt.Sprintf("BINARY %s = ?", filterBy), filterValue).Count(&total)
		} else {
			queryBuilder = queryBuilder.Where(fmt.Sprintf("%s = ?", filterBy), filterValue).Count(&total)
		}
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
		return []DcsStrategies{}, -1, fault.Err("fail to query list", ret.Error, fault.QueryRecord)
	}

	if list == nil {
		return []DcsStrategies{}, 0, nil
	}
	return list, total, nil
}

func (st *DcsStrategies) Update(update map[string]interface{}) *fault.Fault {
	dbCon := db.DbConnect.GetConnect()
	if dbCon.DbConn == nil {
		return fault.ConnectDB
	}

	var err error
	if len(st.UUID) != 0 {
		err = dbCon.DbConn.Model(&DcsStrategies{}).Where("uuid = ?", st.UUID).Updates(update).Error
	} else if len(st.Name) != 0 {
		err = dbCon.DbConn.Model(&DcsStrategies{}).Where("BINARY name = ?", st.Name).Updates(update).Error
	} else {
		err = errors.New("id and name is empty")
	}
	if err != nil {
		return fault.Err(fmt.Sprintf("fail to update relation %s:%s", st.Name, st.UUID), err, fault.UpdateRecord)
	}
	return nil
}
