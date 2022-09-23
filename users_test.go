package zentao

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsersGetByID(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)
	mux.HandleFunc("/api.php/v1/users/1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"id":1,"company":0,"type":"inside","dept":0,"account":"admin","role":"","realname":"admin","pinyin":"","nickname":"","commiter":"","avatar":"","birthday":"0000-00-00","gender":"f","email":"","skype":"","qq":"","mobile":"","phone":"","weixin":"","dingding":"","slack":"","whatsapp":"","address":"","zipcode":"","nature":"","analysis":"","strategy":"","join":"0000-00-00","visits":1,"visions":"rnd,lite","ip":"172.77.77.1","last":"2022-08-31 20:10:29","fails":0,"locked":"0000-00-00 00:00:00","feedback":"0","ranzhi":"","ldap":"","score":0,"scoreLevel":0,"resetToken":"","deleted":"0","clientStatus":"offline","clientLang":"zh-cn"}`)
	})
	expected := &UserProfile{
		UserMeta: UserMeta{
			ID:       1,
			Dept:     0,
			Account:  "admin",
			Realname: "admin",
			Role:     "",
			Email:    "",
			Avatar:   "",
		},
		UserSocial: UserSocial{
			Company:    0,
			Type:       InSideUser,
			Scorelevel: 0,
			Last:       "2022-08-31 20:10:29",
			Deleted:    "0",
			Skype:      "",
			Gender:     WomanGender,
			Address:    "",
			Birthday:   "0000-00-00",
			IP:         "172.77.77.1",
			Visits:     1,
			Join:       "0000-00-00",
			Locked:     "0000-00-00 00:00:00",
			Ranzhi:     "",
		},
		Pinyin:       "",
		Visions:      "rnd,lite",
		Feedback:     "0",
		Clientstatus: "offline",
		Clientlang:   "zh-cn",
	}
	requests, resp, err := client.Users.GetByID(1)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, expected, requests)
}
