package database

import "context"

func Boot(ctx context.Context) context.Context {
	db, err := Database()
	if err != nil {
		panic(err)
	}

	return context.WithValue(ctx, "db", db)
}
