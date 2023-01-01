package http

import (
	"encoding/json"
	"github.com/44nbud1/akademik/akademik/domain"
	"go.uber.org/zap"
	"net/http"
)

type endpointAcademic struct {
	studentService  domain.StudentService
	courseService   domain.CourseService
	lectureService  domain.LectureService
	lecturerService domain.LecturerService
	logger          *zap.Logger
}

func NewEndpointAcademic(studentService domain.StudentService, courseService domain.CourseService, lectureService domain.LectureService, lecturerService domain.LecturerService, logger *zap.Logger) *endpointAcademic {
	return &endpointAcademic{
		courseService:   courseService,
		lectureService:  lectureService,
		lecturerService: lecturerService,
		studentService:  studentService,
		logger:          logger,
	}
}

func (e *endpointAcademic) decodeRequest(r *http.Request, i interface{}) (err error) {
	if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
		return err
	}
	return nil
}
