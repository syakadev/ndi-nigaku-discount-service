package dbquery

// ListDiscount retrieves a paginated list of discounts with optional search
func ListDiscount() string {
	return `
		SELECT 
			id,
			name,
			type,
			value,
			start_date,
			end_date,
			created_at,
			created_by,
			updated_at,
			updated_by
		FROM 
			public.ndi_discount
		WHERE (
			$1 = '' OR 
			name ILIKE '%' || $1 || '%' OR
			type ILIKE '%' || $1 || '%'
		) AND
			deleted_at IS NULL
		LIMIT $2 OFFSET $3;
	`
}

// CountListDiscount counts the total record for pagination
func CountListDiscount() string {
	return `
		SELECT 
			COUNT(*)
		FROM 
			public.ndi_discount
		WHERE (
			$1 = '' OR 
			name ILIKE '%' || $1 || '%' OR
			type ILIKE '%' || $1 || '%'
		) AND 
			deleted_at IS NULL;
	`
}

// GetDiscountByID retrieves a single discount record
func GetDiscountByID() string {
	return `
		SELECT 
			id,
			name,
			type,
			value,
			start_date,
			end_date,
			created_at,
			created_by,
			updated_at,
			updated_by
		FROM
			public.ndi_discount
		WHERE
			id = $1 AND deleted_at IS NULL;
	`
}

// CreateDiscount inserts a new discount record
func CreateDiscount() string {
	return `
		INSERT INTO public.ndi_discount (
			name,
			type,
			value,
			start_date,
			end_date,
			created_by,
			updated_by
		) VALUES (
			$1,  -- name
			$2,  -- type
			$3,  -- value
			$4,  -- start_date
			$5,  -- end_date
			$6,  -- created_by
			$6   -- updated_by
		);
	`
}

// UpdateDiscount updates an existing discount record with COALESCE for optional fields
func UpdateDiscount() string {
	return `
		UPDATE
			public.ndi_discount 
		SET 
			name = COALESCE(NULLIF($2, ''), name),
			type = COALESCE(NULLIF($3, ''), type),
			value = $4,
			start_date = $5,
			end_date = $6,
			updated_at = NOW(),
			updated_by = $7
		WHERE
			id = $1 AND deleted_at IS NULL;
	`
}

// DeleteDiscount performs a soft delete
func DeleteDiscount() string {
	return `
		UPDATE
			public.ndi_discount 
		SET 
			deleted_at = NOW(),
			deleted_by = $2
		WHERE
			id = $1;
	`
}
