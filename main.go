package main

import (
	"fmt"
	"music-app/adapter/database"
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
	if err != nil {
		fmt.Println("error:", err)
	}
}
