package testdata

import (
	"os"
	"server/database/repository"
	"server/tests"
)

type TestdataService struct {
	db *repository.Database
}

func NewTestdataService(db *repository.Database) *TestdataService {
	return &TestdataService{db: db}
}

func (tds *TestdataService) Init() error {
	insertTestData := os.Getenv("InsertTestData")
	if insertTestData == "true" {
		err := tds.insertTestData()
		if err != nil {
			panic(err)
		}
	}
	return nil
}

func (tds *TestdataService) insertTestData() error {
	numBoards, err := tds.db.BoardRepository.Count()
	if err != nil {
		return err
	}
	if numBoards == 0 {
		err := tds.db.BoardRepository.Create(tests.TestBoard)
		if err != nil {
			return err
		}
	}

	numUsers, err := tds.db.UserRepository.Count()
	if err != nil {
		return err
	}
	if numUsers <= 1 {
		err := tests.TestAdminUser.HashPassword()
		if err != nil {
			return err
		}
		err = tests.TestUser.HashPassword()
		if err != nil {
			return err
		}
		err = tds.db.UserRepository.Create(tests.TestAdminUser)
		if err != nil {
			return err
		}
		err = tds.db.UserRepository.Create(tests.TestUser)
		if err != nil {
			return err
		}
	}
	return nil
}
