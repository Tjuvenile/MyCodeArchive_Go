package db

import "MyCodeArchive_Go/logging"

// CreateExampleDbTable 创建一张表。如果表已存在，会出错，可以忽略，如果表不存在，直接创建
func CreateExampleDbTable() {
	dbCon, err1 := GetConnect()
	if err1 != nil {
		logging.Log.Infof("fail to get db session")
	}
	if 0 != dbCon.CreateTableAuto(&ExampleDb{}) {
		logging.Log.Infof("failed to create table ExampleDb")
	}
	//dbCon.DbConn.Commit()
	logging.Log.Infof("end to create ExampleDb table")
}
