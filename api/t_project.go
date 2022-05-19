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

func configTProjectRouter(router *httprouter.Router) {
	router.GET("/tproject", GetAllTProject)
	router.POST("/tproject", AddTProject)
	router.GET("/tproject/:argID", GetTProject)
	router.PUT("/tproject/:argID", UpdateTProject)
	router.DELETE("/tproject/:argID", DeleteTProject)
}

func configGinTProjectRouter(router gin.IRoutes) {
	router.GET("/tproject", ConverHttprouterToGin(GetAllTProject))
	router.POST("/tproject", ConverHttprouterToGin(AddTProject))
	router.GET("/tproject/:argID", ConverHttprouterToGin(GetTProject))
	router.PUT("/tproject/:argID", ConverHttprouterToGin(UpdateTProject))
	router.DELETE("/tproject/:argID", ConverHttprouterToGin(DeleteTProject))
}

// GetAllTProject is a function to get a slice of record(s) from t_project table in the image-labeling database
// @Summary Get list of TProject
// @Tags TProject
// @Description GetAllTProject is a handler to get a slice of record(s) from t_project table in the image-labeling database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.TProject}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /tproject [get]
// http "http://localhost:8080/tproject?page=0&pagesize=20" X-Api-User:user123
func GetAllTProject(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "t_project", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllTProject(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetTProject is a function to get a single record from the t_project table in the image-labeling database
// @Summary Get record from table TProject by  argID
// @Tags TProject
// @ID argID
// @Description GetTProject is a function to get a single record from the t_project table in the image-labeling database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.TProject
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /tproject/{argID} [get]
// http "http://localhost:8080/tproject/1" X-Api-User:user123
func GetTProject(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "t_project", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetTProject(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddTProject add to add a single record to t_project table in the image-labeling database
// @Summary Add an record to t_project table
// @Description add to add a single record to t_project table in the image-labeling database
// @Tags TProject
// @Accept  json
// @Produce  json
// @Param TProject body model.TProject true "Add TProject"
// @Success 200 {object} model.TProject
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /tproject [post]
// echo '{"id": 94,"created_date": "2261-04-10T00:45:47.316840105+03:00","name": "VMJGCjPVxLWLSPLnnUqMuKMff","admin_id": 6,"ımage_set_id": 65}' | http POST "http://localhost:8080/tproject" X-Api-User:user123
func AddTProject(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	tproject := &model.TProject{}

	if err := readJSON(r, tproject); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := tproject.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	tproject.Prepare()

	if err := tproject.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "t_project", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	tproject, _, err = dao.AddTProject(ctx, tproject)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, tproject)
}

// UpdateTProject Update a single record from t_project table in the image-labeling database
// @Summary Update an record in table t_project
// @Description Update a single record from t_project table in the image-labeling database
// @Tags TProject
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  TProject body model.TProject true "Update TProject record"
// @Success 200 {object} model.TProject
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /tproject/{argID} [put]
// echo '{"id": 94,"created_date": "2261-04-10T00:45:47.316840105+03:00","name": "VMJGCjPVxLWLSPLnnUqMuKMff","admin_id": 6,"ımage_set_id": 65}' | http PUT "http://localhost:8080/tproject/1"  X-Api-User:user123
func UpdateTProject(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	tproject := &model.TProject{}
	if err := readJSON(r, tproject); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := tproject.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	tproject.Prepare()

	if err := tproject.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "t_project", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	tproject, _, err = dao.UpdateTProject(ctx,
		argID,
		tproject)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, tproject)
}

// DeleteTProject Delete a single record from t_project table in the image-labeling database
// @Summary Delete a record from t_project
// @Description Delete a single record from t_project table in the image-labeling database
// @Tags TProject
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.TProject
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /tproject/{argID} [delete]
// http DELETE "http://localhost:8080/tproject/1" X-Api-User:user123
func DeleteTProject(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "t_project", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteTProject(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
