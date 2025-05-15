package main

import (
	"context"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	// Simulando uma operação de reserva de hotel
	select {
	// case <-time.After(time.Second * 2):
	case <-time.After(time.Second * 5):
		println("Hotel reservado com sucesso!")
	case <-ctx.Done():
		println("Reserva cancelada:", ctx.Err().Error())
	}
}
