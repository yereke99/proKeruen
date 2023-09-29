package repository

import (
	"context"
	"qkeruen/dto"
	"qkeruen/help"
	"qkeruen/models"
)

type OfferUserDB struct {
	DB PgxIface
}

func NewOfferUserRepository(ds PgxIface) OfferUserDB {
	return OfferUserDB{DB: ds}
}

func (pool OfferUserDB) GetByID(id int64) (*models.DriverModelForUser, error) {
	q := `Select id, phone, firstName, lastName, ava, carNumber, carColor, carModel from driver where id=$1`

	row := pool.DB.QueryRow(context.Background(), q, id)

	var driver models.DriverModelForUser

	err := row.Scan(
		&driver.Id,
		&driver.Phone,
		&driver.FirstName,
		&driver.LastName,
		&driver.Avatar,
		&driver.CarNumber,
		&driver.CarColor,
		&driver.CarModel,
	)

	if err != nil {
		return nil, err
	}

	return &driver, nil
}

func (pool OfferUserDB) Create(id int, offer dto.OfferRequest) error {
	q := `INSERT INTO offer_user(
		    comment,
		    locationFrom,
		    locationTo,
			price,
		    type,
		    customer
	)VALUES($1,$2,$3,$4,$5,$6)`

	_, err := pool.DB.Exec(context.Background(), q, offer.Comment, offer.From, offer.To, offer.Price, offer.Type, id)

	if err != nil {
		return err
	}

	return nil
}

func (pool OfferUserDB) MyOffer(id int64) ([]*dto.OfferResponseUser, error) {
	q := `Select id, comment, locationFrom, locationTo, price, type,  customer From offer_user WHERE customer=$1 ORDER BY id DESC`

	rows, err := pool.DB.Query(context.Background(), q, id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var offers []*dto.OfferResponseUser
	for rows.Next() {
		offer := new(dto.OfferResponseUser)

		err := rows.Scan(
			&offer.Id,
			&offer.Comment,
			&offer.From,
			&offer.To,
			&offer.Price,
			&offer.Type,
			&offer.User,
		)

		if err != nil {
			return nil, err
		}

		offers = append(offers, offer)
	}

	return offers, nil
}

// here ou must to change!
func (pool OfferUserDB) FindAllOffers() ([]*dto.OfferResponseUser, error) {
	q := `Select * From offer_user`
	rows, err := pool.DB.Query(context.Background(), q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return nil, nil
}

func (pool OfferUserDB) Search(from, to, type_ string) ([]*models.OfferUserModel, error) {
	q := `Select * From offer_driver WHERE locationFrom=$1 AND locationTo=$2 ORDER BY id DESC`
	qu := `Select ava from driver where id=$1`

	rows, err := pool.DB.Query(context.Background(), q, from, to)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var offers []*models.OfferUserModel

	c := help.Choose(type_)

	for rows.Next() {
		var ava string

		offer := new(models.OfferUserModel)

		if err := rows.Scan(
			&offer.Id,
			&offer.Comment,
			&offer.From,
			&offer.To,
			&offer.Price,
			&offer.Type,
			&offer.Driver,
		); err != nil {
			return nil, err
		}

		row := pool.DB.QueryRow(context.Background(), qu, offer.Driver)

		if err := row.Scan(&ava); err != nil {
			return nil, err
		}

		offer.DriverAVA = ava

		if help.Choose(offer.Type) == c {
			offers = append(offers, offer)
		}
	}

	return offers, nil
}

func (pool OfferUserDB) SearchA(from, type_ string) ([]*models.OfferUserModel, error) {
	q := `Select * from offer_driver WHERE locationFrom=$1 ORDER BY id DESC`
	qu := `Select ava from driver where id=$1`

	rows, err := pool.DB.Query(context.Background(), q, from)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var offers []*models.OfferUserModel
	c := help.Choose(type_)

	for rows.Next() {
		var ava string

		offer := new(models.OfferUserModel)

		if err := rows.Scan(
			&offer.Id,
			&offer.Comment,
			&offer.From,
			&offer.To,
			&offer.Price,
			&offer.Type,
			&offer.Driver,
		); err != nil {
			return nil, err
		}

		row := pool.DB.QueryRow(context.Background(), qu, offer.Driver)

		if err := row.Scan(&ava); err != nil {
			return nil, err
		}

		offer.DriverAVA = ava

		if help.Choose(offer.Type) == c {
			offers = append(offers, offer)
		}
	}

	return offers, nil
}

func (pool OfferUserDB) Delete(offerId int64) error {
	q := `Delete From offer_user WHERE Id = $1`

	_, err := pool.DB.Exec(context.Background(), q, offerId)

	if err != nil {
		return err
	}

	return nil
}
