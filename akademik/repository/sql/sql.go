package sql

import (
	"context"
	"fmt"
	"github.com/44nbud1/akademik/akademik/domain/model"
	"github.com/44nbud1/akademik/akademik/util"
	"github.com/jinzhu/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewGorm(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) UpdateName(ctx context.Context, lecture model.Lecture, students model.Students) error {

	fmt.Println("HITTTT")
	fmt.Println(util.StructToString(lecture))
	fmt.Println(util.StructToString(students))

	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Save(&lecture).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Save(&students).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
