package services

import (
	"Brennon-Oliveira/TaskMasterPro-ProjectService/internal/models"
	"Brennon-Oliveira/TaskMasterPro-ProjectService/internal/repositories"
)

type ProjectService struct {
	repo *repositories.ProjectRepository
}

func CreateProjectService(repo *repositories.ProjectRepository) *ProjectService {
	return &ProjectService{repo: repo}
}

func (ps *ProjectService) CreateProject(project *models.Project) error {
	return ps.repo.CreateProject(project)
}

func (ps *ProjectService) GetProjectByID(id uint) (*models.Project, error) {
	return ps.repo.GetProjectByID(id)
}

func (ps *ProjectService) GetProjects() ([]models.Project, error) {
	return ps.repo.GetProjects()
}

func (ps *ProjectService) UpdateProject(project *models.Project) error {
	return ps.repo.UpdateProject(project)
}

func (ps *ProjectService) DeleteProject(id uint) error {
	return ps.repo.DeleteProject(id)
}

