package link

import (
	"errors"
	"go/http/pkg/db"

	"gorm.io/gorm/clause"
)

type LinkRepository struct {
	Database *db.Db
}

func NewLinkRepository(database *db.Db) *LinkRepository {
	return &LinkRepository{
		Database: database,
	}
}

func (repo *LinkRepository) Create(link *Link) (*Link, error) {
	key := repo.Database.DB.Create(link)
	if key.Error != nil {
		return nil, key.Error
	}
	return link, nil
}

func (repo *LinkRepository) FindLinkByHash(hash string) (*Link, error) {
	var link Link
	result := repo.Database.DB.First(&link, "hash = ?", hash)

	if result.Error != nil {
		return nil, result.Error
	}

	return &link, nil
}

func (repo *LinkRepository) Update(link *Link) (*Link, error) {
	result := repo.Database.DB.Clauses(clause.Returning{}).Updates(link)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("no id")
	}

	return link, nil
}

func (repo *LinkRepository) Delete(id uint) error {
	result := repo.Database.DB.Delete(&Link{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil

}

func (repo *LinkRepository) FindLinkById(Id uint) (*Link, error) {
	var link Link
	result := repo.Database.DB.First(&link, "id = ?", Id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &link, nil
}
