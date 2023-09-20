package repository

import (
	"context"
	"errors"
	"fmt"
	"qkeruen/models"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

// PgxIface is pgx interface
type PgxIface interface {
	// using pgxconn interface
	// Begin(context.Context) (pgx.Tx, error)
	// Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	// QueryRow(context.Context, string, ...interface{}) pgx.Row
	// Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	// Ping(context.Context) error
	// Prepare(context.Context, string, string) (*pgconn.StatementDescription, error)
	// Close(context.Context) error

	// using pgxpool interface
	Begin(context.Context) (pgx.Tx, error)
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	Ping(context.Context) error
	Close()
}

type Database struct {
	DB PgxIface
}

func NewDatabase(db PgxIface) Database {
	return Database{DB: db}
}

// Check
func (pool Database) CheckFromRepo(contact, role string) (bool, error) {
	switch role {
	case "driver":

		q := `Select phone From driver WHERE phone=$1`
		row := pool.DB.QueryRow(context.Background(), q, contact)

		var phone string

		err := row.Scan(&phone)
		if err != nil {
			switch err.Error() {
			case "no rows in result set":
				return true, nil
			default:
				return false, err
			}
		}

		if phone == contact {
			return false, nil
		}

		return true, nil

	case "user":

		q := `Select phone From customer WHERE phone=$1`
		row := pool.DB.QueryRow(context.Background(), q, contact)

		var phone string

		err := row.Scan(&phone)
		fmt.Println(err)
		if err != nil {
			switch err.Error() {
			case "no rows in result set":
				return true, nil
			default:
				return false, err
			}
		}

		if phone == contact {
			return false, nil
		}

		return true, nil
	default:
		return false, errors.New("Wrong role is given.")
	}

}

func (pool Database) GiveToken(contact, role string) (string, error) {
	switch role {
	case "driver":

		q := `Select token From driver Where phone=$1`
		row := pool.DB.QueryRow(context.Background(), q, contact)

		var token string

		err := row.Scan(&token)

		if err != nil {
			return "", err
		}

		return token, nil
	case "user":

		q := `Select token From customer Where phone=$1`

		row := pool.DB.QueryRow(context.Background(), q, contact)

		var token string

		err := row.Scan(&token)

		if err != nil {
			return "", err
		}

		return token, nil
	default:
		return "", errors.New("Wrong type of role.")
	}
}

// Create method will insert new record to database. 'C' part of the CRUD
func (pool Database) Insert(contact string, code int) error {
	if e := pool.Clean(contact); e != nil {
		return e
	}

	// sql for inserting new record
	q := `INSERT INTO sms_cache(contact, code) VALUES($1,$2)`

	_, err := pool.DB.Exec(context.Background(), q, contact, code)

	if err != nil {
		return err
	}
	defer pool.All()

	return nil
}

func (pool Database) CreateCode(contact string, code int) error {
	q := `Select COUNT(*) FROM sms_cache WHERE contact = $1`
	var count int
	err := pool.DB.QueryRow(context.Background(), q, contact).Scan(&count)

	if err != nil {
		return err
	}

	if count > 0 {
		// Update the existing row
		q = `UPDATE sms_cache SET code = $2 WHERE contact = $1`
		_, err = pool.DB.Exec(context.Background(), q, contact, code)
		if err != nil {
			return err
		}
	} else {
		// Insert a new row
		q = `INSERT INTO sms_cache(contact, code) VALUES ($1, $2)`
		_, err = pool.DB.Exec(context.Background(), q, contact, code)
		if err != nil {
			return err
		}
	}
	return nil
}

func (pool Database) ValidateSMS(contact string) (int, error) {
	q := `Select code From sms_cache WHERE contact=$1`

	row := pool.DB.QueryRow(context.Background(), q, contact)

	var code int

	err := row.Scan(&code)

	if err != nil {
		return 0, err
	}
	return code, nil
}

func (pool Database) Clean(contact string) error {
	q := `DELETE FROM sms_cache x USING sms_cache y WHERE x.id <= y.id AND x.contact = y.contact AND x.contact = $1`

	_, err := pool.DB.Exec(context.Background(), q, contact)

	if err != nil {
		return err
	}

	return nil
}

func (pool Database) All() ([]*models.SMS, error) {
	q := `Select contact,code From sms_cache`

	rows, err := pool.DB.Query(context.Background(), q)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var all []*models.SMS

	for rows.Next() {
		sms := new(models.SMS)

		err := rows.Scan(&sms.Contact, &sms.Code)
		if err != nil {
			return nil, err
		}

		all = append(all, sms)
	}
	for _, v := range all {
		fmt.Println(*v)
	}
	return all, nil
}

func (pool Database) CheckTokenDriver(token string) (*models.DriverModel, error) {
	q := `Select * From driver Where token=$1`
	rows := pool.DB.QueryRow(context.Background(), q, token)

	var modelDriver models.DriverModel
	modelDriver.Type = "driver"

	err := rows.Scan(
		&modelDriver.Id,
		&modelDriver.Phone,
		&modelDriver.FirstName,
		&modelDriver.LastName,
		&modelDriver.Inn,
		&modelDriver.Avatar,
		&modelDriver.CarNumber,
		&modelDriver.CarColor,
		&modelDriver.CarModel,
		&modelDriver.DocsFront,
		&modelDriver.DocsBacks,
		&modelDriver.CarType,
		&modelDriver.Token,
	)

	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	fmt.Println(modelDriver)
	return &modelDriver, nil
}

func (pool Database) CheckTokenUser(token string) (*models.UserModel, error) {
	q := `Select * From customer Where token=$1`
	rows := pool.DB.QueryRow(context.Background(), q, token)

	// you must to change here to user model !
	var modelUser models.UserModel
	modelUser.Type = "user"
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

	return &modelUser, nil
}
