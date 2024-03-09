package repository

import (
	"music-app/adapter/database/model"
	"music-app/usecase/port"

	"gorm.io/gorm"
)

type LocationRepository struct {
	db   *gorm.DB
	ulid port.ULID
}

func NewLocationRepository(
	db *gorm.DB,
	ulid port.ULID,
) port.LocationRepository {
	return &LocationRepository{db: db, ulid: ulid}
}

func (r *LocationRepository) FindByID(LocationId string) (model.Location, error) {
	location := &model.Location{}
	err := r.db.
		Model(&model.Location{}).
		Where("location_id = ?", LocationId).
		First(location).
		Error
	if err != nil {
		return model.Location{}, err
	}
	return model.Location{}, nil
}
