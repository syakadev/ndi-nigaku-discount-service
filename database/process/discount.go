package dbprocess

import (
	"context"

	dbquery "service/discount/api/database/query"
	reqmodel "service/discount/api/model/request"
	resmodel "service/discount/api/model/response"
	"service/discount/api/utils"
)

func ListDiscount(ctx context.Context, exec DBExecutor, request reqmodel.ListRequest) ([]interface{}, *resmodel.PaginationData, error) {
	// Query
	query := dbquery.ListDiscount()
	offset := (request.Page - 1) * request.Size
	rows, err := exec.Query(ctx, query, request.Search, request.Size, offset)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var discounts []interface{}
	for rows.Next() {
		var discount resmodel.Discount
		err := rows.Scan(
			&discount.ID,
			&discount.Name,
			&discount.Type,
			&discount.Value,
			&discount.StartDate,
			&discount.EndDate,
			&discount.Target,
			&discount.IsActive,
			&discount.CreatedAt,
			&discount.CreatedBy,
			&discount.UpdatedAt,
			&discount.UpdatedBy,
		)
		if err != nil {
			return nil, nil, err
		}
		discounts = append(discounts, discount)
	}
	if len(discounts) == 0 {
		return nil, nil, nil
	}

	// Query Total Data
	var total int
	err = exec.QueryRow(ctx, dbquery.CountListPost(), request.Search).Scan(&total)
	if err != nil {
		return nil, nil, err
	}

	pagination := &resmodel.PaginationData{
		Page:      request.Page,
		Size:      request.Size,
		TotalData: total,
	}
	return discounts, pagination, nil
}

func GetDiscountByID(ctx context.Context, exec DBExecutor, discountID string) (*resmodel.Discount, error) {
	// Query
	query := dbquery.GetPostByID()
	row := exec.QueryRow(ctx, query, discountID)
	var discount resmodel.Discount
	err := row.Scan(
		&discount.ID,
		&discount.Name,
		&discount.Type,
		&discount.Value,
		&discount.StartDate,
		&discount.EndDate,
		&discount.Target,
		&discount.IsActive,
		&discount.CreatedAt,
		&discount.CreatedBy,
		&discount.UpdatedAt,
		&discount.UpdatedBy,
	)
	if err != nil {
		return nil, err
	}
	return &discount, nil
}

func CreateDiscount(ctx context.Context, exec DBExecutor, request reqmodel.CreateDiscount) error {
	// Query
	_, err := exec.Exec(ctx, dbquery.CreatePost(),
		request.Name,
		request.Type,
		request.Value,
		request.StartDate,
		request.EndDate,
		request.Target,
		request.IsActive,
		request.AuthUserID,
	)
	if err != nil {
		return err
	}

	// Return Success
	return nil
}

func UpdateDiscount(ctx context.Context, exec DBExecutor, request reqmodel.UpdateDiscount) error {
	// Query
	result, err := exec.Exec(ctx, dbquery.UpdatePost(),
		request.ID,
		request.Name,
		request.Value,
		request.StartDate,
		request.EndDate,
		request.Target,
		request.IsActive,
		request.AuthUserID,
	)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return utils.RequestError{
			StatusCode: 404,
			Message:    "Gagal melakukan perubahan, post tidak ditemukan",
		}
	}

	// Return Success
	return nil
}

func DeleteDiscount(ctx context.Context, exec DBExecutor, discountID, authUserID string) error {
	// Query
	result, err := exec.Exec(ctx, dbquery.DeletePost(),
		discountID,
		authUserID,
	)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return utils.RequestError{
			StatusCode: 404,
			Message:    "Gagal melakukan penghapusan, discount tidak ditemukan",
		}
	}

	// Return Success
	return nil
}
