package http_logic

import (
	"awesome/awe_util"
	"awesome/framework"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type UserData struct {
	Mobile   string `json:"mobile" query:"mobile" gorm:"mobile" xorm:"mobile"`
	PassWord string `json:"pass_word" query:"pass_word" gorm:"pass_word" xorm:"pass_word"`
	HeadImg  string `json:"head_img" query:"head_img" gorm:"head_img" xorm:"head_img"`
	NickName string `json:"nick_name" query:"nick_name" gorm:"nick_name" xorm:"nick_name"`
	Email    string `json:"email" query:"email" gorm:"email" xorm:"email"`
	UserName string `json:"user_name" query:"user_name" gorm:"user_name" xorm:"user_name"`
	UserDesc string `json:"user_desc" query:"user_desc"`
}

//UserPost 注册。
func HandleUserPost(ctx framework.EchoCtx) (err error) {
	log.Println("UserPost.....")
	val := &UserData{}
	if err = ctx.Bind(val); err != nil {
		log.Printf("参数格式错误 %v", err)
		return ctx.JSON(http.StatusOK, &framework.GeneralResponse{Status: framework.ErrorParameterFormatError, Message: err.Error()})
	}
	log.Printf("===注册用户信息== %v", val)


	if len(val.Email) < 1 && len(val.Mobile) < 1 {
		return ctx.JSON(http.StatusOK, &framework.GeneralResponse{Status: framework.ErrorFieldCannotBeEmpty, Message: "邮箱和手机号不能都为空"})
	}
	if len(val.PassWord) < 1 {
		return ctx.JSON(http.StatusOK, &framework.GeneralResponse{Status: framework.ErrorFieldCannotBeEmpty, Message: "密码不能为空"})
	}

	if len(val.NickName) == 0 {
		if len(val.Mobile) > 0 {
			val.NickName = val.Mobile
		} else {
			val.NickName = val.Email
		}
	}
	val.PassWord = awe_util.PasswordMd5(val.PassWord)
	bin,_:=json.Marshal(val)
	framework.SessionSet(ctx,string(bin),7*24*time.Hour)
	return ctx.JSON(http.StatusOK, &framework.GeneralResponse{Status: framework.ErrorOk, Message: "完成注册", Data: val})
}

func HandleUserGet(ctx framework.EchoCtx,sessionData string) error {
	rsp := &UserData{}
	log.Println("sessionData",sessionData)
	if err:=json.Unmarshal([]byte(sessionData),rsp);err!=nil{
		return ctx.JSON(http.StatusBadRequest, &framework.GeneralResponse{Status: framework.ErrorPermissionNotAllowed,Message: fmt.Sprintf("json解析错误:%s",err.Error())})
	}else{
		return ctx.JSON(http.StatusBadRequest, &framework.GeneralResponse{Status: framework.ErrorOk,Message: "成功",Data: rsp})
	}
}

