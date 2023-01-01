package sql

import (
	"context"
	"errors"
	"github.com/44nbud1/akademik/akademik/domain/model"
)

func (r *repository) InsertCourseRepo(ctx context.Context, request model.Course) error {
	if result := r.db.Create(&request); result.Error != nil {
		return errors.New(result.Error.Error())
	}

	return nil
}

func (r *repository) GetCourseRepoByID(ctx context.Context, ID string) (model.Course, error) {
	var course model.Course
	if result := r.db.First(&course, ID); result.Error != nil {
		return model.Course{}, errors.New(result.Error.Error())
	}

	return course, nil
}

func (r *repository) GetCourseRepo(ctx context.Context) ([]model.Course, error) {
	var courses []model.Course
	if result := r.db.Find(&courses); result.Error != nil {
		return []model.Course{}, errors.New(result.Error.Error())
	}

	return courses, nil
}

func (r *repository) UpdateCourseRepo(ctx context.Context, request model.Course, ID string) (model.Course, error) {
	course, err := r.GetCourseRepoByID(ctx, ID)
	if err != nil {
		return model.Course{}, err
	}

	course.UpdateCourse(request)
	if result := r.db.Save(&course); result.Error != nil {
		return model.Course{}, errors.New(result.Error.Error())
	}

	return course, nil
}

func (r *repository) DeleteCourseRepoByID(ctx context.Context, ID string) (model.Course, error) {
	course, err := r.GetCourseRepoByID(ctx, ID)
	if err != nil {
		return model.Course{}, err
	}

	if result := r.db.Unscoped().Delete(&course); result.Error != nil {
		return model.Course{}, result.Error
	}

	return course, nil
}
