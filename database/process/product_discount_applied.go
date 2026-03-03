package dbprocess

import (
	"context"

	dbquery "service/discount/api/database/query"
	reqmodel "service/discount/api/model/request"
	resmodel "service/discount/api/model/response"
	"service/discount/api/utils"
)

func ListProductDiscountApplied(ctx context.Context, exec DBExecutor, request reqmodel.ListRequest) ([]interface{}, *resmodel.PaginationData, error) {
	// Query
	query := dbquery.ListProductDiscountApplied()
	offset := (request.Page - 1) * request.Size
	rows, err := exec.Query(ctx, query, request.Search, request.Size, offset)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var appliedDiscounts []interface{}
	for rows.Next() {
		var appliedDiscount resmodel.ProductDiscountApplied
		err := rows.Scan(
			&appliedDiscount.ID,
			&appliedDiscount.ProductDiscountID,
			&appliedDiscount.CustomerID,
			&appliedDiscount.CustomerName,
			&appliedDiscount.CreatedAt,
			&appliedDiscount.CreatedBy,
			&appliedDiscount.UpdatedAt,
			&appliedDiscount.UpdatedBy,
		)
		if err != nil {
			return nil, nil, err
		}
		appliedDiscounts = append(appliedDiscounts, appliedDiscount)
	}
	if len(appliedDiscounts) == 0 {
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
	return appliedDiscounts, pagination, nil
}
func ListProductDiscountAppliedByID(ctx context.Context, exec DBExecutor,productDiscountAppliedID string,request reqmodel.ListRequest) ([]interface{}, *resmodel.PaginationData, error) {
	// Query
	query := dbquery.GetDiscountProductTargetByID()
	offset := (request.Page - 1) * request.Size
	rows, err := exec.Query(ctx, query, productDiscountAppliedID, request.Size, offset)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var productDiscountApplieds []interface{}
	for rows.Next() {
		var productDiscountApplid resmodel.ProductDiscountApplied
		err := rows.Scan(
			&productDiscountApplid.ID,
			&productDiscountApplid.ProductDiscountID,
			&productDiscountApplid.CustomerID,
			&productDiscountApplid.CustomerName,
			&productDiscountApplid.CreatedAt,
			&productDiscountApplid.CreatedBy,
			&productDiscountApplid.UpdatedAt,
			&productDiscountApplid.UpdatedBy,
		)
		if err != nil {
			return nil, nil, err
		}
		productDiscountApplieds = append(productDiscountApplieds, productDiscountApplid)
	}
	if len(productDiscountApplieds) == 0 {
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
	return productDiscountApplieds, pagination, nil
}


func CreateProductDiscountApplied(ctx context.Context, exec DBExecutor, request reqmodel.CreateProductDiscountApplied) error {
	// Query
	_, err := exec.Exec(ctx, dbquery.CreateProductDiscountApplied(),
		request.DiscountProductTargetID,
		request.CustomerID,
		request.CustomerName,
		request.AuthUserID,
	)
	if err != nil {
		return err
	}

	// Return Success
	return nil
}


func DeleteProductDiscountApplied(ctx context.Context, exec DBExecutor, postID, authUserID string) error {
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
			Message:    "Gagal melakukan penghapusan, data product appplied tidak ditemukan",
		}
	}

	// Return Success
	return nil
}
