package models


type Images struct {
	Id   	 int    `json:"id" orm:"auto"`
	FileName string `json:"file_name"`
	Size int64 `json:"size"`
	ContentType string  `json:"content_type"`
}