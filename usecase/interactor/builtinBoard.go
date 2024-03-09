package interactor

import (
	"music-app/adapter/database/model"
	"music-app/usecase/port"
)

type BuiltinBoardUseCase struct {
	ulid             port.ULID
	builtinBoardRepo port.BuiltinBoardRepository
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

func (u *BuiltinBoardUseCase) Register(register BuiltinBoardRegister) (model.BuiltinBoard, error) {

	newBuiltinBoard := model.BuiltinBoard{
		BuiltinBoardId: u.ulid.GenerateID(),
		ImageUrl:       register.ImageUrl,
		Date:           register.Date,
		LocationId:     register.LocationId,
		ArtistId:       register.ArtistId,
	}

	ret, err := u.builtinBoardRepo.FindByID(newBuiltinBoard.BuiltinBoardId)
	return ret, err
}
