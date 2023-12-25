package dcs

const (
	ZeroLen     = 0
	NameLenMax  = 256
	PageSizeMax = 100

	RelationModule             = "relation"
	StrategyModule             = "strategy"
	CreateRelationFun          = "CreateRelationFun"
	DeleteRelationFun          = "DeleteRelationFun"
	UpdateRelationFun          = "UpdateRelationFun"
	ShowRelationFun            = "ShowRelationFun"
	ListRelationsFun           = "ListRelationsFun"
	ListRelationsStrategiesFun = "ListRelationsStrategiesFun"
	CreateStrategyFun          = "CreateStrategyFun"
	DeleteStrategyFun          = "DeleteStrategyFun"
	UpdateStrategyFun          = "UpdateStrategyFun"
	ShowStrategyFun            = "ShowStrategyFun"
	ListStrategiesFun          = "ListStrategiesFun"
	OrderDesc                  = "DESC"
	OrderAsc                   = "ASC"

	Success = "Success"
)

type Resource struct {
	ID   string
	Name string
}

// RelationParam 复制关系结构体
type RelationParam struct {
	UUID           string   `json:"UUID"`
	Name           string   `json:"Name"`
	MasterIp       string   `json:"MasterIp"`
	SlaveIp        string   `json:"SlaveIp"`
	MasterPool     string   `json:"MasterPool"`
	SlavePool      string   `json:"SlavePool"`
	MasterResource Resource `json:"MasterResource"` // 主端资源
	SlaveResource  Resource `json:"SlaveResource"`  // 从端资源
	ResourceType   string   `json:"ResourceType"`   // 资源类型：块/文件/对象
	StrategyIds    []string `json:"StrategyIds"`    // 策略id列表
	LastSyncTime   int64    `json:"LastSyncTime"`   // 上次同步时间
	LastSyncSnap   string   `json:"LastSyncSnap"`   // 上次同步快照
	CreateTime     int64    `json:"CreateTime"`
	Status         string   `json:"Status"`       // 中间状态
	RunningState   string   `json:"RunningState"` // 运行状态
	HealthState    string   `json:"HealthState"`  // 健康状态
	DataState      string   `json:"DataState"`    // 数据状态
	Role           string   `json:"Role"`         // 记录本地角色（主/从）
	IsConfigSync   bool     `json:"IsConfigSync"` // 是否开启资源配置同步
	RecoverCheck   bool     `json:"RecoverCheck"` // 故障/断开状态恢复标志位，用于处理故障恢复后的一些初始化工作

	OutJson bool `json:"json"` // cli json格式返回结果
}

// StrategyParam 复制关系策略结构体
type StrategyParam struct {
	UUID         string `json:"UUID"` // 策略id
	Name         string `json:"Name"`
	TimePoint    string `json:"TimePoint"`    // 定时同步时间
	Interval     string `json:"Interval"`     // 同步间隔
	StrategyType string `json:"StrategyType"` // 定时策略类型
	Description  string `json:"Description"`
	OutJson      bool   `json:"json"` // cli json格式返回结果
}

type RelationsStrategies struct {
	DcsRelations
	Strategies []DcsStrategies `json:"Strategies"`
}

type FilterParam struct {
	OutJson     bool `json:"json"`
	SearchParam      // pageNumber和size为int类型
}

type FilterParamApi struct {
	OutJson        bool `json:"json"`
	SearchParamApi      // pageNumber和size为string类型
}

type SearchParam struct {
	PageNumber  int    `json:"PageNumber"`
	PageSize    int    `json:"PageSize"`
	SortBy      string `json:"SortBy"`
	Order       string `json:"Order"`
	FilterBy    string `json:"FilterBy"`
	FilterValue string `json:"FilterValue"`
}

type SearchParamApi struct {
	PageNumber  string `json:"PageNumber"`
	PageSize    string `json:"PageSize"`
	SortBy      string `json:"SortBy"`
	Order       string `json:"Order"`
	FilterBy    string `json:"FilterBy"`
	FilterValue string `json:"FilterValue"`
}

type BaseParam struct {
	UUID string `json:"UUID"`
	Name string `json:"Name"`
}

type RelationUpdateParam struct {
	UUID    string `json:"UUID"`
	Name    string `json:"Name"`
	NewName string `json:"NewName"`
}

type StrategyUpdateParam struct {
	UUID        string `json:"UUID"` // 策略id
	Name        string `json:"Name"`
	NewName     string `json:"NewName"`
	Description string `json:"Description"`
}
