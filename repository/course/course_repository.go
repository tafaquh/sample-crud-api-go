package course

import (
	"coba-go/database/schema"
	"coba-go/utils/errors"
)

var (
	CourseRepository courseRepositoryInterface = &courseRepository{}
)

type courseRepositoryInterface interface {
	CourseFileUploadToDB(payload *CourseFileUploadRequest) (*CourseFileUploadResponse, *errors.RestErr)
	CreateCourseToDB(payload *CreateCourseRequest) (*CourseResponse, *errors.RestErr)
	GetCourseByIDFromDB(courseId string) (*CourseResponse, *errors.RestErr)
}
type courseRepository struct{}

func (r *courseRepository) CourseFileUploadToDB(payload *CourseFileUploadRequest) (*CourseFileUploadResponse, *errors.RestErr) {
	fileData := &schema.File{
		CourseId: payload.CourseId,
		Name:     payload.FileName,
		FileUrl:  payload.FileUrl,
	}

	if err := schema.Database.Create(fileData); err.Error != nil {
		return nil, errors.NewInternalServerError("error when trying to save file upload to db: " + err.Error.Error())
	}
	return ToCourseFileUploadResponse(fileData), nil
}

func (r *courseRepository) CreateCourseToDB(payload *CreateCourseRequest) (*CourseResponse, *errors.RestErr) {
	course := &schema.Course{
		CompanyID: payload.CompanyId,
		Name:      payload.Name,
	}

	if err := schema.Database.Create(course); err.Error != nil {
		return nil, errors.NewInternalServerError("error when trying to save course to db: " + err.Error.Error())
	}
	return ToCourseResponse(course), nil
}

func (r *courseRepository) GetCourseByIDFromDB(courseId string) (*CourseResponse, *errors.RestErr) {
	var courseSchema schema.Course
	if err := schema.Database.Preload("Files").Find(&courseSchema, "id = ?", courseId); err.Error != nil {
		return nil, errors.NewInternalServerError("error when trying to get course by id: " + err.Error.Error())
	}

	return ToGetCourseByIDResponse(&courseSchema), nil
}
