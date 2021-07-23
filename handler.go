/**
 * @Author: Resynz
 * @Date: 2021/7/20 17:11
 */
package ws_handler

import (
	"encoding/json"
	"fmt"
	"github.com/rosbit/go-wget"
	"net/http"
)

type WsHandler struct {
	BaseUrl string `json:"base_url"`
}

type BaseRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (s *WsHandler) GetWsUrl(userId string) (string, error) {
	reqUrl := fmt.Sprintf("%s/api/ws-url?user_id=%s", s.BaseUrl, userId)
	method := "GET"
	status, content, _, err := wget.Wget(reqUrl, method, nil, nil)
	if err != nil {
		return "", err
	}
	if status != http.StatusOK {
		return "", fmt.Errorf("bad http status:%d", status)
	}

	type res struct {
		*BaseRes
		Data struct {
			WsUrl string `json:"ws_url"`
		} `json:"data"`
	}

	var r res
	if err = json.Unmarshal(content, &r); err != nil {
		return "", err
	}
	if r.Code != 0 {
		return "", fmt.Errorf("%s", r.Message)
	}
	return r.Data.WsUrl, nil
}

func (s *WsHandler) GetOnlineCount() (int64, error) {
	reqUrl := fmt.Sprintf("%s/api/online-count", s.BaseUrl)
	method := "GET"
	status, content, _, err := wget.Wget(reqUrl, method, nil, nil)
	if err != nil {
		return 0, err
	}
	if status != http.StatusOK {
		return 0, fmt.Errorf("bad http status:%d", status)
	}

	type res struct {
		*BaseRes
		Data struct {
			Count int64 `json:"count"`
		} `json:"data"`
	}

	var r res
	if err = json.Unmarshal(content, &r); err != nil {
		return 0, err
	}
	if r.Code != 0 {
		return 0, fmt.Errorf("%s", r.Message)
	}
	return r.Data.Count, nil
}

type ClientObj struct {
	ClientId   string       `json:"client_id"`
	CreateTime int64        `json:"create_time"`
	Platform   PlatformType `json:"platform"`
}

type UserInfo struct {
	UserId  string       `json:"user_id"`
	Clients []*ClientObj `json:"clients"`
}

func (s *WsHandler) GetUserInfo(userId string) (*UserInfo, error) {
	reqUrl := fmt.Sprintf("%s/api/info?user_id=%s", s.BaseUrl, userId)
	method := "GET"
	status, content, _, err := wget.Wget(reqUrl, method, nil, nil)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, fmt.Errorf("bad http status:%d", status)
	}

	type res struct {
		*BaseRes
		Data struct {
			Info *UserInfo `json:"info"`
		} `json:"data"`
	}

	var r res
	if err = json.Unmarshal(content, &r); err != nil {
		return nil, err
	}
	if r.Code != 0 {
		return nil, fmt.Errorf("%s", r.Message)
	}
	return r.Data.Info, nil
}

func (s *WsHandler) CheckIsOnline(userId string) (bool, error) {
	reqUrl := fmt.Sprintf("%s/api/is-online?user_id=%s", s.BaseUrl, userId)
	method := "GET"
	status, content, _, err := wget.Wget(reqUrl, method, nil, nil)
	if err != nil {
		return false, err
	}
	if status != http.StatusOK {
		return false, fmt.Errorf("bad http status:%d", status)
	}

	type res struct {
		*BaseRes
		Data struct {
			Result bool `json:"result"`
		} `json:"data"`
	}

	var r res
	if err = json.Unmarshal(content, &r); err != nil {
		return false, err
	}
	if r.Code != 0 {
		return false, fmt.Errorf("%s", r.Message)
	}
	return r.Data.Result, nil
}

func (s *WsHandler) SendMsg(userIds, messages, clientIds []string) (bool, error) {
	reqUrl := fmt.Sprintf("%s/api/send-msg", s.BaseUrl)
	method := "POST"
	params := map[string]interface{}{
		"user_id_list":   userIds,
		"msg_list":       messages,
		"client_id_list": clientIds,
	}
	status, content, _, err := wget.PostJson(reqUrl, method, params, nil)
	if err != nil {
		return false, err
	}
	if status != http.StatusOK {
		return false, fmt.Errorf("bad http status:%d", status)
	}
	type res struct {
		*BaseRes
		Data struct {
			Result bool `json:"result"`
		} `json:"data"`
	}

	var r res
	if err = json.Unmarshal(content, &r); err != nil {
		return false, err
	}
	if r.Code != 0 {
		return false, fmt.Errorf("%s", r.Message)
	}
	return r.Data.Result, nil
}

func (s *WsHandler) Broadcast(messages []string) (bool, error) {
	reqUrl := fmt.Sprintf("%s/api/broadcast", s.BaseUrl)
	method := "POST"
	params := map[string]interface{}{
		"msg_list": messages,
	}
	status, content, _, err := wget.PostJson(reqUrl, method, params, nil)
	if err != nil {
		return false, err
	}
	if status != http.StatusOK {
		return false, fmt.Errorf("bad http status:%d", status)
	}
	type res struct {
		*BaseRes
		Data struct {
			Result bool `json:"result"`
		} `json:"data"`
	}

	var r res
	if err = json.Unmarshal(content, &r); err != nil {
		return false, err
	}
	if r.Code != 0 {
		return false, fmt.Errorf("%s", r.Message)
	}
	return r.Data.Result, nil
}
