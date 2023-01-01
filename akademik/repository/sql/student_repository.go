package sql

import (
	"context"
	"errors"
	"github.com/44nbud1/akademik/akademik/domain/model"
)

func (r *repository) InsertStudentRepo(ctx context.Context, request model.Students) error {
	if result := r.db.Create(&request); result.Error != nil {
		return errors.New(result.Error.Error())
	}

	return nil
}

func (r *repository) GetStudentRepoByID(ctx context.Context, ID string) (model.Students, error) {
	var student model.Students
	if result := r.db.First(&student, ID); result.Error != nil {
		return model.Students{}, errors.New(result.Error.Error())
	}

	return student, nil
}

func (r *repository) GetStudentRepo(ctx context.Context) ([]model.Students, error) {
	var student []model.Students
	if result := r.db.Find(&student); result.Error != nil {
		return []model.Students{}, errors.New(result.Error.Error())
	}

	return student, nil
}

func (r *repository) UpdateStudentRepo(ctx context.Context, request model.Students, ID string) (model.Students, error) {
	student, err := r.GetStudentRepoByID(ctx, ID)
	if err != nil {
		return model.Students{}, err
	}

	student.UpdateStudent(request)
	if result := r.db.Save(&student); result.Error != nil {
		return model.Students{}, errors.New(result.Error.Error())
	}

	return student, nil
}

func (r *repository) DeleteStudentRepoByID(ctx context.Context, ID string) (model.Students, error) {
	student, err := r.GetStudentRepoByID(ctx, ID)
	if err != nil {
		return model.Students{}, err
	}

	if result := r.db.Unscoped().Delete(&student); result.Error != nil {
		return model.Students{}, result.Error
	}

	return student, nil
}
