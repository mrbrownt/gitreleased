#!/bin/sh

set -eo pipefail

[ "${TRACE}" ] && set -x

if [ -n "${CI}" ]; then
    if [ -z "${DOCKER_HOST}" ] && [ "${KUBERNETES_PORT}" ]; then
        export DOCKER_HOST='tcp://localhost:2375'
    fi
fi

build() {
    if [ -n "${GITLAB_CI}" ]; then
        docker login \
            -p "${CI_REGISTRY_PASSWORD}" \
            -u "${CI_REGISTRY_USER}" \
            "${CI_REGISTRY}"
    fi

    case ${1} in
        backend)
            cd backend
            backendBuild
            ;;
        frontend)
            cd frontend
            ;;
        *)
            echo "unknown build product"
            exit 1
            ;;
    esac
}

backendBuild() {
    docker build . -t stuffs
}

case ${1} in
    build)
        build "${2}"
        ;;
    *)
        echo "pipeline task was not specified"
        exit 1
        ;;
esac
