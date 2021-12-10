/*
 * @Author: CuiYao
 * @Date: 2021-09-06 15:17:26
 * @Last Modified by: CuiYao
 * @Last Modified time: 2021-09-07 09:44:53
 */
package request

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
)

type GetRequestResponse interface {
	Errorcode() int
}

//获取微信公众号AccessToken
func GetAccessToken(client *http.Client, appid, Secret string, resData GetRequestResponse) error {
	url := "https://api.weixin.qq.com/cgi-bin/token"
	//将token拼接至url
	newUrl := appendTokenToUrl(url, appid, Secret)
	//发起get请求
	resp, err := client.Get(newUrl)
	if err != nil {
		log.Println("Get request failed: ", url, err)
		return errors.New("Request failed")
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Println("Get request failed: ", url, err)
		return errors.New("Request failed")
	}
	//解析返回数据
	json.NewDecoder(resp.Body).Decode(resData)
	log.Println(resData)
	return nil
}
func appendTokenToUrl(u, Appid, Secret string) string {
	parsedUrl, _ := url.Parse(u)
	q, _ := url.ParseQuery(parsedUrl.RawQuery)
	q.Add("grant_type", "client_credential")
	q.Add("appid", Appid)
	q.Add("secret", Secret)
	parsedUrl.RawQuery = q.Encode()
	return parsedUrl.String()
}

//获取Jsapi_Ticket
func GetWeiXinJsapi_Ticket(client *http.Client, accessToken string, resData GetRequestResponse) error {
	URL := "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=" + accessToken + "&type=jsapi"
	resp, err := client.Get(URL)
	if err != nil {
		log.Println("Get request failed: ", URL, err)
		return errors.New("Request failed")
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Println("Get request failed: ", URL, err)
		return errors.New("Request failed")
	}
	//解析返回数据
	json.NewDecoder(resp.Body).Decode(resData)
	log.Println(resData)
	return nil
}
