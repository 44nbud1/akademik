package sql

import (
	"context"
	"errors"
	"github.com/44nbud1/akademik/akademik/domain/model"
)

func (r *repository) InsertLectureRepo(ctx context.Context, request model.Lecture) error {

	if result := r.db.Create(&request); result.Error != nil {
		return errors.New(result.Error.Error())
	}

	return nil
}

func (r *repository) GetLectureRepoByID(ctx context.Context, ID string) (model.Lecture, error) {
	var lecture model.Lecture
	if result := r.db.First(&lecture, ID); result.Error != nil {
		return model.Lecture{}, errors.New(result.Error.Error())
	}

	return lecture, nil
}

func (r *repository) GetLectureRepo(ctx context.Context) ([]model.Lecture, error) {
	var lecture []model.Lecture
	if result := r.db.Find(&lecture); result.Error != nil {
		return []model.Lecture{}, errors.New(result.Error.Error())
	}

	return lecture, nil
}

func (r *repository) UpdateScoreLecturerRepo(ctx context.Context, request model.Lecture, ID string) error {
	lecture, err := r.GetLectureRepoByID(ctx, ID)
	if err != nil {
		return err
	}

	lecture.UpdateLecture(request)
	if result := r.db.Save(&lecture); result.Error != nil {
		return errors.New(result.Error.Error())
	}

	return nil
}

func (r *repository) DeleteLectureRepoByID(ctx context.Context, ID string) (model.Lecture, error) {
	lecture, err := r.GetLectureRepoByID(ctx, ID)
	if err != nil {
		return model.Lecture{}, err
	}

	if result := r.db.Unscoped().Delete(&lecture); result.Error != nil {
		return model.Lecture{}, result.Error
	}

	return lecture, nil
}

func (r *repository) UpdateNameLecturerRepo(ctx context.Context, lectureRequest model.Lecture, studentRequest model.Students) error {

	return r.UpdateName(ctx, lectureRequest, studentRequest)
}
