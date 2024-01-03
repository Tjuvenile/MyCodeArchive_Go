package db

import (
	"MyCodeArchive_Go/utils/fault"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/agiledragon/gomonkey"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"reflect"
	"testing"
	"time"
)

// getDBMock sqlmock 更适合普通sql，不适合gorm
func getDBMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	// mock 一个 *sql.DB 对象，不需要连接真实的数据库
	sqlDb, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	// 在测试时使用 gorm.Open 来创建 GORM 数据库连接，避免真正连接到数据库
	gdb, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlDb, // 使用上面创建的 mock *sql.DB 对象
		DriverName:                "mysql",
		SkipInitializeWithVersion: true, // 如果此选项不设置，gorm初始化时会调用select version()，如果sqlmock没有mock这个调用则会报错。
	}), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	return gdb, mock, nil
}

// TestDcsRelations_Create 更适合普通sql，不适合gorm
func TestDcsRelations_Create(t *testing.T) {
	gdb, mock, err := getDBMock()
	if err != nil {
		t.Errorf("Error get gorm connection and sql mock: %v", err)
	}
	fmt.Println(gdb == nil)
	mock.ExpectBegin()
	mock.ExpectCommit()
	repo := &DbRepo{}
	repoType := reflect.TypeOf(repo)
	patch := gomonkey.ApplyMethod(repoType, "GetConnect", func(_ *DbRepo) *DbRepo {
		return &DbRepo{
			DbConn: gdb,
		}
	})
	defer patch.Reset()

	tests := []struct {
		name       string
		args       ExampleDb
		want       *fault.Fault
		mockExpect func(db ExampleDb)
	}{
		{
			name: "create relation1",
			args: ExampleDb{
				Name:       "relation1",
				CreateTime: time.Now(),
				UpdateTime: time.Now(),
			},
			want: nil,
			mockExpect: func(args ExampleDb) {
				// 如果不加begin，会报错，显示begin语句未被捕捉到。
				mock.ExpectBegin()
				// args为空时可以匹配所有参数。 如果想让单测更敏感，可以把参数都精确的填写下来。 如果你的sql语句需要参数，但是你没有withargs，会报错
				mock.ExpectExec("^INSERT INTO `example` (.*) VALUES (.*)").WithArgs(args.Name).WillReturnError(nil)
				mock.ExpectCommit()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			re := &ExampleDb{
				Name:       tt.args.Name,
				CreateTime: tt.args.CreateTime,
				UpdateTime: tt.args.UpdateTime,
			}
			tt.mockExpect(tt.args)
			if got := re.Create(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Expectations were not met: %v", err)
	}
}
