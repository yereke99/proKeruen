package repository

import (
	"context"
	"fmt"
	"qkeruen/models"
)

type UserDB struct {
	DB PgxIface
}

func NewUserRepository(ds PgxIface) UserDB {
	return UserDB{
		DB: ds,
	}
}

func (pool UserDB) Insert(data models.UserRegister) (*models.UserModel, error) {
	var user models.UserModel

	q := `INSERT INTO customer(
		phone,
        firstName,
		lastName,
		ava,
		token
	)VALUES($1,$2,$3,$4,$5)`
	qu := `Select * from customer where token=$1`

	_, err := pool.DB.Exec(context.Background(), q, data.Phone, data.FirstName, data.LastName, data.Avatar, data.Token)
	if err != nil {
		return &models.UserModel{}, err
	}

	row := pool.DB.QueryRow(context.Background(), qu, data.Token)

	if err := row.Scan(
		&user.Id,
		&user.Phone,
		&user.FirstName,
		&user.LastName,
		&user.Avatar,
		&user.Token,
	); err != nil {
		return &models.UserModel{}, err
	}

	user.Type = "user"
	return &user, nil
}

func (pool UserDB) CheckTokenUser(token string) (*models.UserModel, error) {
	q := `Select * From customer Where token=$1`
	rows := pool.DB.QueryRow(context.Background(), q, token)

	// you must to change here to user model !
	var modelUser models.UserModel

	err := rows.Scan(
		&modelUser.Id,
		&modelUser.Phone,
		&modelUser.FirstName,
		&modelUser.LastName,
		&modelUser.Avatar,
		&modelUser.Token,
	)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	modelUser.Type = "user"

	return &modelUser, nil

}

func (pool UserDB) Update(update models.UserModel) (*models.UserModel, error) {

	if update.Phone != "" {
		q := `Update customer SET phone=$1 WHERE token=$2;`
		_, err := pool.DB.Exec(context.Background(), q, update.Phone, update.Token)

		if err != nil {
			return nil, err
		}
	}

	if update.FirstName != "" {
		q := `Update customer SET firstName=$1 WHERE token=$2;`
		_, err := pool.DB.Exec(context.Background(), q, update.FirstName, update.Token)

		if err != nil {
			return nil, err
		}
	}

	if update.FirstName != "" {
		q := `Update customer SET lastName=$1 WHERE token=$2;`
		_, err := pool.DB.Exec(context.Background(), q, update.LastName, update.Token)

		if err != nil {
			return nil, err
		}
	}

	if update.Avatar != "" {
		q := `Update customer SET ava=$1 WHERE token=$2;`
		_, err := pool.DB.Exec(context.Background(), q, update.Avatar, update.Token)

		if err != nil {
			return nil, err
		}
	}

	q := `Select * From customer WHERE token=$1;`

	row := pool.DB.QueryRow(context.Background(), q, update.Token)

	u := new(models.UserModel)
	err := row.Scan(&u.Id, &u.Phone, &u.FirstName, &u.LastName, &u.Avatar, &u.Token)

	if err != nil {
		return nil, err
	}

	return u, nil

}

func (pool UserDB) Delete(id int64) error {
	q := `Delete From customer WHERE id=$1;`

	_, err := pool.DB.Exec(context.Background(), q, id)

	if err != nil {
		return err
	}

	return nil
}
