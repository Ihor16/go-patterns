package singleton

import "patterns/singleton/database"

type IssueService struct {
}

func NewIssueService() *IssueService {
	return &IssueService{}
}

func (s IssueService) GetIssueById() {
	db := database.GetInstance()
	db.Query("issue")
}
