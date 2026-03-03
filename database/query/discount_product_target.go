package dbquery

func ListDiscountProductTarget() string {
	return `
		SELECT
			id,
			discount_id,
			target_id,
			product_name,
			max_total_quota,
			price_before_discount,
			total_discount,
			price_after_discount,
			created_at,
			created_by,
			updated_at,
			updated_by
		FROM
			public.ndi_discount_transaction_target
		WHERE (
			$1 = '' OR
			discount_id::text = $1 OR
			product_name ILIKE '%' || $1 || '%' OR
			max_total_quota::text = $1 OR
			price_before_discount::text = $1 OR
			total_discount::text = $1 OR
			price_after_discount::text = $1 OR
			target_id::text = $1
			is_active::text = $1
		) AND
			deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3;
	`
}

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

func GetDiscountProductTargetByID() string {
	return `
		SELECT
			id,
			discount_id,
			target_id,
			product_name,
			max_total_quota,
			price_before_discount,
			total_discount,
			price_after_discount,
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

func GetDiscountProductTargetByDiscountID() string {
	return `
		SELECT
			id,
			discount_id,
			target_id,
			product_name,
			max_total_quota,
			price_before_discount,
			total_discount,
			price_after_discount,
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

func CreateDiscountProductTarget() string {
	return `
		INSERT INTO public.ndi_discount_transaction_target (
			discount_id,
			target_id,
			product_name
			max_total_quota,
			price_before_discount,
			total_discount,
			price_after_discount,
			created_at,
			created_by,
			updated_at,
			updated_by
		) VALUES (
		    $1,  -- discount_id
			$2,  -- target_id
			$3,  -- product_name
			$4,  -- max_total_quota
			$5,  -- price_before_discount
			$6,  -- total_discount
			$7,  -- price_after_discount
			NOW(), -- created_at
			$8,  -- created_by
			NOW(), -- updated_at
			$8   -- updated_by

		);
	`
}

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