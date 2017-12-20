FROM centos
MAINTAINER aoaolion <274912665@qq.com>

RUN rpm --rebuilddb && \
	yum install -y go
RUN yum install -y git
RUN mkdir -p /root/go

ENV GOPATH /root/go

RUN go get github.com/aoaolion/webserv
WORKDIR /root/go/src/github.com/aoaolion/webserv
RUN go build
CMD ./webserv

EXPOSE 80