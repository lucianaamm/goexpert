package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*3) // espera 3 segundos para o processo de reserva, se passar disso, o contexto é cancelado
	//ctx, cancel := context.WithCancel(ctx)
	defer cancel() //boa prática para evitar vazamento de goroutines
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Hotel booking cancelled. Timeout reached.")
		return
		//case <-time.After(5 * time.Second):
	case <-time.After(1 * time.Second): //espera 1 segundo para simular o processo de reserva
		fmt.Println("Hotel booked.")
	}
}
