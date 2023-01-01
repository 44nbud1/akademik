package http

import (
	"context"
	"fmt"
	"github.com/44nbud1/akademik/akademik/domain/model"
	"github.com/44nbud1/akademik/pkg/delivery/pkghttp"
	"github.com/44nbud1/akademik/pkg/delivery/pkghttp/commonresp"
	"github.com/44nbud1/akademik/pkg/pkgservice"
	"github.com/bwmarrin/snowflake"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	"time"
)

type EndpointLecture interface {
	CreateLecture(snowflakeNode *snowflake.Node) pkghttp.Endpoint
	GetLectureByID() pkghttp.Endpoint
	GetLectures() pkghttp.Endpoint
	UpdateLecture() pkghttp.Endpoint
	DeleteLecture() pkghttp.Endpoint
	UpdateName() pkghttp.Endpoint
}

func (e endpointAcademic) CreateLecture(snowflakeNode *snowflake.Node) pkghttp.Endpoint {
	return func(ctx context.Context, r *http.Request) (response interface{}) {

		var lecture Lecture
		if err := e.decodeRequest(r, &lecture); err != nil {
			e.logger.Error(fmt.Sprintf("[Endpoint_lecture] Failed to decode request, err: %v", err))

			return commonresp.NewResponseData(commonresp.NewResponseCode(pkgservice.NewErrorService(pkgservice.ErrBadRequest)), lecture)
		}

		// get student
		student, errx := e.studentService.GetStudentByID(ctx, lecture.StudentID)
		if errx != nil {
			return commonresp.NewResponseData(commonresp.NewResponseCode(errx), lecture)
		}

		// get lecturer
		lecturer, errx := e.lecturerService.GetLecturerByID(ctx, lecture.LecturerID)
		if errx != nil {
			return commonresp.NewResponseData(commonresp.NewResponseCode(errx), lecture)
		}

		// get course
		course, errx := e.courseService.GetCourseByID(ctx, lecture.CourseID)
		if errx != nil {
			return commonresp.NewResponseData(commonresp.NewResponseCode(errx), lecture)
		}

		epochNow := time.Now().UnixMilli()
		result, errx := e.lectureService.CreateLecture(ctx, model.Lecture{
			ID:         snowflakeNode.Generate().String(),
			StudentID:  student.ID,
			LecturerID: lecturer.ID,
			CourseID:   course.ID,
			Score:      lecture.Score,
			Name:       getName(lecture.Name, student.Name),
			CreatedAt:  epochNow,
			UpdatedAt:  epochNow,
		})

		return commonresp.NewResponseData(commonresp.NewResponseCode(errx), e.getLectureResponse(result))
	}
}

func getName(reqName string, dataName string) string {

	if !strings.EqualFold(reqName, "") {
		return reqName
	}

	return dataName
}

func (e endpointAcademic) GetLectureByID() pkghttp.Endpoint {
	return func(ctx context.Context, r *http.Request) (response interface{}) {
		id := mux.Vars(r)["id"]
		result, errx := e.lectureService.GetLectureByID(ctx, id)
		return commonresp.NewResponseData(commonresp.NewResponseCode(errx), e.getLectureResponse(result))
	}
}

func (e endpointAcademic) GetLectures() pkghttp.Endpoint {
	return func(ctx context.Context, r *http.Request) (response interface{}) {
		lectures, errx := e.lectureService.GetLectures(ctx)
		var lectureArray []Lecture
		if lectures != nil {
			for _, s := range *lectures {
				lectureArray = append(lectureArray, Lecture{
					ID:         s.ID,
					StudentID:  s.StudentID,
					LecturerID: s.LecturerID,
					Score:      s.Score,
					Name:       s.Name,
					CourseID:   s.CourseID,
					CreatedAt:  s.CreatedAt,
					UpdatedAt:  s.UpdatedAt,
				})
			}
		}

		return commonresp.NewResponseData(commonresp.NewResponseCode(errx), &lectureArray)
	}
}

func (e endpointAcademic) UpdateLecture() pkghttp.Endpoint {
	return func(ctx context.Context, r *http.Request) (response interface{}) {
		var lecture Lecture
		// Decode request from partner
		if err := e.decodeRequest(r, &lecture); err != nil {
			e.logger.Error(fmt.Sprintf("[Endpoint_lecture] Failed to decode request, err: %v", err))

			return commonresp.NewResponseData(commonresp.NewResponseCode(pkgservice.NewErrorService(pkgservice.ErrBadRequest)), lecture)
		}

		id := mux.Vars(r)["id"]
		result, errx := e.lectureService.UpdateLecture(ctx, model.Lecture{
			Score: lecture.Score,
		}, id)
		return commonresp.NewResponseData(commonresp.NewResponseCode(errx), e.getLectureResponse(result))
	}
}

func (e endpointAcademic) DeleteLecture() pkghttp.Endpoint {
	return func(ctx context.Context, r *http.Request) (response interface{}) {
		id := mux.Vars(r)["id"]
		result, errx := e.lectureService.DeleteLecture(ctx, id)
		return commonresp.NewResponseData(commonresp.NewResponseCode(errx), e.getLectureResponse(result))
	}
}

func (e endpointAcademic) UpdateName() pkghttp.Endpoint {
	return func(ctx context.Context, r *http.Request) (response interface{}) {
		var lecture Lecture
		// Decode request from partner
		if err := e.decodeRequest(r, &lecture); err != nil {
			e.logger.Error(fmt.Sprintf("[Endpoint_lecture] Failed to decode request, err: %v", err))

			return commonresp.NewResponseData(commonresp.NewResponseCode(pkgservice.NewErrorService(pkgservice.ErrBadRequest)), lecture)
		}

		id := mux.Vars(r)["id"]
		result, errx := e.lectureService.UpdateName(ctx, model.Lecture{Name: lecture.Name}, id)
		return commonresp.NewResponseData(commonresp.NewResponseCode(errx), e.getLectureResponse(result))
	}
}

func (e endpointAcademic) getLectureResponse(result *model.Lecture) *Lecture {

	if result != nil {
		return &Lecture{
			ID:         result.ID,
			StudentID:  result.StudentID,
			LecturerID: result.LecturerID,
			Score:      result.Score,
			Name:       result.Name,
			CourseID:   result.CourseID,
			CreatedAt:  result.CreatedAt,
			UpdatedAt:  result.UpdatedAt,
		}
	}

	return &Lecture{}
}
