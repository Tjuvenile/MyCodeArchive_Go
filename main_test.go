package main

import (
	"MyCodeArchive_Go/request/dcs"
	"MyCodeArchive_Go/utils/db"
	"MyCodeArchive_Go/utils/fault"
	"MyCodeArchive_Go/utils/strings_"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/agiledragon/gomonkey"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"reflect"
	"testing"
	"time"
)

func getDBMock() (sqlmock.Sqlmock, *gomonkey.Patches, error) {
	// mock 一个 *sql.DB 对象，不需要连接真实的数据库
	sqlDb, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		return nil, nil, err
	}

	// 在测试时使用 gorm.Open 来创建 GORM 数据库连接，避免真正连接到数据库
	gdb, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlDb, // 使用上面创建的 mock *sql.DB 对象
		DriverName:                "mysql",
		SkipInitializeWithVersion: true, // 如果此选项不设置，gorm初始化时会调用select version()，如果sqlmock没有mock这个调用则会报错。
	}), &gorm.Config{
		// silent是最高级别，只记录slient及以上的日志
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, nil, err
	}

	repo := &db.DbRepo{}
	repoType := reflect.TypeOf(repo)
	patch := gomonkey.ApplyMethod(repoType, "GetConnect", func(_ *db.DbRepo) *db.DbRepo {
		return &db.DbRepo{
			DbConn: gdb,
		}
	})

	return mock, patch, nil
}

func TestDcsRelations_Create(t *testing.T) {
	mock, funcPatch, err := getDBMock()
	if err != nil {
		t.Errorf("Error get gorm connection and sql mock: %v", err)
	}
	defer funcPatch.Reset()

	tests := []struct {
		name       string
		args       dcs.BgrRelations
		want       *fault.Fault
		mockExpect func(dcs.BgrRelations)
	}{
		{
			name: "create relation1",
			args: dcs.BgrRelations{
				UUID:             uuid.New().String(),
				Name:             "relation1",
				MasterPool:       "pool_test",
				SlavePool:        "slave_test",
				MasterResourceId: "master_resource_id_test",
				SlaveResourceId:  "slave_resource_id_test",
				ResourceType:     "resource_type_test",
				StrategyIds:      "strategy_ids_test",
				LastSyncTime:     111111111,
				LastSyncSnap:     "last_sync_snap_test",
				State:            "state_test",
				RunningState:     "running_state_test",
				HealthState:      "health_state_test",
				DataState:        "data_state_test",
				Role:             "role_test",
				IsConfigSync:     true,
				CreateTime:       time.Now(),
				UpdateTime:       time.Now(),
			},
			want: nil,
			mockExpect: func(args dcs.BgrRelations) {
				// 如果不加begin，会报错，显示begin语句未被捕捉到。
				mock.ExpectBegin()
				// args为空时可以匹配所有参数。 如果想让单测更敏感，可以把参数都精确的填写下来。 如果你的sql语句需要参数，但是你没有withargs，会报错
				// WillReturnResult的作用是对某些sql函数的返回值进行修改。
				// NewResult(1,1)时，当你查看id，RollAffected函数时，返回的就是1,1
				// 如果改成WillReturnError的话，正常来说，会改变ret.Error的值。
				mock.ExpectExec("INSERT INTO `dcs_relations` (`uuid`,`name`,`master_pool`,`slave_pool`,`master_resource_id`,`slave_resource_id`,`resource_type`,`strategy_ids`,`last_sync_time`,`last_sync_snap`,`state`,`running_state`,`health_state`,`data_state`,`role`,`is_config_sync`,`create_time`,`update_time`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)").WithArgs(
					args.UUID, args.Name, args.MasterPool, args.SlavePool, args.MasterResourceId, args.SlaveResourceId,
					args.ResourceType, args.StrategyIds, args.LastSyncTime, args.LastSyncSnap, args.State, args.RunningState,
					args.HealthState, args.DataState, args.Role, args.IsConfigSync, args.CreateTime, args.UpdateTime).
					WillReturnResult(sqlmock.NewResult(1, 1))
				// 如果是willReturnError的话，下面的commit需要改成ExpectRollback
				// mock.ExpectRollback()
				mock.ExpectCommit()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockExpect(tt.args)
			if got := tt.args.Create(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Expectations were not met: %v", err)
	}
}

func TestUnquoteString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "unquote string",
			args: args{str: `\u0026slkdfjkdslfjs(*&#*(@$@\\u\///`},
			want: `&slkdfjkdslfjs(*&#*(@$@\\u\\///`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strings_.UnquoteString(tt.args.str); got != tt.want {
				t.Errorf("UnquoteString() = %v, want %v", got, tt.want)
			}
		})
	}
}
