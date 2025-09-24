package Postmans

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Postman_Func(ctx context.Context, wg *sync.WaitGroup, transfer_Point chan<- string, n int, mail string) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Почтальон закончил")
			return
		default:
			fmt.Println("Я почтальон номер", n, "взял письмо")
			time.Sleep(1 * time.Second)
			fmt.Println("Я донёс писмо", mail, "до почты")

			transfer_Point <- mail
		}
	}
}

func Postman_Pool(ctx context.Context, postman_Count int) <-chan string {
	wg := &sync.WaitGroup{}
	mail_Transfer_Point := make(chan string)
	for i := 1; i <= postman_Count; i++ {
		wg.Add(1)
		go Postman_Func(ctx, wg, mail_Transfer_Point, i, postman_To_Mail(i))
	}
	go func() {
		wg.Wait()
		close(mail_Transfer_Point)
	}()

	return mail_Transfer_Point
}

func postman_To_Mail(postman_Number int) string {
	ptm := map[int]string{
		1: "Газета",
		2: "Информация",
		3: "Писмо",
	}
	mail, ok := ptm[postman_Number]
	if !ok {
		return "Лотерея"
	}
	return mail
}
