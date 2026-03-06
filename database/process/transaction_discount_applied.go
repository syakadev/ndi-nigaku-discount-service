package dbprocess

import (
	"context"

	dbquery "service/discount/api/database/query"
	reqmodel "service/discount/api/model/request"
	resmodel "service/discount/api/model/response"
	"service/discount/api/utils"
)

func ListTransactionDiscountApplied(ctx context.Context, exec DBExecutor, request reqmodel.ListRequest) ([]interface{}, *resmodel.PaginationData, error) {
	// default value page and size
	if request.Page <= 0 {
		request.Page = 1
	}

	if request.Size <= 0 {
		request.Size = 10
	}

	// Query
	query := dbquery.ListTransactionDiscountApplied()
	offset := (request.Page - 1) * request.Size
	rows, err := exec.Query(ctx, query, request.Search, request.Size, offset)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var transactionDiscountApplieds []interface{}
	for rows.Next() {
		var transactionDiscountApplied resmodel.TransactionDiscountApplied
		err := rows.Scan(
			&transactionDiscountApplied.ID,
			&transactionDiscountApplied.DiscountTransactionTargetID,
			&transactionDiscountApplied.TargetID, // transaksi id target
			&transactionDiscountApplied.PriceBeforeDiscount,
			&transactionDiscountApplied.TotalDiscount,
			&transactionDiscountApplied.PriceAfterDiscount,
			&transactionDiscountApplied.CreatedAt,
			&transactionDiscountApplied.CreatedBy,
		)
		if err != nil {
			return nil, nil, err
		}
		transactionDiscountApplieds = append(transactionDiscountApplieds, transactionDiscountApplied)
	}
	if len(transactionDiscountApplieds) == 0 {
		return nil, nil, nil
	}

	// Query Total Data
	var total int
	err = exec.QueryRow(ctx, dbquery.CountListTransactionDiscountApplied(), request.Search).Scan(&total)
	if err != nil {
		return nil, nil, err
	}

	pagination := &resmodel.PaginationData{
		Page:      request.Page,
		Size:      request.Size,
		TotalData: total,
	}
	return transactionDiscountApplieds, pagination, nil
}
func ListTransactionDiscountAppliedByDiscountID(ctx context.Context, exec DBExecutor, transactionDiscountTargetID string, request reqmodel.ListRequest) ([]interface{}, *resmodel.PaginationData, error) {
	// default value page and size
	if request.Page <= 0 {
		request.Page = 1
	}

	if request.Size <= 0 {
		request.Size = 10
	}
	// Query
	query := dbquery.GetListTransctionDiscountAppliedByDiscountID()
	offset := (request.Page - 1) * request.Size
	rows, err := exec.Query(ctx, query, transactionDiscountTargetID, request.Size, offset)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var transactionDiscountApplieds []interface{}
	for rows.Next() {
		var transactionDiscountApplid resmodel.TransactionDiscountApplied
		err := rows.Scan(
			&transactionDiscountApplid.ID,
			&transactionDiscountApplid.DiscountTransactionTargetID,
			&transactionDiscountApplid.TargetID, // transaksi id target
			&transactionDiscountApplid.PriceBeforeDiscount,
			&transactionDiscountApplid.TotalDiscount,
			&transactionDiscountApplid.PriceAfterDiscount,
			&transactionDiscountApplid.CreatedAt,
			&transactionDiscountApplid.CreatedBy,
		)
		if err != nil {
			return nil, nil, err
		}
		transactionDiscountApplieds = append(transactionDiscountApplieds, transactionDiscountApplid)
	}
	if len(transactionDiscountApplieds) == 0 {
		return nil, nil, nil
	}

	// Query Total Data
	var total int
	err = exec.QueryRow(ctx, dbquery.CountListTransactionDiscountAppliedByDiscountID(), transactionDiscountTargetID).Scan(&total)
	if err != nil {
		return nil, nil, err
	}

	pagination := &resmodel.PaginationData{
		Page:      request.Page,
		Size:      request.Size,
		TotalData: total,
	}
	return transactionDiscountApplieds, pagination, nil
}

func CreateTransactionDiscountApplied(ctx context.Context, exec DBExecutor, request reqmodel.CreateTransactionDiscountApplied) error {
	// Query
	_, err := exec.Exec(ctx, dbquery.CreateTransactionDiscountApplied(),
		request.DiscountTransactionTargetID,
		request.TargetID, // transaksi id target
		request.PriceBeforeDiscount,
		request.TotalDiscount,
		request.PriceAfterDiscount,
		request.AuthUserID,
	)
	if err != nil {
		return err
	}

	// Return Success
	return nil
}

func DeleteTransactionDiscountApplied(ctx context.Context, exec DBExecutor, transactionDiscountAppliedID, authUserID string) error {
	// Query
	result, err := exec.Exec(ctx, dbquery.DeleteTransactionDiscountApplied(),
		transactionDiscountAppliedID,
		authUserID,
	)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return utils.RequestError{
			StatusCode: 404,
			Message:    "Gagal melakukan penghapusan, data diskon transaksi yang diterapkan tidak ditemukan",
		}
	}

	// Return Success
	return nil
}
