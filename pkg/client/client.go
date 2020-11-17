package client

import (
	"encoding/json"
	"fmt"
	"github.com/Solar-2020/Account-Backend/pkg/models"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
	"strconv"
)

type Client interface {
	GetUserByUid(userID int) (user models.User, err error)
	GetUserByEmail(email string) (user models.User, err error)
	GetYandexUser(userToken string) (user models.User, err error)
	CreateUser(request models.User) (userID int, err error)
}

type client struct {
	host   string
	secret string
}

func NewClient(host string, secret string) Client {
	return &client{host: host, secret: secret}
}

func (c *client) GetUserByUid(userID int) (user models.User, err error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.URI().SetScheme("http")
	req.URI().SetHost(c.host)
	req.URI().SetPath(fmt.Sprintf("api/internal/account/by-user/%s", strconv.Itoa(userID)))

	req.Header.SetMethod(fasthttp.MethodGet)
	req.Header.Set("Authorization", c.secret)

	err = fasthttp.Do(req, resp)
	if err != nil {
		return
	}

	switch resp.StatusCode() {
	case fasthttp.StatusOK:
		var response models.User
		err = json.Unmarshal(resp.Body(), &response)
		return response, err
	case fasthttp.StatusBadRequest:
		var httpErr httpError
		err = json.Unmarshal(resp.Body(), &httpErr)
		if err != nil {
			return
		}
		return user, errors.New(httpErr.Error)
	default:
		return user, ResponseError{
			StatusCode: resp.StatusCode(),
			Message:    InternalServerStatus,
			Err:        errors.Errorf(ErrorUnknownStatusCode, resp.StatusCode()),
		}
	}
}

func (c *client) GetUserByEmail(email string) (user models.User, err error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.URI().SetScheme("http")
	req.URI().SetHost(c.host)
	req.URI().SetPath(fmt.Sprintf("api/internal/account/by-email/%s", email))

	req.Header.SetMethod(fasthttp.MethodGet)
	req.Header.Set("Authorization", c.secret)

	err = fasthttp.Do(req, resp)
	if err != nil {
		return
	}

	switch resp.StatusCode() {
	case fasthttp.StatusOK:
		var response models.User
		err = json.Unmarshal(resp.Body(), &response)
		return response, err
	case fasthttp.StatusBadRequest:
		var httpErr httpError
		err = json.Unmarshal(resp.Body(), &httpErr)
		if err != nil {
			return
		}
		return user, errors.New(httpErr.Error)
	default:
		return user, ResponseError{
			StatusCode: resp.StatusCode(),
			Message:    InternalServerStatus,
			Err:        errors.Errorf(ErrorUnknownStatusCode, resp.StatusCode()),
		}
	}
}

func (c *client) GetYandexUser(userToken string) (user models.User, err error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod(fasthttp.MethodGet)
	req.Header.Add("Authorization", c.secret)

	req.URI().SetScheme("http")
	req.URI().SetHost(c.host)
	req.URI().SetPath("api/internal/account/yandex/" + userToken)

	err = fasthttp.Do(req, resp)
	if err != nil {
		return
	}

	switch resp.StatusCode() {
	case fasthttp.StatusOK:
		var response models.User
		err = json.Unmarshal(resp.Body(), &response)
		return response, err
	case fasthttp.StatusBadRequest:
		var httpErr httpError
		err = json.Unmarshal(resp.Body(), &httpErr)
		if err != nil {
			return
		}
		return user, errors.New(httpErr.Error)
	default:
		return user, ResponseError{
			StatusCode: resp.StatusCode(),
			Message:    InternalServerStatus,
			Err:        errors.Errorf(ErrorUnknownStatusCode, resp.StatusCode()),
		}
	}
}

func (c *client) CreateUser(request models.User) (userID int, err error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	body, err := json.Marshal(request)
	if err != nil {
		return
	}

	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.Add("Authorization", c.secret)

	req.URI().SetScheme("http")
	req.URI().SetHost(c.host)
	req.URI().SetPath("api/internal/account/user")

	req.SetBody(body)

	err = fasthttp.Do(req, resp)
	if err != nil {
		return
	}

	switch resp.StatusCode() {
	case fasthttp.StatusOK:
		var response models.User
		err = json.Unmarshal(resp.Body(), &response)
		return response.ID, err
	case fasthttp.StatusBadRequest:
		var httpErr httpError
		err = json.Unmarshal(resp.Body(), &httpErr)
		if err != nil {
			return
		}
		return userID, errors.New(httpErr.Error)
	default:
		return userID, ResponseError{
			StatusCode: resp.StatusCode(),
			Message:    InternalServerStatus,
			Err:        errors.Errorf(ErrorUnknownStatusCode, resp.StatusCode()),
		}
	}
}
