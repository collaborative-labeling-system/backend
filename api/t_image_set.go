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

func configTImageSetRouter(router *httprouter.Router) {
	router.GET("/timageset", GetAllTImageSet)
	router.POST("/timageset", AddTImageSet)
	router.GET("/timageset/:argID", GetTImageSet)
	router.PUT("/timageset/:argID", UpdateTImageSet)
	router.DELETE("/timageset/:argID", DeleteTImageSet)
}

func configGinTImageSetRouter(router gin.IRoutes) {
	router.GET("/timageset", ConverHttprouterToGin(GetAllTImageSet))
	router.POST("/timageset", ConverHttprouterToGin(AddTImageSet))
	router.GET("/timageset/:argID", ConverHttprouterToGin(GetTImageSet))
	router.PUT("/timageset/:argID", ConverHttprouterToGin(UpdateTImageSet))
	router.DELETE("/timageset/:argID", ConverHttprouterToGin(DeleteTImageSet))
}

// GetAllTImageSet is a function to get a slice of record(s) from t_image_set table in the image-labeling database
// @Summary Get list of TImageSet
// @Tags TImageSet
// @Description GetAllTImageSet is a handler to get a slice of record(s) from t_image_set table in the image-labeling database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.TImageSet}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /timageset [get]
// http "http://localhost:8080/timageset?page=0&pagesize=20" X-Api-User:user123
func GetAllTImageSet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "t_image_set", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllTImageSet(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetTImageSet is a function to get a single record from the t_image_set table in the image-labeling database
// @Summary Get record from table TImageSet by  argID
// @Tags TImageSet
// @ID argID
// @Description GetTImageSet is a function to get a single record from the t_image_set table in the image-labeling database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.TImageSet
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /timageset/{argID} [get]
// http "http://localhost:8080/timageset/1" X-Api-User:user123
func GetTImageSet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "t_image_set", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetTImageSet(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddTImageSet add to add a single record to t_image_set table in the image-labeling database
// @Summary Add an record to t_image_set table
// @Description add to add a single record to t_image_set table in the image-labeling database
// @Tags TImageSet
// @Accept  json
// @Produce  json
// @Param TImageSet body model.TImageSet true "Add TImageSet"
// @Success 200 {object} model.TImageSet
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /timageset [post]
// echo '{"id": 42,"created_date": "2082-11-23T17:54:52.713906225+03:00","image_count": 47,"is_used": true,"name": "CQvdNOntKAdoAxeWIZPxVdkID","project_id": 77,"user_id": 16}' | http POST "http://localhost:8080/timageset" X-Api-User:user123
func AddTImageSet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	timageset := &model.TImageSet{}

	if err := readJSON(r, timageset); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := timageset.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	timageset.Prepare()

	if err := timageset.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "t_image_set", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	timageset, _, err = dao.AddTImageSet(ctx, timageset)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, timageset)
}

// UpdateTImageSet Update a single record from t_image_set table in the image-labeling database
// @Summary Update an record in table t_image_set
// @Description Update a single record from t_image_set table in the image-labeling database
// @Tags TImageSet
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  TImageSet body model.TImageSet true "Update TImageSet record"
// @Success 200 {object} model.TImageSet
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /timageset/{argID} [put]
// echo '{"id": 42,"created_date": "2082-11-23T17:54:52.713906225+03:00","image_count": 47,"is_used": true,"name": "CQvdNOntKAdoAxeWIZPxVdkID","project_id": 77,"user_id": 16}' | http PUT "http://localhost:8080/timageset/1"  X-Api-User:user123
func UpdateTImageSet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	timageset := &model.TImageSet{}
	if err := readJSON(r, timageset); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := timageset.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	timageset.Prepare()

	if err := timageset.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "t_image_set", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	timageset, _, err = dao.UpdateTImageSet(ctx,
		argID,
		timageset)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, timageset)
}

// DeleteTImageSet Delete a single record from t_image_set table in the image-labeling database
// @Summary Delete a record from t_image_set
// @Description Delete a single record from t_image_set table in the image-labeling database
// @Tags TImageSet
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.TImageSet
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /timageset/{argID} [delete]
// http DELETE "http://localhost:8080/timageset/1" X-Api-User:user123
func DeleteTImageSet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "t_image_set", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteTImageSet(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
