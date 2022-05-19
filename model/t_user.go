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


Table: t_user
[ 0] id                                             INT8                 null: false  primary: true   isArray: false  auto: false  col: INT8            len: -1      default: []
[ 1] email                                          VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] name                                           VARCHAR(25)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 25      default: []
[ 3] password                                       VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 4] surname                                        VARCHAR(25)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 25      default: []
[ 5] username                                       VARCHAR(25)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 25      default: []


JSON Sample
-------------------------------------
{    "id": 86,    "email": "TfJoXBuPsGfABQtJRdaqHQvRI",    "name": "uiXNZgSjrjyyDGIAmqrKxVsqT",    "password": "WSLudKTljKmSbAkyQUVjiiAEi",    "surname": "bbHWNEZYTvkbILotXrReMnZKr",    "username": "JAlhEffcVWRwINcosQqcxjZKN"}



*/

// TUser struct is a row record of the t_user table in the image-labeling database
type TUser struct {
	//[ 0] id                                             INT8                 null: false  primary: true   isArray: false  auto: false  col: INT8            len: -1      default: []
	ID int64 `gorm:"primary_key;column:id;type:INT8;" json:"id"`
	//[ 1] email                                          VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Email null.String `gorm:"column:email;type:VARCHAR;size:255;" json:"email"`
	//[ 2] name                                           VARCHAR(25)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 25      default: []
	Name string `gorm:"column:name;type:VARCHAR;size:25;" json:"name"`
	//[ 3] password                                       VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Password string `gorm:"column:password;type:VARCHAR;size:255;" json:"password"`
	//[ 4] surname                                        VARCHAR(25)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 25      default: []
	Surname string `gorm:"column:surname;type:VARCHAR;size:25;" json:"surname"`
	//[ 5] username                                       VARCHAR(25)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 25      default: []
	Username string `gorm:"column:username;type:VARCHAR;size:25;" json:"username"`
}

var t_userTableInfo = &TableInfo{
	Name: "t_user",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT8",
			DatabaseTypePretty: "INT8",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT8",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "int64",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "email",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Email",
			GoFieldType:        "null.String",
			JSONFieldName:      "email",
			ProtobufFieldName:  "email",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(25)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       25,
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "password",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Password",
			GoFieldType:        "string",
			JSONFieldName:      "password",
			ProtobufFieldName:  "password",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "surname",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(25)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       25,
			GoFieldName:        "Surname",
			GoFieldType:        "string",
			JSONFieldName:      "surname",
			ProtobufFieldName:  "surname",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "username",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(25)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       25,
			GoFieldName:        "Username",
			GoFieldType:        "string",
			JSONFieldName:      "username",
			ProtobufFieldName:  "username",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},
	},
}

// TableName sets the insert table name for this struct type
func (t *TUser) TableName() string {
	return "t_user"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (t *TUser) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (t *TUser) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (t *TUser) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (t *TUser) TableInfo() *TableInfo {
	return t_userTableInfo
}
