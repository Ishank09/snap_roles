package model

type Db_Connection struct {
	Server       string `json:"server"`
	Port         int64  `json:"port"`
	Password     string `json:"password"`
	UserId       string `json:"userId"`
	DatabaseName string `json:"databaseName"`
}
