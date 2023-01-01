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

func (s *academicService) CreateStudent(ctx context.Context, students model.Students) (*model.Students, *pkgservice.ErrorService) {

	// flush data from redis
	if err := s.redisClient.FlushData(); err != nil {
		s.log.Error(fmt.Sprintf("[Service_create_student] Failed to create redis data: %v", err.Error()))
		return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
	}

	err := s.studentRepo.InsertStudentRepo(ctx, students)
	if err != nil {
		s.log.Error(fmt.Sprintf("[Service_create_student] Failed to store data student, err: %v", err.Error()))
		return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
	}

	return &students, nil
}

func (s *academicService) GetStudentByID(ctx context.Context, ID string) (*model.Students, *pkgservice.ErrorService) {

	// check to redis first
	data, err := s.redisClient.GetData("GET_STUDENT_" + ID)
	if err != nil {
		if !strings.EqualFold(err.Error(), "redis: nil") {
			s.log.Error(fmt.Sprintf("[Service_get_student] Failed to get data from redis, err: %v", err.Error()))

			return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
		}
	}

	if !strings.EqualFold("", data) {
		s.log.Info("[Service_get_student] Success get data from Redis")

		var student model.Students
		err := json.Unmarshal([]byte(data), &student)
		if err != nil {
			return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
		}

		return &student, nil
	}

	student, err := s.studentRepo.GetStudentRepoByID(ctx, ID)
	if err != nil {
		if strings.EqualFold(err.Error(), "record not found") {
			return nil, pkgservice.NewErrorService(pkgservice.ErrNotFound)
		}

		s.log.Error(fmt.Sprintf("[Service_get_student] Failed to get student, err: %v", err.Error()))
		return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
	}

	// save to redis
	if err := s.redisClient.SetDataWithExpiry("GET_STUDENT_"+ID, util.StructToString(student), 5); err != nil {
		s.log.Error(fmt.Sprintf("[Service_get_student] Failed to save student to redis, err: %v", err.Error()))

		return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
	}

	return &student, nil
}

func (s *academicService) GetStudents(ctx context.Context) (*[]model.Students, *pkgservice.ErrorService) {
	// check to redis first
	data, err := s.redisClient.GetData("GET_ALL_STUDENTS")
	if err != nil {
		if !strings.EqualFold(err.Error(), "redis: nil") {
			s.log.Error(fmt.Sprintf("[Service_get_student] Failed to get data from redis, err: %v", err.Error()))

			return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
		}
	}

	if !strings.EqualFold("", data) {
		s.log.Info("[Service_get_all_student] Success get data from Redis")

		var students []model.Students
		err := json.Unmarshal([]byte(data), &students)
		if err != nil {
			return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
		}

		return &students, nil
	}

	listStudent, err := s.studentRepo.GetStudentRepo(ctx)
	if err != nil {
		s.log.Error(fmt.Sprintf("[Service_get_student] Failed to get student, err: %v", err.Error()))
		return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
	}

	// save to redis
	if err := s.redisClient.SetDataWithExpiry("GET_ALL_STUDENTS", util.StructToString(listStudent), 5); err != nil {
		s.log.Error(fmt.Sprintf("[Service_get_student] Failed to save student to redis, err: %v", err.Error()))

		return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
	}

	return &listStudent, nil
}

func (s *academicService) UpdateStudent(ctx context.Context, request model.Students, ID string) (*model.Students, *pkgservice.ErrorService) {

	// flush data from redis
	if err := s.redisClient.FlushData(); err != nil {
		s.log.Error(fmt.Sprintf("[Service_update_student] Failed to clear redis data: %v", err.Error()))
		return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
	}

	student, err := s.studentRepo.UpdateStudentRepo(ctx, request, ID)
	if err != nil {
		if strings.Contains(err.Error(), "does not exist") || strings.Contains(err.Error(), "record not found") {
			return nil, pkgservice.NewErrorService(pkgservice.ErrNotFound)
		}
		s.log.Error(fmt.Sprintf("[Service_get_student] Failed to update student, err: %v", err.Error()))
		return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
	}

	return &student, nil
}

func (s *academicService) DeleteStudent(ctx context.Context, id string) (*model.Students, *pkgservice.ErrorService) {
	student, err := s.studentRepo.DeleteStudentRepoByID(ctx, id)
	if err != nil {
		if strings.Contains(err.Error(), "does not exist") || strings.Contains(err.Error(), "record not found") {
			return nil, pkgservice.NewErrorService(pkgservice.ErrNotFound)
		}
		s.log.Error(fmt.Sprintf("[Service_delete_student] Failed to delete student, err: %v", err.Error()))
		return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
	}
	return &student, nil
}
