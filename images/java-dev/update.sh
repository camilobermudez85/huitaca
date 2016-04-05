#!/bin/bash
#
# Utility to generate docker files and build java images.
#
# To generate docker files run './update.sh gen'
# To build images run './update.sh build'
#

set -eo pipefail

huitaca_java_dev_versions=('openjdk-7' 'openjdk-8')

build_images() {
    for version in "${huitaca_java_dev_versions[@]}"; do
        docker build -t huitaca/java-dev:${version} ${version}
    done
}

generate_dockerfiles() {
    for version in "${huitaca_java_dev_versions[@]}"; do
        mkdir -p "${version}"
	
	cat > "${version}/Dockerfile" <<-EOF
#
# This Dockerfile is generated via "images/jre-dev/update.sh"
#
# Please DO NOT edit it directly.
#

FROM huitaca/java:${version}

MAINTAINER Camilo Bermúdez <camilobermudez85@gmail.com>

LABEL io.k8s.description="Image to support java development environments" \\
      io.k8s.display-name="Java Development Builder" \\
      io.openshift.tags="builder,java,development,${version}"
EOF

	cat >> "${version}/Dockerfile" <<-"EOF"

EXPOSE ${JDWP_ADDRESS}

COPY ./.s2i/bin/* ${STI_SCRIPTS_PATH}/
EOF
        cp -r template/.s2i "./${version}" 
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
