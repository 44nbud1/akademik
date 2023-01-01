package sql

import (
	"context"
	"errors"
	"github.com/44nbud1/akademik/akademik/domain/model"
)

func (r *repository) InsertLecturerRepo(ctx context.Context, request model.Lecturer) error {
	if result := r.db.Create(&request); result.Error != nil {
		return errors.New(result.Error.Error())
	}

	return nil
}

func (r *repository) GetLecturerRepoByID(ctx context.Context, ID string) (model.Lecturer, error) {
	var lecturer model.Lecturer
	if result := r.db.First(&lecturer, ID); result.Error != nil {
		return model.Lecturer{}, errors.New(result.Error.Error())
	}

	return lecturer, nil
}

func (r *repository) GetLecturerRepo(ctx context.Context) ([]model.Lecturer, error) {
	var lecturers []model.Lecturer
	if result := r.db.Find(&lecturers); result.Error != nil {
		return []model.Lecturer{}, errors.New(result.Error.Error())
	}

	return lecturers, nil
}

func (r *repository) UpdateLecturerRepo(ctx context.Context, lecturer model.Lecturer, ID string) (model.Lecturer, error) {

	lecturerData, err := r.GetLecturerRepoByID(ctx, ID)
	if err != nil {
		return model.Lecturer{}, err
	}

	lecturerData.UpdateLecturer(lecturer)
	if result := r.db.Save(&lecturerData); result.Error != nil {
		return model.Lecturer{}, errors.New(result.Error.Error())
	}

	return lecturerData, nil
}

func (r *repository) DeleteLecturerRepoByID(ctx context.Context, ID string) (model.Lecturer, error) {
	lecturer, err := r.GetLecturerRepoByID(ctx, ID)
	if err != nil {
		return model.Lecturer{}, err
	}

	if result := r.db.Unscoped().Delete(&lecturer); result.Error != nil {
		return model.Lecturer{}, result.Error
	}

	return lecturer, nil
}
