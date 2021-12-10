#!/bin/sh
# wait-for-postgres.sh

set -e
  
uri="$1"
shift
cmd="$@"
  
until PGPASSWORD=$POSTGRES_PASSWORD psql "$uri" -c '\q'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done
  
>&2 echo "Postgres is up - executing command"
exec $cmd