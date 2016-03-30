#
# This Dockerfile is generated via "images/jre/update.sh"
#
# Please DO NOT edit it directly.
#

FROM java:openjdk-8-jdk

MAINTAINER Camilo Bermúdez <camilobermudez85@gmail.com>

ENV JOLOKIA_VERSION "1.3.2"
ENV JOLOKIA_DOWNLOAD "http://search.maven.org/remotecontent?filepath=org/jolokia/jolokia-jvm/1.3.2/jolokia-jvm-1.3.2-agent.jar"
ENV JOLOKIA_HOST "0.0.0.0"
ENV JOLOKIA_CONTEXT "jmx"
ENV JOLOKIA_DISCOVERY "false"
ENV JOLOKIA_PORT "7777"

ENV J4LOG_VERSION "v0.1.0-alpha"
ENV J4LOG_DOWNLOAD "https://github.com/camilobermudez85/j4log/releases/download/v0.1.0-alpha/j4log-agent-v0.1.0-alpha.jar"

ENV MAVEN_VERSION "3.3.9"
ENV MAVEN_HOME "/usr/share/maven"
ENV MAVEN_DOWNLOAD "http://apache.osuosl.org/maven/maven-3/3.3.9/binaries/apache-maven-3.3.9-bin.tar.gz"

ENV GRADLE_VERSION "2.6"
ENV GRADLE_HOME "/usr/share/gradle"
ENV GRADLE_DOWNLOAD "https://services.gradle.org/distributions/gradle-2.6-all.zip"

EXPOSE "7777"

RUN set -x \
    && mkdir -p "${MAVEN_HOME}" \
    && curl -fsSL "${MAVEN_DOWNLOAD}" \
       | tar -xzC "${MAVEN_HOME}" --strip-components=1 \
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
    && cd "${JAVA_HOME}/bin" \
        && mv java java-exec \
        && echo "\${JAVA_HOME}/bin/java-exec \"\$@\" -javaagent:\${JAVA_HOME}/bin/jolokia-jvm-${JOLOKIA_VERSION}-agent.jar=host=\${JOLOKIA_HOST},port=\${JOLOKIA_PORT},agentContext=\${JOLOKIA_CONTEXT},discoveryEnabled=\${JOLOKIA_DISCOVERY} -javaagent:\${JAVA_HOME}/bin/j4log-agent-${J4LOG_VERSION}.jar" > java \
        && chmod 755 java \
