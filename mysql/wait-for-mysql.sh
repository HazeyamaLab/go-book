#!/bin/sh

set -e

host="$1"
shift
cmd="$2"

echo "Waiting for mysql"
until mysqladmin ping -h $host --silent; do
  echo 'waiting for mysqld to be connectable...'
  sleep 2
done

echo "app is starting...!"

exec go-book