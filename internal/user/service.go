package user

import (
	"database/sql"
	"github.com/flejz/cp-server/internal/repository"
)

type User struct {
	usr  string
	pwd  string
	salt string
}

type UserService struct {
	Repository repository.Repository
}

func (self UserService) Get(usr string) (*User, error) {
	fieldList := []string{"pwd", "salt"}
	whereMap := map[string]interface{}{
		"usr": usr,
	}

	user := &User{usr, "", ""}
	row := self.Repository.QueryRow(fieldList, whereMap)

	if err := row.Err(); err != nil {
		return nil, err
	}

	if err := row.Scan(&user.pwd, &user.salt); err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, nil
		default:
			return nil, err
		}
	}

	return user, nil
}

func (self UserService) Register(usr, pwd string) error {
	user, err := self.Get(usr)
	if err != nil {
		return err
	}

	if user != nil {
		return ErrInvalidCredentials
	}

	salt := Salt()
	hash := Hash(pwd, salt)

	fieldMap := map[string]interface{}{
		"usr":  usr,
		"pwd":  hash,
		"salt": salt,
	}

	if _, err = self.Repository.Insert(fieldMap); err != nil {
		return err
	}

	return nil
}

func (self UserService) Validate(usr, pwd string) error {
	user, err := self.Get(usr)
	if err != nil {
		return err
	}

	if user == nil {
		return ErrInvalidCredentials
	}

	if hash := Hash(pwd, user.salt); hash != user.pwd {
		return ErrInvalidCredentials
	}

	return nil
}
