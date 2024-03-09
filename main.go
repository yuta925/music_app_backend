package main

import (
	"fmt"
	"music-app/adapter/api/router.go"
	"music-app/adapter/authentication"
	"music-app/adapter/clock"
	"music-app/adapter/database"
	"music-app/adapter/database/repository"
	"music-app/adapter/ulid"
	"music-app/usecase/interactor"
	"os"
)

func main() {

	db, err := database.NewMySQLDB()
	if err != nil {
		fmt.Errorf("エラーが発生しました: %v", err)
		return
	}

	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			fmt.Errorf("エラーが発生しました: %v", err)
		}
		err = sqlDB.Close()
		if err != nil {
			fmt.Errorf("エラーが発生しました: %v", err)
		}
	}()

	err = database.Migrate(db)

	userAuth := authentication.NewUserAuth()

	ulidDriver := ulid.NewULID()
	clockDriver := clock.New()

	userRepo := repository.NewUserRepository(db, ulidDriver)
	builtinBoardRepo := repository.NewBuiltinBoardRepository(db, ulidDriver)

	userUC := interactor.NewUserUseCase(clockDriver, ulidDriver, userAuth, userRepo)
	builtinBoardUC := interactor.NewBuiltinBoardUseCase(ulidDriver, builtinBoardRepo)

	s := router.NewServer(
		userUC,
		builtinBoardUC,
	)
	if err := s.Start(":80"); err != nil {
		fmt.Printf("エラーが発生しました: %v\n", err)
		os.Exit(1)
	}

}
