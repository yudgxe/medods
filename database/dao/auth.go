package dao

import (
	"medods/database"
	"medods/database/model"

	"github.com/go-pg/pg/v10/orm"
)

var auth AuthDao = nil

type AuthDao interface {
	CreateOrUpdate(auth *model.Auth) error
	IsExistsUUID(uuid string) (bool, error)
	IsExistsToken(token string) (bool, error)
}

type authDao struct {
	db orm.DB
}

func Auth() AuthDao {
	if auth == nil {
		auth = &authDao{db: database.GetDatabase()}
	}

	return auth
}

// for tests
func SetupAuthDao(dao AuthDao) {
	auth = dao
}

func (this *authDao) CreateOrUpdate(auth *model.Auth) error {
	exists, err := this.IsExistsUUID(auth.Uuid)
	if err != nil {
		return err
	}

	if exists {
		if _, err := this.db.Model(auth).WherePK().Update(); err != nil {
			return err
		}

		return nil
	}

	_, err = this.db.Model(auth).Insert()

	return err
}

func (this *authDao) IsExistsUUID(uuid string) (bool, error) {
	var exists struct {
		Exists bool `pg:"exists"`
	}
	_, err := this.db.QueryOne(&exists, "SELECT EXISTS(SELECT uuid FROM auth WHERE uuid = ?)", uuid)

	return exists.Exists, err
}

func (this *authDao) IsExistsToken(token string) (bool, error) {
	var exists struct {
		Exists bool `pg:"exists"`
	}
	_, err := this.db.QueryOne(&exists, "SELECT EXISTS(SELECT refresh_token FROM auth WHERE refresh_token = ?)", token)

	return exists.Exists, err
}
