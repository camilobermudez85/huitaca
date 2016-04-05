#!/bin/bash
#
# Utility to generate docker files and build java images.
#
# To generate docker files run './update.sh gen'
# To build images run './update.sh build'
#

set -eo pipefail

HOTSWAP_ZIP="HotswapAgent-0.3.zip"
HOTSWAP_BASE="https://github.com/HotswapProjects/HotswapAgent/releases/download/RELEASE-0.3"
huitaca_java_dev_versions=('openjdk-7' 'openjdk-8')

declare -A dcevm_download=(
	['openjdk-7']='https://github.com/dcevm/dcevm/releases/download/full-jdk7u79%2B7/DCEVM-full-7u79-installer.jar'
	['openjdk-8']='https://github.com/dcevm/dcevm/releases/download/light-jdk8u66%2B5/DCEVM-light-8u66-installer.jar'
)

JDWP_ADDRESS="9009"
JDWP_DEFAULT_SERVER="y"
JDWP_DEFAULT_SUSPEND="n"

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

ENV DCEVM_DOWNLOAD="${dcevm_download[${version}]}" \\
    HOTSWAP_ZIP="${HOTSWAP_ZIP}" \\
    HOTSWAP_BASE="${HOTSWAP_BASE}" \\
    JDWP_SERVER="${JDWP_DEFAULT_SERVER}" \\
    JDWP_SUSPEND="${JDWP_DEFAULT_SUSPEND}" \\
    JDWP_ADDRESS="${JDWP_ADDRESS}"

LABEL io.k8s.description="Image for building Java development environments" \\
      io.k8s.display-name="Java Development Builder" \\
      io.openshift.tags="builder,java,development,${version}"

EXPOSE "${JDWP_ADDRESS}"
EOF

	cat >> "${version}/Dockerfile" <<-"EOF"

RUN set -x \
    && curl -OL "${DCEVM_DOWNLOAD}" \
    && whoami \
    && unzip DCEVM-*-installer.jar "linux_amd64_compiler2/product/libjvm.so" \
    && mv "linux_amd64_compiler2/product/libjvm.so" "${JAVA_HOME}/jre/lib/amd64/server/" \
    && rm -rf linux_amd64_compiler2 DCEVM* \
    && curl -OL "${HOTSWAP_BASE}/${HOTSWAP_ZIP}" \
    && unzip "${HOTSWAP_ZIP}" \
    && mv hotswap-agent.jar "${JAVA_HOME}/bin" \
    && rm "${HOTSWAP_ZIP}"

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
