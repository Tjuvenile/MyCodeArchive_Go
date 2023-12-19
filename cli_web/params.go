package cli_web

/**
关于参数相关的校验和工具方法，放到这里。
比如检查参数的长短，检查参数是否重复等。
*/

import (
	"MyCodeArchive_Go/db"
	"MyCodeArchive_Go/fault"
	"MyCodeArchive_Go/logging"
	mystring "MyCodeArchive_Go/utils/strings"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"regexp"
	"strconv"
	"strings"
)

const (
	ZeroLen       = 0
	NameLenMax    = 128
	PageSizeMax   = 100
	RO            = "RO"
	RW            = "RW"
	OrderDesc     = "DESC"
	OrderAsc      = "ASC"
	ExampleDbName = "ExampleDb"
)

type SearchApiParam struct {
	PageNumber  string `json:"PageNumber"`
	PageSize    string `json:"PageSize"`
	SortBy      string `json:"SortBy"`
	Order       string `json:"Order"`
	FilterBy    string `json:"FilterBy"`
	FilterValue string `json:"FilterValue"`
}

type SearchParam struct {
	PageNumber  int    `json:"PageNumber"`
	PageSize    int    `json:"PageSize"`
	SortBy      string `json:"SortBy"`
	Order       string `json:"Order"`
	FilterBy    string `json:"FilterBy"`
	FilterValue string `json:"FilterValue"`
}

// PerformCommonSetup 通用的获取参数，绑定到对应结构体上
func PerformCommonSetup(args []byte, funcName string) (interface{}, *fault.Fault) {
	params, err := getParam(args, funcName)
	if err != nil {
		logging.Log.Errorf("fail to parse payload: %v %s", err, args)
		return nil, fault.ParseData
	}
	logging.Log.Infof("In progress, params: %+v", params)
	return params, nil
}

// 即便unmarshal失败了，也不会有空指针问题，都会变成类型零值
func getParam(args []byte, funcName string) (interface{}, error) {
	switch funcName {
	case "Create":
		fallthrough
	case "Delete":
		params := struct{}{}
		err := json.Unmarshal(args, &params)
		return params, err
	case "List":
		params := SearchApiParam{}
		params.PageSize, params.PageNumber = "-1", "-1"
		err := json.Unmarshal(args, &params)
		listParams := SearchParam{
			SortBy:      params.SortBy,
			Order:       params.Order,
			FilterBy:    params.FilterBy,
			FilterValue: params.FilterValue,
		}
		listParams.PageNumber, listParams.PageSize, err = mystring.ConvertPageToInt(params.PageNumber, params.PageSize)
		return params, err
	default:
		return nil, errors.New(fmt.Sprintf("invalid funcName %s", funcName))
	}
}

// CheckAccessType 检查accessType参数 RO RW
func CheckAccessType(accessType string) *fault.Fault {
	if strings.ToUpper(accessType) != RO && strings.ToUpper(accessType) != RW {
		return fault.InvalidParam("accessType", accessType)
	}
	return nil
}

// CheckNameParam 大小规格，检测name只能以大小写字母和数字开头，并且后续字符只能是字母，数字，中划线，下划线
func CheckNameParam(name string) *fault.Fault {
	if len(name) == ZeroLen {
		return fault.ParamLen("name", false, ZeroLen, NameLenMax)
	}
	if len(name) > NameLenMax {
		return fault.ParamLen("name", true, ZeroLen, NameLenMax)
	}
	match, _ := regexp.MatchString(`^[a-zA-Z\d][a-zA-Z\d\-_]*$`, name)
	if !match {
		logging.Log.Errorf("name match failed %s", name)
		return fault.MatchName
	}
	return nil
}

// CheckExampleExisted 检测example是否存在。 expectIsExisted：期望的是否存在，如果不符合期望，会报错。最终返回id值。
func CheckExampleExisted(name string, expectIsExisted bool) (uint, *fault.Fault) {
	var ex db.ExampleDb
	isExisted, err := ex.IsExist(name)
	if err != nil {
		return 0, err
	}
	logging.Log.Infof("name %s, expectIsExisted is %t, isExisted is %t", name, expectIsExisted, isExisted)
	if expectIsExisted && !isExisted {
		return 0, fault.NotExist("example", name)
	} else if !expectIsExisted && isExisted {
		return 0, fault.Existed("example", name)
	}
	return ex.Id, nil
}

func CheckExampleCount() *fault.Fault {
	var ex db.ExampleDb
	var total int64
	if err := ex.QueryCount(&total); err != nil {
		return err
	}
	if total >= 20 {
		return fault.ParamCount("example", 20)
	}
	return nil
}

// CheckSearchParam 检测Searchparam的相关内容，比如判空。FilterBy是一种限制，类似于精确查询。
func CheckSearchParam(params SearchParam, dbName string) *fault.Fault {
	err := CheckOrderParam(params.Order, params.SortBy, dbName)
	if err != nil {
		return err
	}

	// PageNumber和PageSize默认会被赋值为-1，表示没有被传值
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
		if err = CheckNameParam(params.FilterValue); err != nil {
			return err
		}
	}
	return nil
}

// CheckOrderParam sortBy是需要对一些参数进行校验的，如果不是这些参数，需要报错
func CheckOrderParam(order, sort string, dbName string) *fault.Fault {
	if strings.ToUpper(order) != OrderDesc && strings.ToUpper(order) != OrderAsc && order != "" {
		return fault.InvalidParam("order", order)
	}
	if order != "" {
		if len(sort) == 0 {
			return fault.ParamEmpty("sort")
		}
		sortByLower := strings.ToLower(sort)
		switch dbName {
		case ExampleDbName:
			switch sortByLower {
			case "name":
			case "phone":
			default:
				return fault.InvalidParam("sort", sort)
			}
		}
	} else {
		if sort != "" {
			return fault.ParamEmpty("order")
		}
	}
	return nil
}

func CheckOrderParam2(order, sort string) *fault.Fault {
	if strings.ToUpper(order) != OrderDesc && strings.ToUpper(order) != OrderAsc && order != "" {
		return fault.InvalidParam("order", order)
	}
	if len(order) != 0 && len(sort) == 0 {
		return fault.ParamEmpty("sort")
	}
	if len(sort) != 0 && len(order) == 0 {
		return fault.ParamEmpty("order")
	}
	return nil
}

func GetUUID() uuid.UUID {
	return uuid.New()
}

func GetUUIDString() string {
	return uuid.NewString()
}
