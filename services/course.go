package services

import (
	"coba-go/repository/course"
	"coba-go/utils/errors"
)

var (
	CourseService courseServiceInterface = &courseService{}
)

type courseService struct {
}

type courseServiceInterface interface {
	CourseFileUpload(*course.CourseFileUploadRequest) (interface{}, *errors.RestErr)
	CreateCourse(*course.CreateCourseRequest) (interface{}, *errors.RestErr)
	GetCourseByID(string) (interface{}, *errors.RestErr)
}

func (s *courseService) CourseFileUpload(payload *course.CourseFileUploadRequest) (interface{}, *errors.RestErr) {
	result, err := course.CourseRepository.CourseFileUploadToDB(payload)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *courseService) CreateCourse(payload *course.CreateCourseRequest) (interface{}, *errors.RestErr) {
	result, err := course.CourseRepository.CreateCourseToDB(payload)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *courseService) GetCourseByID(courseId string) (interface{}, *errors.RestErr) {
	res, err := course.CourseRepository.GetCourseByIDFromDB(courseId)
	if err != nil {
		return nil, err
	}
	return res, nil
}
