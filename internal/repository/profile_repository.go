package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/purisaurabh/database-connection/internal/pkg/constants"
)

type Repository interface {
	PostProfileData(context.Context, PostProfileRepo) error
	ListProfileData(context.Context) ([]ListProfileData, error)
	UpdateProfileData(context.Context, UpdateProfileData) error
	DeleteProfileData(context.Context, int) error
}

type RepositoryStruct struct {
	DB *sql.DB
}

func (r *RepositoryStruct) PostProfileData(ctx context.Context, req PostProfileRepo) error {
	insertQuery, args, err := squirrel.Insert("profiles").Columns(constants.PostRequestColumns...).Values(req.Name, req.Email, req.Mobile, req.Created_At, req.Updated_At).ToSql()
	if err != nil {
		fmt.Println("Error in building insert query:", err)
		return err
	}

	result, err := r.DB.ExecContext(ctx, insertQuery, args...)
	if err != nil {
		fmt.Println("Error in executing insert query:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error in getting rows affected:", err)
		return err
	}

	if rowsAffected == 0 {
		fmt.Println("No rows affected")
		return err
	}

	return nil
}

func (r *RepositoryStruct) ListProfileData(ctx context.Context) ([]ListProfileData, error) {
	var resp []ListProfileData
	query, args, err := squirrel.Select(constants.ListRequestColumns...).From("profiles").ToSql()
	if err != nil {
		fmt.Println("Error in building select query:", err)
		return resp, err
	}

	rows, err := r.DB.QueryContext(ctx, query, args...)
	if err != nil {
		fmt.Println("Error in executing select query:", err)
		return resp, err
	}

	for rows.Next() {
		var data ListProfileData
		err = rows.Scan(&data.ID, &data.Name, &data.Email, &data.Mobile, &data.Created_At, &data.Updated_At)
		if err != nil {
			fmt.Println("Error in scanning rows:", err)
			return resp, err
		}
		resp = append(resp, data)
	}

	return resp, nil
}

func (r *RepositoryStruct) UpdateProfileData(ctx context.Context, req UpdateProfileData) error {
	updateQuery, args, err := squirrel.Update("profiles").SetMap(map[string]interface{}{
		"name":       req.Name,
		"email":      req.Email,
		"mobile":     req.Mobile,
		"updated_at": req.Updated_At,
	}).Where(squirrel.Eq{"id": req.ID}).ToSql()
	if err != nil {
		fmt.Println("Error in building update query:", err)
		return err
	}

	result, err := r.DB.ExecContext(ctx, updateQuery, args...)
	if err != nil {
		fmt.Println("Error in executing update query:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error in getting rows affected:", err)
		return err
	}

	if rowsAffected == 0 {
		fmt.Println("No rows affected")
		return err
	}

	return nil
}

func (r *RepositoryStruct) DeleteProfileData(ctx context.Context, id int) error {
	deleteQuery, args, err := squirrel.Delete("profiles").Where(squirrel.Eq{"id": id}).ToSql()
	if err != nil {
		fmt.Println("Error in building delete query:", err)
		return err
	}

	result, err := r.DB.ExecContext(ctx, deleteQuery, args...)
	if err != nil {
		fmt.Println("Error in executing delete query:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error in getting rows affected:", err)
		return err
	}

	if rowsAffected == 0 {
		fmt.Println("No rows affected")
		return nil
	}

	return nil
}
