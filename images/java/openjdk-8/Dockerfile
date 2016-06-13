#
# This Dockerfile is generated via "images/jre/update.sh"
#
# Please DO NOT edit it directly.
#

FROM openshift/base-centos7

MAINTAINER Camilo Bermúdez <camilobermudez85@gmail.com>

ENV JAVA_VERSION="java-1.8.0" \
    JAVA_HOME="/usr/lib/jvm/java-1.8.0-openjdk" \
    JOLOKIA_VERSION="1.3.2" \
    JOLOKIA_DOWNLOAD="http://search.maven.org/remotecontent?filepath=org/jolokia/jolokia-jvm/1.3.2/jolokia-jvm-1.3.2-agent.jar" \
    JOLOKIA_HOST="0.0.0.0" \
    JOLOKIA_CONTEXT="/" \
    JOLOKIA_DISCOVERY="false" \
    JOLOKIA_PORT="7777" \
    J4LOG_VERSION="v0.1.0-alpha" \
    J4LOG_DOWNLOAD="https://github.com/camilobermudez85/j4log/releases/download/v0.1.0-alpha/j4log-agent-v0.1.0-alpha.jar" \
    MAVEN_VERSION="3.3.3" \
    MAVEN_HOME="/usr/share/maven" \
    MAVEN_DOWNLOAD="http://archive.apache.org/dist/maven/maven-3/3.3.3/binaries/apache-maven-3.3.3-bin.tar.gz" \
    GRADLE_VERSION="2.6" \
    GRADLE_HOME="/usr/share/gradle" \
    GRADLE_DOWNLOAD="https://services.gradle.org/distributions/gradle-2.6-all.zip" \
    BUILDER_VERSION="1.0.0" \
    DCEVM_DOWNLOAD="https://github.com/dcevm/dcevm/releases/download/light-jdk8u66%2B5/DCEVM-light-8u66-installer.jar" \
    HOTSWAP_ZIP="HotswapAgent-0.3.zip" \
    HOTSWAP_BASE="https://github.com/HotswapProjects/HotswapAgent/releases/download/RELEASE-0.3" \
    JDWP_SERVER="y" \
    JDWP_SUSPEND="n" \
    JDWP_ADDRESS="9009" \
    BUILD_OUTPUT_DIR="/opt/app-root/build-output-dir"

LABEL io.k8s.description="Image for building Java applications" \
      io.k8s.display-name="Java Builder" \
      io.openshift.tags="builder,java,java-1.8.0"

EXPOSE "7777"

RUN set -x \
    && yum -y update && yum -y install "${JAVA_VERSION}-openjdk-devel" && yum -y clean all \
    && mkdir -p "${MAVEN_HOME}" \
    && curl -fsSL "${MAVEN_DOWNLOAD}" | tar -xzC "${MAVEN_HOME}" --strip-components=1 \
    && ln -s "${MAVEN_HOME}/bin/mvn" "/usr/bin/mvn" \
    && mkdir -p "${GRADLE_HOME}" && rm -rf "${GRADLE_HOME}" \
    && curl -sLO "${GRADLE_DOWNLOAD}" \
    && unzip "gradle-${GRADLE_VERSION}-all.zip" \
    && mv "gradle-${GRADLE_VERSION}" "$GRADLE_HOME" \
    && rm "gradle-${GRADLE_VERSION}-all.zip" \
    && ln -s "${GRADLE_HOME}/bin/gradle" "/usr/bin/gradle" \
    && curl -OL "${JOLOKIA_DOWNLOAD}" \
    && mv "jolokia-jvm-${JOLOKIA_VERSION}-agent.jar" "${JAVA_HOME}/bin" \
    && rm -f "jolokia-jvm-${JOLOKIA_VERSION}-agent.jar" \
    && curl -OL "${J4LOG_DOWNLOAD}" \
    && mv "j4log-agent-${J4LOG_VERSION}.jar" "${JAVA_HOME}/bin" \
    && rm -f "j4log-agent-${J4LOG_VERSION}.jar" \
    && curl -OL "${DCEVM_DOWNLOAD}" \
    && unzip DCEVM-*-installer.jar "linux_amd64_compiler2/product/libjvm.so" \
    && mv "linux_amd64_compiler2/product/libjvm.so" "${JAVA_HOME}/jre/lib/amd64/server/" \
    && rm -rf linux_amd64_compiler2 DCEVM* \
    && curl -OL "${HOTSWAP_BASE}/${HOTSWAP_ZIP}" \
    && unzip "${HOTSWAP_ZIP}" \
    && mv hotswap-agent.jar "${JAVA_HOME}/bin" \
    && rm "${HOTSWAP_ZIP}"  \
    && chown -R 1001:1001 /opt/app-root

COPY ./.s2i/bin/ ${STI_SCRIPTS_PATH}

USER 1001

CMD ["usage"]
