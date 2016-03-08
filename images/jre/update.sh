#!/bin/bash
#
# Utility to generate docker files and build java images.
#
# To generate docker files run './update.sh generate'
# To build images run './update.sh build'
#

set -eo pipefail

JOLOKIA_VERSION="1.3.2"
JOLOKIA_DOWNLOAD="http://search.maven.org/remotecontent?filepath=org/jolokia/jolokia-jvm/${JOLOKIA_VERSION}/jolokia-jvm-${JOLOKIA_VERSION}-agent.jar"
JOLOKIA_PORT="7777"
JOLOKIA_DEFAULT_HOST="0.0.0.0"
JOLOKIA_DEFAULT_CONTEXT="jmx"
JOLOKIA_DEFAULT_DISCOVERY="false"

J4LOG_VERSION="v0.1.0-alpha"
J4LOG_DOWNLOAD="https://github.com/camilobermudez85/j4log/releases/download/${J4LOG_VERSION}/j4log-agent-${J4LOG_VERSION}.jar"

declare -A java_tags=(
	['openjdk-7']='openjdk-7-jre'
	['openjdk-8']='openjdk-8-jre'
)

JDWP_ADDRESS="9009"
JDWP_DEFAULT_SERVER="y"
JDWP_DEFAULT_SUSPEND="n"

build_images() {
    for version in "${!java_tags[@]}"; do
	docker build -t huitaca/jre:${version} ${version}
    done
}

generate_dockerfiles() {
    for version in "${!java_tags[@]}"; do
	mkdir -p "${version}"
	
	cat > "${version}/Dockerfile" <<-EOF
		#
		# This Dockerfile is generated via "images/jre/update.sh"
		#
		# Please DO NOT edit it directly.
		#

		FROM java:${java_tags[${version}]}

		ENV JOLOKIA_VERSION "${JOLOKIA_VERSION}"
		ENV JOLOKIA_DOWNLOAD "${JOLOKIA_DOWNLOAD}"
		ENV JOLOKIA_HOST "${JOLOKIA_DEFAULT_HOST}"
		ENV JOLOKIA_CONTEXT "${JOLOKIA_DEFAULT_CONTEXT}"
		ENV JOLOKIA_DISCOVERY "${JOLOKIA_DEFAULT_DISCOVERY}"
		ENV JOLOKIA_PORT "${JOLOKIA_PORT}"

		ENV J4LOG_VERSION "${J4LOG_VERSION}"
		ENV J4LOG_DOWNLOAD "${J4LOG_DOWNLOAD}"

		EXPOSE "${JOLOKIA_PORT}"
EOF

	cat >> "${version}/Dockerfile" <<-"EOF"

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
EOF
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
