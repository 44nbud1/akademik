package http

import (
	"github.com/44nbud1/akademik/pkg/delivery/pkghttp"
	"github.com/bwmarrin/snowflake"
	"github.com/gorilla/mux"
	"net/http"
)

func NewAcademicHandler(p *mux.Router, es EndpointStudent, ec EndpointCourse, el EndpointLecture, elr EndpointLecturer, snowflakeNode *snowflake.Node) {
	r := p.NewRoute().Subrouter()

	// Student
	r.Methods(http.MethodPost).Path("/v1.0/students").Handler(pkghttp.NewServer(es.CreateStudent(snowflakeNode)))
	r.Methods(http.MethodGet).Path("/v1.0/students/{id}").Handler(pkghttp.NewServer(es.GetStudentByID()))
	r.Methods(http.MethodGet).Path("/v1.0/students").Handler(pkghttp.NewServer(es.GetStudents()))
	r.Methods(http.MethodPut).Path("/v1.0/students/{id}").Handler(pkghttp.NewServer(es.UpdateStudent()))
	r.Methods(http.MethodDelete).Path("/v1.0/students/{id}").Handler(pkghttp.NewServer(es.DeleteStudent()))

	// Course
	r.Methods(http.MethodPost).Path("/v1.0/courses").Handler(pkghttp.NewServer(ec.CreateCourse(snowflakeNode)))
	r.Methods(http.MethodGet).Path("/v1.0/courses/{id}").Handler(pkghttp.NewServer(ec.GetCourseByID()))
	r.Methods(http.MethodGet).Path("/v1.0/courses").Handler(pkghttp.NewServer(ec.GetCourses()))
	r.Methods(http.MethodPut).Path("/v1.0/courses/{id}").Handler(pkghttp.NewServer(ec.UpdateCourse()))
	r.Methods(http.MethodDelete).Path("/v1.0/courses/{id}").Handler(pkghttp.NewServer(ec.DeleteCourse()))

	// Lecture
	r.Methods(http.MethodPost).Path("/v1.0/lectures").Handler(pkghttp.NewServer(el.CreateLecture(snowflakeNode)))
	r.Methods(http.MethodGet).Path("/v1.0/lectures/{id}").Handler(pkghttp.NewServer(el.GetLectureByID()))
	r.Methods(http.MethodGet).Path("/v1.0/lectures").Handler(pkghttp.NewServer(el.GetLectures()))
	r.Methods(http.MethodPatch).Path("/v1.0/lectures/{id}").Handler(pkghttp.NewServer(el.UpdateLecture()))
	r.Methods(http.MethodDelete).Path("/v1.0/lectures/{id}").Handler(pkghttp.NewServer(el.DeleteLecture()))
	r.Methods(http.MethodPatch).Path("/v1.0/lectures-name/{id}").Handler(pkghttp.NewServer(el.UpdateName()))

	// Lecturer
	r.Methods(http.MethodPost).Path("/v1.0/lecturers").Handler(pkghttp.NewServer(elr.CreateLecturer(snowflakeNode)))
	r.Methods(http.MethodGet).Path("/v1.0/lecturers/{id}").Handler(pkghttp.NewServer(elr.GetLecturerByID()))
	r.Methods(http.MethodGet).Path("/v1.0/lecturers").Handler(pkghttp.NewServer(elr.GetLecturers()))
	r.Methods(http.MethodPut).Path("/v1.0/lecturers/{id}").Handler(pkghttp.NewServer(elr.UpdateLecturer()))
	r.Methods(http.MethodDelete).Path("/v1.0/lecturers/{id}").Handler(pkghttp.NewServer(elr.DeleteLecturer()))
}
