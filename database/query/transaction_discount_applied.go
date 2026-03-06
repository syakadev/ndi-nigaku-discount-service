package dbquery

// ListDiscount retrieves a paginated list of discounts with optional search
func ListTransactionDiscountApplied() string {
	return `
		SELECT
			id,
			discount_transaction_target_id,
			target_id, -- transaksi id target
			price_before_discount,
			total_discount,
			price_after_discount,
			created_at,
			created_by
		FROM
			public.ndi_discount_transaction_applied
		WHERE (
			$1 = '' OR
			price_before_discount::text = $1 OR
			total_discount::text ILIKE '%' || $1 || '%' OR
			price_after_discount::text ILIKE '%' || $1 || '%'
		) AND
			deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3;
	`
}

// CountListDiscount counts the total record for pagination
func CountListTransactionDiscountApplied() string {
	return `
		SELECT
			COUNT(*)
		FROM
			public.ndi_discount_transaction_applied
		WHERE (
			$1 = '' OR
			price_before_discount::text = $1 OR
			total_discount::text ILIKE '%' || $1 || '%' OR
			total_discount::text ILIKE '%' || $1 || '%' OR
			price_after_discount::text ILIKE '%' || $1 || '%'
		) AND
			deleted_at IS NULL;
	`
}

func CountListTransactionDiscountAppliedByDiscountID() string {
	return `
		SELECT
			COUNT(*)
		FROM
			public.ndi_discount_transaction_applied
		WHERE (
			discount_transaction_target_id = $1
		) AND
			deleted_at IS NULL;
	`
}

// GetDiscountByID retrieves a single discount record
func GetListTransctionDiscountAppliedByDiscountID() string {
	return `
		SELECT
			id,
			discount_transaction_target_id,
			target_id, -- transaksi id target
			price_before_discount,
			total_discount,
			price_after_discount,
			created_at,
			created_by
		FROM
			public.ndi_discount_transaction_applied
		WHERE
			discount_transaction_target_id = $1
		AND
			deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3;
	`
}

// CreateDiscount inserts a new discount record
func CreateTransactionDiscountApplied() string {
	return `
		INSERT INTO public.ndi_discount_transaction_applied (
			discount_transaction_target_id,
			target_id, -- transaksi id target
			price_before_discount,
			total_discount,
			price_after_discount,
			created_by
		) VALUES (
			$1,  -- discount_transaction_target_id
			$2,  -- target_id
			$3,  -- price_before_discount
			$4,  -- total_discount
			$5,  -- price_after_discount
			$6  -- created_by
		);
	`
}

// DeleteDiscount performs a soft delete
func DeleteTransactionDiscountApplied() string {
	return `
		UPDATE
			public.ndi_discount_transaction_applied
		SET
			deleted_at = NOW(),
			deleted_by = $2
		WHERE
			id = $1;
	`
}
