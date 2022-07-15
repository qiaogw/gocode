package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gocode/config"
	"gorm.io/gorm"
)

var (
	GenDB      *gorm.DB
	GEN_DBList map[string]*gorm.DB
	GEN_REDIS  *redis.Client
	GenViper   *viper.Viper
	GEN_LOG    *zap.Logger
	GenConfig  *config.APP
)
