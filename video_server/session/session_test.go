package session

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
	"video_server/weberrors"
)

func TestGetSession(t *testing.T) {
	var group sync.WaitGroup
	for i := 0; i < 20; i++ {
		group.Add(1)
		go testSync(i, &group)
	}
	group.Wait()
}

func testSync(n int, group *sync.WaitGroup)  {
	defer group.Done()
	session, err := NewSession(time.Second * 10)
	if err != nil {
		fmt.Println(err)
	}
	m := make(map[string]string)
	m["name"] = "yx" + strconv.FormatInt(int64(n), 10)
	m["id"] = strconv.FormatInt(int64(n), 10)
	session.PutData(m)
	if session.Get("name") != m["name"] {
		fmt.Println("get name error")
	}
	time.Sleep(time.Second * 20)
	_, err = GetSession(session.GetId())
	if weberrors.Type(err) != weberrors.NOT_FOUND {
		fmt.Println("GC is not sucess")
	}
}

