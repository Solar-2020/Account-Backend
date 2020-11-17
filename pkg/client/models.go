package client

var (
	InternalServerStatus   = "Ошибка сервера, повторите попытку позже"
	ErrorUnknownStatusCode = "Unknown status code %v"
)

type ResponseError struct {
	StatusCode int
	Message    string
	Err        error
}

func (re ResponseError) Error() string {
	return re.Message
}

type CreateUser struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	AvatarURL string `json:"avatarURL"`
}

type httpError struct {
	Error string `json:"error"`
}
