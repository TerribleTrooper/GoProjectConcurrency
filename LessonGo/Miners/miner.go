package Miners

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Miner(
	ctx context.Context,
	wg *sync.WaitGroup,
	transfer_point chan<- int,
	number_of_miners int,
	power int,
) {

	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Шахтёр закончил", number_of_miners)
			return
		default:
			fmt.Println("Miner ", number_of_miners)
			time.Sleep(1 * time.Second)
			fmt.Println("I collect ", power, "coal")
			transfer_point <- power
		}
	}
}

func Miner_Pool(ctx context.Context, miner_Count int) chan int {
	coal_Transfer_Point := make(chan int)
	wg := &sync.WaitGroup{}
	for i := 1; i <= miner_Count; i++ {
		wg.Add(1)
		go Miner(ctx, wg, coal_Transfer_Point, i, i*10)
	}
	go func() {
		wg.Wait()
		close(coal_Transfer_Point)
	}()
	return coal_Transfer_Point
}
