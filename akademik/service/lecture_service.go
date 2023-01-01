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

func (s *academicService) CreateLecture(ctx context.Context, lecture model.Lecture) (*model.Lecture, *pkgservice.ErrorService) {

	err := s.lectureRepo.InsertLectureRepo(ctx, lecture)
	if err != nil {
		s.log.Error(fmt.Sprintf("[Service_create_lecture] Failed to store data lecture, err: %v", err.Error()))
		return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
	}

	// flush data from redis
	if err := s.redisClient.FlushData(); err != nil {
		s.log.Error(fmt.Sprintf("[Service_create_lecture] Failed to flush redis data: %v", err.Error()))
		return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
	}

	return &lecture, nil
}

func (s *academicService) GetLectureByID(ctx context.Context, ID string) (*model.Lecture, *pkgservice.ErrorService) {

	// check to redis first
	data, err := s.redisClient.GetData("GET_LECTURE_" + ID)
	if err != nil {
		if !strings.EqualFold(err.Error(), "redis: nil") {
			s.log.Error(fmt.Sprintf("[Service_get_lecture] Failed to get data from redis, err: %v", err.Error()))

			return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
		}
	}

	if !strings.EqualFold("", data) {
		s.log.Info("[Service_get_lecture] Success get data from Redis")

		var lecture model.Lecture
		err := json.Unmarshal([]byte(data), &lecture)
		if err != nil {
			return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
		}

		return &lecture, nil
	}

	lecture, err := s.lectureRepo.GetLectureRepoByID(ctx, ID)
	if err != nil {

		if strings.EqualFold(err.Error(), "record not found") {
			return nil, pkgservice.NewErrorService(pkgservice.ErrNotFound)
		}

		s.log.Error(fmt.Sprintf("[Service_get_lecture] Failed to get lecture, err: %v", err.Error()))
		return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
	}

	// save to redis
	if err := s.redisClient.SetDataWithExpiry("GET_LECTURE_"+ID, util.StructToString(lecture), 5); err != nil {
		s.log.Error(fmt.Sprintf("[Service_get_lecture] Failed to save lecture to redis, err: %v", err.Error()))

		return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
	}

	return &lecture, nil
}

func (s *academicService) GetLectures(ctx context.Context) (*[]model.Lecture, *pkgservice.ErrorService) {

	// check to redis first
	data, err := s.redisClient.GetData("GET_LECTURE_ALL")
	if err != nil {
		if !strings.EqualFold(err.Error(), "redis: nil") {
			s.log.Error(fmt.Sprintf("[Service_get_lecture] Failed to get data from redis, err: %v", err.Error()))

			return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
		}
	}

	if !strings.EqualFold("", data) {
		s.log.Info("[Service_get_all_lecture] Success get data from Redis")

		var lecture []model.Lecture
		err := json.Unmarshal([]byte(data), &lecture)
		if err != nil {
			return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
		}

		return &lecture, nil
	}

	listLecture, err := s.lectureRepo.GetLectureRepo(ctx)
	if err != nil {
		s.log.Error(fmt.Sprintf("[Service_get_lecture] Failed to get lecture, err: %v", err.Error()))
		return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
	}

	// save to redis
	if err := s.redisClient.SetDataWithExpiry("GET_LECTURE_ALL", util.StructToString(listLecture), 5); err != nil {
		s.log.Error(fmt.Sprintf("[Service_get_lecture] Failed to save lecture to redis, err: %v", err.Error()))

		return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
	}

	return &listLecture, nil
}

func (s *academicService) UpdateLecture(ctx context.Context, request model.Lecture, ID string) (*model.Lecture, *pkgservice.ErrorService) {

	if err := s.lectureRepo.UpdateScoreLecturerRepo(ctx, request, ID); err != nil {
		s.log.Error(fmt.Sprintf("[Service_get_lecture] Failed to update lecture, err: %v", err.Error()))
		return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
	}

	// flush data from redis
	if err := s.redisClient.FlushData(); err != nil {
		s.log.Error(fmt.Sprintf("[Service_update_lecture] Failed to flush redis data: %v", err.Error()))
		return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
	}

	return &request, nil
}

func (s *academicService) UpdateName(ctx context.Context, request model.Lecture, ID string) (*model.Lecture, *pkgservice.ErrorService) {

	lecture, err := s.lectureRepo.GetLectureRepoByID(ctx, ID)
	if err != nil {
		return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
	}

	lecture.UpdateLectureName(request)
	student, err := s.studentRepo.GetStudentRepoByID(ctx, lecture.StudentID)
	if err != nil {
		return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
	}

	student.UpdateStudentName(request.Name)

	if err := s.lectureRepo.UpdateNameLecturerRepo(ctx, lecture, student); err != nil {
		if strings.Contains(err.Error(), "does not exist") || strings.Contains(err.Error(), "record not found") {
			return nil, pkgservice.NewErrorService(pkgservice.ErrNotFound)
		}

		s.log.Error(fmt.Sprintf("[Service_get_lecture] Failed to update lecture, err: %v", err.Error()))
		return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
	}

	// flush data from redis
	if err := s.redisClient.FlushData(); err != nil {
		s.log.Error(fmt.Sprintf("[Service_update_lecture] Failed to flush redis data: %v", err.Error()))
		return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
	}

	return &lecture, nil
}

func (s *academicService) DeleteLecture(ctx context.Context, id string) (*model.Lecture, *pkgservice.ErrorService) {
	lecture, err := s.lectureRepo.DeleteLectureRepoByID(ctx, id)
	if err != nil {
		if strings.Contains(err.Error(), "does not exist") || strings.Contains(err.Error(), "record not found") {
			return nil, pkgservice.NewErrorService(pkgservice.ErrNotFound)
		}
		s.log.Error(fmt.Sprintf("[Service_delete_lecture] Failed to delete lecture, err: %v", err.Error()))
		return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
	}

	// flush data from redis
	if err := s.redisClient.FlushData(); err != nil {
		s.log.Error(fmt.Sprintf("[Service_delete_lecture] Failed to flush redis data: %v", err.Error()))
		return nil, pkgservice.NewErrorService(pkgservice.ErrInternal)
	}

	return &lecture, nil
}
