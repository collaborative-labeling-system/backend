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

func configTUserRouter(router *httprouter.Router) {
	router.GET("/tuser", GetAllTUser)
	router.POST("/tuser", AddTUser)
	router.GET("/tuser/:argID", GetTUser)
	router.PUT("/tuser/:argID", UpdateTUser)
	router.DELETE("/tuser/:argID", DeleteTUser)
}

func configGinTUserRouter(router gin.IRoutes) {
	router.GET("/tuser", ConverHttprouterToGin(GetAllTUser))
	router.POST("/tuser", ConverHttprouterToGin(AddTUser))
	router.GET("/tuser/:argID", ConverHttprouterToGin(GetTUser))
	router.PUT("/tuser/:argID", ConverHttprouterToGin(UpdateTUser))
	router.DELETE("/tuser/:argID", ConverHttprouterToGin(DeleteTUser))
}

// GetAllTUser is a function to get a slice of record(s) from t_user table in the image-labeling database
// @Summary Get list of TUser
// @Tags TUser
// @Description GetAllTUser is a handler to get a slice of record(s) from t_user table in the image-labeling database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.TUser}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /tuser [get]
// http "http://localhost:8080/tuser?page=0&pagesize=20" X-Api-User:user123
func GetAllTUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "t_user", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllTUser(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetTUser is a function to get a single record from the t_user table in the image-labeling database
// @Summary Get record from table TUser by  argID
// @Tags TUser
// @ID argID
// @Description GetTUser is a function to get a single record from the t_user table in the image-labeling database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.TUser
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /tuser/{argID} [get]
// http "http://localhost:8080/tuser/1" X-Api-User:user123
func GetTUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "t_user", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetTUser(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddTUser add to add a single record to t_user table in the image-labeling database
// @Summary Add an record to t_user table
// @Description add to add a single record to t_user table in the image-labeling database
// @Tags TUser
// @Accept  json
// @Produce  json
// @Param TUser body model.TUser true "Add TUser"
// @Success 200 {object} model.TUser
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /tuser [post]
// echo '{"id": 86,"email": "TfJoXBuPsGfABQtJRdaqHQvRI","name": "uiXNZgSjrjyyDGIAmqrKxVsqT","password": "WSLudKTljKmSbAkyQUVjiiAEi","surname": "bbHWNEZYTvkbILotXrReMnZKr","username": "JAlhEffcVWRwINcosQqcxjZKN"}' | http POST "http://localhost:8080/tuser" X-Api-User:user123
func AddTUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	tuser := &model.TUser{}

	if err := readJSON(r, tuser); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := tuser.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	tuser.Prepare()

	if err := tuser.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "t_user", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	tuser, _, err = dao.AddTUser(ctx, tuser)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, tuser)
}

// UpdateTUser Update a single record from t_user table in the image-labeling database
// @Summary Update an record in table t_user
// @Description Update a single record from t_user table in the image-labeling database
// @Tags TUser
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  TUser body model.TUser true "Update TUser record"
// @Success 200 {object} model.TUser
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /tuser/{argID} [put]
// echo '{"id": 86,"email": "TfJoXBuPsGfABQtJRdaqHQvRI","name": "uiXNZgSjrjyyDGIAmqrKxVsqT","password": "WSLudKTljKmSbAkyQUVjiiAEi","surname": "bbHWNEZYTvkbILotXrReMnZKr","username": "JAlhEffcVWRwINcosQqcxjZKN"}' | http PUT "http://localhost:8080/tuser/1"  X-Api-User:user123
func UpdateTUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	tuser := &model.TUser{}
	if err := readJSON(r, tuser); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := tuser.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	tuser.Prepare()

	if err := tuser.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "t_user", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	tuser, _, err = dao.UpdateTUser(ctx,
		argID,
		tuser)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, tuser)
}

// DeleteTUser Delete a single record from t_user table in the image-labeling database
// @Summary Delete a record from t_user
// @Description Delete a single record from t_user table in the image-labeling database
// @Tags TUser
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.TUser
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /tuser/{argID} [delete]
// http DELETE "http://localhost:8080/tuser/1" X-Api-User:user123
func DeleteTUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "t_user", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteTUser(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
