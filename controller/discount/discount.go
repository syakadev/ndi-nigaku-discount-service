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

type DiscountController struct {
	contextTimeout time.Duration
	pgxConn        *pgxpool.Pool
	log            *utils.AppLogger
}

func NewDiscountController(
	conn *pgxpool.Pool,
	timeout time.Duration,
	log *utils.AppLogger,
) *DiscountController {
	return &DiscountController{
		pgxConn:        conn,
		contextTimeout: timeout,
		log:            log,
	}
}

func (c *DiscountController) ListDiscount(ctx context.Context, request reqmodel.ListRequest) ([]interface{}, interface{}, error) {
	// Context Timeout
	reqCtx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	// Get
	discounts, pagination, err := dbprocess.ListDiscount(reqCtx, c.pgxConn, request)
	if err != nil {
		c.log.Error("gagal memproses daftar diskon:", err)
		return nil, nil, utils.StandardError{
			Code:    http.StatusInternalServerError,
			Message: "Gagal mengambil daftar diskon",
		}
	}
	if len(discounts) == 0 {
		return nil, nil, utils.StandardError{
			Code:    http.StatusNotFound,
			Message: "Tidak ada data diskon yang tersedia",
		}
	}

	// Return Success
	return discounts, pagination, nil
}

func (c *DiscountController) GetDiscountByID(ctx context.Context, discountID string) (*resmodel.Discount, error) {
	// Context Timeout
	reqCtx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	// Get
	discount, err := dbprocess.GetDiscountByID(reqCtx, c.pgxConn, discountID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, utils.StandardError{
				Code:    http.StatusNotFound,
				Message: "Diskon tidak ditemukan",
			}
		}
		c.log.Error("gagal memproses diskon:", err)
		return nil, utils.StandardError{
			Code:    http.StatusInternalServerError,
			Message: "Gagal mengambil data diskon",
		}
	}

	// Return Success
	return discount, nil
}

func (c *DiscountController) CreateDiscount(ctx context.Context, request reqmodel.CreateDiscount) error {
	// Context Timeout
	reqCtx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	// Create
	err := dbprocess.CreateDiscount(reqCtx, c.pgxConn, request)
	if err != nil {
		c.log.Error("gagal memproses diskon:", err)
		return utils.StandardError{
			Code:    http.StatusInternalServerError,
			Message: "Gagal membuat diskon",
		}
	}

	// Return Success
	return nil
}

func (c *DiscountController) UpdateDiscount(ctx context.Context, request reqmodel.UpdateDiscount) error {
	// Context Timeout
	reqCtx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	// Update
	err := dbprocess.UpdateDiscount(reqCtx, c.pgxConn, request)
	if err != nil {
		if reqErr, ok := err.(utils.RequestError); ok {
			return utils.StandardError{
				Code:    reqErr.StatusCode,
				Message: reqErr.Message,
			}
		}
		c.log.Error("gagal memproses diskon:", err)
		return utils.StandardError{
			Code:    http.StatusInternalServerError,
			Message: "Gagal mengubah data diskon",
		}
	}

	// Return Success
	return nil
}

func (c *DiscountController) DeleteDiscount(ctx context.Context, discountID, authUserID string) error {
	// Context Timeout
	reqCtx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	// Delete
	err := dbprocess.DeleteDiscount(reqCtx, c.pgxConn, discountID, authUserID)
	if err != nil {
		if reqErr, ok := err.(utils.RequestError); ok {
			return utils.StandardError{
				Code:    reqErr.StatusCode,
				Message: reqErr.Message,
			}
		}
		c.log.Error("gagal memproses diskon:", err)
		return utils.StandardError{
			Code:    http.StatusInternalServerError,
			Message: "Gagal menghapus data diskon",
		}
	}

	// Return Success
	return nil
}
