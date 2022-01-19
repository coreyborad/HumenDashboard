FROM golang:1.16.2-stretch

WORKDIR $GOPATH/src/stock

# Instal package
RUN set -x \
  && apt-get update \
  && apt-get upgrade -y \
  && apt-get install -y --no-install-recommends \
  git unzip

# Set locale
RUN apt-get update -y
RUN apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y locales

RUN sed -i -e 's/# en_US.UTF-8 UTF-8/en_US.UTF-8 UTF-8/' /etc/locale.gen && \
    dpkg-reconfigure --frontend=noninteractive locales && \
    update-locale LANG=en_US.UTF-8

ENV LANG en_US.UTF-8

RUN apt-get install -y fonts-wqy-microhei ttf-wqy-microhei fonts-wqy-zenhei ttf-wqy-zenhei

RUN fc-cache -f -v

# Set Timezone
ARG TZ=UTC
ENV TZ ${TZ}
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# User Aliases
COPY ./aliases.sh /root/aliases.sh
RUN sed -i 's/\r//g' /root/aliases.sh && \
  echo "" >> ~/.bashrc && \
  echo "# Load Custom Aliases" >> ~/.bashrc && \
  echo "source /root/aliases.sh" >> ~/.bashrc && \
  echo "" >> ~/.bashrc

# golint
RUN go get -u golang.org/x/lint/golint

# Wire
RUN go get -u github.com/google/wire/cmd/wire

# protoc-gen-go
RUN go get -u google.golang.org/protobuf/cmd/protoc-gen-go
RUN go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

# protoc
RUN curl -L https://github.com/protocolbuffers/protobuf/releases/download/v3.12.3/protoc-3.12.3-linux-x86_64.zip -o /tmp/protoc.zip && \
  unzip -o /tmp/protoc.zip -d /usr/local bin/protoc && \
  unzip -o /tmp/protoc.zip -d /usr/local include/*

# wkhtml
RUN apt-get update -y
RUN apt-get install -y wkhtmltopdf xvfb xauth xfonts-base xfonts-75dpi fontconfig

# Clean up
RUN apt-get clean \
  && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

