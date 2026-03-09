package discount_ctrl

import (
	"context"
	"net/http"
	"time"

	"errors"
	dbprocess "service/discount/api/database/process"
	reqmodel "service/discount/api/model/request"
	"service/discount/api/utils"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductDiscountAppliedController struct {
	contextTimeout time.Duration
	pgxConn        *pgxpool.Pool
	log            *utils.AppLogger
}

func NewProductDiscountAppliedController(
	conn *pgxpool.Pool,
	timeout time.Duration,
	log *utils.AppLogger,
) *ProductDiscountAppliedController {
	return &ProductDiscountAppliedController{
		pgxConn:        conn,
		contextTimeout: timeout,
		log:            log,
	}
}

func (c *ProductDiscountAppliedController) ListProductDiscountApplied(ctx context.Context, request reqmodel.ListRequest) ([]interface{}, interface{}, error) {
	// Context Timeout
	reqCtx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	// Get
	appliedDiscounts, pagination, err := dbprocess.ListProductDiscountApplied(reqCtx, c.pgxConn, request)
	if err != nil {
		c.log.Error("gagal memproses daftar diskon produk yang diterapkan:", err)
		return nil, nil, utils.StandardError{
			Code:    http.StatusInternalServerError,
			Message: "Gagal mengambil daftar diskon produk yang diterapkan",
		}
	}
	if len(appliedDiscounts) == 0 {
		return nil, nil, utils.StandardError{
			Code:    http.StatusNotFound,
			Message: "Tidak ada data diskon produk yang diterapkan yang tersedia",
		}
	}

	// Return Success
	return appliedDiscounts, pagination, nil
}

func (c *ProductDiscountAppliedController) ListProductDiscountAppliedByProductDiscountID(ctx context.Context, productDiscountID string, request reqmodel.ListRequest) ([]interface{}, interface{}, error) {
	// Context Timeout
	reqCtx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	// Get
	appliedDiscounts, pagination, err := dbprocess.ListProductDiscountAppliedByProductDiscountID(reqCtx, c.pgxConn, productDiscountID, request)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil, utils.StandardError{
				Code:    http.StatusNotFound,
				Message: "Diskon produk yang diterapkan tidak ditemukan",
			}
		}
		c.log.Error("gagal memproses diskon produk yang diterapkan:", err)
		return nil, nil, utils.StandardError{
			Code:    http.StatusInternalServerError,
			Message: "Gagal mengambil data diskon produk yang diterapkan",
		}
	}

	if len(appliedDiscounts) == 0 {
		return nil, nil, utils.StandardError{
			Code:    http.StatusNotFound,
			Message: "Tidak ada data diskon produk yang diterapkan yang tersedia",
		}
	}

	// Return Success
	return appliedDiscounts, pagination, nil
}

func (c *ProductDiscountAppliedController) CreateProductDiscountApplied(ctx context.Context, requests []reqmodel.CreateProductDiscountApplied) error {
	// Context Timeout
	reqCtx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	// Validasi nama customer untuk setiap request
	for i := range requests {
		customerName, err := utils.FetchCustomerName(requests[i].CustomerID)
		if err != nil {
			c.log.Error("gagal mengambil data customer:", err)
			return utils.StandardError{
				Code:    http.StatusBadGateway,
				Message: "Gagal memvalidasi data customer: " + err.Error(),
			}
		}
		requests[i].CustomerName = customerName
	}

	// Create
	err := dbprocess.CreateProductDiscountApplied(reqCtx, c.pgxConn, requests)
	if err != nil {
		if reqErr, ok := err.(utils.RequestError); ok {
			return utils.StandardError{
				Code:    reqErr.StatusCode,
				Message: reqErr.Message,
			}
		}
		c.log.Error("gagal memproses diskon produk yang diterapkan:", err)
		return utils.StandardError{
			Code:    http.StatusInternalServerError,
			Message: "Gagal membuat diskon produk yang diterapkan",
		}
	}

	// Return Success
	return nil
}

func (c *ProductDiscountAppliedController) DeleteProductDiscountApplied(ctx context.Context, productDiscountAppliedID, authUserID string) error {
	// Context Timeout
	reqCtx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	// Delete
	err := dbprocess.DeleteProductDiscountApplied(reqCtx, c.pgxConn, productDiscountAppliedID, authUserID)
	if err != nil {
		if reqErr, ok := err.(utils.RequestError); ok {
			return utils.StandardError{
				Code:    reqErr.StatusCode,
				Message: reqErr.Message,
			}
		}
		c.log.Error("gagal memproses penghapusan diskon produk yang diterapkan:", err)
		return utils.StandardError{
			Code:    http.StatusInternalServerError,
			Message: "Gagal menghapus data diskon produk yang diterapkan",
		}
	}

	// Return Success
	return nil
}
