package dcs

import (
	"MyCodeArchive_Go/utils/fault"
	"MyCodeArchive_Go/utils/logging"
	"MyCodeArchive_Go/utils/strings_"
	"MyCodeArchive_Go/utils/tool"
	"MyCodeArchive_Go/utils/tool/db"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func ParamWrap(mapData map[string]interface{}, funcName string) (interface{}, *fault.Fault) {
	apiParam := mapData
	byteParam, err := json.Marshal(apiParam)
	if err != nil {
		logging.Log.Error(err)
		return nil, fault.ParseData
	}
	anyParams, err := GetParam(byteParam, funcName)
	if err != nil {
		return nil, fault.ParseData
	}
	return anyParams, nil
}

func GetParam(args []byte, funcName string) (interface{}, error) {
	params, err := getParam(args, funcName)
	if err != nil {
		logging.Log.Errorf("fail to parse payload: %v %s", err, args)
		return nil, err
	}
	logging.Log.Infof("In progress, params: %+v", params)
	return params, nil
}

func getParam(args []byte, funcName string) (interface{}, error) {
	switch funcName {
	case CreateRelationFun:
		params := RelationParam{}
		err := json.Unmarshal(args, &params)
		return params, err
	case UpdateRelationFun:
		params := RelationUpdateParam{}
		err := json.Unmarshal(args, &params)
		return params, err
	case DeleteRelationFun:
		fallthrough
	case ShowRelationFun:
		params := BaseParam{}
		err := json.Unmarshal(args, &params)
		return params, err
	case ListRelationsFun:
		fallthrough
	case ListRelationsStrategiesFun:
		fallthrough
	case ListStrategiesFun:
		params := FilterParamApi{}
		params.PageSize, params.PageNumber = "-1", "-1"
		err := json.Unmarshal(args, &params)
		if err != nil {
			return FilterParam{}, err
		}
		listParams := FilterParam{
			OutJson: params.OutJson,
			SearchParam: SearchParam{
				SortBy:      params.SortBy,
				Order:       params.Order,
				FilterBy:    params.FilterBy,
				FilterValue: params.FilterValue,
			},
		}
		listParams.PageNumber, listParams.PageSize, err = strings_.ConvertPageToInt(params.PageNumber, params.PageSize)
		return listParams, err
	case CreateStrategyFun:
		params := StrategyParam{}
		err := json.Unmarshal(args, &params)
		return params, err
	case UpdateStrategyFun:
		params := StrategyUpdateParam{}
		err := json.Unmarshal(args, &params)
		return params, err
	case DeleteStrategyFun:
		fallthrough
	case ShowStrategyFun:
		params := BaseParam{}
		err := json.Unmarshal(args, &params)
		return params, err
	default:
		return nil, errors.New(fmt.Sprintf("invalid funcName %s", funcName))
	}
}

func checkNameParam(name string) *fault.Fault {
	return tool.CheckNameParam(name)
}

func checkExisted(name, uuid, moduleName string, expectIsExisted bool) *fault.Fault {
	var isExisted bool
	var err *fault.Fault
	switch moduleName {
	case RelationModule:
		isExisted, err = IsExisted(name, uuid, DcsRelationsName())
	case StrategyModule:
		isExisted, err = IsExisted(name, uuid, DcsStrategiesName())
	default:
		return fault.Wrap(
			fmt.Sprintf("not found the module name %q", moduleName), fmt.Sprintf("没有找到这个模块%q", moduleName))
	}
	if err != nil {
		return err
	}

	logging.Log.Infof("name %s,uuid %s,  expectIsExisted is %t, isExisted is %t", name, uuid, expectIsExisted, isExisted)
	if expectIsExisted && !isExisted {
		return fault.NotExist(name, moduleName)
	} else if !expectIsExisted && isExisted {
		return fault.Existed(name, moduleName)
	}
	return nil
}

func IsExisted(name, uuid, tableName string) (bool, *fault.Fault) {
	dbCon := db.DbConnect.GetConnect()
	if dbCon.DbConn == nil {
		logging.Log.Error("fail to get db connect")
		return false, fault.ConnectDB
	}

	queryBuilder := dbCon.DbConn.Table(tableName)
	if name != "" {
		queryBuilder = queryBuilder.Where("BINARY name = ?", name)
	}
	if uuid != "" {
		queryBuilder = queryBuilder.Where("uuid = ?", uuid)
	}
	if name == "" && uuid == "" {
		return false, fault.Err(fmt.Sprintf("fail to query %s:%s", name, uuid), errors.New("uuid and name is empty"), fault.QueryRecord)
	}

	var count int64
	ret := queryBuilder.Select("COUNT(*)").Count(&count)
	if ret.Error != nil {
		return false, fault.Err(fmt.Sprintf("fail to query %s:%s", name, uuid), ret.Error, fault.QueryRecord)
	}
	if count > 0 {
		logging.Log.Infof("found relation %s:%s, %d", name, uuid, count)
		return true, nil
	}
	logging.Log.Infof("not found relation %s:%s, affected rows %d", name, uuid, count)
	return false, nil
}

func checkIDAndNameParam(uuid, name string) *fault.Fault {
	if len(uuid) == 0 && len(name) == 0 {
		return fault.Wrap("ID and Name is empty", "ID和Name参数都是空的")
	}
	if len(name) != 0 {
		return checkNameParam(name)
	}
	return nil
}

func CheckSearchParam(params SearchParam) *fault.Fault {
	err := CheckOrderParam(params.Order, params.SortBy)
	if err != nil {
		return err
	}

	if params.PageNumber < -1 || params.PageNumber == 0 {
		return fault.InvalidParam("pageNumber", strconv.Itoa(params.PageNumber))
	}
	if params.PageSize < -1 || params.PageSize == 0 {
		return fault.InvalidParam("pageSize", strconv.Itoa(params.PageSize))
	} else if params.PageSize > PageSizeMax {
		return fault.ParamCount("pageSize", PageSizeMax)
	}

	if params.FilterBy != "Name" && params.FilterBy != "" {
		return fault.InvalidParam("filterBy", params.FilterBy)
	}
	if params.FilterValue != "" {
		if params.FilterBy == "" {
			return fault.ParamEmpty("filterBy")
		}
		if err = checkNameParam(params.FilterValue); err != nil {
			return err
		}
	}
	return nil
}

func CheckOrderParam(order, sort string) *fault.Fault {
	if strings.ToUpper(order) != OrderDesc && strings.ToUpper(order) != OrderAsc && order != "" {
		return fault.InvalidParam("order", order)
	}
	if len(order) != 0 && len(sort) == 0 {
		return fault.ParamEmpty("sort")
	} else if len(sort) != 0 && len(order) == 0 {
		return fault.ParamEmpty("order")
	}
	return nil
}
