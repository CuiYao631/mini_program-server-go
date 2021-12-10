/*
 * @Author: CuiYao
 * @Date: 2021-12-10 16:32:00
 * @Last Modified by:   CuiYao
 * @Last Modified time: 2021-12-10 16:32:00
 */

package entity

type User struct {
	ID       string   `json:"id"`
	Name     string   `json:"Name"`
	Password string   `json:"password"`
	UserType UserType `json:"userType"`
}
