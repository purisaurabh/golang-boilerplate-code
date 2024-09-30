package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/purisaurabh/database-connection/internal/pkg/constants"
)

type Repository interface {
	PostProfileData(context.Context, *sql.Tx, PostProfileRepo) error
	ListProfileData(context.Context, *sql.Tx) ([]ListProfileData, error)
	UpdateProfileData(context.Context, *sql.Tx, UpdateProfileData) error
	DeleteProfileData(context.Context, *sql.Tx, int) error
	BeginTransaction(context.Context) (*sql.Tx, error)
	HandlerTransaction(context.Context, *sql.Tx, error) error
}

type RepositoryStruct struct {
	DB *sql.DB
}

// Transaction Part
func (r *RepositoryStruct) BeginTransaction(ctx context.Context) (*sql.Tx, error) {
	tx, err := r.DB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		fmt.Println("Error in beginning transaction:", err)
		return nil, err
	}
	return tx, nil
}

func (r *RepositoryStruct) HandlerTransaction(ctx context.Context, tx *sql.Tx, inComingErr error) error {
	if inComingErr != nil {
		fmt.Println("Error in incoming error:", inComingErr)
		err := tx.Rollback()
		if err != nil {
			fmt.Println("Error in rolling back transaction:", err)
			return err
		}
		return inComingErr
	}
	err := tx.Commit()
	if err != nil {
		fmt.Println("Error in committing transaction in handler:", err)
		return err
	}
	return nil
}

func (r *RepositoryStruct) PostProfileData(ctx context.Context, tx *sql.Tx, req PostProfileRepo) error {
	insertQuery, args, err := squirrel.Insert("profiles").Columns(constants.PostRequestColumns...).Values(req.Name, req.Email, req.Mobile, req.Created_At, req.Updated_At).ToSql()
	if err != nil {
		fmt.Println("Error in building insert query:", err)
		return err
	}

	if tx == nil {
		fmt.Println("Transaction is nil")
		return fmt.Errorf("internal server error : transaction is nil")
	}

	result, err := tx.ExecContext(ctx, insertQuery, args...)
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

func (r *RepositoryStruct) ListProfileData(ctx context.Context, tx *sql.Tx) ([]ListProfileData, error) {
	var resp []ListProfileData
	query, args, err := squirrel.Select(constants.ListRequestColumns...).From("profiles").ToSql()
	if err != nil {
		fmt.Println("Error in building select query:", err)
		return resp, err
	}

	rows, err := tx.QueryContext(ctx, query, args...)
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

func (r *RepositoryStruct) UpdateProfileData(ctx context.Context, tx *sql.Tx, req UpdateProfileData) error {
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

	if tx == nil {
		fmt.Println("Transaction is nil")
		return fmt.Errorf("internal server error : transaction is nil")
	}

	result, err := tx.ExecContext(ctx, updateQuery, args...)
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

func (r *RepositoryStruct) DeleteProfileData(ctx context.Context, tx *sql.Tx, id int) error {
	deleteQuery, args, err := squirrel.Delete("profiles").Where(squirrel.Eq{"id": id}).ToSql()
	if err != nil {
		fmt.Println("Error in building delete query:", err)
		return err
	}

	if tx == nil {
		fmt.Println("Transaction is nil")
		return fmt.Errorf("internal server error : transaction is nil")
	}

	result, err := tx.ExecContext(ctx, deleteQuery, args...)
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
