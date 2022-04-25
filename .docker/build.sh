#!/bin/sh
cd "$(dirname "$0")"

BASE="${PWD}/../"

if [ ! -z "${BUILD_IMAGE}" ]
then
  ENV_BUILD_IMAGE="${BUILD_IMAGE}"
fi

if [ -f ".env.default" ]
then
  export $(cat .env.default | sed 's/#.*//g' | xargs)
fi

if [ -f ".env" ]
then
  export $(cat .env | sed 's/#.*//g' | xargs)
fi

if [ ! -z "${ENV_BUILD_IMAGE}" ]
then
  BUILD_IMAGE="${ENV_BUILD_IMAGE}"
  unset ENV_BUILD_IMAGE
fi

if [ -z "${1}" ] && [ ! -z "${SERVER_TAG}" ]
then
  ./build base
  ./build ${SERVER_TAG}
  exit 0
fi


build_client() {
  docker build -t neko-dev-client -f base/Dockerfile --target client "${BASE}"
  docker run --rm \
    --user "$(id -u):$(id -g)" \
    -v "${BASE}client/dist:/tmp/dist" \
    neko-dev-client sh -c "rm -rf /tmp/dist/*; cp -r /src/dist/* /tmp/dist"
}

build_server() {
  docker build -t neko-dev-server -f base/Dockerfile --target server "${BASE}"
  docker run --rm \
    --user "$(id -u):$(id -g)" \
    -v "${BASE}server/bin:/tmp/bin" \
    neko-dev-server sh -c "rm -rf /tmp/bin/neko; cp /src/bin/neko /tmp/bin"
}

build() {
  if [ "$1" = "base" ]
  then
    docker build -t "${BUILD_IMAGE}:base" -f base/Dockerfile "${BASE}"
  else
    docker build -t "${BUILD_IMAGE}:$1" --build-arg="BASE_IMAGE=${BUILD_IMAGE}:base" -f "$1/Dockerfile" "$1/"
  fi
}

case $1 in
  client) build_client;;
  server) build_server;;

  arm-*) build_arm "${1#arm-}";;

  *) build "$1";;
esac
