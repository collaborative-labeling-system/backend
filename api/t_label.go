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

func configTLabelRouter(router *httprouter.Router) {
	router.GET("/tlabel", GetAllTLabel)
	router.POST("/tlabel", AddTLabel)
	router.GET("/tlabel/:argID", GetTLabel)
	router.PUT("/tlabel/:argID", UpdateTLabel)
	router.DELETE("/tlabel/:argID", DeleteTLabel)
}

func configGinTLabelRouter(router gin.IRoutes) {
	router.GET("/tlabel", ConverHttprouterToGin(GetAllTLabel))
	router.POST("/tlabel", ConverHttprouterToGin(AddTLabel))
	router.GET("/tlabel/:argID", ConverHttprouterToGin(GetTLabel))
	router.PUT("/tlabel/:argID", ConverHttprouterToGin(UpdateTLabel))
	router.DELETE("/tlabel/:argID", ConverHttprouterToGin(DeleteTLabel))
}

// GetAllTLabel is a function to get a slice of record(s) from t_label table in the image-labeling database
// @Summary Get list of TLabel
// @Tags TLabel
// @Description GetAllTLabel is a handler to get a slice of record(s) from t_label table in the image-labeling database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.TLabel}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /tlabel [get]
// http "http://localhost:8080/tlabel?page=0&pagesize=20" X-Api-User:user123
func GetAllTLabel(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "t_label", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllTLabel(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetTLabel is a function to get a single record from the t_label table in the image-labeling database
// @Summary Get record from table TLabel by  argID
// @Tags TLabel
// @ID argID
// @Description GetTLabel is a function to get a single record from the t_label table in the image-labeling database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.TLabel
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /tlabel/{argID} [get]
// http "http://localhost:8080/tlabel/1" X-Api-User:user123
func GetTLabel(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "t_label", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetTLabel(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddTLabel add to add a single record to t_label table in the image-labeling database
// @Summary Add an record to t_label table
// @Description add to add a single record to t_label table in the image-labeling database
// @Tags TLabel
// @Accept  json
// @Produce  json
// @Param TLabel body model.TLabel true "Add TLabel"
// @Success 200 {object} model.TLabel
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /tlabel [post]
// echo '{"id": 51,"comment": "krDBLCmxUlAGZPrEiLRRhYCoR","created_date": "2040-04-09T11:40:32.6710092+03:00","height": "fsqnFahEdqyKwgejxOpkIKtRM","width": "NtLsicIFjXxUTVQNpSGirQfJq","x": "uXRMgWyXXXkoaoFOTOiVfRGjx","y": "CjZsKIFBXjdULMVexdnERnUdW","image_id": 60,"user_id": 46}' | http POST "http://localhost:8080/tlabel" X-Api-User:user123
func AddTLabel(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	tlabel := &model.TLabel{}

	if err := readJSON(r, tlabel); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := tlabel.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	tlabel.Prepare()

	if err := tlabel.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "t_label", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	tlabel, _, err = dao.AddTLabel(ctx, tlabel)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, tlabel)
}

// UpdateTLabel Update a single record from t_label table in the image-labeling database
// @Summary Update an record in table t_label
// @Description Update a single record from t_label table in the image-labeling database
// @Tags TLabel
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  TLabel body model.TLabel true "Update TLabel record"
// @Success 200 {object} model.TLabel
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /tlabel/{argID} [put]
// echo '{"id": 51,"comment": "krDBLCmxUlAGZPrEiLRRhYCoR","created_date": "2040-04-09T11:40:32.6710092+03:00","height": "fsqnFahEdqyKwgejxOpkIKtRM","width": "NtLsicIFjXxUTVQNpSGirQfJq","x": "uXRMgWyXXXkoaoFOTOiVfRGjx","y": "CjZsKIFBXjdULMVexdnERnUdW","image_id": 60,"user_id": 46}' | http PUT "http://localhost:8080/tlabel/1"  X-Api-User:user123
func UpdateTLabel(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	tlabel := &model.TLabel{}
	if err := readJSON(r, tlabel); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := tlabel.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	tlabel.Prepare()

	if err := tlabel.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "t_label", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	tlabel, _, err = dao.UpdateTLabel(ctx,
		argID,
		tlabel)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, tlabel)
}

// DeleteTLabel Delete a single record from t_label table in the image-labeling database
// @Summary Delete a record from t_label
// @Description Delete a single record from t_label table in the image-labeling database
// @Tags TLabel
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.TLabel
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /tlabel/{argID} [delete]
// http DELETE "http://localhost:8080/tlabel/1" X-Api-User:user123
func DeleteTLabel(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "t_label", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteTLabel(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
