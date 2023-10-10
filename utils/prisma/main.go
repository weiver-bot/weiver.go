package prisma

import (
	"context"

	"github.com/y2hO0ol23/weiver/db"
)

var (
	client *db.PrismaClient
	ctx    context.Context
)

func init() {
	client = db.NewClient()
	ctx = context.Background()
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}
}
