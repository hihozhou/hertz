package models

import "hertz/datebase"

type Menu struct {
	datebase.Model
	Id         int
	Name       string
	Parent_id  string
	Url        string
	Remark     string
	Type       string
	Icon_class string
	datebase.SoftDeletes
	Child []*Menu
}
