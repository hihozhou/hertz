package config

import (
	"fmt"
	"hertz/app/models"
)

var (
	Menus []models.Menu
)

func init() {
	//todo 通过数据库获取
	Menus = []models.Menu{
		{Name: "首页", Url: "/admin/index/console"},
		{Name: "管理员", Url: "",
			Child: []*models.Menu{
				{Name: "管理员列表", Url: "/admin/admin"},
			},
		},
	}
	fmt.Println(Menus)
}
