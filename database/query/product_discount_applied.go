package dbquery

// ListDiscount retrieves a paginated list of discounts with optional search
func ListProductDiscountApplied() string {
	return `
		SELECT
			id,
			discount_product_target_id,
			customer_id,
			customer_name,
			transaction_date,
			created_at,
			created_by
		FROM
			public.ndi_discount_product_applied
		WHERE (
			$1 = '' OR
			discount_product_target_id::text = $1 OR
			customer_id::text = $1 OR
			customer_name ILIKE '%' || $1 || '%' OR
			transaction_date::text = $1
		) AND
			deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3;
	`
}

// CountListDiscount counts the total record for pagination
func CountListProductDiscountAppliedDiscount() string {
	return `
		SELECT
			COUNT(*)
		FROM
			public.ndi_discount_product_applied
		WHERE (
			$1 = '' OR
			discount_product_target_id::text = $1 OR
			customer_id::text = $1 OR
			customer_name ILIKE '%' || $1 || '%' OR
			transaction_date::text = $1
		) AND
			deleted_at IS NULL;
	`
}

// GetDiscountByID retrieves a single discount record
func GetListProductDiscountAppliedByProductDiscountID() string {
	return `
		SELECT
			id,
			discount_product_target_id,
			customer_id,
			customer_name,
			transaction_date,
			created_at,
			created_by
		FROM
			public.ndi_discount_product_applied
		WHERE
			discount_product_target_id = $1 AND deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3;
	`
}

// CountListProductDiscountAppliedByProductDiscountID counts the total record for pagination
func CountListProductDiscountAppliedByProductDiscountID() string {
	return `
		SELECT
			COUNT(*)
		FROM
			public.ndi_discount_product_applied
		WHERE
			discount_product_target_id = $1
		AND
			deleted_at IS NULL;
	`
}

// CreateDiscount inserts a new discount record
func CreateProductDiscountApplied() string {
	return `
		INSERT INTO public.ndi_discount_product_applied (
			discount_product_target_id,
			customer_id,
			customer_name,
			transaction_date,
			created_by
		) VALUES (
		    $1,  -- discount_product_target_id
			$2,  -- price_before_discount
			$3,  -- total_discount
			$4,  -- price_after_discount
			$5   -- created_by
		);
	`
}

// DeleteDiscount performs a soft delete
func DeleteProductDiscountApplied() string {
	return `
		UPDATE
			public.ndi_discount_product_applied
		SET
			deleted_at = NOW(),
			deleted_by = $2
		WHERE
			id = $1;
	`
}
