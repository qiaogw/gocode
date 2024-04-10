package gormx

import (
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

const (
	//gormSpanKey        = "__gorm_span"
	callBackAfterZeroGorm = "logger:afterZeroGorm"
)

func afterZeroGorm(db *gorm.DB) {
	// Error
	if db.Error != nil {
		logx.Error(db.Error)
	}
	// sql
	logx.Infof(db.Statement.SQL.String(), db.Statement.Vars...)

	return
}

type ZeroGorm struct{}

func (op *ZeroGorm) Name() string {
	return "zeroGorm"
}

func (op *ZeroGorm) Initialize(db *gorm.DB) (err error) {
	// 开始前

	// 结束后
	db.Callback().Create().After("gorm:after_create").Register(callBackAfterZeroGorm, afterZeroGorm)
	db.Callback().Query().After("gorm:after_query").Register(callBackAfterZeroGorm, afterZeroGorm)
	db.Callback().Delete().After("gorm:after_delete").Register(callBackAfterZeroGorm, afterZeroGorm)
	db.Callback().Update().After("gorm:after_update").Register(callBackAfterZeroGorm, afterZeroGorm)
	db.Callback().Row().After("gorm:row").Register(callBackAfterZeroGorm, afterZeroGorm)
	db.Callback().Raw().After("gorm:raw").Register(callBackAfterZeroGorm, afterZeroGorm)
	return
}

var _ gorm.Plugin = &ZeroGorm{}
