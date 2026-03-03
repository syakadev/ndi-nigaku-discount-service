package dbquery

// ListDiscount retrieves a paginated list of discounts with optional search
func ListDiscountTransactionTarget() string {
	return `
		SELECT
			id,
			discount_id,
			max_total_quota,
			is_active,
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
func CountListDiscountTransactionTarget() string {
	return `
		SELECT 
			COUNT(*)
		FROM 
			public.ndi_discount_transaction_target
		WHERE (
			$1 = '' OR
			max_total_quota ILIKE '%' || $1 || '%' OR
			discount_id::text = $1
		) AND
			deleted_at IS NULL;
	`
}

func GetDiscountTransactionTargetByID() string {
	return `
		SELECT
			id,
			discount_id,
			max_total_quota,
			is_active,
			created_at,
			created_by,
			updated_at,
			updated_by
		FROM
			public.ndi_discount_transaction_target
		WHERE
			id = $1 AND deleted_at IS NULL;
	`
}

// CreateDiscount inserts a new discount record
func CreateDiscountTransactionTarget() string {
	return `
		INSERT INTO public.ndi_discount_transaction_target (
			discount_id,
			max_total_quota,
			is_active,
			created_at,
			created_by,
			updated_at,
			updated_by
		) VALUES (
		    $1,  -- discount_id
			$2,  -- max_total_quota
			$3,  -- is_active
			NOW(), -- created_at
			$4,  -- created_by
			NOW(), -- updated_at
			$4   -- updated_by
		);
	`
}

// UpdateDiscount updates an existing discount record with COALESCE for optional fields
func UpdateDiscountTransactionTarget() string {
	return `
		UPDATE
			public.ndi_discount_transaction_target
		SET
			max_total_quota = $2,
			is_active = $3,
			updated_at = NOW(),
			updated_by = $4
		WHERE
			id = $1
		AND
		deleted_at IS NULL;
	`
}

// DeleteDiscount performs a soft delete
func DeleteDiscountTransactionTarget() string {
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