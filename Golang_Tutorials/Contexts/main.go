package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("Hey there! i am going to learn today context")
	start := time.Now()
	ctx := context.Background()
	response, err := ExternalCallWithTimeout(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result - ", response)
	fmt.Println("total time took - ", time.Since(start))
}

type ValueStruct struct {
	val string
	err error
}

func ExternalCallWithTimeout(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*1500)
	defer cancel()
	resultch := make(chan ValueStruct)
	go func() {
		val, err := thirdPartyCall()
		result := ValueStruct{
			val: val,
			err: err,
		}
		resultch <- result
	}()
	for {
		select {
		case <-ctx.Done():
			return "", fmt.Errorf("error - ", ctx.Err())
		case res := <-resultch:
			return res.val, res.err
		}
	}
}

func thirdPartyCall() (string, error) {
	time.Sleep(time.Millisecond * 750)
	return "dummy response", nil
}
