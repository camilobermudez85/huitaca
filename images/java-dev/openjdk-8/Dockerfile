#
# This Dockerfile is generated via "images/jre-dev/update.sh"
#
# Please DO NOT edit it directly.
#

FROM huitaca/java:openjdk-8

MAINTAINER Camilo Bermúdez <camilobermudez85@gmail.com>

LABEL io.k8s.description="Image to support java development environments" \
      io.k8s.display-name="Java Development Builder" \
      io.openshift.tags="builder,java,development,openjdk-8"

EXPOSE ${JDWP_ADDRESS}

COPY ./.s2i/bin/* ${STI_SCRIPTS_PATH}/
