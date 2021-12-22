/*
 * @Author: CuiYao
 * @Date: 2021-12-10 16:32:00
 * @Last Modified by:   CuiYao
 * @Last Modified time: 2021-12-10 16:32:00
 */

package entity

import "time"

type User struct {
	ID       string   `json:"id"`
	Name     string   `json:"Name"`
	Password string   `json:"password"`
	UserType UserType `json:"userType"`
}
type Resources struct {
	ID       string    `json:"id"`
	Icon     string    `json:"icon"`
	Name     string    `json:"name"`
	Tag      string    `json:"tag"`
	Desc     string    `json:"desc"`
	Url      string    `json:"url"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
}
