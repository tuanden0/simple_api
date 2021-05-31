package services

import (
	"context"
	"fmt"

	"github.com/tuanden0/simple_api/internal/api"
	"github.com/tuanden0/simple_api/internal/models"
	"github.com/tuanden0/simple_api/internal/repository"
	"google.golang.org/protobuf/types/known/emptypb"
)

type StudentGRPCService struct {
	student repository.StudentObject
	api.UnimplementedStudentServiceServer
}

func NewStudentGRPCService(s repository.StudentObject) *StudentGRPCService {
	return &StudentGRPCService{student: s}
}

func mutateModelToStudentResponse(m *models.Student) *api.Student {
	return &api.Student{
		Id:   uint32(m.Id),
		Name: m.Name,
		Gpa:  float32(m.GPA),
	}
}

func mutateListModelToStudentResponse(s []*models.Student) *api.ListResponse {

	res := &api.ListResponse{}
	for _, v := range s {
		res.Student = append(res.Student, mutateModelToStudentResponse(v))
	}
	return res
}

func (srv *StudentGRPCService) Create(ctx context.Context, in *api.CreateRequest) (*api.Student, error) {

	studentCreate := &models.Student{
		Name: in.Name,
		GPA:  float64(in.Gpa),
	}
	if err := srv.student.Create(studentCreate); err != nil {
		return nil, err
	}

	return mutateModelToStudentResponse(studentCreate), nil
}

func (srv *StudentGRPCService) Retrieve(ctx context.Context, in *api.RetrieveRequest) (*api.Student, error) {

	student, err := srv.student.Retrieve(fmt.Sprint(in.Id))
	if err != nil {
		return nil, err
	}
	return mutateModelToStudentResponse(student), nil
}

func (srv *StudentGRPCService) Update(ctx context.Context, in *api.UpdateRequest) (*api.Student, error) {

	studentUpdate := models.Student{
		Name: in.Name,
		GPA:  float64(in.Gpa),
	}
	student, err := srv.student.Update(fmt.Sprint(in.Id), studentUpdate)
	if err != nil {
		return nil, err
	}
	return mutateModelToStudentResponse(student), nil
}

func (srv *StudentGRPCService) Delete(ctx context.Context, in *api.DeleteRequest) (*emptypb.Empty, error) {
	err := srv.student.Delete(fmt.Sprint(in.Id))
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (srv *StudentGRPCService) List(ctx context.Context, in *emptypb.Empty) (*api.ListResponse, error) {

	page := repository.NewPagination(1, 5)
	sort := repository.NewSort("id", "asc")

	students, err := srv.student.List(*page, *sort)
	if err != nil {
		return nil, err
	}
	return mutateListModelToStudentResponse(students), nil
}
