package logic

import (

	"awesome/anet"
	"awesome/defs"
	"awesome"

	"awesome/alog"
)

type AwesomeImplement struct{}



func (AwesomeImplement)OnDispatchLogicMessage( roomCode defs.RoomCode, room *awesome.Room, user *awesome.Player, msg *anet.PackHead) (err error){
	l,e:=user.SendPackage(msg)
	alog.Debug(l,e)
	user.LogInfo("OnDispatchLogicMessage====>",msg.Cmd,"the cmd:", string(msg.Body))
	//alog.Trace("OnDispatchLogicMessage===>", string(msg.Body),"the cmd:", msg.Cmd)

	return nil
}

func (AwesomeImplement)OnDispatchHallMessage(msg *anet.PackHead){

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


func (AwesomeImplement)OnParseRoomCode(msg *anet.PackHead) (roomCode defs.RoomCode,err error){
	alog.Debug("OnParseRoomCode===>", string(msg.Body),"the cmd:", msg.Cmd)
	return 0,nil
}


