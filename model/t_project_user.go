package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
	"github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
	_ = uuid.UUID{}
)

/*
DB Table Details
-------------------------------------


Table: t_project_user
[ 0] project_id                                     INT8                 null: false  primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 1] user_id                                        INT8                 null: false  primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []


JSON Sample
-------------------------------------
{    "project_id": 63,    "user_id": 87}


Comments
-------------------------------------
[ 0] Warning table: t_project_user does not have a primary key defined, setting col position 1 project_id as primary key




*/

// TProjectUser struct is a row record of the t_project_user table in the image-labeling database
type TProjectUser struct {
	//[ 0] project_id                                     INT8                 null: false  primary: true   isArray: false  auto: false  col: INT8            len: -1      default: []
	ProjectID int64 `gorm:"primary_key;column:project_id;type:INT8;" json:"project_id"`
	//[ 1] user_id                                        INT8                 null: false  primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
	UserID int64 `gorm:"column:user_id;type:INT8;" json:"user_id"`
}

var t_project_userTableInfo = &TableInfo{
	Name: "t_project_user",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "project_id",
			Comment: ``,
			Notes: `Warning table: t_project_user does not have a primary key defined, setting col position 1 project_id as primary key
`,
			Nullable:           false,
			DatabaseTypeName:   "INT8",
			DatabaseTypePretty: "INT8",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT8",
			ColumnLength:       -1,
			GoFieldName:        "ProjectID",
			GoFieldType:        "int64",
			JSONFieldName:      "project_id",
			ProtobufFieldName:  "project_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "user_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT8",
			DatabaseTypePretty: "INT8",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT8",
			ColumnLength:       -1,
			GoFieldName:        "UserID",
			GoFieldType:        "int64",
			JSONFieldName:      "user_id",
			ProtobufFieldName:  "user_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},
	},
}

// TableName sets the insert table name for this struct type
func (t *TProjectUser) TableName() string {
	return "t_project_user"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (t *TProjectUser) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (t *TProjectUser) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (t *TProjectUser) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (t *TProjectUser) TableInfo() *TableInfo {
	return t_project_userTableInfo
}
