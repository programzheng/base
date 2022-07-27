if [ -z "$1" ]; then
    goose -h
    exit
fi

godotenv -f .env goose -dir ./migrations $1