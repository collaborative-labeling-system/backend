package dao

import (
	"context"
	"time"

	"backend/model"

	"github.com/guregu/null"
	"github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = null.Bool{}
	_ = uuid.UUID{}
)

// GetAllTUser is a function to get a slice of record(s) from t_user table in the image-labeling database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllTUser(ctx context.Context, page, pagesize int64, order string) (results []*model.TUser, totalRows int, err error) {

	resultOrm := DB.Model(&model.TUser{})
	resultOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		resultOrm = resultOrm.Offset(offset).Limit(pagesize)
	} else {
		resultOrm = resultOrm.Limit(pagesize)
	}

	if order != "" {
		resultOrm = resultOrm.Order(order)
	}

	if err = resultOrm.Find(&results).Error; err != nil {
		err = ErrNotFound
		return nil, -1, err
	}

	return results, totalRows, nil
}

// GetTUser is a function to get a single record from the t_user table in the image-labeling database
// error - ErrNotFound, db Find error
func GetTUser(ctx context.Context, argID int64) (record *model.TUser, err error) {
	record = &model.TUser{}
	if err = DB.First(record, argID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddTUser is a function to add a single record to t_user table in the image-labeling database
// error - ErrInsertFailed, db save call failed
func AddTUser(ctx context.Context, record *model.TUser) (result *model.TUser, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateTUser is a function to update a single record from t_user table in the image-labeling database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateTUser(ctx context.Context, argID int64, updated *model.TUser) (result *model.TUser, RowsAffected int64, err error) {

	result = &model.TUser{}
	db := DB.First(result, argID)
	if err = db.Error; err != nil {
		return nil, -1, ErrNotFound
	}

	if err = Copy(result, updated); err != nil {
		return nil, -1, ErrUpdateFailed
	}

	db = db.Save(result)
	if err = db.Error; err != nil {
		return nil, -1, ErrUpdateFailed
	}

	return result, db.RowsAffected, nil
}

// DeleteTUser is a function to delete a single record from t_user table in the image-labeling database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteTUser(ctx context.Context, argID int64) (rowsAffected int64, err error) {

	record := &model.TUser{}
	db := DB.First(record, argID)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
