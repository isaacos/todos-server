package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	ErrNotFound = errors.New("models: resource not found")
)

func NewListService(connectionInfo string) (*ListService, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	return &ListService{
		db: db,
	}, nil
}

type ListService struct {
	db *gorm.DB
}

//ByID will look up by the id provided.
//1 - list, nil
//2 -nil, ErrNotFound
//3 - nil, otherError
func (lis *ListService) ByID(id uint) (*List, error) {
	var list List
	err := lis.db.Where("id = ?", id).First(&list).Error
	switch err {
	case nil:
		return &list, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

//Create will create the provided list and backfill data
// like the ID, CreatedAt, and UpdatedAt fields.
func (lis *ListService) Create(list *List) error {
	return lis.db.Create(list).Error
}

//Closes list DB connection
func (lis *ListService) Close() error {
	return lis.db.Close()
}

func (lis *ListService) DestructiveReset() error {
	if err := lis.db.DropTableIfExists(&List{}).Error; err != nil {
		return err
	}
	return lis.AutoMigrate()
}

//AutoMigrate attempts to migrate the Lists table
func (lis *ListService) AutoMigrate() error {
	if err := lis.db.AutoMigrate(&List{}).Error; err != nil {
		return err
	}
	return nil
}

type List struct {
	gorm.Model
	Title string
}
