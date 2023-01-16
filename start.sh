if test -f ".env"; then
  . ./.env
fi

docker-compose -p $DOCKER_APP_NAME up -d --build