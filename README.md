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

very early development, nothing usable yet.

## license

mit
