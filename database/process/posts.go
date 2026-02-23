package dbprocess

import (
	"context"
	"fmt"

	dbquery "template/service/auth/database/query"
	reqmodel "template/service/auth/model/request"
	resmodel "template/service/auth/model/response"
	"template/service/auth/utils"
)

func ListPost(ctx context.Context, exec DBExecutor, request reqmodel.ListRequest) ([]interface{}, *resmodel.PaginationData, error) {
	// Query
	query := dbquery.ListPost()
	offset := (request.Page - 1) * request.Size
	rows, err := exec.Query(ctx, query, request.Search, request.Size, offset)
	if err != nil {
		fmt.Println("testinggg")
		return nil, nil, err
	}
	defer rows.Close()

	var posts []interface{}
	for rows.Next() {
		var post resmodel.Post
		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.IsActive,
			&post.CreatedAt,
			&post.CreatedBy,
			&post.UpdatedAt,
			&post.UpdatedBy,
		)
		if err != nil {
			return nil, nil, err
		}
		posts = append(posts, post)
	}
	if len(posts) == 0 {
		return nil, nil, nil
	}

	// Query Total Data
	var total int
	err = exec.QueryRow(ctx, dbquery.CountListPost(), request.Search).Scan(&total)
	if err != nil {
		return nil, nil, err
	}

	pagination := &resmodel.PaginationData{
		Page:      request.Page,
		Size:      request.Size,
		TotalData: total,
	}
	return posts, pagination, nil
}

func GetPostByID(ctx context.Context, exec DBExecutor, postID string) (*resmodel.Post, error) {
	// Query
	query := dbquery.GetPostByID()
	row := exec.QueryRow(ctx, query, postID)
	var post resmodel.Post
	err := row.Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.IsActive,
		&post.CreatedAt,
		&post.CreatedBy,
		&post.UpdatedAt,
		&post.UpdatedBy,
	)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func CreatePost(ctx context.Context, exec DBExecutor, request reqmodel.CreatePost) error {
	// Query
	_, err := exec.Exec(ctx, dbquery.CreatePost(),
		request.Title,
		request.Content,
		request.IsActive,
		request.AuthUserID,
	)
	if err != nil {
		return err
	}

	// Return Success
	return nil
}

func UpdatePost(ctx context.Context, exec DBExecutor, request reqmodel.UpdatePost) error {
	// Query
	result, err := exec.Exec(ctx, dbquery.UpdatePost(),
		request.ID,
		request.Title,
		request.Content,
		request.IsActive,
		request.AuthUserID,
	)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return utils.RequestError{
			StatusCode: 404,
			Message:    "Gagal melakukan perubahan, post tidak ditemukan",
		}
	}

	// Return Success
	return nil
}

func DeletePost(ctx context.Context, exec DBExecutor, postID, authUserID string) error {
	// Query
	result, err := exec.Exec(ctx, dbquery.DeletePost(),
		postID,
		authUserID,
	)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return utils.RequestError{
			StatusCode: 404,
			Message:    "Gagal melakukan penghapusan, post tidak ditemukan",
		}
	}

	// Return Success
	return nil
}
