package dbcrud

import (
	"database/sql"
	"school-management/services/models"

	"github.com/google/uuid"
)

type CrudEntry struct {
	dba *sql.DB
}

type IDbCrud interface {
	Create(payload models.School) (*uuid.UUID, error)
	Update(schoolId string, payload models.School) (*bool, error)
	Get() (*[]models.School, error)
	GetById(schoolId string) (*[]models.School, error)
	Delete(schoolId string) (*bool, error)
}

func NewDbCrudEntry(dba *sql.DB) IDbCrud {
	return &CrudEntry{dba: dba}
}
