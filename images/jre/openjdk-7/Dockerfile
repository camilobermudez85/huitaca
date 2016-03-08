#
# This Dockerfile is generated via "images/jre/update.sh"
#
# Please DO NOT edit it directly.
#

FROM java:openjdk-7-jre

ENV JOLOKIA_VERSION "1.3.2"
ENV JOLOKIA_DOWNLOAD "http://search.maven.org/remotecontent?filepath=org/jolokia/jolokia-jvm/1.3.2/jolokia-jvm-1.3.2-agent.jar"
ENV JOLOKIA_HOST "0.0.0.0"
ENV JOLOKIA_CONTEXT "jmx"
ENV JOLOKIA_DISCOVERY "false"
ENV JOLOKIA_PORT "7777"

ENV J4LOG_VERSION "v0.1.0-alpha"
ENV J4LOG_DOWNLOAD "https://github.com/camilobermudez85/j4log/releases/download/v0.1.0-alpha/j4log-agent-v0.1.0-alpha.jar"

EXPOSE "7777"

RUN set -x \
    && curl -OL "${JOLOKIA_DOWNLOAD}" \
    && mv "jolokia-jvm-${JOLOKIA_VERSION}-agent.jar" "${JAVA_HOME}/bin" \
    && rm -f jolokia-jvm-${JOLOKIA_VERSION}-agent.jar \
    && curl -OL "${J4LOG_DOWNLOAD}" \
    && mv "j4log-agent-${J4LOG_VERSION}.jar" "${JAVA_HOME}/bin" \
    && rm -f j4log-agent-${J4LOG_VERSION}.jar \
    && cd "${JAVA_HOME}/bin" \
        && mv java java-exec \
        && echo "\${JAVA_HOME}/bin/java-exec \"\$@\" -javaagent:\${JAVA_HOME}/bin/jolokia-jvm-${JOLOKIA_VERSION}-agent.jar=host=\${JOLOKIA_HOST},port=\${JOLOKIA_PORT},agentContext=\${JOLOKIA_CONTEXT},discoveryEnabled=\${JOLOKIA_DISCOVERY} -javaagent:\${JAVA_HOME}/bin/j4log-agent-${J4LOG_VERSION}.jar" > java \
        && chmod 755 java \
