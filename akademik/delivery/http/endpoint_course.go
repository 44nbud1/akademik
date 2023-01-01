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
	"time"
)

type EndpointCourse interface {
	CreateCourse(snowflakeNode *snowflake.Node) pkghttp.Endpoint
	GetCourseByID() pkghttp.Endpoint
	GetCourses() pkghttp.Endpoint
	UpdateCourse() pkghttp.Endpoint
	DeleteCourse() pkghttp.Endpoint
}

func (e endpointAcademic) CreateCourse(snowflakeNode *snowflake.Node) pkghttp.Endpoint {
	return func(ctx context.Context, r *http.Request) (response interface{}) {

		var course Course
		// Decode request from partner
		if err := e.decodeRequest(r, &course); err != nil {
			e.logger.Error(fmt.Sprintf("[Endpoint_course] Failed to decode request, err: %v", err))

			return commonresp.NewResponseData(commonresp.NewResponseCode(pkgservice.NewErrorService(pkgservice.ErrBadRequest)), course)

		}

		epochNow := time.Now().UnixMilli()
		result, errx := e.courseService.CreateCourse(ctx, model.Course{
			ID:            snowflakeNode.Generate().String(),
			Name:          course.Name,
			CreditsCourse: course.CreditsCourse,
			CreatedAt:     epochNow,
			UpdatedAt:     epochNow,
		})

		return commonresp.NewResponseData(commonresp.NewResponseCode(errx), e.getCourseResponse(result))
	}

}

func (e endpointAcademic) GetCourseByID() pkghttp.Endpoint {
	return func(ctx context.Context, r *http.Request) (response interface{}) {
		id := mux.Vars(r)["id"]
		course, errx := e.courseService.GetCourseByID(ctx, id)
		return commonresp.NewResponseData(commonresp.NewResponseCode(errx), e.getCourseResponse(course))
	}
}

func (e endpointAcademic) GetCourses() pkghttp.Endpoint {
	return func(ctx context.Context, r *http.Request) (response interface{}) {
		course, errx := e.courseService.GetCourses(ctx)
		var courses []Course
		if course != nil {
			for _, s := range *course {
				courses = append(courses, Course{
					ID:            s.ID,
					Name:          s.Name,
					CreditsCourse: s.CreditsCourse,
					CreatedAt:     s.CreatedAt,
					UpdatedAt:     s.UpdatedAt,
				})
			}
		}

		return commonresp.NewResponseData(commonresp.NewResponseCode(errx), &courses)
	}
}

func (e endpointAcademic) UpdateCourse() pkghttp.Endpoint {
	return func(ctx context.Context, r *http.Request) (response interface{}) {
		var course Course
		// Decode request from partner
		if err := e.decodeRequest(r, &course); err != nil {
			e.logger.Error(fmt.Sprintf("[Endpoint_course] Failed to decode request, err: %v", err))

			return commonresp.NewResponseData(commonresp.NewResponseCode(pkgservice.NewErrorService(pkgservice.ErrBadRequest)), course)
		}

		id := mux.Vars(r)["id"]
		result, errx := e.courseService.UpdateCourse(ctx, model.Course{
			Name: course.Name,
		}, id)
		return commonresp.NewResponseData(commonresp.NewResponseCode(errx), e.getCourseResponse(result))
	}
}

func (e endpointAcademic) DeleteCourse() pkghttp.Endpoint {
	return func(ctx context.Context, r *http.Request) (response interface{}) {
		id := mux.Vars(r)["id"]
		course, errx := e.courseService.DeleteCourse(ctx, id)
		return commonresp.NewResponseData(commonresp.NewResponseCode(errx), e.getCourseResponse(course))
	}
}

func (e endpointAcademic) getCourseResponse(course *model.Course) *Course {

	if course != nil {
		return &Course{
			ID:            course.ID,
			Name:          course.Name,
			CreditsCourse: course.CreditsCourse,
			CreatedAt:     course.CreatedAt,
			UpdatedAt:     course.UpdatedAt,
		}
	}

	return &Course{}
}
