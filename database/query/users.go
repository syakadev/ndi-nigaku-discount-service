package dbquery

func ListPost() string {
	return `
		SELECT 
			id,
			title,
			content,
			is_active,
			created_at,
			created_by,
			updated_at,
			updated_by
		FROM 
			ndi_posts
		WHERE (
			$1 = '' OR 
            title ILIKE '%' || $1 || '%' OR
            content ILIKE '%' || $1 || '%'
		) AND
			deleted_at IS NULL
		LIMIT $2 OFFSET $3;
	`
}

func CountListPost() string {
	return `
		SELECT 
			COUNT(*)
		FROM 
			ndi_posts
		WHERE (
			$1 = '' OR 
            title ILIKE '%' || $1 || '%' OR
            content ILIKE '%' || $1 || '%'
		) AND 
			deleted_at IS NULL;
	`
}

func GetPostByID() string {
	return `
		SELECT 
			id,
			title,
			content,
			is_active,
			created_at,
			created_by,
			updated_at,
			updated_by
		FROM 
			ndi_posts
		WHERE
			id = $1 AND deleted_at IS NULL
	`
}

func CreatePost() string {
	return `
		INSERT INTO ndi_posts (
			title,
			content,
			is_active,
			created_at,
			created_by,
			updated_at,
			updated_by
		) VALUES (
			$1,  -- title
			$2,  -- content
			$3,  -- is_active
			NOW(),  -- created_at
			$4  -- created_by
			NOW(),  -- updated_at
			$4  -- updated_by
		)
	`
}

func UpdatePost() string {
	return `
		UPDATE
			ndi_posts 
		SET 
			title = COALESCE(NULLIF($2, ''), title),
			content = $3,
			is_active = $4,
			updated_at = NOW(),
			updated_by = $5
		WHERE
			id = $1
	`
}

func DeletePost() string {
	return `
		UPDATE
			ndi_posts 
		SET 
			deleted_at = NOW(),
			deleted_by = $2
		WHERE
			id = $1
	`
}
