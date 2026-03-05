package discount_ctrl

import (
	"context"
	"net/http"
	"time"

	"errors"
	dbprocess "service/discount/api/database/process"
	reqmodel "service/discount/api/model/request"
	resmodel "service/discount/api/model/response"
	"service/discount/api/utils"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DiscountTransactionTargetController struct {
	contextTimeout time.Duration
	pgxConn        *pgxpool.Pool
	log            *utils.AppLogger
}

func NewDiscountTransactionTargetController(
	conn *pgxpool.Pool,
	timeout time.Duration,
	log *utils.AppLogger,
) *DiscountTransactionTargetController {
	return &DiscountTransactionTargetController{
		pgxConn:        conn,
		contextTimeout: timeout,
		log:            log,
	}
}

func (c *DiscountTransactionTargetController) ListDiscountTransactionTarget(ctx context.Context, request reqmodel.ListRequest) ([]interface{}, interface{}, error) {
	// Context Timeout
	reqCtx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	// Get
	transactionsTarget, pagination, err := dbprocess.ListDiscountTransactionTarget(reqCtx, c.pgxConn, request)
	if err != nil {
		c.log.Error("gagal memproses daftar diskon transaksi:", err)
		return nil, nil, utils.StandardError{
			Code:    http.StatusInternalServerError,
			Message: "Gagal mengambil daftar diskon transaksi",
		}
	}
	if len(transactionsTarget) == 0 {
		return nil, nil, utils.StandardError{
			Code:    http.StatusNotFound,
			Message: "Tidak ada data diskon transaksi yang tersedia",
		}
	}

	// Return Success
	return transactionsTarget, pagination, nil
}

func (c *DiscountTransactionTargetController) GetDiscountTransactionTargetByID(ctx context.Context, discountTransactionTargetID string) (*resmodel.DiscountTransactionTarget, error) {
	// Context Timeout
	reqCtx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	// Get
	transactionTarget, err := dbprocess.GetDiscountTransactionTargetByID(reqCtx, c.pgxConn, discountTransactionTargetID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, utils.StandardError{
				Code:    http.StatusNotFound,
				Message: "Diskon transaksi tidak ditemukan",
			}
		}
		c.log.Error("gagal memproses diskon transaksi:", err)
		return nil, utils.StandardError{
			Code:    http.StatusInternalServerError,
			Message: "Gagal mengambil data diskon transaksi",
		}
	}

	// Return Success
	return transactionTarget, nil
}

func (c *DiscountTransactionTargetController) CreateDiscountTransactionTarget(ctx context.Context, request reqmodel.CreateDiscountTransactionTarget) error {
	// Context Timeout
	reqCtx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	// Create
	err := dbprocess.CreateDiscountTransactionTarget(reqCtx, c.pgxConn, request)
	if err != nil {
		c.log.Error("gagal memproses diskon transaksi:", err)
		return utils.StandardError{
			Code:    http.StatusInternalServerError,
			Message: "Gagal membuat diskon transaksi",
		}
	}

	// Return Success
	return nil
}

func (c *DiscountTransactionTargetController) UpdateDiscountTransactionTarget(ctx context.Context, request reqmodel.UpdateDiscountTransactionTarget) error {
	// Context Timeout
	reqCtx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	// Update
	err := dbprocess.UpdateDiscountTransactionTarget(reqCtx, c.pgxConn, request)
	if err != nil {
		if reqErr, ok := err.(utils.RequestError); ok {
			return utils.StandardError{
				Code:    reqErr.StatusCode,
				Message: reqErr.Message,
			}
		}
		c.log.Error("gagal memproses diskon transaksi:", err)
		return utils.StandardError{
			Code:    http.StatusInternalServerError,
			Message: "Gagal mengubah data diskon transaksi",
		}
	}

	// Return Success
	return nil
}

func (c *DiscountTransactionTargetController) DeleteDiscountTransactionTarget(ctx context.Context, discountTransactionTargetID, authUserID string) error {
	// Context Timeout
	reqCtx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	// Delete
	err := dbprocess.DeleteDiscountTransactionTarget(reqCtx, c.pgxConn, discountTransactionTargetID, authUserID)
	if err != nil {
		if reqErr, ok := err.(utils.RequestError); ok {
			return utils.StandardError{
				Code:    reqErr.StatusCode,
				Message: reqErr.Message,
			}
		}
		c.log.Error("gagal memproses diskon transaksi:", err)
		return utils.StandardError{
			Code:    http.StatusInternalServerError,
			Message: "Gagal menghapus data diskon transaksi",
		}
	}

	// Return Success
	return nil
}
