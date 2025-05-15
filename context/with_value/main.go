package main

import "context"

func main() {
	ctx := context.Background()
	// ctx = context.WithValue(ctx, "UserID", 123456)
	ctx = context.WithValue(ctx, "UserID", nil)
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	userID := ctx.Value("UserID")
	if userID == nil {
		panic("UserID not found in context")
	}
	println(userID.(int))
}
