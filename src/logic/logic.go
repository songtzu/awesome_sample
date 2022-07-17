package logic

import (
	"awesome"
	"awesome/alog"
	"awesome/anet"
	"awesome/defs"
	"awesome/framework"
	"awesome/pb_protocol"
	"fmt"
	"http_logic"
	"log"
)

type AwesomeImplement struct{}

func (AwesomeImplement) OnMatchPlayers(players []*framework.PlayerImpl, isTimeout bool) {
	log.Printf("自动匹配的回调事件,列表%v, 是否超时:%v", players, isTimeout)
}

func (AwesomeImplement)OnDispatchLogicMessage( roomCode defs.RoomCode, room *awesome.Room, user *awesome.Player, msg *anet.PackHead) (err error){
	l,e:=user.SendPackage(msg)
	alog.Debug(l,e)
	user.LogInfo("OnDispatchLogicMessage====>",msg.Cmd,"the cmd:", string(msg.Body))
	return nil
}


func (AwesomeImplement)OnCreateRoom(msg *anet.PackHead) ( extension interface{}, error error){
	alog.Debug("OnCreateRoom===>", msg.Cmd, string(msg.Body))
	return nil,nil
}

func (AwesomeImplement)OnDispatchSystemMessage(room interface{}, msg *anet.PackHead) (err error){
	alog.Debug("OnDispatchSystemMessage", string(msg.Body), msg.Cmd)
	return nil
}


func (AwesomeImplement)OnError(msg *anet.PackHead){

}


func (AwesomeImplement)OnParseUser(msg *anet.PackHead) (roomCode defs.RoomCode, userId defs.TypeUserId, err error){
	alog.Debug("OnParseRoomCode===>", string(msg.Body),"the cmd:", msg.Cmd)
	return 0, 0, nil
}


func (AwesomeImplement)OnInit() {
	fmt.Println("awesome implement onInit")

}

type UserData struct {
	UID int64
	NickName string
}

func (AwesomeImplement) OnRegisterHttpRouters(e framework.Echo) {
	log.Printf("OnRegisterHttpRouters===>")
	login := pb_protocol.UserLogin{}
	err := framework.RegisterHttpPostHandle("/api/user", http_logic.HandleUserPost )
	err = framework.RegisterHttpGetWithSessionHandle("/api/test", http_logic.HandleUserGet, 30)
	if err != nil {
		return
	}
	log.Println(login)
}
//	OnParseMatch(msg *anet.PackHead)(match *MatchRule, userId defs.TypeUserId)
func (AwesomeImplement)	OnParseMatch(msg *anet.PackHead)(match *framework.MatchRule, userId defs.TypeUserId) {
	log.Printf("解析是否是匹配请求, cmd:%d", msg.Cmd)
	return nil, 0
}

