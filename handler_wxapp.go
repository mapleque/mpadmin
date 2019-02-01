package mpadmin

import (
	"github.com/mapleque/kelp/http"
	"github.com/mapleque/kelp/mysql"
)

type CreateWXAppParam struct {
	Appid           string `json:"appid" valid:"(0,),message=invalid appid"`
	Secret          string `json:"secret" valid:"(0,),message=invalid secret"`
	Host            string `json:"host"`
	HostToken       string `json:"host_token"`
	MessageKey      string `json:"message_key"`
	MessageApi      string `json:"message_api"`
	MessageApiToken string `json:"message_api_token"`
	Additional      string `json:"additional"`
}

func (this *Server) WXAppCreate(in *CreateWXAppParam, out interface{}, c *http.Context) *http.Status {
	userId := c.MustGet("user_id").(int64)
	count := &mysql.Count{}
	if err := this.conn.QueryOne(
		count,
		"SELECT COUNT(*) AS total FROM `app_user` WHERE `user_id` = ?",
		userId,
	); err != nil {
		return http.STATUS_ERROR_DB
	}
	if count.Total > MAX_APP_PER_USER {
		return http.StatusDiy("can not create more app, up to max app per user limited")
	}

	trans, err := this.conn.Begin()
	if err != nil {
		return http.STATUS_ERROR_DB
	}
	id, err := trans.Insert(
		"INSERT INTO `app` (appid,secret,host,host_token,message_key,message_api,message_api_token,additional) VALUES (?,?,?,?,?,?,?,?)",
		in.Appid,
		in.Secret,
		in.Host,
		in.HostToken,
		in.MessageKey,
		in.MessageApi,
		in.MessageApiToken,
		in.Additional,
	)
	if err != nil || id == 0 {
		trans.Rollback()
		return http.STATUS_ERROR_DB
	}
	if _, err := trans.Insert(
		"INSERT INTO `app_user` (app_id,user_id) VALUES (?,?)",
		id,
		userId,
	); err != nil {
		trans.Rollback()
		return http.STATUS_ERROR_DB
	}
	trans.Commit()
	return nil
}

type UpdateWXAppParam struct {
	Id   int64             `json:"id" valid:"(0,),message=invalid id"`
	Data *CreateWXAppParam `json:"data"`
}

func (this *Server) WXAppUpdate(in *UpdateWXAppParam, out interface{}, c *http.Context) *http.Status {
	userId := c.MustGet("user_id").(int64)

	if err := this.conn.QueryOne(
		nil,
		"SELECT * FROM `app_user` WHERE `app_id`=? AND `user_id` = ?",
		in.Id,
		userId,
	); err != nil {
		return http.STATUS_ERROR_DB
	}
	if _, err := this.conn.Execute(
		"UPDATE `app` SET secret=?,host=?,host_token=?,message_key=?,message_api=?,message_api_token=?,addtional=? WHERE id=?",
		in.Data.Secret,
		in.Data.Host,
		in.Data.HostToken,
		in.Data.MessageKey,
		in.Data.MessageApi,
		in.Data.MessageApiToken,
		in.Data.Additional,
		in.Id,
	); err != nil {
		return http.STATUS_ERROR_DB
	}
	return nil
}

type DeleteWXAppParam struct {
	Id int64 `json:"id" valid:"(0,),message=invalid id"`
}

func (this *Server) WXAppDelete(in *DeleteWXAppParam, out interface{}, c *http.Context) *http.Status {
	userId := c.MustGet("user_id").(int64)

	if err := this.conn.QueryOne(
		nil,
		"SELECT * FROM `app_user` WHERE `app_id`=? AND `user_id` = ?",
		in.Id,
		userId,
	); err != nil {
		return http.STATUS_ERROR_DB
	}
	if _, err := this.conn.Execute(
		"DELETE FROM `app` WHERE id=? LIMIT 1",
		in.Id,
	); err != nil {
		return http.STATUS_ERROR_DB
	}
	return nil
}

type RetrieveWXAppResponse struct {
	List []*App `json:"list"`
}

func (this *Server) WXAppRetrieve(in interface{}, out *RetrieveWXAppResponse, c *http.Context) *http.Status {
	userId := c.MustGet("user_id").(int64)
	out.List = []*App{}
	if err := this.conn.Query(
		&out.List,
		"SELECT * FROM `app` WHERE `id` IN (SELECT `app_id` FROM `app_user` WHERE `user_id` = ?)",
		userId,
	); err != nil {
		return http.STATUS_ERROR_DB
	}
	return nil
}
