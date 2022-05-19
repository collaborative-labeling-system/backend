package api

import (
	"net/http"

	"backend/dao"
	"backend/model"

	"github.com/gin-gonic/gin"
	"github.com/guregu/null"
	"github.com/julienschmidt/httprouter"
)

var (
	_ = null.Bool{}
)

func configLabelTypeRouter(router *httprouter.Router) {
	router.GET("/labeltype", GetAllLabelType)
	router.POST("/labeltype", AddLabelType)
	router.GET("/labeltype/:argID", GetLabelType)
	router.PUT("/labeltype/:argID", UpdateLabelType)
	router.DELETE("/labeltype/:argID", DeleteLabelType)
}

func configGinLabelTypeRouter(router gin.IRoutes) {
	router.GET("/labeltype", ConverHttprouterToGin(GetAllLabelType))
	router.POST("/labeltype", ConverHttprouterToGin(AddLabelType))
	router.GET("/labeltype/:argID", ConverHttprouterToGin(GetLabelType))
	router.PUT("/labeltype/:argID", ConverHttprouterToGin(UpdateLabelType))
	router.DELETE("/labeltype/:argID", ConverHttprouterToGin(DeleteLabelType))
}

// GetAllLabelType is a function to get a slice of record(s) from label_type table in the image-labeling database
// @Summary Get list of LabelType
// @Tags LabelType
// @Description GetAllLabelType is a handler to get a slice of record(s) from label_type table in the image-labeling database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.LabelType}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /labeltype [get]
// http "http://localhost:8080/labeltype?page=0&pagesize=20" X-Api-User:user123
func GetAllLabelType(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	pagesize, err := readInt(r, "pagesize", 20)
	if err != nil || pagesize <= 0 {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	order := r.FormValue("order")

	if err := ValidateRequest(ctx, r, "label_type", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllLabelType(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetLabelType is a function to get a single record from the label_type table in the image-labeling database
// @Summary Get record from table LabelType by  argID
// @Tags LabelType
// @ID argID
// @Description GetLabelType is a function to get a single record from the label_type table in the image-labeling database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.LabelType
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /labeltype/{argID} [get]
// http "http://localhost:8080/labeltype/1" X-Api-User:user123
func GetLabelType(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "label_type", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetLabelType(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddLabelType add to add a single record to label_type table in the image-labeling database
// @Summary Add an record to label_type table
// @Description add to add a single record to label_type table in the image-labeling database
// @Tags LabelType
// @Accept  json
// @Produce  json
// @Param LabelType body model.LabelType true "Add LabelType"
// @Success 200 {object} model.LabelType
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /labeltype [post]
// echo '{"id": 64,"name": "LdHScGOXvVsWIxiPKcnSnZjeW","project_id": 62}' | http POST "http://localhost:8080/labeltype" X-Api-User:user123
func AddLabelType(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	labeltype := &model.LabelType{}

	if err := readJSON(r, labeltype); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := labeltype.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	labeltype.Prepare()

	if err := labeltype.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "label_type", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	labeltype, _, err = dao.AddLabelType(ctx, labeltype)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, labeltype)
}

// UpdateLabelType Update a single record from label_type table in the image-labeling database
// @Summary Update an record in table label_type
// @Description Update a single record from label_type table in the image-labeling database
// @Tags LabelType
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  LabelType body model.LabelType true "Update LabelType record"
// @Success 200 {object} model.LabelType
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /labeltype/{argID} [put]
// echo '{"id": 64,"name": "LdHScGOXvVsWIxiPKcnSnZjeW","project_id": 62}' | http PUT "http://localhost:8080/labeltype/1"  X-Api-User:user123
func UpdateLabelType(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	labeltype := &model.LabelType{}
	if err := readJSON(r, labeltype); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := labeltype.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	labeltype.Prepare()

	if err := labeltype.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "label_type", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	labeltype, _, err = dao.UpdateLabelType(ctx,
		argID,
		labeltype)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, labeltype)
}

// DeleteLabelType Delete a single record from label_type table in the image-labeling database
// @Summary Delete a record from label_type
// @Description Delete a single record from label_type table in the image-labeling database
// @Tags LabelType
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.LabelType
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /labeltype/{argID} [delete]
// http DELETE "http://localhost:8080/labeltype/1" X-Api-User:user123
func DeleteLabelType(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "label_type", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteLabelType(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
