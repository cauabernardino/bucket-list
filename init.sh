#!/bin/sh

set -e

#!/bin/sh

if [ "$DATABASE" = "postgres" ]
then
    echo "Waiting for postgres..."

    while ! nc -z $DB_HOST $DB_PORT; do
      sleep 0.1
    done

    echo "PostgreSQL started"
fi

# Comment below lines if you don't want to flush the db every time the container is started
# python3 manage.py flush --no-input
# python3 manage.py migrate


echo "run db migration"
/app/migrate -path /app/migration \
    -database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${DB_HOST}:${DB_PORT}/${POSTGRES_DB}?sslmode=disable" \
     -verbose up

echo "start the app"
exec "$@"