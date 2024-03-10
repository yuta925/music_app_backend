package interactor

import (
	"log"
	"music-app/adapter/database/model"
	"music-app/usecase/port"
	"time"
)

type BuiltinBoardUseCase struct {
	ulid             port.ULID
	builtinBoardRepo port.BuiltinBoardRepository
}
type BuiltinBoardSearch struct {
	ArtistId   string
	LocationId string
	Date       time.Time
	Skip       int
	Limit      int
}

func NewBuiltinBoardUseCase(
	ulid port.ULID,
	builtinBoardRepo port.BuiltinBoardRepository,
) IBuiltinBoardUseCase {
	return &BuiltinBoardUseCase{
		ulid:             ulid,
		builtinBoardRepo: builtinBoardRepo,
	}
}

func (u *BuiltinBoardUseCase) Register(register BuiltinBoardRegister) model.BuiltinBoard {

	newBuiltinBoard := model.BuiltinBoard{
		BuiltinBoardId: u.ulid.GenerateID(),
		ImageUrl:       register.ImageUrl,
		Date:           register.Date,
		LocationId:     register.LocationId,
		ArtistId:       register.ArtistId,
	}
	err := u.builtinBoardRepo.Create(newBuiltinBoard)
	if err != nil {
		log.Println("Error:", err)
	}

	return newBuiltinBoard
}

func (u *BuiltinBoardUseCase) Search(builtinBoardSearch BuiltinBoardSearch) ([]model.BuiltinBoard, error) {
	return u.builtinBoardRepo.Search(port.BuiltinBoardSearchQuery{
		ArtistId:   builtinBoardSearch.ArtistId,
		LocationId: builtinBoardSearch.LocationId,
		Date:       builtinBoardSearch.Date,
		Skip:       builtinBoardSearch.Skip,
		Limit:      builtinBoardSearch.Limit,
	})
}
