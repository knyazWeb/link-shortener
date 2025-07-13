package link

import "go/http/pkg/db"

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
