package singleton

import "patterns/singleton/database"

type ProjectService struct {
}

func NewProjectService() *ProjectService {
	return &ProjectService{}
}

func (s ProjectService) GetProjectById() {
	db := database.GetInstance()
	db.Query("project")
}
