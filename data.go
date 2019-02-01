package mpadmin

const (
	MAX_APP_PER_USER = 5
)

type App struct {
	Id              int64  `json:"id"`
	Appid           string `json:"appid"`
	Secret          string `json:"secret"`
	Host            string `json:"host"`
	HostToken       string `json:"host_token"`
	MessageKey      string `json:"message_key"`
	MessageApi      string `json:"message_api"`
	MessageApiToken string `json:"message_api_token"`
	Additional      string `json:"additional"`
}

type AppUser struct {
	Id     int64  `json:"id"`
	Appid  string `json:"appid" valid:"(0,),message=invalid appid"`
	UserId int64  `json:"user_id"`
}
