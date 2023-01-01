package sql

import (
	"context"
	"github.com/44nbud1/akademik/akademik/domain/model"
)

type StudentRepo interface {
	InsertStudentRepo(ctx context.Context, request model.Students) error
	GetStudentRepoByID(ctx context.Context, ID string) (model.Students, error)
	GetStudentRepo(ctx context.Context) ([]model.Students, error)
	UpdateStudentRepo(ctx context.Context, request model.Students, ID string) (model.Students, error)
	DeleteStudentRepoByID(ctx context.Context, ID string) (model.Students, error)
}

type CourseRepo interface {
	InsertCourseRepo(ctx context.Context, request model.Course) error
	GetCourseRepoByID(ctx context.Context, ID string) (model.Course, error)
	GetCourseRepo(ctx context.Context) ([]model.Course, error)
	UpdateCourseRepo(ctx context.Context, request model.Course, ID string) (model.Course, error)
	DeleteCourseRepoByID(ctx context.Context, ID string) (model.Course, error)
}

type LectureRepo interface {
	InsertLectureRepo(ctx context.Context, request model.Lecture) error
	GetLectureRepoByID(ctx context.Context, ID string) (model.Lecture, error)
	GetLectureRepo(ctx context.Context) ([]model.Lecture, error)
	UpdateScoreLecturerRepo(ctx context.Context, request model.Lecture, ID string) error
	UpdateNameLecturerRepo(ctx context.Context, lectureRequest model.Lecture, studentRequest model.Students) error
	DeleteLectureRepoByID(ctx context.Context, ID string) (model.Lecture, error)
}

type LecturerRepo interface {
	InsertLecturerRepo(ctx context.Context, request model.Lecturer) error
	GetLecturerRepoByID(ctx context.Context, ID string) (model.Lecturer, error)
	GetLecturerRepo(ctx context.Context) ([]model.Lecturer, error)
	UpdateLecturerRepo(ctx context.Context, lecture model.Lecturer, ID string) (model.Lecturer, error)
	DeleteLecturerRepoByID(ctx context.Context, ID string) (model.Lecturer, error)
}
