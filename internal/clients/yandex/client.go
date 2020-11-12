package yandex

import (
	"encoding/json"
	"errors"
	"github.com/Solar-2020/Account-Backend/internal/models"
	"github.com/valyala/fasthttp"
)

type Client interface {
	GetUserInfo(userToken string) (user models.YandexUser, err error)
}

type client struct {
}

func NewClient() Client {
	return &client{}
}

func (c *client) GetUserInfo(userToken string) (user models.YandexUser, err error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.Add("Authorization", "OAuth "+userToken)

	req.URI().SetScheme("https")
	req.URI().SetHost("login.yandex.ru")
	req.URI().SetPath("info")

	err = fasthttp.Do(req, resp)
	if err != nil {
		return
	}

	switch resp.StatusCode() {
	case fasthttp.StatusOK:
		var response models.YandexUser
		err = json.Unmarshal(resp.Body(), &response)
		return response, err
	default:
		return user, errors.New("Unexpected Server Error")
	}
}
