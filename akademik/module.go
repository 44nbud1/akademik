package akademik

import (
	"github.com/44nbud1/akademik/akademik/delivery/http"
	"github.com/44nbud1/akademik/akademik/repository/rdb"
	"github.com/44nbud1/akademik/akademik/repository/sql"
	"github.com/44nbud1/akademik/akademik/service"
	"github.com/bwmarrin/snowflake"
	"github.com/go-redis/redis/v7"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type Module struct{}

func NewAcademicModule(db *gorm.DB, snowflakeNode *snowflake.Node, router *mux.Router, logger *zap.Logger, client *redis.Client) {
	sqlDB := sql.NewGorm(db)
	// migrate db
	sqlDB.AutoMigrate()

	redisDB := rdb.NewRedisRepo(client)
	redisDB.Ping()
	svc := service.NewAcademicService(sqlDB, sqlDB, snowflakeNode, logger, sqlDB, sqlDB, redisDB)
	endpoint := http.NewEndpointAcademic(svc, svc, svc, svc, logger)
	http.NewAcademicHandler(router, endpoint, endpoint, endpoint, endpoint, snowflakeNode)
}
