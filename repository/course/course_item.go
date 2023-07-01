package course

import (
	"coba-go/database/schema"
	"time"
)

type CreateCourseRequest struct {
	CompanyId string `json:"company_id"`
	Name      string `json:"name"`
}

type JoinCourseRequest struct {
	RoomId  string `json:"room_id"`
	UserId  string `json:"user_id"`
	ByAdmin bool   `json:"by_admin"`
}

type JoinCourseResponse struct {
	IsSuccess bool `json:"is_success"`
}

type CourseResponse struct {
	CourseId  string                      `json:"course_id"`
	CompanyId string                      `json:"company_id"`
	Name      string                      `json:"name"`
	CreatedAt time.Time                   `json:"created_at"`
	CreatedBy string                      `json:"created_by"`
	UpdateAt  time.Time                   `json:"updated_at"`
	UpdatedBy string                      `json:"updated_by"`
	File      []*CourseFileUploadResponse `json:"file"`
}

type CourseFileUploadRequest struct {
	CourseId string `json:"course_id"`
	FileName string `json:"file_name"`
	FileUrl  string `json:"file_url"`
}

type CourseFileUploadResponse struct {
	FileId    string    `json:"file_id"`
	CourseId  string    `json:"course_id"`
	FileName  string    `json:"file_name"`
	FileUrl   string    `json:"file_url"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdateAt  time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
}

func ToCourseFileUploadResponse(file *schema.File) *CourseFileUploadResponse {
	return &CourseFileUploadResponse{
		FileId:    file.ID,
		CourseId:  file.CourseId,
		FileName:  file.Name,
		FileUrl:   file.FileUrl,
		CreatedAt: file.CreatedAt,
		CreatedBy: file.CreatedBy,
		UpdateAt:  file.UpdatedAt,
		UpdatedBy: file.UpdatedBy,
	}
}

func ToCourseResponse(course *schema.Course) *CourseResponse {
	return &CourseResponse{
		CourseId:  course.ID,
		CompanyId: course.CompanyID,
		Name:      course.Name,
		CreatedAt: course.CreatedAt,
		CreatedBy: course.CreatedBy,
		UpdateAt:  course.UpdatedAt,
		UpdatedBy: course.UpdatedBy,
	}
}

func ToGetCourseByIDResponse(courseSchema *schema.Course) *CourseResponse {
	course := CourseResponse{
		CourseId:  courseSchema.ID,
		CompanyId: courseSchema.CompanyID,
		Name:      courseSchema.Name,
		CreatedAt: courseSchema.CreatedAt,
		CreatedBy: courseSchema.CreatedBy,
		UpdateAt:  courseSchema.UpdatedAt,
		UpdatedBy: courseSchema.UpdatedBy,
	}
	for _, file := range courseSchema.Files {
		course.File = append(course.File, &CourseFileUploadResponse{
			FileId:    file.ID,
			CourseId:  file.CourseId,
			FileName:  file.Name,
			FileUrl:   file.FileUrl,
			CreatedAt: file.CreatedAt,
			CreatedBy: file.CreatedBy,
			UpdateAt:  file.UpdatedAt,
			UpdatedBy: file.UpdatedBy,
		})
	}
	return &course
}
