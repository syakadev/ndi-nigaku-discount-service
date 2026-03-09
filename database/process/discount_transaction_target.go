package dbprocess

import (
	"context"

	dbquery "service/discount/api/database/query"
	reqmodel "service/discount/api/model/request"
	resmodel "service/discount/api/model/response"
	"service/discount/api/utils"
)

func ListDiscountTransactionTarget(ctx context.Context, exec DBExecutor, request reqmodel.ListRequest) ([]interface{}, *resmodel.PaginationData, error) {
	// default value page and size
	if request.Page <= 0 {
		request.Page = 1
	}

	if request.Size <= 0 {
		request.Size = 10
	}

	// Query
	query := dbquery.ListDiscountTransactionTarget()
	offset := (request.Page - 1) * request.Size
	rows, err := exec.Query(ctx, query, request.Search, request.Size, offset)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var discTransactionTargets []interface{}
	for rows.Next() {
		var discTransactionTarget resmodel.DiscountTransactionTarget
		err := rows.Scan(
			&discTransactionTarget.ID,
			&discTransactionTarget.DiscountID,
			&discTransactionTarget.MaxTotalQuota,
			&discTransactionTarget.IsActive,
			&discTransactionTarget.CreatedAt,
			&discTransactionTarget.CreatedBy,
			&discTransactionTarget.UpdatedAt,
			&discTransactionTarget.UpdatedBy,
		)
		if err != nil {
			return nil, nil, err
		}
		discTransactionTargets = append(discTransactionTargets, discTransactionTarget)
	}
	if len(discTransactionTargets) == 0 {
		return nil, nil, nil
	}

	// Query Total Data
	var total int
	err = exec.QueryRow(ctx, dbquery.CountListDiscountTransactionTarget(), request.Search).Scan(&total)
	if err != nil {
		return nil, nil, err
	}

	pagination := &resmodel.PaginationData{
		Page:      request.Page,
		Size:      request.Size,
		TotalData: total,
	}
	return discTransactionTargets, pagination, nil
}

func GetDiscountTransactionTargetByID(ctx context.Context, exec DBExecutor, discTransactionTargetID string) (*resmodel.DiscountTransactionTarget, error) {
	// Query
	query := dbquery.GetDiscountTransactionTargetByID()
	row := exec.QueryRow(ctx, query, discTransactionTargetID)
	var discTransactionTarget resmodel.DiscountTransactionTarget
	err := row.Scan(
		&discTransactionTarget.ID,
		&discTransactionTarget.DiscountID,
		&discTransactionTarget.MaxTotalQuota,
		&discTransactionTarget.IsActive,
		&discTransactionTarget.CreatedAt,
		&discTransactionTarget.CreatedBy,
		&discTransactionTarget.UpdatedAt,
		&discTransactionTarget.UpdatedBy,
	)
	if err != nil {
		return nil, err
	}
	return &discTransactionTarget, nil
}

func CreateDiscountTransactionTarget(ctx context.Context, exec DBExecutor, request reqmodel.CreateDiscountTransactionTarget) error {
	// Check if discount exists
	_, errDiscount := GetDiscountByID(ctx, exec, request.DiscountID)
	if errDiscount != nil {
		return utils.RequestError{
			StatusCode: 404,
			Message:    "Data diskon tidak ditemukan",
		}
	}

	// Query
	_, err := exec.Exec(ctx, dbquery.CreateDiscountTransactionTarget(),
		request.DiscountID,
		request.MaxTotalQuota,
		request.IsActive,
		request.AuthUserID,
	)
	if err != nil {
		return err
	}

	// Return Success
	return nil
}

func UpdateDiscountTransactionTarget(ctx context.Context, exec DBExecutor, request reqmodel.UpdateDiscountTransactionTarget) error {
	// Query
	result, err := exec.Exec(ctx, dbquery.UpdateDiscountTransactionTarget(),
		request.ID,
		request.MaxTotalQuota,
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

func DeleteDiscountTransactionTarget(ctx context.Context, exec DBExecutor, postID, authUserID string) error {
	// Query
	result, err := exec.Exec(ctx, dbquery.DeleteDiscountTransactionTarget(),
		postID,
		authUserID,
	)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return utils.RequestError{
			StatusCode: 404,
			Message:    "Gagal melakukan penghapusan, post tidak ditemukan",
		}
	}

	// Return Success
	return nil
}
