package main

import (
	"context"
	"fmt"
)

func doBalue(ctx context.Context) {
	fmt.Println(ctx.Value("group"), ctx.Value("name"))
}

func main() {
	ctx := context.WithValue(context.Background(), "group", "admin")
	ctxuser := context.WithValue(ctx, "name", "Kirill")
	doBalue(ctxuser)
}
