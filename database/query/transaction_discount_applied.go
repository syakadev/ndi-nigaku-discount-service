package dbquery

// ListDiscount retrieves a paginated list of discounts with optional search
func ListTransactionDiscountApplied() string {
	return `
		SELECT
			id,
			discount_target_id,
			target_id,
			price_before_discount,
			total_discount,
			price_after_discount,
			created_at,
			created_by,
			updated_at,
			updated_by
		FROM
			public.ndi_transaction_discount_applied
		WHERE (
			$1 = '' OR
			price_before_discount ILIKE '%' || $1 || '%' OR
			total_discount ILIKE '%' || $1 || '%' OR
			price_after_discount ILIKE '%' || $1 || '%'
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
			public.ndi_applied_discount
		WHERE (
			$1 = '' OR
			price_before_discount ILIKE '%' || $1 || '%' OR
			total_discount ILIKE '%' || $1 || '%' OR
			price_after_discount ILIKE '%' || $1 || '%'
		) AND
			deleted_at IS NULL;
	`
}

// GetDiscountByID retrieves a single discount record
func GetTransctionDiscountAppliedID() string {
	return `
		SELECT
			id,
			discount_target_id,
			price_before_discount,
			total_discount,
			price_after_discount,
			created_at,
			created_by,
			updated_at,
			updated_by
		FROM
			public.ndi_applied_discount
		WHERE
			discount_id = $1 AND deleted_at IS NULL;
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3;
	`
} 

// CreateDiscount inserts a new discount record
func CreateTransactionDiscountApplied() string {
	return `
		INSERT INTO public.ndi_applied_discount (
			discount_target_id,
			price_before_discount,
			total_discount,
			price_after_discount,
			created_at,
			created_by,
			updated_at,
			updated_by
		) VALUES (
		    $1,  -- discount_target_id
			$2,  -- price_before_discount
			$3,  -- total_discount
			$4,  -- price_after_discount
			NOW(), -- created_at
			$5,  -- created_by
			NOW(), -- updated_at
			$6   -- updated_by
		);
	`
}


// DeleteDiscount performs a soft delete
func DeleteTransactionDiscountApplied() string {
	return `
		UPDATE
			public.ndi_applied_discount
		SET
			deleted_at = NOW(),
			deleted_by = $2
		WHERE
			id = $1;
	`
}