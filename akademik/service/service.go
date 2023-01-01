package service

import (
	"github.com/44nbud1/akademik/akademik/repository/rdb"
	"github.com/44nbud1/akademik/akademik/repository/sql"
	"github.com/bwmarrin/snowflake"
	"go.uber.org/zap"
)

type academicService struct {
	lectureRepo   sql.LectureRepo
	lecturerRepo  sql.LecturerRepo
	courseRepo    sql.CourseRepo
	studentRepo   sql.StudentRepo
	snowflakeNode *snowflake.Node
	log           *zap.Logger
	redisClient   rdb.CacheRepository
}

func NewAcademicService(repo sql.StudentRepo, courseRepo sql.CourseRepo, snowflakeNode *snowflake.Node, log *zap.Logger, lectureRepo sql.LectureRepo,
	lecturerRepo sql.LecturerRepo, client rdb.CacheRepository) *academicService {
	return &academicService{
		studentRepo:   repo,
		snowflakeNode: snowflakeNode,
		log:           log,
		courseRepo:    courseRepo,
		lectureRepo:   lectureRepo,
		lecturerRepo:  lecturerRepo,
		redisClient:   client,
	}
}
