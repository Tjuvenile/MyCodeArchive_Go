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

type BgrRepLinkMgt struct {
	UUID               string `json:"UUID"`
	Name               string `json:"Name"`
	LocalNodePoolId    int    `json:"LocalNodePoolId"`
	RemoteNodePoolId   int    `json:"RemoteNodePoolId"`
	RemoteDeviceName   string `json:"RemoteDeviceName"`
	RemoteNodePoolName string `json:"RemoteNodePoolName"`
	RemoteCtrlIp       string `json:"RemoteCtrlIp"`
	Status             string `json:"Status"`
}

type BgrRelations struct {
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

type BgrStrategies struct {
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

func (re *BgrRelations) Create() *fault.Fault {
	dbCon := db.DbConnect.GetConnect()
	if dbCon.DbConn == nil {
		logging.Log.Error("fail to get db connect")
		return fault.ConnectDB
	}

	ret := dbCon.DbConn.Model(&BgrRelations{}).Create(re)
	if ret.Error != nil {
		return fault.Err(fmt.Sprintf("failed to create %s:%s into database", re.Name, re.UUID), ret.Error, fault.InsertRecord)
	}
	return nil
}

func (re *BgrRelations) Update(update map[string]interface{}) *fault.Fault {
	dbCon := db.DbConnect.GetConnect()
	if dbCon.DbConn == nil {
		logging.Log.Error("fail to get db connect")
		return fault.ConnectDB
	}

	var err error
	if len(re.UUID) != 0 {
		err = dbCon.DbConn.Model(&BgrRelations{}).Where("uuid = ?", re.UUID).Updates(update).Error
	} else if len(re.Name) != 0 {
		err = dbCon.DbConn.Model(&BgrRelations{}).Where("BINARY name = ?", re.Name).Updates(update).Error
	} else {
		err = errors.New("id and name is empty")
	}
	if err != nil {
		return fault.Err(fmt.Sprintf("fail to update relation %s:%s", re.Name, re.UUID), err, fault.UpdateRecord)
	}
	return nil
}

func (re *BgrRelations) Delete() *fault.Fault {
	dbCon := db.DbConnect.GetConnect()
	if dbCon.DbConn == nil {
		logging.Log.Error("fail to get db connect")
		return fault.ConnectDB
	}

	var err error
	if len(re.UUID) != 0 {
		err = dbCon.DbConn.Where("uuid = ?", re.UUID).Delete(&BgrRelations{}).Error
	} else if len(re.Name) != 0 {
		err = dbCon.DbConn.Where("BINARY name = ?", re.Name).Delete(&BgrRelations{}).Error
	} else {
		err = errors.New("id and name is empty")
	}
	if err != nil {
		return fault.Err(fmt.Sprintf("fail to delete relation %s:%s", re.Name, re.UUID), err, fault.DeleteRecord)
	}
	return nil
}

func (re *BgrRelations) QueryById() *fault.Fault {
	dbCon := db.DbConnect.GetConnect()
	if dbCon.DbConn == nil {
		logging.Log.Error("fail to get db connect")
		return fault.ConnectDB
	}

	out := dbCon.DbConn.Model(&BgrRelations{}).Where("uuid = ?", re.UUID).Find(re)
	if out.Error != nil {
		return fault.Err(fmt.Sprintf("fail to query %s", re.UUID), out.Error, fault.QueryRecord)
	}
	if out.RowsAffected == 0 {
		return fault.NotExist("relation uuid", re.UUID)
	}
	return nil
}

func (re *BgrRelations) QueryByName() *fault.Fault {
	dbCon := db.DbConnect.GetConnect()
	if dbCon.DbConn == nil {
		logging.Log.Error("fail to get db connect")
		return fault.ConnectDB
	}

	out := dbCon.DbConn.Model(&BgrRelations{}).Where("BINARY name = ?", re.Name).Find(re)
	if out.Error != nil {
		return fault.Err(fmt.Sprintf("fail to query %s", re.Name), out.Error, fault.QueryRecord)
	}
	if out.RowsAffected == 0 {
		return fault.NotExist("relation name", re.Name)
	}
	return nil
}

func (re *BgrRelations) List(filterBy, filterValue, order, sortBy string, pageSize, pageNumber int) (list []BgrRelations, total int64, err *fault.Fault) {
	dbCon := db.DbConnect.GetConnect()
	if dbCon.DbConn == nil {
		logging.Log.Error("fail to get db connect")
		return []BgrRelations{}, 0, fault.ConnectDB
	}

	queryBuilder := dbCon.DbConn.Model(&BgrRelations{})
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
		return []BgrRelations{}, -1, fault.Err("fail to query list", ret.Error, fault.QueryRecord)
	}

	if list == nil {
		return []BgrRelations{}, 0, nil
	}
	return list, total, nil
}

func (st *BgrStrategies) QueryByIds(ids []string) (list []BgrStrategies, err *fault.Fault) {
	dbCon := db.DbConnect.GetConnect()
	if dbCon.DbConn == nil {
		logging.Log.Error("fail to get db connect")
		return []BgrStrategies{}, fault.ConnectDB
	}

	ret := dbCon.DbConn.Model(&BgrStrategies{}).Find(&list, ids)
	if ret.Error != nil {
		return []BgrStrategies{}, fault.Err("fail to query list", ret.Error, fault.QueryRecord)
	}
	if list == nil {
		return []BgrStrategies{}, nil
	}
	return list, nil
}

func (st *BgrStrategies) Create() *fault.Fault {
	dbCon := db.DbConnect.GetConnect()
	if dbCon.DbConn == nil {
		return fault.ConnectDB
	}

	ret := dbCon.DbConn.Model(&BgrStrategies{}).Create(st)
	if ret.Error != nil {
		return fault.Err(fmt.Sprintf("failed to create %s:%s into database", st.Name, st.UUID), ret.Error, fault.InsertRecord)
	}
	return nil
}

func (st *BgrStrategies) Delete() *fault.Fault {
	dbCon := db.DbConnect.GetConnect()
	if dbCon.DbConn == nil {
		return fault.ConnectDB
	}

	var err error
	if len(st.UUID) != 0 {
		err = dbCon.DbConn.Where("uuid = ?", st.UUID).Delete(&BgrStrategies{}).Error
	} else if len(st.Name) != 0 {
		err = dbCon.DbConn.Where("BINARY name = ?", st.Name).Delete(&BgrStrategies{}).Error
	} else {
		err = errors.New("id and name is empty")
	}
	if err != nil {
		return fault.Err(fmt.Sprintf("fail to delete relation %s:%s", st.Name, st.UUID), err, fault.DeleteRecord)
	}
	return nil
}

func (st *BgrStrategies) List(filterBy, filterValue, order, sortBy string, pageSize, pageNumber int) (list []BgrStrategies, total int64, err *fault.Fault) {
	dbCon := db.DbConnect.GetConnect()
	if dbCon.DbConn == nil {
		return []BgrStrategies{}, 0, fault.ConnectDB
	}

	queryBuilder := dbCon.DbConn.Model(&BgrStrategies{})
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
		return []BgrStrategies{}, -1, fault.Err("fail to query list", ret.Error, fault.QueryRecord)
	}

	if list == nil {
		return []BgrStrategies{}, 0, nil
	}
	return list, total, nil
}

func (st *BgrStrategies) Update(update map[string]interface{}) *fault.Fault {
	dbCon := db.DbConnect.GetConnect()
	if dbCon.DbConn == nil {
		return fault.ConnectDB
	}

	var err error
	if len(st.UUID) != 0 {
		err = dbCon.DbConn.Model(&BgrStrategies{}).Where("uuid = ?", st.UUID).Updates(update).Error
	} else if len(st.Name) != 0 {
		err = dbCon.DbConn.Model(&BgrStrategies{}).Where("BINARY name = ?", st.Name).Updates(update).Error
	} else {
		err = errors.New("id and name is empty")
	}
	if err != nil {
		return fault.Err(fmt.Sprintf("fail to update relation %s:%s", st.Name, st.UUID), err, fault.UpdateRecord)
	}
	return nil
}

type RepLinkMgt struct {
	uuid               string // 复制链路的uuid
	Name               string // 复制链路名称
	LocalNodePoolId    int    // 本端节点池id
	RemoteNodePoolId   int    // 远端节点池id
	RemoteDeviceName   string // 远端设备名称
	RemoteNodePoolName string // 远端节点池名称
	RemoteCtrlIp       string // 远端控制节点复制网ip
	Status             string
}
