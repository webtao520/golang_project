package golog

import (
	"log"
	"testing"
)

func TestLevel(t *testing.T) {
	logex := New("test")
	logex.colorFile = NewColorFile()
	logex.colorFile.Load("color_sample.json")
	logex.enableColor = true
	logex.Debugf("%d %s %v", 1, "hello", "world")
	logex.Errorln("hello1")
	logex.Errorln("2")
	logex.Infoln("no")

}

func TestMyLog(t *testing.T) {
	logex := New("test2")
	logex.colorFile = NewColorFile()
	logex.colorFile.Load("color_sample.json")
	logex.enableColor = true
	logex.Debugln("hello1")
	logex.DebugColorln("blue", "hello2")
	logex.Debugln("hello3")
}

func TestSystemLog(t *testing.T) {
	log.Println("hello1")
	log.Println("hello2")
	log.Println("hello3")
}
