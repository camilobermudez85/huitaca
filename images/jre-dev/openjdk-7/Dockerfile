#
# This Dockerfile is generated via "images/jre-dev/update.sh"
#
# Please DO NOT edit it directly.
#

FROM huitaca/jre:openjdk-7

ENV DCEVM_DOWNLOAD "https://github.com/dcevm/dcevm/releases/download/full-jdk7u79%2B7/DCEVM-full-7u79-installer.jar"
ENV HOTSWAP_ZIP "HotswapAgent-0.3.zip"
ENV HOTSWAP_BASE "https://github.com/HotswapProjects/HotswapAgent/releases/download/RELEASE-0.3"

# JDWP variables
ENV JDWP_SERVER "y"
ENV JDWP_SUSPEND "n"
ENV JDWP_ADDRESS "9009"

EXPOSE "9009"

RUN set -x \
    && curl -OL "${DCEVM_DOWNLOAD}" \
    && unzip DCEVM-*-installer.jar "linux_amd64_compiler2/product/libjvm.so" \
    && mv "linux_amd64_compiler2/product/libjvm.so" "${JAVA_HOME}/lib/amd64/server/" \
    && rm -rf linux_amd64_compiler2 DCEVM* \
    && curl -OL "${HOTSWAP_BASE}/${HOTSWAP_ZIP}" \
    && unzip "${HOTSWAP_ZIP}" \
    && mv hotswap-agent.jar "${JAVA_HOME}/bin" \
    && rm "${HOTSWAP_ZIP}" \
    && cd "${JAVA_HOME}/bin" \
        && EXTRA_OPTIONS="-javaagent:\${JAVA_HOME}/bin/hotswap-agent.jar -agentlib:jdwp=transport=dt_socket,server=\${JDWP_SERVER},suspend=\${JDWP_SUSPEND},address=${JDWP_ADDRESS}" \
        && echo "$(cat java) ${EXTRA_OPTIONS}" > java
