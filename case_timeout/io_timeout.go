package case_timeout

import (
	"context"
	"errors"
	"log"
	"time"
)

var (
	ErrTIMEOUT = errors.New("timeout")
)

type IWriter interface {
	Process() error
}

// 确保已实现改接口
var _ IWriter = timeout1{}
var _ IWriter = timeout2{}

/*

 */

type timeout1 struct {
}

func (t timeout1) Process() error {
	log.Println("timeout1 ...")
	return nil
}

type timeout2 struct {
}

func (t timeout2) Process() error {
	log.Println("timeout2 ...")
	time.Sleep(time.Second * 3)
	return nil
}

func Q1(writer IWriter) error {
	doneCh := make(chan error)
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*2)
	defer cancel()

	go func() {
		var err error
		defer func() { doneCh <- err }()
		err = writer.Process()
	}()
	select {
	case <-ctx.Done():
		return ErrTIMEOUT
	case err := <-doneCh:
		return err
	}
}
