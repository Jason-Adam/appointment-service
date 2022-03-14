package appointment

import (
	"time"
	_ "time/tzdata"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jason-adam/appointment-service/internal/util"
)

type Option func(*repository)

func WithConnectionPool(dbPool *pgxpool.Pool) Option {
	return Option(func(r *repository) {
		util.Ensure(dbPool != nil, "connection pool cannot be nil")

		r.dbPool = dbPool
	})
}

func WithLocation(tz string) Option {
	return Option(func(r *repository) {
		loc, err := time.LoadLocation(tz)
		if err != nil {
			panic(err)
		}

		r.loc = loc
	})
}
