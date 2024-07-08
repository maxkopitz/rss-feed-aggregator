#! /bin/sh

if [ -f .env ]; then
    source .env
fi

cd sql/schema
goose $DATABASE_URL $DATABASE_URL up


