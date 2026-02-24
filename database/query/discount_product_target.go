package dbquery

// ListDiscount retrieves a paginated list of discounts with optional search
func ListDiscountProductTarget() string {
	return `
		SELECT
			id,
			discount_id,
			target_id,
			max_total_quota,
			created_at,
			created_by,
			updated_at,
			updated_by
		FROM
			public.ndi_discount_transaction_target
		WHERE (
			$1 = '' OR
			max_total_quota ILIKE '%' || $1 || '%'
		) AND
			deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3;
	`
}

// CountListDiscount counts the total record for pagination
func CountListDiscountProductTarget() string {
	return `
		SELECT
			COUNT(*)
		FROM
			public.ndi_discount_transaction_target
		WHERE (
			$1 = '' OR
			max_total_quota ILIKE '%' || $1 || '%'
		) AND
			deleted_at IS NULL;
	`
}

// GetDiscountByID retrieves a single discount record
func GetListDiscountProductTargetByDiscountID() string {
	return `
		SELECT
			id,
			discount_id,
			target_id,
			max_total_quota,
			created_at,
			created_by,
			updated_at,
			updated_by
		FROM
			public.ndi_discount_transaction_target
		WHERE
			discount_id = $1 AND deleted_at IS NULL;
	`
} 

// CreateDiscount inserts a new discount record
func CreateDiscountPorductTarget() string {
	return `
		INSERT INTO public.ndi_discount_transaction_target (
			discount_id,
			target_id,
			max_total_quota,
			created_at,
			created_by,
			updated_at,
			updated_by
		) VALUES (
		    $1,  -- discount_id
			$2,  -- max_total_quota
			$3,  -- target_id
			NOW(), -- created_at
			$4,  -- created_by
			NOW(), -- updated_at
			$5   -- updated_by
		);
	`
}

// UpdateDiscount updates an existing discount record with COALESCE for optional fields
func UpdateDiscountProductTarget() string {
	return `
		UPDATE
			public.ndi_discount_transaction_target
		SET
			target_id = $2,
			max_total_quota = $3,
			updated_at = NOW(),
			updated_by = $4
		WHERE
			id = $1
		AND
		deleted_at IS NULL;
	`
}

// DeleteDiscount performs a soft delete
func DeleteDiscountProductTarget() string {
	return `
		UPDATE
			public.ndi_discount_transaction_target
		SET
			deleted_at = NOW(),
			deleted_by = $2
		WHERE
			id = $1;
	`
}