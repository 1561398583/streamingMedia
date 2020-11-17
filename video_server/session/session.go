package session

import (
	"sync"
	"video_server/weberrors"
	"time"
	"video_server/utils"
)

func init() {
	go GC()
}

//session容器
var (
	Slock sync.RWMutex
	Sessions = make(map[string]*SessionInfo)
)

type SessionInfo struct {
	id string
	properties map[string]string
	expirationTime time.Time
}

type Session struct {
	id string
	properties map[string]string
}


func (session *Session) Get(name string) string {
	return session.properties[name]
}

func (session *Session) GetId() string {
	return session.id
}

func (session *Session) PutData(data map[string]string) error{
	//存到session中
	for k, v := range data {
		session.properties[k] = v
	}
	//在存一份到session容器中
	Slock.Lock()
	defer Slock.Unlock()
	sessionInfo, ok := Sessions[session.id]
	if !ok {
		return weberrors.New("session is not exist" ,2)
	}
	for k, v := range data {
		sessionInfo.properties[k] = v
	}
	return nil
}

func GetSession(id string) (*Session, error) {
	Slock.RLock()
	defer Slock.RUnlock()
	sessionInfo, ok := Sessions[id]
	if !ok {
		return nil, weberrors.New("this session is not exist", weberrors.NOT_FOUND)
	}
	session := Session{id: id}
	//为了保证数据并发安全，这里把sessionInfo.properties的内容复制一份返回
	for k, v := range sessionInfo.properties {
		session.properties[k] = v
	}
	return &session, nil
}




func NewSession(dur time.Duration)  (*Session,error){
	id, err := utils.GetUUID()
	if err != nil {
		return nil, weberrors.New(err.Error(), weberrors.UNKNOW_ERROR)
	}
	expire := time.Now().Add(dur)
	sessionInfo := SessionInfo{id: id, expirationTime: expire, properties: make(map[string]string)}
	Slock.Lock()
	Sessions[id] = &sessionInfo
	Slock.Unlock()
	return &Session{id: id, properties: make(map[string]string)}, nil
}

func GC()  {
	for  {
		time.Sleep(time.Second * 1)
		now := time.Now()
		Slock.Lock()
		for k, v := range Sessions {
			if now.After(v.expirationTime){
				delete(Sessions, k)
			}
		}
		Slock.Unlock()
	}
}
