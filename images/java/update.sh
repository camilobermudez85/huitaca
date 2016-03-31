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

MAVEN_VERSION="3.3.9"
MAVEN_HOME="/usr/share/maven"
MAVEN_DOWNLOAD="http://apache.osuosl.org/maven/maven-3/${MAVEN_VERSION}/binaries/apache-maven-${MAVEN_VERSION}-bin.tar.gz"

GRADLE_VERSION="2.6"
GRADLE_HOME="/usr/share/gradle"
GRADLE_DOWNLOAD="https://services.gradle.org/distributions/gradle-${GRADLE_VERSION}-all.zip"

declare -A java_tags=(
	['openjdk-7']='java-1.7.0'
	['openjdk-8']='java-1.8.0'
)

JDWP_ADDRESS="9009"
JDWP_DEFAULT_SERVER="y"
JDWP_DEFAULT_SUSPEND="n"

build_images() {
    for version in "${!java_tags[@]}"; do
	docker build -t huitaca/java:${version} ${version}
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

		FROM openshift/base-centos7

		MAINTAINER Camilo Bermúdez <camilobermudez85@gmail.com>

		ENV JAVA_VERSION "${java_tags[${version}]}"

		ENV JOLOKIA_VERSION "${JOLOKIA_VERSION}" \\
		    JOLOKIA_DOWNLOAD "${JOLOKIA_DOWNLOAD}" \\
		    JOLOKIA_HOST "${JOLOKIA_DEFAULT_HOST}" \\
		    JOLOKIA_CONTEXT "${JOLOKIA_DEFAULT_CONTEXT}" \\
		    JOLOKIA_DISCOVERY "${JOLOKIA_DEFAULT_DISCOVERY}" \\
		    JOLOKIA_PORT "${JOLOKIA_PORT}"

		ENV J4LOG_VERSION "${J4LOG_VERSION}" \\
		    J4LOG_DOWNLOAD "${J4LOG_DOWNLOAD}"

		ENV MAVEN_VERSION "${MAVEN_VERSION}" \\
		    MAVEN_HOME "${MAVEN_HOME}" \\
		    MAVEN_DOWNLOAD "${MAVEN_DOWNLOAD}"

		ENV GRADLE_VERSION "${GRADLE_VERSION}" \\
		    GRADLE_HOME "${GRADLE_HOME}" \\
		    GRADLE_DOWNLOAD "${GRADLE_DOWNLOAD}"

		EXPOSE "${JOLOKIA_PORT}"
EOF

	cat >> "${version}/Dockerfile" <<-"EOF"

		RUN set -x \
	    && yum update -y && yum install -y "${JAVA_VERSION}-openjdk-devel" && yum clean all -y \
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
