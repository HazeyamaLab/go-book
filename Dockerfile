FROM golang:1.12

LABEL maintainer = "Yutsuki Miyashita <j148015n@st.u-gakugei.ac.jp>"
LABEL description = "ginAPI"

RUN apt-get update -qq && \
    apt-get install -y mysql-client

WORKDIR /
ENV GOPATH /go
ENV APIDIR ${GOPATH}/src/github.com/mytheta/gin_api

COPY mysql/wait-for-mysql.sh /wait-for-mysql.sh
COPY . ${APIDIR}
RUN cd ${APIDIR} && GO111MODULE=on go build -o bin/gin_api ./main.go
RUN cd ${APIDIR} && cp bin/gin_api /usr/local/bin/

