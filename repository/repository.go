/*
 * @Author: CuiYao
 * @Date: 2021-12-10 16:31:15
 * @Last Modified by: CuiYao
 * @Last Modified time: 2021-12-10 16:51:30
 */

package repository

import "github.com/CuiYao631/mini_program-server-go/ent"

type Repository interface {
	User
}
type repository struct {
	db *ent.Client
}

func MakeRepository(db *ent.Client) *repository {
	return &repository{db: db}
}
