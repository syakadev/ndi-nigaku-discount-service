package dbprocess

import (
	"context"

	dbquery "service/discount/api/database/query"
	reqmodel "service/discount/api/model/request"
	resmodel "service/discount/api/model/response"
	"service/discount/api/utils"
)

func ListDiscountProductTarget(ctx context.Context, exec DBExecutor, request reqmodel.ListRequest) ([]interface{}, *resmodel.PaginationData, error) {
	// Query
	query := dbquery.ListDiscountProductTarget()
	offset := (request.Page - 1) * request.Size
	rows, err := exec.Query(ctx, query, request.Search, request.Size, offset)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var discProductTargets []interface{}
	for rows.Next() {
		var discProductTarget resmodel.DiscountProductTarget
		err := rows.Scan(
			&discProductTarget.ID,
			&discProductTarget.DiscountID,
			&discProductTarget.TargetID,
			&discProductTarget.MaxTotalQuota,
			&discProductTarget.IsActive,
			&discProductTarget.CreatedAt,
			&discProductTarget.CreatedBy,
			&discProductTarget.UpdatedAt,
			&discProductTarget.UpdatedBy,
		)
		if err != nil {
			return nil, nil, err
		}
		discProductTargets = append(discProductTargets, discProductTarget)
	}
	if len(discProductTargets) == 0 {
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
	return discProductTargets, pagination, nil
}

func GetDiscountProductTargetByID(ctx context.Context, exec DBExecutor, discProductTargetID string) (*resmodel.DiscountProductTarget, error) {
	// Query
	query := dbquery.GetDiscountProductTargetByID()
	row := exec.QueryRow(ctx, query, discProductTargetID)
	var discProductTarget resmodel.DiscountProductTarget
	err := row.Scan(
		&discProductTarget.ID,
		&discProductTarget.DiscountID,
		&discProductTarget.TargetID,
		&discProductTarget.MaxTotalQuota,
		&discProductTarget.IsActive,
		&discProductTarget.CreatedAt,
		&discProductTarget.CreatedBy,
		&discProductTarget.UpdatedAt,
		&discProductTarget.UpdatedBy,
	)
	if err != nil {
		return nil, err
	}
	return &discProductTarget, nil
}

func CreateDiscountProductTarget(ctx context.Context, exec DBExecutor, request reqmodel.CreateDiscountProductTarget) error {
	// Query
	_, err := exec.Exec(ctx, dbquery.CreateDiscountProductTarget(),
		request.DiscountID,
		request.TargetType,
		request.TargetID,
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

func UpdateDiscountProductTarget(ctx context.Context, exec DBExecutor, request reqmodel.UpdateDiscountProductTarget) error {
	// Query
	result, err := exec.Exec(ctx, dbquery.UpdateDiscountProductTarget(),
		request.ID,
		request.DiscountID,
		request.TargetType,
		request.TargetID,
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

func DeleteDiscountProductTarget(ctx context.Context, exec DBExecutor, postID, authUserID string) error {
	// Query
	result, err := exec.Exec(ctx, dbquery.DeletePost(),
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
