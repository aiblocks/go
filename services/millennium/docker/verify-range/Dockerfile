FROM ubuntu:18.04

MAINTAINER Bartek Nowotarski <bartek@aiblocks.io>

ENV AIBLOCKS_CORE_VERSION 12.2.0-1088-3f8a5356
# to remove tzdata interactive flow
ENV DEBIAN_FRONTEND=noninteractive

ADD dependencies /
RUN ["chmod", "+x", "dependencies"]
RUN /dependencies

ADD aiblocks-core.cfg /

ADD start /
RUN ["chmod", "+x", "start"]

ENTRYPOINT ["/start"]