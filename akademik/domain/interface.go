package domain

import (
	"context"
	"github.com/44nbud1/akademik/akademik/domain/model"
	"github.com/44nbud1/akademik/pkg/pkgservice"
)

type StudentService interface {
	CreateStudent(ctx context.Context, students model.Students) (*model.Students, *pkgservice.ErrorService)
	GetStudentByID(ctx context.Context, ID string) (*model.Students, *pkgservice.ErrorService)
	GetStudents(ctx context.Context) (*[]model.Students, *pkgservice.ErrorService)
	UpdateStudent(ctx context.Context, request model.Students, ID string) (*model.Students, *pkgservice.ErrorService)
	DeleteStudent(ctx context.Context, id string) (*model.Students, *pkgservice.ErrorService)
}

type CourseService interface {
	CreateCourse(ctx context.Context, course model.Course) (*model.Course, *pkgservice.ErrorService)
	GetCourseByID(ctx context.Context, ID string) (*model.Course, *pkgservice.ErrorService)
	GetCourses(ctx context.Context) (*[]model.Course, *pkgservice.ErrorService)
	UpdateCourse(ctx context.Context, request model.Course, ID string) (*model.Course, *pkgservice.ErrorService)
	DeleteCourse(ctx context.Context, id string) (*model.Course, *pkgservice.ErrorService)
}

type LectureService interface {
	CreateLecture(ctx context.Context, lecture model.Lecture) (*model.Lecture, *pkgservice.ErrorService)
	GetLectureByID(ctx context.Context, ID string) (*model.Lecture, *pkgservice.ErrorService)
	GetLectures(ctx context.Context) (*[]model.Lecture, *pkgservice.ErrorService)
	UpdateLecture(ctx context.Context, request model.Lecture, ID string) (*model.Lecture, *pkgservice.ErrorService)
	DeleteLecture(ctx context.Context, id string) (*model.Lecture, *pkgservice.ErrorService)
	UpdateName(ctx context.Context, request model.Lecture, ID string) (*model.Lecture, *pkgservice.ErrorService)
}

type LecturerService interface {
	CreateLecturer(ctx context.Context, lecturer model.Lecturer) (*model.Lecturer, *pkgservice.ErrorService)
	GetLecturerByID(ctx context.Context, ID string) (*model.Lecturer, *pkgservice.ErrorService)
	GetLecturers(ctx context.Context) (*[]model.Lecturer, *pkgservice.ErrorService)
	UpdateLecturer(ctx context.Context, request model.Lecturer, ID string) (*model.Lecturer, *pkgservice.ErrorService)
	DeleteLecturer(ctx context.Context, id string) (*model.Lecturer, *pkgservice.ErrorService)
}
