#!/bin/sh

set -e

[ "${TRACE}" ] && set -x

if [ -n "${CI}" ]; then
    if [ -z "${DOCKER_HOST}" ] && [ "${KUBERNETES_PORT}" ]; then
        export DOCKER_HOST='tcp://localhost:2375'
    fi
fi

testApp() {
    case ${1} in
        backend)
            cd backend
            go test ./...
            ;;
        frontend)
            cd frontend
            ;;
        *)
            echo "unknown test suite"
            exit 1
            ;;
    esac
}

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
            BACKEND_TAG="${CI_REGISTRY}/${CI_PROJECT_PATH}/backend:${CI_COMMIT_SHA}"
            docker build . -t "${BACKEND_TAG}"
            docker push "${BACKEND_TAG}"
            ;;
        frontend)
            cd frontend
            FRONTEND_TAG="${CI_REGISTRY}/${CI_PROJECT_PATH}/frontend:${CI_COMMIT_SHA}"
            docker build . -t "${FRONTEND_TAG}"
            docker push "${FRONTEND_TAG}"
            ;;
        *)
            echo "unknown build product"
            exit 1
            ;;
    esac
}

case ${1} in
    build)
        build "${2}"
        ;;
    test)
        testApp "${2}"
        ;;
    *)
        echo "pipeline task was not specified"
        exit 1
        ;;
esac
