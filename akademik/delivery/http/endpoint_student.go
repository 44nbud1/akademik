package http

import (
	"context"
	"fmt"
	"github.com/44nbud1/akademik/akademik/domain/model"
	"github.com/44nbud1/akademik/akademik/util"
	"github.com/44nbud1/akademik/pkg/delivery/pkghttp"
	"github.com/44nbud1/akademik/pkg/delivery/pkghttp/commonresp"
	"github.com/44nbud1/akademik/pkg/pkgservice"
	"github.com/bwmarrin/snowflake"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type EndpointStudent interface {
	CreateStudent(snowflakeNode *snowflake.Node) pkghttp.Endpoint
	GetStudentByID() pkghttp.Endpoint
	GetStudents() pkghttp.Endpoint
	UpdateStudent() pkghttp.Endpoint
	DeleteStudent() pkghttp.Endpoint
}

func (e endpointAcademic) CreateStudent(snowflakeNode *snowflake.Node) pkghttp.Endpoint {
	return func(ctx context.Context, r *http.Request) (response interface{}) {

		var student Students
		// Decode request from partner
		if err := e.decodeRequest(r, &student); err != nil {
			e.logger.Error(fmt.Sprintf("[Endpoint_student] Failed to decode request, err: %v", err))

			return commonresp.NewResponseData(commonresp.NewResponseCode(pkgservice.NewErrorService(pkgservice.ErrBadRequest)), student)

		}

		var gender util.Gender
		enum, err := gender.Enum(student.Gender)
		if err != nil {
			return commonresp.NewResponseData(commonresp.NewResponseCode(pkgservice.NewErrorService(pkgservice.ErrBadRequest)), student)
		}

		epochNow := time.Now().UnixMilli()
		result, errx := e.studentService.CreateStudent(ctx, model.Students{
			ID:        snowflakeNode.Generate().String(),
			Name:      student.Name,
			Address:   student.Address,
			Gender:    enum.String(),
			Email:     student.Email,
			Phone:     student.Phone,
			CreatedAt: epochNow,
			UpdatedAt: epochNow,
		})

		return commonresp.NewResponseData(commonresp.NewResponseCode(errx), e.getStudentResponse(result))
	}

}

func (e endpointAcademic) GetStudentByID() pkghttp.Endpoint {
	return func(ctx context.Context, r *http.Request) (response interface{}) {
		id := mux.Vars(r)["id"]
		result, errx := e.studentService.GetStudentByID(ctx, id)
		return commonresp.NewResponseData(commonresp.NewResponseCode(errx), e.getStudentResponse(result))
	}
}

func (e endpointAcademic) GetStudents() pkghttp.Endpoint {
	return func(ctx context.Context, r *http.Request) (response interface{}) {
		students, errx := e.studentService.GetStudents(ctx)

		var studentArray []Students
		if students != nil {
			for _, s := range *students {
				studentArray = append(studentArray, Students{
					ID:        s.ID,
					Name:      s.Name,
					Address:   s.Address,
					Gender:    s.Gender,
					Email:     s.Email,
					Phone:     s.Phone,
					CreatedAt: s.CreatedAt,
					UpdatedAt: s.UpdatedAt,
				})
			}
		}

		return commonresp.NewResponseData(commonresp.NewResponseCode(errx), &studentArray)
	}
}

func (e endpointAcademic) UpdateStudent() pkghttp.Endpoint {
	return func(ctx context.Context, r *http.Request) (response interface{}) {
		var student Students
		// Decode request from partner
		if err := e.decodeRequest(r, &student); err != nil {
			e.logger.Error(fmt.Sprintf("[Endpoint_student] Failed to decode request, err: %v", err))

			return commonresp.NewResponseData(commonresp.NewResponseCode(pkgservice.NewErrorService(pkgservice.ErrBadRequest)), student)

		}

		var gender util.Gender
		enum, err := gender.Enum(student.Gender)
		if err != nil {
			return commonresp.NewResponseData(commonresp.NewResponseCode(pkgservice.NewErrorService(pkgservice.ErrBadRequest)), student)
		}

		id := mux.Vars(r)["id"]
		result, errx := e.studentService.UpdateStudent(ctx, model.Students{
			Name:    student.Name,
			Address: student.Address,
			Gender:  enum.String(),
			Email:   student.Email,
			Phone:   student.Phone,
		}, id)

		return commonresp.NewResponseData(commonresp.NewResponseCode(errx), e.getStudentResponse(result))
	}
}

func (e endpointAcademic) DeleteStudent() pkghttp.Endpoint {
	return func(ctx context.Context, r *http.Request) (response interface{}) {
		id := mux.Vars(r)["id"]
		result, errx := e.studentService.DeleteStudent(ctx, id)
		return commonresp.NewResponseData(commonresp.NewResponseCode(errx), e.getStudentResponse(result))
	}
}

func (e endpointAcademic) getStudentResponse(result *model.Students) *Students {

	if result != nil {
		return &Students{
			ID:        result.ID,
			Name:      result.Name,
			Address:   result.Address,
			Gender:    result.Gender,
			Email:     result.Email,
			Phone:     result.Phone,
			CreatedAt: result.CreatedAt,
			UpdatedAt: result.UpdatedAt,
		}
	}

	return &Students{}
}
