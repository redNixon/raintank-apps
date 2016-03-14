package api

import (
	"github.com/raintank/raintank-apps/task-server/model"
	"github.com/raintank/raintank-apps/task-server/sqlstore"
)

func GetMetrics(ctx *Context, query model.GetMetricsQuery) {
	query.Owner = ctx.Owner
	metrics, err := sqlstore.GetMetrics(&query)
	if err != nil {
		log.Error(err)
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(200, metrics)
}