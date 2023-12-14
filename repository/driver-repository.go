package repository

import (
	"context"
	"fmt"
	"qkeruen/models"
)

type DriverDB struct {
	DB PgxIface
}

func NewDriverRepository(ds PgxIface) DriverDB {
	return DriverDB{DB: ds}
}

func (pool DriverDB) InsertDriverData(data models.DriverRegister) (*models.DriverModel, error) {
	var driver models.DriverModel
	q := `
	    INSERT INTO driver(
		  phone,
		  firstName,
		  lastName,
		  inn,
		  ava,
		  carNumber,
		  carColor,
		  carModel,
		  docsfront,
		  docsback,
		  carType,
		  token
		)VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)`
	_, err := pool.DB.Exec(
		context.Background(),
		q,
		data.Phone,
		data.FirstName,
		data.LastName,
		data.Inn,
		data.Avatar,
		data.CarNumber,
		data.CarColor,
		data.CarModel,
		data.DocsFront,
		data.DocsBacks,
		data.CarType,
		data.Token,
	)
	if err != nil {
		return &models.DriverModel{}, err
	}

	qu := `Select * from driver WHERE token=$1`
	row := pool.DB.QueryRow(context.Background(), qu, data.Token)

	if err := row.Scan(
		&driver.Id,
		&driver.Phone,
		&driver.FirstName,
		&driver.LastName,
		&driver.Inn,
		&driver.Avatar,
		&driver.CarNumber,
		&driver.CarColor,
		&driver.CarModel,
		&driver.DocsFront,
		&driver.DocsBacks,
		&driver.CarType,
		&driver.Token,
	); err != nil {
		return &models.DriverModel{}, err
	}

	driver.Type = "driver"

	return &driver, nil
}

func (pool DriverDB) GetDriverProfile(token string) (*models.DriverModel, error) {

	q := `Select * From driver Where token=$1`
	rows := pool.DB.QueryRow(context.Background(), q, token)

	var modelDriver models.DriverModel

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
	modelDriver.Type = "driver"
	fmt.Println(modelDriver)

	return &modelDriver, nil

}

func (pool DriverDB) UpdateDriver(model models.DriverModel) (*models.DriverModel, error) {
	if model.Phone != "" {
		q := `Update driver SET phone=$1 WHERE token=$2`
		_, err := pool.DB.Exec(context.Background(), q, model.Phone, model.Token)
		if err != nil {
			return nil, err
		}
	}

	if model.FirstName != "" {
		q := `Update driver SET firstName=$1 WHERE token=$2`
		_, err := pool.DB.Exec(context.Background(), q, model.FirstName, model.Token)
		if err != nil {
			return nil, err
		}
	}

	if model.LastName != "" {
		q := `Update driver SET lastName=$1 WHERE token=$2`
		_, err := pool.DB.Exec(context.Background(), q, model.LastName, model.Token)
		if err != nil {
			return nil, err
		}
	}
	if model.Inn != "" {
		q := `Update driver SET inn=$1 WHERE token=$2`
		_, err := pool.DB.Exec(context.Background(), q, model.Inn, model.Token)
		if err != nil {
			return nil, err
		}
	}
	if model.Avatar != "" {
		q := `Update driver SET ava=$1 WHERE token=$2;`
		_, err := pool.DB.Exec(context.Background(), q, model.Avatar, model.Token)
		if err != nil {
			return nil, err
		}
	}

	if model.CarNumber != "" {
		q := `Update driver SET carNumber=$1 WHERE token=$2`
		_, err := pool.DB.Exec(context.Background(), q, model.CarNumber, model.Token)
		if err != nil {
			return nil, err
		}
	}
	if model.CarColor != "" {
		q := `Update driver SET carColor=$1 WHERE token=$2`
		_, err := pool.DB.Exec(context.Background(), q, model.CarColor, model.Token)
		if err != nil {
			return nil, err
		}
	}

	if model.CarModel != "" {
		q := `Update driver set carModel=$1 WHERE token=$2`
		_, err := pool.DB.Exec(context.Background(), q, model.CarModel, model.Token)
		if err != nil {
			return nil, err
		}
	}
	if model.DocsFront != "" {
		q := `Update driver set docsfront=$1 WHERE token=$2`
		_, err := pool.DB.Exec(context.Background(), q, model.DocsFront, model.Token)
		if err != nil {
			return nil, err
		}
	}

	if model.DocsBacks != "" {
		q := `Update driver set docsback=$1 WHERE token=$2`
		_, err := pool.DB.Exec(context.Background(), q, model.DocsBacks, model.Token)
		if err != nil {
			return nil, err
		}
	}

	if model.CarType != "" {
		q := `Update driver set carType=$1 WHERE token=$2`
		_, err := pool.DB.Exec(context.Background(), q, model.CarType, model.Token)
		if err != nil {
			return nil, err
		}
	}

	q := `Select * From driver WHERE token=$1`
	rows := pool.DB.QueryRow(context.Background(), q, model.Token)

	u := new(models.DriverModel)

	err := rows.Scan(
		&u.Id,
		&u.Phone,
		&u.FirstName,
		&u.LastName,
		&u.Inn,
		&u.Avatar,
		&u.CarNumber,
		&u.CarColor,
		&u.CarModel,
		&u.DocsFront,
		&u.DocsBacks,
		&u.CarType,
		&u.Token,
	)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return u, nil
}

func (pool DriverDB) Delete(id int64) error {
	q := `Delete From driver WHERE id=$1`
	_, err := pool.DB.Exec(context.Background(), q, id)
	if err != nil {
		return err
	}

	return nil
}
