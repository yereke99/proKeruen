package repository

import (
	"context"
	"qkeruen/models"
)

type HistoryDB struct {
	DB PgxIface
}

func NewHistoryRepository(ds PgxIface) HistoryDB {
	return HistoryDB{DB: ds}
}

func (pool HistoryDB) GetUserHistory(user_id int64) ([]*models.History, error) {
	q := `Select * From history WHERE userId=$1`

	rows, err := pool.DB.Query(context.Background(), q, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hList []*models.History
	for rows.Next() {
		h := new(models.History)

		err := rows.Scan(
			&h.Id,
			&h.OrderId,
			&h.DriverId,
			&h.UserId,
			&h.StartDate,
			&h.FinishedDate,
		)

		if err != nil {
			return nil, err
		}

		hList = append(hList, h)
	}
	return hList, nil
}

func (pool HistoryDB) GetDriverHistory(user_id int64) ([]*models.History, error) {
	q := `Select * From history WHERE driverId=$1`
	
	rows, err := pool.DB.Query(context.Background(), q, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hList []*models.History
	for rows.Next() {
		h := new(models.History)

		err := rows.Scan(
			&h.Id,
			&h.OrderId,
			&h.DriverId,
			&h.UserId,
			&h.StartDate,
			&h.FinishedDate,
		)

		if err != nil {
			return nil, err
		}

		hList = append(hList, h)
	}
	return hList, nil
}
