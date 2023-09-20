package repository

import (
	"context"
	"qkeruen/models"
)

type SecurityDB struct {
	DB PgxIface
}

func NewSecurityService(ds PgxIface) SecurityDB {
	security := SecurityDB{
		DB: ds,
	}
	return security
}

func (pool *SecurityDB) Insert(data models.Security) error {
	q := `INSERT INTO security(
		    userId,
		    firstName,
		    lastName, 
		    A,
		    B,
		    fiod,
		    phone, 
		    carNumber,
			timeStart,
			status
		)VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	_, err := pool.DB.Exec(
		context.Background(),
		q,
		data.UserId,
		data.FirsrtName,
		data.LastName,
		data.From,
		data.To,
		data.FioD,
		data.Phone,
		data.CarNumber,
		data.TimeStart,
		data.Check,
	)
	if err != nil {
		return err
	}

	return nil
}

func (pool *SecurityDB) Finish(id int64, time string) (*models.Security, error) {
	q := `
	Update security SET 
	timeFinish=$2, 
	status=$3
	WHERE id=$1
	`

	status := "true"

	_, err := pool.DB.Exec(context.Background(), q, id, time, status)

	if err != nil {
		return nil, err
	}

	qu := `Select userId, firstName, lastName, A, B, fiod, phone, carNumber, timeStart, timeFinish  from security where id=$1`

	row, err := pool.DB.Query(context.Background(), qu, id)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var data models.Security

	for row.Next() {
		if err := row.Scan(
			&data.UserId,
			&data.FirsrtName,
			&data.LastName,
			&data.From,
			&data.To,
			&data.FioD,
			&data.Phone,
			&data.CarNumber,
			&data.TimeStart,
			&data.TimeFinish,
		); err != nil {
			return nil, err
		}

	}

	return &data, nil

}

func (pool *SecurityDB) GetMyHistory(id int64) ([]*models.Security, error) {
	q := `SELECT id, userId, firstName, lastName, A, B, fiod, phone, carNumber, timeStart, timeFinish FROM security WHERE userId=$1`

	rows, err := pool.DB.Query(context.Background(), q, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res []*models.Security

	for rows.Next() {
		data := new(models.Security)
		if err := rows.Scan(
			&data.Id,
			&data.UserId,
			&data.FirsrtName,
			&data.LastName,
			&data.From,
			&data.To,
			&data.FioD,
			&data.Phone,
			&data.CarNumber,
			&data.TimeStart,
			&data.TimeFinish,
		); err != nil {
			return nil, err
		}

		res = append(res, data)
	}

	return res, nil
}
