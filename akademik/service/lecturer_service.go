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

func (s *academicService) CreateLecturer(ctx context.Context, lecturer model.Lecturer) (*model.Lecturer, *pkgservice.ErrorService) {

	err := s.lecturerRepo.InsertLecturerRepo(ctx, lecturer)
	if err != nil {
		s.log.Error(fmt.Sprintf("[Service_create_lecturer] Failed to store data lecturer, err: %v", err.Error()))
		return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
	}

	// flush data from redis
	if err := s.redisClient.FlushData(); err != nil {
		s.log.Error(fmt.Sprintf("[Service_create_lecturer] Failed to flush redis data: %v", err.Error()))
		return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
	}

	return &lecturer, nil
}

func (s *academicService) GetLecturerByID(ctx context.Context, ID string) (*model.Lecturer, *pkgservice.ErrorService) {

	// check to redis first
	data, err := s.redisClient.GetData("GET_LECTURER_" + ID)
	if err != nil {
		if !strings.EqualFold(err.Error(), "redis: nil") {
			s.log.Error(fmt.Sprintf("[Service_get_lecturer] Failed to get data from redis, err: %v", err.Error()))

			return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
		}
	}

	if !strings.EqualFold("", data) {
		s.log.Info("[Service_get_lecture] Success get data from Redis")

		var lecturer model.Lecturer
		err := json.Unmarshal([]byte(data), &lecturer)
		if err != nil {
			return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
		}

		return &lecturer, nil
	}

	lecturer, err := s.lecturerRepo.GetLecturerRepoByID(ctx, ID)
	if err != nil {

		if strings.EqualFold(err.Error(), "record not found") {
			return nil, pkgservice.NewErrorService(pkgservice.ErrNotFound)
		}

		s.log.Error(fmt.Sprintf("[Service_get_lecturer] Failed to get lecturer, err: %v", err.Error()))
		return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
	}

	// save to redis
	if err := s.redisClient.SetDataWithExpiry("GET_LECTURER_"+ID, util.StructToString(lecturer), 5); err != nil {
		s.log.Error(fmt.Sprintf("[Service_get_lecturer] Failed to save lecturer to redis, err: %v", err.Error()))

		return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
	}

	return &lecturer, nil
}

func (s *academicService) GetLecturers(ctx context.Context) (*[]model.Lecturer, *pkgservice.ErrorService) {

	// check to redis first
	data, err := s.redisClient.GetData("GET_ALL_LECTURER")
	if err != nil {
		if !strings.EqualFold(err.Error(), "redis: nil") {
			s.log.Error(fmt.Sprintf("[Service_get_lecturer] Failed to get data from redis, err: %v", err.Error()))

			return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
		}
	}

	if !strings.EqualFold("", data) {
		s.log.Info("[Service_get_lecture] Success get data from Redis")

		var lecturer []model.Lecturer
		err := json.Unmarshal([]byte(data), &lecturer)
		if err != nil {
			return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
		}

		return &lecturer, nil
	}
	listLecturer, err := s.lecturerRepo.GetLecturerRepo(ctx)
	if err != nil {
		s.log.Error(fmt.Sprintf("[Service_get_lecturer] Failed to get lecturer, err: %v", err.Error()))
		return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
	}

	// save to redis
	if err := s.redisClient.SetDataWithExpiry("GET_ALL_LECTURER", util.StructToString(listLecturer), 5); err != nil {
		s.log.Error(fmt.Sprintf("[Service_get_lecturer] Failed to save lecturer to redis, err: %v", err.Error()))

		return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
	}

	return &listLecturer, nil
}

func (s *academicService) UpdateLecturer(ctx context.Context, request model.Lecturer, ID string) (*model.Lecturer, *pkgservice.ErrorService) {
	lecturer, err := s.lecturerRepo.UpdateLecturerRepo(ctx, request, ID)
	if err != nil {
		if strings.Contains(err.Error(), "does not exist") || strings.Contains(err.Error(), "record not found") {
			return nil, pkgservice.NewErrorService(pkgservice.ErrNotFound)
		}

		s.log.Error(fmt.Sprintf("[Service_get_lecturer] Failed to update lecturer, err: %v", err.Error()))
		return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
	}

	// flush data from redis
	if err := s.redisClient.FlushData(); err != nil {
		s.log.Error(fmt.Sprintf("[Service_update_lecturer] Failed to flush redis data: %v", err.Error()))
		return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
	}

	return &lecturer, nil
}

func (s *academicService) DeleteLecturer(ctx context.Context, id string) (*model.Lecturer, *pkgservice.ErrorService) {
	lecturer, err := s.lecturerRepo.DeleteLecturerRepoByID(ctx, id)
	if err != nil {
		if strings.Contains(err.Error(), "does not exist") || strings.Contains(err.Error(), "record not found") {
			return nil, pkgservice.NewErrorService(pkgservice.ErrNotFound)
		}
		s.log.Error(fmt.Sprintf("[Service_delete_lecturer] Failed to delete lecturer, err: %v", err.Error()))
		return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
	}

	// flush data from redis
	if err := s.redisClient.FlushData(); err != nil {
		s.log.Error(fmt.Sprintf("[Service_delete_lecturer] Failed to flush redis data: %v", err.Error()))
		return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
	}

	return &lecturer, nil
}
