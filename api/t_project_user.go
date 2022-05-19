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

func configTProjectUserRouter(router *httprouter.Router) {
	router.GET("/tprojectuser", GetAllTProjectUser)
	router.POST("/tprojectuser", AddTProjectUser)
	router.GET("/tprojectuser/:argProjectID", GetTProjectUser)
	router.PUT("/tprojectuser/:argProjectID", UpdateTProjectUser)
	router.DELETE("/tprojectuser/:argProjectID", DeleteTProjectUser)
}

func configGinTProjectUserRouter(router gin.IRoutes) {
	router.GET("/tprojectuser", ConverHttprouterToGin(GetAllTProjectUser))
	router.POST("/tprojectuser", ConverHttprouterToGin(AddTProjectUser))
	router.GET("/tprojectuser/:argProjectID", ConverHttprouterToGin(GetTProjectUser))
	router.PUT("/tprojectuser/:argProjectID", ConverHttprouterToGin(UpdateTProjectUser))
	router.DELETE("/tprojectuser/:argProjectID", ConverHttprouterToGin(DeleteTProjectUser))
}

// GetAllTProjectUser is a function to get a slice of record(s) from t_project_user table in the image-labeling database
// @Summary Get list of TProjectUser
// @Tags TProjectUser
// @Description GetAllTProjectUser is a handler to get a slice of record(s) from t_project_user table in the image-labeling database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.TProjectUser}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /tprojectuser [get]
// http "http://localhost:8080/tprojectuser?page=0&pagesize=20" X-Api-User:user123
func GetAllTProjectUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "t_project_user", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllTProjectUser(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetTProjectUser is a function to get a single record from the t_project_user table in the image-labeling database
// @Summary Get record from table TProjectUser by  argProjectID
// @Tags TProjectUser
// @ID argProjectID
// @Description GetTProjectUser is a function to get a single record from the t_project_user table in the image-labeling database
// @Accept  json
// @Produce  json
// @Param  argProjectID path int64 true "project_id"
// @Success 200 {object} model.TProjectUser
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /tprojectuser/{argProjectID} [get]
// http "http://localhost:8080/tprojectuser/1" X-Api-User:user123
func GetTProjectUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argProjectID, err := parseInt64(ps, "argProjectID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "t_project_user", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetTProjectUser(ctx, argProjectID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddTProjectUser add to add a single record to t_project_user table in the image-labeling database
// @Summary Add an record to t_project_user table
// @Description add to add a single record to t_project_user table in the image-labeling database
// @Tags TProjectUser
// @Accept  json
// @Produce  json
// @Param TProjectUser body model.TProjectUser true "Add TProjectUser"
// @Success 200 {object} model.TProjectUser
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /tprojectuser [post]
// echo '{"project_id": 63,"user_id": 87}' | http POST "http://localhost:8080/tprojectuser" X-Api-User:user123
func AddTProjectUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	tprojectuser := &model.TProjectUser{}

	if err := readJSON(r, tprojectuser); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := tprojectuser.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	tprojectuser.Prepare()

	if err := tprojectuser.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "t_project_user", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	tprojectuser, _, err = dao.AddTProjectUser(ctx, tprojectuser)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, tprojectuser)
}

// UpdateTProjectUser Update a single record from t_project_user table in the image-labeling database
// @Summary Update an record in table t_project_user
// @Description Update a single record from t_project_user table in the image-labeling database
// @Tags TProjectUser
// @Accept  json
// @Produce  json
// @Param  argProjectID path int64 true "project_id"
// @Param  TProjectUser body model.TProjectUser true "Update TProjectUser record"
// @Success 200 {object} model.TProjectUser
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /tprojectuser/{argProjectID} [put]
// echo '{"project_id": 63,"user_id": 87}' | http PUT "http://localhost:8080/tprojectuser/1"  X-Api-User:user123
func UpdateTProjectUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argProjectID, err := parseInt64(ps, "argProjectID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	tprojectuser := &model.TProjectUser{}
	if err := readJSON(r, tprojectuser); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := tprojectuser.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	tprojectuser.Prepare()

	if err := tprojectuser.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "t_project_user", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	tprojectuser, _, err = dao.UpdateTProjectUser(ctx,
		argProjectID,
		tprojectuser)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, tprojectuser)
}

// DeleteTProjectUser Delete a single record from t_project_user table in the image-labeling database
// @Summary Delete a record from t_project_user
// @Description Delete a single record from t_project_user table in the image-labeling database
// @Tags TProjectUser
// @Accept  json
// @Produce  json
// @Param  argProjectID path int64 true "project_id"
// @Success 204 {object} model.TProjectUser
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /tprojectuser/{argProjectID} [delete]
// http DELETE "http://localhost:8080/tprojectuser/1" X-Api-User:user123
func DeleteTProjectUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argProjectID, err := parseInt64(ps, "argProjectID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "t_project_user", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteTProjectUser(ctx, argProjectID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
