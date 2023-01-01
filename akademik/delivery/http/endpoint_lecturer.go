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

type EndpointLecturer interface {
	CreateLecturer(snowflakeNode *snowflake.Node) pkghttp.Endpoint
	GetLecturerByID() pkghttp.Endpoint
	GetLecturers() pkghttp.Endpoint
	UpdateLecturer() pkghttp.Endpoint
	DeleteLecturer() pkghttp.Endpoint
}

func (e endpointAcademic) CreateLecturer(snowflakeNode *snowflake.Node) pkghttp.Endpoint {
	return func(ctx context.Context, r *http.Request) (response interface{}) {

		var lecturer Lecturer
		// Decode request from partner
		if err := e.decodeRequest(r, &lecturer); err != nil {
			e.logger.Error(fmt.Sprintf("[Endpoint_lecturer] Failed to decode request, err: %v", err))

			return commonresp.NewResponseData(commonresp.NewResponseCode(pkgservice.NewErrorService(pkgservice.ErrBadRequest)), lecturer)
		}

		epochNow := time.Now().UnixMilli()
		result, errx := e.lecturerService.CreateLecturer(ctx, model.Lecturer{
			ID:        snowflakeNode.Generate().String(),
			Name:      lecturer.Name,
			CreatedAt: epochNow,
			UpdatedAt: epochNow,
		})

		return commonresp.NewResponseData(commonresp.NewResponseCode(errx), e.getLecturerResponse(result))
	}

}

func (e endpointAcademic) GetLecturerByID() pkghttp.Endpoint {
	return func(ctx context.Context, r *http.Request) (response interface{}) {
		id := mux.Vars(r)["id"]
		result, errx := e.lecturerService.GetLecturerByID(ctx, id)
		return commonresp.NewResponseData(commonresp.NewResponseCode(errx), e.getLecturerResponse(result))
	}
}

func (e endpointAcademic) GetLecturers() pkghttp.Endpoint {
	return func(ctx context.Context, r *http.Request) (response interface{}) {
		lecturers, errx := e.lecturerService.GetLecturers(ctx)

		var lectureArray []Lecture
		if lecturers != nil {
			for _, s := range *lecturers {
				lectureArray = append(lectureArray, Lecture{
					StudentID: s.ID,
					Name:      s.Name,
					CreatedAt: s.CreatedAt,
					UpdatedAt: s.UpdatedAt,
				})
			}
		}

		return commonresp.NewResponseData(commonresp.NewResponseCode(errx), &lectureArray)
	}
}

func (e endpointAcademic) UpdateLecturer() pkghttp.Endpoint {
	return func(ctx context.Context, r *http.Request) (response interface{}) {
		var lecturer Lecturer
		// Decode request from partner
		if err := e.decodeRequest(r, &lecturer); err != nil {
			e.logger.Error(fmt.Sprintf("[Endpoint_lecturer] Failed to decode request, err: %v", err))

			return commonresp.NewResponseData(commonresp.NewResponseCode(pkgservice.NewErrorService(pkgservice.ErrBadRequest)), lecturer)

		}

		id := mux.Vars(r)["id"]
		result, errx := e.lecturerService.UpdateLecturer(ctx, model.Lecturer{
			Name: lecturer.Name,
		}, id)

		return commonresp.NewResponseData(commonresp.NewResponseCode(errx), e.getLecturerResponse(result))
	}
}

func (e endpointAcademic) DeleteLecturer() pkghttp.Endpoint {
	return func(ctx context.Context, r *http.Request) (response interface{}) {
		id := mux.Vars(r)["id"]
		result, errx := e.lecturerService.DeleteLecturer(ctx, id)
		return commonresp.NewResponseData(commonresp.NewResponseCode(errx), e.getLecturerResponse(result))
	}
}

func (e endpointAcademic) getLecturerResponse(result *model.Lecturer) *Lecturer {

	if result != nil {
		return &Lecturer{
			ID:        result.ID,
			Name:      result.Name,
			CreatedAt: result.CreatedAt,
			UpdatedAt: result.UpdatedAt,
		}
	}

	return &Lecturer{}
}
