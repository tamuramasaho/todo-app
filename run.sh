#!/bin/sh

# wait cloudsql-proxy ready
sleep 5

if [ "$DB_AUTO_MIGRATION" != '' ]; then
    ./sql-migrate up
fi

./sirius

