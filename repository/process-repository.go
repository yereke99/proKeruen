package repository

import (
	"context"
	"qkeruen/dto"
	"qkeruen/models"
	"strconv"
	"time"
)

type ProcessDB struct {
	DB PgxIface
}

func NewProcessRepository(ds PgxIface) ProcessDB {
	return ProcessDB{DB: ds}
}

func (pool ProcessDB) GetOrder(orderId int64) (*dto.OrderResponse, error) {
	q := `Select * From order_process WHERE id=$1`

	row := pool.DB.QueryRow(context.Background(), q, orderId)

	var order dto.OrderResponse

	err := row.Scan(
		&order.Id,
		&order.UserId,
		&order.LatitudeFrom,
		&order.LongitudeFrom,
		&order.LatitudeTo,
		&order.LongitudeTo,
		&order.Comments,
		&order.Price,
		&order.Type,
		&order.OrderStatus,
	)

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (pool ProcessDB) FinishUpdate(t_, orderId int64, finishedDate int64) error {
	q := `Update order SET orderStatus=$1, finishedDate=$2 WHERE id=$3`
	_, err := pool.DB.Exec(context.Background(), q, t_, strconv.FormatInt(finishedDate, 10), orderId)

	if err != nil {
		return err
	}

	return nil
}

func (pool ProcessDB) Update(t_, orderId int64) error {
	q := `Update order_process SET orderStatus=$1 WHERE id=$2`
	_, err := pool.DB.Exec(context.Background(), q, t_, orderId)

	if err != nil {
		return err
	}

	return nil
}

func (pool ProcessDB) Delete(orderId int64) error {
	delete := `Delete from order_process WHERE id=$1`

	_, err := pool.DB.Exec(context.Background(), delete, orderId)

	if err != nil {
		return err
	}

	return nil
}

func (pool ProcessDB) AcceptOrder(driverId, orderId int64) (interface{}, error) {
	order, err := pool.GetOrder(orderId)

	if err != nil {
		return nil, err
	}

	if order.OrderStatus != 0 {
		return "Order already in process", nil
	}

	errSave := pool.Update(1, orderId)

	if err != nil {
		return "can not change order status.", errSave
	}

	save := `INSERT INTO history(
		orderId,
		driverId,
		userId,
		startDate  
	)VALUES($1,$2,$3,$4)`

	row := pool.DB.QueryRow(context.Background(), save, orderId, driverId, order.UserId, strconv.FormatInt(time.Now().UnixMilli(), 10))

	var orderProcess models.ProcessModel

	if err := row.Scan(
		&orderProcess.Id,
		&orderProcess.OrderId,
		&orderProcess.DriverId,
		&orderProcess.UserId,
		&orderProcess.StartDate,
	); err != nil {
		return nil, err
	}

	return orderProcess, nil

}

func (pool ProcessDB) CancellOrder(orderId int64) (interface{}, error) {
	order, err := pool.GetOrder(orderId)

	if err != nil {
		return "Order not found", err
	}

	if order.OrderStatus != 1 {
		return "order is not busy.", nil
	}
	if order.OrderStatus == 2 {
		return "Order is not in process", nil
	}
	errSave := pool.Update(0, orderId)

	if errSave != nil {
		return "error in update status order method.", errSave
	}

	if errDelete := pool.Delete(orderId); errDelete != nil {
		return "error in delete order process service", errDelete
	}

	return "", nil
}

func (pool ProcessDB) GetOrdersInProcessDriver(driverId int64) ([]*models.ProcessModel, error) {
	q := `Select * From order_process WHERE driverId=$1;`

	rows, err := pool.DB.Query(context.Background(), q, driverId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*models.ProcessModel

	for rows.Next() {
		order := new(models.ProcessModel)

		err := rows.Scan(
			&order.Id,
			&order.OrderId,
			&order.DriverId,
			&order.UserId,
			&order.StartDate,
		)

		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (pool ProcessDB) GetOrdersInProcessUser(userId int64) ([]*models.ProcessModel, error) {
	q := `Select * From order_process WHERE userId=$1;`

	rows, err := pool.DB.Query(context.Background(), q, userId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*models.ProcessModel

	for rows.Next() {
		order := new(models.ProcessModel)

		err := rows.Scan(
			&order.Id,
			&order.OrderId,
			&order.DriverId,
			&order.UserId,
			&order.StartDate,
		)

		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (pool ProcessDB) FinishOrder(driverId, orderId int64) (interface{}, error) {
	order, err := pool.GetOrder(orderId)

	if err != nil {
		return "error in get order.", err
	}

	if order.OrderStatus == 3 {
		return "order already finished.", nil
	}

	if errSave := pool.FinishUpdate(3, orderId, time.Now().UnixMilli()); errSave != nil {
		return "error in finish order service", errSave
	}

	if errDelete := pool.Delete(orderId); errDelete != nil {
		return "error in finish order delete repository.", errDelete
	}

	return "finished.", nil
}
