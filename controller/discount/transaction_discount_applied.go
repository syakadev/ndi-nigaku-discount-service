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

type TransactionDiscountAppliedController struct {
	contextTimeout time.Duration
	pgxConn        *pgxpool.Pool
	log            *utils.AppLogger
}

func NewTransactionDiscountAppliedController(
	conn *pgxpool.Pool,
	timeout time.Duration,
	log *utils.AppLogger,
) *TransactionDiscountAppliedController {
	return &TransactionDiscountAppliedController{
		pgxConn:        conn,
		contextTimeout: timeout,
		log:            log,
	}
}

func (c *TransactionDiscountAppliedController) ListTransactionDiscountApplied(ctx context.Context, request reqmodel.ListRequest) ([]interface{}, interface{}, error) {
	// Context Timeout
	reqCtx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	// Get
	transactionDiscountApplieds, pagination, err := dbprocess.ListTransactionDiscountApplied(reqCtx, c.pgxConn, request)
	if err != nil {
		c.log.Error("gagal memproses daftar diskon transaksi yang diterapkan:", err)
		return nil, nil, utils.StandardError{
			Code:    http.StatusInternalServerError,
			Message: "Gagal mengambil daftar diskon transaksi yang diterapkan",
		}
	}
	if len(transactionDiscountApplieds) == 0 {
		return nil, nil, utils.StandardError{
			Code:    http.StatusNotFound,
			Message: "Tidak ada data diskon transaksi yang diterapkan yang tersedia",
		}
	}

	// Return Success
	return transactionDiscountApplieds, pagination, nil
}

func (c *TransactionDiscountAppliedController) ListTransactionDiscountAppliedByTargetID(ctx context.Context, transactionDiscountTargetID string, request reqmodel.ListRequest) ([]interface{}, interface{}, error) {
	// Context Timeout
	reqCtx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	// Get
	transactionDiscountApplieds, pagination, err := dbprocess.ListTransactionDiscountAppliedByDiscountID(reqCtx, c.pgxConn, transactionDiscountTargetID, request)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil, utils.StandardError{
				Code:    http.StatusNotFound,
				Message: "Diskon transaksi yang diterapkan tidak ditemukan",
			}
		}
		c.log.Error("gagal memproses diskon transaksi yang diterapkan:", err)
		return nil, nil, utils.StandardError{
			Code:    http.StatusInternalServerError,
			Message: "Gagal mengambil data diskon transaksi yang diterapkan",
		}
	}

	if len(transactionDiscountApplieds) == 0 {
		return nil, nil, utils.StandardError{
			Code:    http.StatusNotFound,
			Message: "Tidak ada data diskon transaksi yang diterapkan yang tersedia",
		}
	}

	// Return Success
	return transactionDiscountApplieds, pagination, nil
}

func (c *TransactionDiscountAppliedController) CreateTransactionDiscountApplied(ctx context.Context, request reqmodel.CreateTransactionDiscountApplied) error {
	// Context Timeout
	reqCtx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	// Create
	err := dbprocess.CreateTransactionDiscountApplied(reqCtx, c.pgxConn, request)
	if err != nil {
		c.log.Error("gagal memproses diskon transaksi yang diterapkan:", err)
		return utils.StandardError{
			Code:    http.StatusInternalServerError,
			Message: "Gagal membuat diskon transaksi yang diterapkan",
		}
	}

	// Return Success
	return nil
}

func (c *TransactionDiscountAppliedController) DeleteTransactionDiscountApplied(ctx context.Context, transactionDiscountAppliedID, authUserID string) error {
	// Context Timeout
	reqCtx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	// Delete
	err := dbprocess.DeleteTransactionDiscountApplied(reqCtx, c.pgxConn, transactionDiscountAppliedID, authUserID)
	if err != nil {
		if reqErr, ok := err.(utils.RequestError); ok {
			return utils.StandardError{
				Code:    reqErr.StatusCode,
				Message: reqErr.Message,
			}
		}
		c.log.Error("gagal memproses penghapusan diskon transaksi yang diterapkan:", err)
		return utils.StandardError{
			Code:    http.StatusInternalServerError,
			Message: "Gagal menghapus data diskon transaksi yang diterapkan",
		}
	}

	// Return Success
	return nil
}
