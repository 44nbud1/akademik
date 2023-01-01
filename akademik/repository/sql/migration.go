package sql

import (
	"github.com/44nbud1/akademik/akademik/domain/model"
	"github.com/44nbud1/akademik/akademik/util"
	"log"
)

func (r *repository) AutoMigrate() {
	db := r.db.AutoMigrate(&model.Students{}, &model.Course{}, &model.Lecturer{}, &model.Lecture{})
	if db != nil && db.Error != nil {
		log.Print("Failed to migrate, err: ", util.StructToString(db.Error))
		return
	}

	r.db.Model(&model.Lecture{}).Related(&model.Students{}, "ID")
	r.db.LogMode(true)
}
