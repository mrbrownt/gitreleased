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
    BACKEND_TAG="${CI_REGISTRY}/backend:${CI_COMMIT_REF_SLUG}"

    docker build . -t "${BACKEND_TAG}"
    docker push "${BACKEND_TAG}"
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
