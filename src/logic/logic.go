package logic

import (
	"awesome"
	"awesome/alog"
	"awesome/anet"
	"awesome/defs"
	"awesome/framework"
	"awesome/pb_protocol"
	"fmt"
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


func (AwesomeImplement)OnParseRoomCodeAndUser(msg *anet.PackHead) (roomCode defs.RoomCode, userId defs.TypeUserId ,err error){
	alog.Debug("OnParseRoomCode===>", string(msg.Body),"the cmd:", msg.Cmd)
	return 0, 0,nil
}


func (AwesomeImplement)OnInit() {
	fmt.Println("awesome implement onInit")
}

type UserData struct {
	UID int64
	NickName string
}

func (i AwesomeImplement) OnRegisterHttpRouters(e framework.Echo) {
	login := pb_protocol.UserLogin{}
	log.Println(login)
}

//
//func (i AwesomeImplement) OnDispatchLogicMessage(roomCode defs.RoomCode, IRoom *framework.Room, user *framework.PlayerImpl, msg *anet.PackHead) (err error) {
//	roomData, ok := IRoom.GetRoomData().(*room.Room)
//	if !ok {
//		//glog.Errorf("OnRoomMsgDispatch错误的参数类型 %v", reflect.TypeOf(IRoom.GetRoomData()), "msg.Cmd:", msg.Cmd)
//		return errors.New("OnRoomMsgDispatch 房间数据类型不匹配")
//	}
//	if isCmdRouterExist(msg.Cmd) {
//		hd := getRouterFunc(msg.Cmd)
//		t := getCmdRouterProto(msg.Cmd)
//		v := reflect.New(t)
//		if err := proto.Unmarshal(msg.Body, v.Interface().(proto.Message)); err == nil {
//			res := hd.Call([]reflect.Value{reflect.ValueOf(roomData), v, reflect.ValueOf(user)})
//			if !res[0].IsNil() {
//				framework.SendUserMsg(user, res[1].Interface().(int), res[0].Interface())
//			}
//		} else {
//			glog.Errorln("protocol  unmarshal fail: ", err)
//		}
//	} else {
//		glog.Infof("not found command %d %s", msg.Cmd, pb_protocol.PbCmd(msg.Cmd).String())
//	}
//	return nil
//}