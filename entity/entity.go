/*
 * @Author: CuiYao
 * @Date: 2021-12-10 16:32:00
 * @Last Modified by:   CuiYao
 * @Last Modified time: 2021-12-10 16:32:00
 */

package entity

import (
	"time"
)

type User struct {
	ID       string   `json:"id"`
	Name     string   `json:"Name"`
	Password string   `json:"password"`
	UserType UserType `json:"userType"`
}

//轮播图

type RotationImage struct {
	ID          string `json:"id"`
	Url         string `json:"url"`
	ResourcesID string `json:"resourcesId"`
}

//工具资源
type Resources struct {
	ID       string    `json:"id"`
	Icon     string    `json:"icon"`
	Name     string    `json:"name"`
	Tag      string    `json:"tag"`
	Desc     string    `json:"desc"`
	Explain  string    `json:"explain"` //说明
	Url      string    `json:"url"`
	Topping  bool      `json:"topping"` //是否置顶
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
}

type Wallpaper struct {
	Links []string `json:"links"`
}
