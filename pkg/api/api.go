package api

import (
	"encoding/json"
	"github.com/Solar-2020/Account-Backend/pkg/models"
	service "github.com/Solar-2020/GoUtils/http"
	"net/url"
	"strconv"
)

// POST /account/user
type CreateRequest struct {
	models.User
}

type CreateResponse struct {
	models.User
}

func (r *CreateResponse) Decode(src []byte) (err error) {
	err = json.Unmarshal(src, r)
	return
}

type GetUserRequest struct {
	Uid int `json:"uid"`
	Email string `json:"email"`
}
func (r *GetUserRequest) Encode() (queryString string, err error) {
	if r.Email != "" {
		queryString = url.QueryEscape(r.Email)
		return
	}
	if r.Uid != 0 {
		queryString = url.QueryEscape(strconv.Itoa(r.Uid))
		return
	}
	return
}


type AccountServiceInterface interface {
	GetUserByUid(userID int) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	UidToEmail(userID int) (string, error)
	CreateUser(request CreateRequest) (CreateResponse, error)
}

type AccountClient struct {
	service.Service
	Addr string
}
func (c *AccountClient) Address () string { return c.Addr }

func (c *AccountClient) UidToEmail(userID int) (email string, err error) {
	user, err := c.GetUserByUid(userID)
	if err != nil {
		return
	}
	email = user.Email
	return
}

func (c *AccountClient) GetUserByUid(userID int) (models.User, error) {
	return c.GetUser(userID, "", "/api/internal/account/by-user")
}
func (c *AccountClient) GetUserByEmail(email string) (models.User, error){
	return c.GetUser(0, email, "/api/internal/account/by-email")
}
func (c *AccountClient) GetUser(userID int, email string, address string) (res models.User, err error) {
	req := GetUserRequest{
		Uid:     userID,
		Email: email,
	}

	endpoint := service.ServiceEndpoint{
		Service:  c,
		Endpoint: address,
		Method:   "GET",
	}
	err = endpoint.Send(&req, &res)
	return
}

func (c *AccountClient) CreateUser(request CreateRequest) (resp CreateResponse, err error) {
	endpoint := service.ServiceEndpoint{
		Service:  c,
		Endpoint: "/api/internal/account/user",
		Method:   "POST",
	}
	err = endpoint.Send(&request, &resp)
	return
}
