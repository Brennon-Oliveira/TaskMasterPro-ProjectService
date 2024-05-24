package repositories

import (
	"Brennon-Oliveira/TaskMasterPro-ProjectService/internal/models"

	"github.com/jinzhu/gorm"
)

type ProjectRepository struct {
	db *gorm.DB
}

func CreateProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}

func (pr *ProjectRepository) CreateProject(project *models.Project) error {
	return pr.db.Create(project).Error
}

func (pr *ProjectRepository) GetProjectByID(id uint) (*models.Project, error) {
	project := &models.Project{}
	err := pr.db.First(project, id).Error
	return project, err
}

func (pr *ProjectRepository) GetProjects() ([]models.Project, error) {
	var projects []models.Project
	err := pr.db.Find(&projects).Error
	return projects, err
}

func (pr *ProjectRepository) UpdateProject(project *models.Project) error {
	return pr.db.Save(project).Error
}

func (pr *ProjectRepository) DeleteProject(id uint) error {
	return pr.db.Delete(&models.Project{}, id).Error
}
