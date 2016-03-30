#!/bin/bash
#
# Pulls the tomcat official images and tweaks them to make them use
# huitaca's java development images.
#

TOMCAT_DOCKER_REPO="https://github.com/docker-library/tomcat.git"
DOCKER_REPO="huitaca"
IMAGE_NAME="tomcat"

generate_dockerfiles() {
    find . -type d -regex ".*/[1-9]-java[1-9]" -prune -exec rm -rf {} \;
    git clone ${TOMCAT_DOCKER_REPO}
    cd tomcat

    for version in $(find . -type d -regex "./[1-9]-jre[1-9]" -printf "%f\n" -prune); do

	sed -i 's/FROM java:7-jre/FROM huitaca\/java-dev:openjdk-7/g' "${version}/Dockerfile"
	sed -i 's/FROM java:8-jre/FROM huitaca\/java-dev:openjdk-8/g' "${version}/Dockerfile"

	mv ${version} ../${version/jre/java}

    done

    cd ..
    rm -rf tomcat
}

build_images() {

    for tag in $(find . -type d -regex ".*/[1-9]-java[1-9]" -printf '%P\n' -prune); do
	docker build -t "${DOCKER_REPO}/${IMAGE_NAME}:${tag}" "${tag}"
    done
}

cmd=$1
shift
case $cmd in
    gen) generate_dockerfiles "$@";
	;;
    build) build_images "$@"
	;;
esac
