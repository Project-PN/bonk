# bonk

bonk postgres until it returns row

a postgres orm for go, inspired by prisma.

## about

i wanted something similar to prisma that works natively with go, without having to generate a client or maintain a separate schema.

bonk uses structs, reflection and pgx to handle queries without needing to write sql for everything.

## planned features

- model registration using structs and tags
- querying, filtering and pagination
- crud operations
- relations and eager loading
- transactions
- migrations
- raw sql when you actually need it

## status

very early development. connection pools work; orm features aren't implemented yet.

## connections

connection code lives in `src/postgres/`. import `github.com/Project-PN/bonk/src/postgres`
and call `postgres.Open(ctx, connString)` to open a `*pgxpool.Pool` and check connectivity.
the caller owns the pool and must call `pool.Close()` when finished.

for pool settings or connection hooks, use `pgxpool.ParseConfig`, adjust the config,
then pass it to `postgres.OpenConfig(ctx, config)`. don't change the config afterward.
use a context with a deadline to bound connection setup.

## tests

run `go test ./...`. postgres tests use `BONK_TEST_DATABASE_URL` and skip when it
isn't set. point it at a test database; the tests only open connections and run
read-only queries.

## license

mit
