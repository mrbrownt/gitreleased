#!/bin/sh

set -e

[ "${TRACE}" ] && set -x

export DOCKER_HOST="tcp://docker:2375"

setupGCP() {
    echo "${GCP_JSON}" | base64 -d >/gcp.json
    gcloud auth activate-service-account --key-file /gcp.json
    gcloud auth configure-docker
}

setupGitlabDocker() {
    if [ -n "${GITLAB_CI}" ]; then
        echo "${CI_REGISTRY_PASSWORD}" |
            docker login \
                --password-stdin \
                -u "${CI_REGISTRY_USER}" \
                "${CI_REGISTRY}"
    fi
}

testApp() {
    case ${1} in
    auth)
        cd auth
        go mod vendor
        go test ./...
        ;;
    backend)
        cd backend
        go mod vendor
        go test ./...
        ;;
    frontend)
        cd frontend
        yarn install --pure-lockfile --cache-folder .yarn-cache
        ;;
    *)
        echo "unknown test suite"
        exit 1
        ;;
    esac
}

build() {
    setupGitlabDocker

    case ${1} in
    auth)
        cd auth
        AUTH_TAG="${CI_REGISTRY}/${CI_PROJECT_PATH}/auth:${CI_COMMIT_SHA}"
        docker build . -t "${AUTH_TAG}"
        docker push "${AUTH_TAG}"
        ;;
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

deploy() {

    case ${1} in
    auth)
        setupGitlabDocker
        setupGCP

        AUTH_TAG="${CI_REGISTRY}/${CI_PROJECT_PATH}/auth:${CI_COMMIT_SHA}"
        docker pull "${AUTH_TAG}"
        docker tag "${AUTH_TAG}" "us.gcr.io/spheric-subject-165900/gitreleased/auth:${CI_COMMIT_REF_NAME}"
        docker push "us.gcr.io/spheric-subject-165900/gitreleased/auth:${CI_COMMIT_REF_NAME}"
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
deploy)
    deploy "${2}"
    ;;
*)
    echo "pipeline task was not specified"
    exit 1
    ;;
esac
