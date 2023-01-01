package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/44nbud1/akademik/akademik/domain/model"
	"github.com/44nbud1/akademik/akademik/util"
	"github.com/44nbud1/akademik/pkg/pkgservice"
	"strings"
)

func (s *academicService) CreateCourse(ctx context.Context, course model.Course) (*model.Course, *pkgservice.ErrorService) {

	err := s.courseRepo.InsertCourseRepo(ctx, course)
	if err != nil {
		s.log.Error(fmt.Sprintf("[Service_create_course] Failed to store data course, err: %v", err.Error()))
		return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
	}

	// flush data from redis
	if err := s.redisClient.FlushData(); err != nil {
		s.log.Error(fmt.Sprintf("[Service_create_course] Failed to flush redis data: %v", err.Error()))
		return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
	}

	return &course, nil
}

func (s *academicService) GetCourseByID(ctx context.Context, ID string) (*model.Course, *pkgservice.ErrorService) {

	// check to redis first
	data, err := s.redisClient.GetData("GET_COURSE_" + ID)
	if err != nil {
		if !strings.EqualFold(err.Error(), "redis: nil") {
			s.log.Error(fmt.Sprintf("[Service_get_course] Failed to get data from redis, err: %v", err.Error()))

			return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
		}
	}

	if !strings.EqualFold("", data) {
		s.log.Info("[Service_get_course] Success get data from Redis")

		var course model.Course
		err := json.Unmarshal([]byte(data), &course)
		if err != nil {
			return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
		}

		return &course, nil
	}

	course, err := s.courseRepo.GetCourseRepoByID(ctx, ID)
	if err != nil {

		if strings.EqualFold(err.Error(), "record not found") {
			return nil, pkgservice.NewErrorService(pkgservice.ErrNotFound)
		}

		s.log.Error(fmt.Sprintf("[Service_get_course] Failed to get course, err: %v", err.Error()))
		return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
	}

	// save to redis
	if err := s.redisClient.SetDataWithExpiry("GET_COURSE_"+ID, util.StructToString(course), 5); err != nil {
		s.log.Error(fmt.Sprintf("[Service_get_course] Failed to save course to redis, err: %v", err.Error()))

		return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
	}

	return &course, nil
}

func (s *academicService) GetCourses(ctx context.Context) (*[]model.Course, *pkgservice.ErrorService) {

	// check to redis first
	data, err := s.redisClient.GetData("GET_ALL_COURSE")
	if err != nil {
		if !strings.EqualFold(err.Error(), "redis: nil") {
			s.log.Error(fmt.Sprintf("[Service_get_course] Failed to get data from redis, err: %v", err.Error()))

			return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
		}
	}

	if !strings.EqualFold("", data) {
		s.log.Info("[Service_get_course] Success get data from Redis")

		var course []model.Course
		err := json.Unmarshal([]byte(data), &course)
		if err != nil {
			return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
		}

		return &course, nil
	}

	listCourse, err := s.courseRepo.GetCourseRepo(ctx)
	if err != nil {
		s.log.Error(fmt.Sprintf("[Service_get_course] Failed to get course, err: %v", err.Error()))
		return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
	}

	// save to redis
	if err := s.redisClient.SetDataWithExpiry("GET_ALL_COURSE", util.StructToString(listCourse), 5); err != nil {
		s.log.Error(fmt.Sprintf("[Service_get_course] Failed to save course to redis, err: %v", err.Error()))

		return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
	}

	return &listCourse, nil
}

func (s *academicService) UpdateCourse(ctx context.Context, request model.Course, ID string) (*model.Course, *pkgservice.ErrorService) {

	course, err := s.courseRepo.UpdateCourseRepo(ctx, request, ID)
	if err != nil {
		s.log.Error(fmt.Sprintf("[Service_get_course] Failed to update course, err: %v", err.Error()))
		return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
	}

	// flush data from redis
	if err := s.redisClient.FlushData(); err != nil {
		s.log.Error(fmt.Sprintf("[Service_update_course] Failed to flush redis data: %v", err.Error()))
		return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
	}

	return &course, nil
}

func (s *academicService) DeleteCourse(ctx context.Context, id string) (*model.Course, *pkgservice.ErrorService) {

	course, err := s.courseRepo.DeleteCourseRepoByID(ctx, id)
	if err != nil {
		if strings.Contains(err.Error(), "does not exist") {
			return nil, pkgservice.NewErrorService(pkgservice.ErrNotFound)
		}
		s.log.Error(fmt.Sprintf("[Service_delete_course] Failed to delete course, err: %v", err.Error()))
		return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
	}

	// flush data from redis
	if err := s.redisClient.FlushData(); err != nil {
		s.log.Error(fmt.Sprintf("[Service_delete_course] Failed to flush redis data: %v", err.Error()))
		return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
	}

	return &course, nil
}
