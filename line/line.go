package line

import (
	"go_line_bot/entity"
	"net/http"
	"net/url"
)

const baseURL = "https://access.line.me"

const urlPath = "/oauth2/v2.1/authorize"

type OAuth struct {
	admin      *entity.Admin
	httpClient *http.Client
}

func New(admin *entity.Admin) *OAuth {
	oauth := &OAuth{admin, &http.Client{}}
	return oauth
}

// 別のパッケージで呼び出す場合は関数名を大文字で始める必要がある
func (oauth *OAuth) AuthUri(state string) string {
	params := map[string]string{
		"response_type": "code",
		"client_id":     "1656382872",                                      //ToDo: 今は直書きだがDBからの取り出しに変更
		"redirect_uri":  "https://1aa1-153-187-96-194.ngrok.io/1/callback", //仮
		"state":         state,
		"scope":         "openid",
		"prompt":        "consent",
		"bot_prompt":    "aggressive",
	}
	endpoint := baseURL + urlPath
	u, _ := url.Parse(endpoint)
	q := u.Query()
	for key, value := range params {
		q.Set(key, value)
	}
	u.RawQuery = q.Encode()

	return u.String()
}
