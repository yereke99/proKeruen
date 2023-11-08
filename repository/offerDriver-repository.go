package repository

import (
	"context"
	"qkeruen/dto"
	"qkeruen/help"
	"qkeruen/models"
)

type OfferDriverDB struct {
	DB PgxIface
}

func NewOfferDriverRepository(ds PgxIface) OfferDriverDB {
	return OfferDriverDB{
		DB: ds,
	}
}

func (pool OfferDriverDB) GetByID(id int64) (*models.UserModelForDriver, error) {
	q := `Select id, phone, firstName, lastName, ava from customer where id=$1`

	row := pool.DB.QueryRow(context.Background(), q, id)

	var user models.UserModelForDriver

	if err := row.Scan(
		&user.Id,
		&user.Phone,
		&user.FirstName,
		&user.LastName,
		&user.Avatar,
	); err != nil {
		return nil, err
	}

	return &user, nil
}

func (pool OfferDriverDB) Create(id int, offer dto.OfferRequest) error {
	q := `INSERT INTO offer_driver(
		comment,
		locationFrom,
		locationTo,
		price,
		type,
		driver
	)VALUES($1,$2,$3,$4,$5,$6)`

	_, err := pool.DB.Exec(context.Background(), q, offer.Comment, offer.From, offer.To, offer.Price, offer.Type, id)

	if err != nil {
		return err
	}

	return nil
}

func (pool OfferDriverDB) MyOffer(id int64) ([]*dto.OfferResponseDriver, error) {
	q := `Select id, comment, locationFrom, locationTo, price, driver  From offer_driver WHERE driver=$1 ORDER BY id DESC`

	rows, err := pool.DB.Query(context.Background(), q, id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var offers []*dto.OfferResponseDriver
	for rows.Next() {
		offer := new(dto.OfferResponseDriver)

		err := rows.Scan(
			&offer.Id,
			&offer.Comment,
			&offer.From,
			&offer.To,
			&offer.Price,
			&offer.Driver,
		)

		if err != nil {
			return nil, err
		}

		offers = append(offers, offer)
	}

	return offers, nil
}

func (pool OfferDriverDB) FindAllOffers() ([]*models.OfferDriverModel, error) {
	q := `Select * From offer_driver ORDER BY id DESC`
	rows, err := pool.DB.Query(context.Background(), q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return nil, nil
}

func (pool OfferDriverDB) Search(to, from, type_ string) ([]*models.OfferDriverModel, error) {
	q := `SELECT * FROM offer_user WHERE locationFrom=$1 AND locationTo=$2 ORDER BY id DESC`

	qu := `Select ava from customer where id=$1`

	rows, err := pool.DB.Query(context.Background(), q, from, to)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var offers []*models.OfferDriverModel

	c := help.Choose(type_)

	for rows.Next() {
		var ava string

		offer := new(models.OfferDriverModel)

		if err := rows.Scan(&offer.Id,
			&offer.Comment,
			&offer.From,
			&offer.To,
			&offer.Price,
			&offer.Type,
			&offer.User,
		); err != nil {
			return nil, err
		}

		row := pool.DB.QueryRow(context.Background(), qu, offer.User)

		if err := row.Scan(&ava); err != nil {
			return nil, err
		}

		offer.UserAVA = ava

		if help.Choose(offer.Type) == c {
			offers = append(offers, offer)
		}

	}

	return offers, nil
}

func (pool OfferDriverDB) SearchOneSide(from, type_ string) ([]*models.OfferDriverModel, error) {
	q := `Select * from offer_user WHERE locationFrom=$1 ORDER BY id DESC`

	qu := `Select ava from customer where id=$1`

	rows, err := pool.DB.Query(context.Background(), q, from)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var offers []*models.OfferDriverModel
	c := help.Choose(type_)

	for rows.Next() {
		var ava string

		offer := new(models.OfferDriverModel)

		if err := rows.Scan(
			&offer.Id,
			&offer.Comment,
			&offer.From,
			&offer.To,
			&offer.Price,
			&offer.Type,
			&offer.User,
		); err != nil {
			return nil, err
		}

		row := pool.DB.QueryRow(context.Background(), qu, offer.User)

		if err := row.Scan(&ava); err != nil {
			return nil, err
		}

		offer.UserAVA = ava

		if help.Choose(offer.Type) == c {
			offers = append(offers, offer)
		}
	}

	return offers, nil
}

func (pool OfferDriverDB) Delete(offerId int64) error {
	q := `Delete from offer_driver WHERE id=$1`

	_, err := pool.DB.Exec(context.Background(), q, offerId)

	if err != nil {
		return err
	}

	return nil
}
