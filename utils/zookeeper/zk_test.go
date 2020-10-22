package zookeeper

import (
	"github.com/go-zookeeper/zk"
	"gogs.sharkgulf.cn/sg/library/qlog"
	"testing"
)

var log = qlog.WithField("app", "zkTest")

func TestName(t *testing.T) {
	path, err := zkClient.Create("/test", []byte("123"), 1, zk.WorldACL(zk.PermAll))
	if err != nil {
		log.Errorf("Create:%s", err)
	}
	log.Debugf("path: %s", path)
	go func() {
		w, stat, event, err := zkClient.GetW("/test")
		if err != nil {
			log.Errorf("getW:%s", err)
			return
		}
		log.Debugf("data:%s,stat:%+v", w, stat)
		select {
		case val := <-event:
			log.Debugf("event:%+v", val)
		}
	}()
	go func() {
		for {
			v, _, event, err := zkClient.GetW("/test")
			if err != nil {
				log.Errorf("getW:%s", err)
				return
			}
			log.Debugf("value of path[%s]=[%s].", path, v)
			for {
				select {
				case val := <-event:
					{
						if val.Type == zk.EventNodeCreated {
							log.Debugf("has new node[%s] create", val.Path)
						} else if val.Type == zk.EventNodeDeleted {
							log.Debugf("has node[%s] delete", val.Path)
						} else if val.Type == zk.EventNodeDataChanged {
							data, _, err := zkClient.Get("/test")
							if err != nil {
								log.Errorf("get:%s", err)
							}
							log.Debugf("data:%s", data)
							continue
						} else if val.Type == zk.EventNodeChildrenChanged {
							log.Debugf("children node change%+v", val.Path)
						}
					}
				}
				log.Debug("end")
				break
			}
		}
	}()
	go func() {
		stat, err := zkClient.Set("/test", []byte("ok"), 1)
		if err != nil {
			log.Errorf("set:%s", err)
			return
		}
		log.Debugf("stat:%v", stat)
	}()
	select {}
}
