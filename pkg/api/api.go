package api

import (
	"encoding/json"
	"github.com/Solar-2020/Account-Backend/pkg/models"
)

// POST /account/user
type CreateRequest struct {
	User models.User `json:"user"`
}

type CreateResponse struct {
	User models.User `json:"user"`
}

func (r *CreateResponse) Decode(src []byte) (err error) {
	err = json.Unmarshal(src, r)
	return
}
