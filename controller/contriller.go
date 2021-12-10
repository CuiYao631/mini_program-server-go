/*
 * @Author: CuiYao
 * @Date: 2021-12-10 16:31:29
 * @Last Modified by: CuiYao
 * @Last Modified time: 2021-12-10 17:06:25
 */

package controller

import "github.com/CuiYao631/mini_program-server-go/usecase"

type Controller interface {
	//创建资源
	//更新资源
	//获取资源
	//删除资源
}
type controller struct {
	uc usecase.Usecase
}

func MakeController(uc usecase.Usecase) *controller {
	return &controller{uc: uc}
}
