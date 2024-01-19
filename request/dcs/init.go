package dcs

import (
	"MyCodeArchive_Go/utils/db"
	"MyCodeArchive_Go/utils/logging"
)

func init() {
	CreateTable()
}

func CreateTable() {
	dbCon := db.DbConnect.GetConnect()
	if dbCon.DbConn == nil {
		logging.Log.Error("fail to get db connect")
		return
	}

	if dbCon.CreateTableAuto(&BgrRelations{}) != 0 {
		logging.Log.Error("failed to create table BgrRelations")
		return
	}

	if dbCon.CreateTableAuto(&BgrStrategies{}) != 0 {
		logging.Log.Error("failed to create table BgrStrategies")
		return
	}

	if dbCon.CreateTableAuto(&BgrRepLinkMgt{}) != 0 {
		logging.Log.Error("failed to create table BgrRepLinkMgt")
		return
	}

	if dbCon.CreateTable(&NodePoolMgt{}) != 0 {
		logging.Log.Errorf("failed to create table NodePoolMgt")
		return
	}

	logging.Log.Infof("end to create DCS table")
}
