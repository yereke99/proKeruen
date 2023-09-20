package repository

import "context"

type SearchDB struct {
	DB PgxIface
}

func NewSearchRepository(ds PgxIface) SearchDB {
	return SearchDB{DB: ds}
}

func (s *SearchDB) Check(place string) (bool, error) {
	q := `select target from places where target=$1`
	row := s.DB.QueryRow(context.Background(), q, place)

	var target_ string
	if err := row.Scan(&target_); err != nil {
		switch err.Error() {
		case "no rows in result set":
			return true, nil
		default:
			return false, err
		}
	}

	if target_ == place {
		return false, nil
	}

	return false, nil
}

func (s *SearchDB) Create(places string) error {
	q := `INSERT INTO places(target)VALUES($1)`

	_, err := s.DB.Exec(context.Background(), q)

	if err != nil {
		return err
	}

	return nil
}

func (s *SearchDB) CheckGeo(places string) (bool, error) {
	return false, nil
}

func (s *SearchDB) CreateGeo(places string) error {
	return nil
}
