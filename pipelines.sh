#!/bin/sh

set -e

[ "${TRACE}" ] && set -x

export DOCKER_HOST="tcp://docker:2375"

GCR_IMAGE_BASE=us.gcr.io/spheric-subject-165900/gitreleased
GCR_FRONTEND_IMG="$GCR_IMAGE_BASE/frontend:${CI_COMMIT_SHA}"
GCR_BACKEND_IMG="${GCR_IMAGE_BASE}/backend:${CI_COMMIT_SHA}"

setupGCP() {
    echo "${GCP_JSON}" | base64 -d >/gcp.json
    gcloud auth activate-service-account --key-file /gcp.json
    gcloud auth configure-docker
    gcloud config set account gitreleased-cloud-run@spheric-subject-165900.iam.gserviceaccount.com
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
    backend)
        export GOPROXY=https://proxy.golang.org
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
        setupGCP
        gcloud builds submit --tag="${GCR_BACKEND_IMG}" .
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
    frontend)
        setupGCP

        gcloud beta run deploy gitreleased-frontend \
            --project spheric-subject-165900 \
            --region us-central1 \
            --image "${GCR_FRONTEND_IMG}"
        ;;
    backend)
        setupGCP

        gcloud beta run deploy gitreleased-backend \
            --project spheric-subject-165900 \
            --region us-central1 \
            --image "${GCR_BACKEND_IMG}" \
            --set-env-vars "GITHUB_KEY=${GITHUB_KEY}" \
            --set-env-vars "GITHUB_SECRET=${GITHUB_SECRET}" \
            --set-env-vars "GITLAB_USER=mrbrownt" \
            --set-env-vars "GITLAB_ACCESS_TOKEN=${GITLAB_ACCESS_TOKEN}" \
            --set-env-vars "ENVIRONMENT=production" \
            --set-env-vars "SESSION_SECRET=${SESSION_SECRET}" \
            --set-env-vars "CLOUDSQL=yes" \
            --set-env-vars "DB_HOST=spheric-subject-165900:us-central1:gitreleased" \
            --set-env-vars "DB_PASS=${DB_PASS}" \
            --set-env-vars "BASE_URL=api.gitreleased.app"
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
