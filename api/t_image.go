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

func configTImageRouter(router *httprouter.Router) {
	router.GET("/timage", GetAllTImage)
	router.POST("/timage", AddTImage)
	router.GET("/timage/:argID", GetTImage)
	router.PUT("/timage/:argID", UpdateTImage)
	router.DELETE("/timage/:argID", DeleteTImage)
}

func configGinTImageRouter(router gin.IRoutes) {
	router.GET("/timage", ConverHttprouterToGin(GetAllTImage))
	router.POST("/timage", ConverHttprouterToGin(AddTImage))
	router.GET("/timage/:argID", ConverHttprouterToGin(GetTImage))
	router.PUT("/timage/:argID", ConverHttprouterToGin(UpdateTImage))
	router.DELETE("/timage/:argID", ConverHttprouterToGin(DeleteTImage))
}

// GetAllTImage is a function to get a slice of record(s) from t_image table in the image-labeling database
// @Summary Get list of TImage
// @Tags TImage
// @Description GetAllTImage is a handler to get a slice of record(s) from t_image table in the image-labeling database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.TImage}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /timage [get]
// http "http://localhost:8080/timage?page=0&pagesize=20" X-Api-User:user123
func GetAllTImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "t_image", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllTImage(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetTImage is a function to get a single record from the t_image table in the image-labeling database
// @Summary Get record from table TImage by  argID
// @Tags TImage
// @ID argID
// @Description GetTImage is a function to get a single record from the t_image table in the image-labeling database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.TImage
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /timage/{argID} [get]
// http "http://localhost:8080/timage/1" X-Api-User:user123
func GetTImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "t_image", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetTImage(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddTImage add to add a single record to t_image table in the image-labeling database
// @Summary Add an record to t_image table
// @Description add to add a single record to t_image table in the image-labeling database
// @Tags TImage
// @Accept  json
// @Produce  json
// @Param TImage body model.TImage true "Add TImage"
// @Success 200 {object} model.TImage
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /timage [post]
// echo '{"id": 17,"name": "DeCIWfDYDMlCqntafiUZXKOSB","url": "mYpToxqLXJlPYUoXyfqGdCuUP","image_set_id": 79,"user_id": 15}' | http POST "http://localhost:8080/timage" X-Api-User:user123
func AddTImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	timage := &model.TImage{}

	if err := readJSON(r, timage); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := timage.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	timage.Prepare()

	if err := timage.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "t_image", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	timage, _, err = dao.AddTImage(ctx, timage)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, timage)
}

// UpdateTImage Update a single record from t_image table in the image-labeling database
// @Summary Update an record in table t_image
// @Description Update a single record from t_image table in the image-labeling database
// @Tags TImage
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  TImage body model.TImage true "Update TImage record"
// @Success 200 {object} model.TImage
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /timage/{argID} [put]
// echo '{"id": 17,"name": "DeCIWfDYDMlCqntafiUZXKOSB","url": "mYpToxqLXJlPYUoXyfqGdCuUP","image_set_id": 79,"user_id": 15}' | http PUT "http://localhost:8080/timage/1"  X-Api-User:user123
func UpdateTImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	timage := &model.TImage{}
	if err := readJSON(r, timage); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := timage.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	timage.Prepare()

	if err := timage.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "t_image", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	timage, _, err = dao.UpdateTImage(ctx,
		argID,
		timage)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, timage)
}

// DeleteTImage Delete a single record from t_image table in the image-labeling database
// @Summary Delete a record from t_image
// @Description Delete a single record from t_image table in the image-labeling database
// @Tags TImage
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.TImage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /timage/{argID} [delete]
// http DELETE "http://localhost:8080/timage/1" X-Api-User:user123
func DeleteTImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "t_image", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteTImage(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
