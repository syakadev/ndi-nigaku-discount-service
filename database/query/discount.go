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
			target,
			total_quota,
			created_at,
			created_by,
			updated_at,
			updated_by
		FROM 
			discount
		WHERE (
			$1 = '' OR 
			name ILIKE '%' || $1 || '%' OR
			target ILIKE '%' || $1 || '%'
		) AND
			deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3;
	`
}

// CountListDiscount counts the total record for pagination
func CountListDiscount() string {
	return `
		SELECT 
			COUNT(*)
		FROM 
			discount
		WHERE (
			$1 = '' OR 
			name ILIKE '%' || $1 || '%' OR
			target ILIKE '%' || $1 || '%'
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
			target,
			total_quota,
			created_at,
			created_by,
			updated_at,
			updated_by
		FROM 
			discount
		WHERE
			id = $1 AND deleted_at IS NULL;
	`
}

// CreateDiscount inserts a new discount record
func CreateDiscount() string {
	return `
		INSERT INTO discount (
			name,
			type,
			value,
			start_date,
			end_date,
			target,
			total_quota,
			created_at,
			created_by,
			updated_at,
			updated_by
		) VALUES (
			$1,  -- name
			$2,  -- type
			$3,  -- value
			$4,  -- start_date
			$5,  -- end_date
			$6,  -- target
			$7,  -- total_quota
			NOW(), -- created_at
			$8,  -- created_by
			NOW(), -- updated_at
			$8   -- updated_by
		) RETURNING id;
	`
}

// UpdateDiscount updates an existing discount record with COALESCE for optional fields
func UpdateDiscount() string {
	return `
		UPDATE
			discount 
		SET 
			name = COALESCE(NULLIF($2, ''), name),
			type = COALESCE(NULLIF($3, ''), type),
			value = $4,
			start_date = $5,
			end_date = $6,
			target = $7,
			total_quota = $8,
			updated_at = NOW(),
			updated_by = $9
		WHERE
			id = $1 AND deleted_at IS NULL;
	`
}

// DeleteDiscount performs a soft delete
func DeleteDiscount() string {
	return `
		UPDATE
			discount 
		SET 
			deleted_at = NOW(),
			deleted_by = $2
		WHERE
			id = $1;
	`
}
