package port

import "music-app/adapter/database/model"

type LocationRepository interface {
	FindByID(LocationId string) (model.Location, error)
}
