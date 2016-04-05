#
# This Dockerfile is generated via "images/jre-dev/update.sh"
#
# Please DO NOT edit it directly.
#

FROM huitaca/java:openjdk-8

MAINTAINER Camilo Bermúdez <camilobermudez85@gmail.com>

ENV DCEVM_DOWNLOAD="https://github.com/dcevm/dcevm/releases/download/light-jdk8u66%2B5/DCEVM-light-8u66-installer.jar" \
    HOTSWAP_ZIP="HotswapAgent-0.3.zip" \
    HOTSWAP_BASE="https://github.com/HotswapProjects/HotswapAgent/releases/download/RELEASE-0.3" \
    JDWP_SERVER="y" \
    JDWP_SUSPEND="n" \
    JDWP_ADDRESS="9009"

LABEL io.k8s.description="Image for building Java development environments" \
      io.k8s.display-name="Java Development Builder" \
      io.openshift.tags="builder,java,development,openjdk-8"

EXPOSE "9009"

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
