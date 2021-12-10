/*
 * @Author: CuiYao
 * @Date: 2021-09-06 15:24:35
 * @Last Modified by: CuiYao
 * @Last Modified time: 2021-09-07 09:55:42
 */
package request

import "math/rand"

type GetAccTokenResponse struct {
	Errcode     int    `json:"errcode"`
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
	Ticket      string `json:"ticket"`
}

func (r *GetAccTokenResponse) Errorcode() int {
	return r.Errcode
}

//生成随机字符串
func GetRandomString() string {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	var result []byte
	for i := 0; i < 15; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}
func CreatenTimestamp() {

}
