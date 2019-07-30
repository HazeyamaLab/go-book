FROM golang:1.12

LABEL maintainer = "Yutsuki Miyashita <j148015n@st.u-gakugei.ac.jp>"
LABEL description = "gi-simple-application"

RUN apt-get update -qq && \
    apt-get install -y default-mysql-client

WORKDIR /
ENV GOPATH /go
ENV APIDIR ${GOPATH}/src/github.com/HazeyamaLab/go-book

COPY mysql/wait-for-mysql.sh /wait-for-mysql.sh
COPY . ${APIDIR}
RUN cd ${APIDIR} && GO111MODULE=on go build -o bin/go-book ./main.go
RUN cd ${APIDIR} && cp bin/go-book /usr/local/bin/
RUN cd ${APIDIR} && cp -r ./ /

