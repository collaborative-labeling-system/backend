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

// GetAllLabelType is a function to get a slice of record(s) from label_type table in the image-labeling database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllLabelType(ctx context.Context, page, pagesize int64, order string) (results []*model.LabelType, totalRows int, err error) {

	resultOrm := DB.Model(&model.LabelType{})
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

// GetLabelType is a function to get a single record from the label_type table in the image-labeling database
// error - ErrNotFound, db Find error
func GetLabelType(ctx context.Context, argID int64) (record *model.LabelType, err error) {
	record = &model.LabelType{}
	if err = DB.First(record, argID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddLabelType is a function to add a single record to label_type table in the image-labeling database
// error - ErrInsertFailed, db save call failed
func AddLabelType(ctx context.Context, record *model.LabelType) (result *model.LabelType, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateLabelType is a function to update a single record from label_type table in the image-labeling database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateLabelType(ctx context.Context, argID int64, updated *model.LabelType) (result *model.LabelType, RowsAffected int64, err error) {

	result = &model.LabelType{}
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

// DeleteLabelType is a function to delete a single record from label_type table in the image-labeling database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteLabelType(ctx context.Context, argID int64) (rowsAffected int64, err error) {

	record := &model.LabelType{}
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
