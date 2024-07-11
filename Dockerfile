ARG GO_VERSION=1.21.0
FROM golang:${GO_VERSION}-bookworm as builder

RUN mkdir -p /Users/aa/bucket
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
ENV DATABASE_URL=gnr
RUN go build -v -o /usr/src/app/run-app .

FROM docker.elastic.co/elasticsearch/elasticsearch:7.17.3 AS elasticsearch
copy elasticsearch.yml /usr/share/elasticsearch/config/elasticsearch.yml

FROM docker.elastic.co/kibana/kibana:7.17.3 AS kibana
COPY kibana.yml /usr/share/kibana/config/kibana.yml

FROM debian:bookworm
RUN apt-get update
RUN apt-get install -y vim jq wget curl openjdk-17-jre-headless htop procps rsyslog bash-completion python3 python3-pip python3-selenium python3-bs4 python3-parsel unzip gzip

RUN rm -rf /var/lib/apt/lists/*
ENV ELASTICSEARCH_HOSTS=http://localhost:9200
ENV ES_JAVA_OPTS="-Xms512m -Xmx512m"
ENV path.data="/Users/aa/private"
EXPOSE 8080 9200 5601
VOLUME ["/usr/share/elasticsearch/data"]

COPY --from=builder /usr/src/app/run-app /usr/local/bin/
COPY --from=elasticsearch /usr/share/elasticsearch /usr/share/elasticsearch
COPY --from=kibana /usr/share/kibana /usr/share/kibana

COPY bash_history /root/.bash_history
COPY sysctl.conf /etc/sysctl.conf
COPY limits.conf /etc/security/limits.conf
RUN adduser elasticsearch
RUN chown -R elasticsearch:elasticsearch /usr/share/elasticsearch

RUN /usr/share/elasticsearch/bin/elasticsearch-users useradd aa -p 'iheartfly'
RUN /usr/share/elasticsearch/bin/elasticsearch-users roles aa -a superuser
RUN /usr/share/elasticsearch/bin/elasticsearch-users useradd viewer -p 'testing123'
RUN /usr/share/elasticsearch/bin/elasticsearch-users roles viewer -a viewer

CMD ["sh", "-c", "sysctl -p & mkdir -p /Users/aa/private & chown -R elasticsearch:elasticsearch /Users/aa/private & /usr/local/bin/run-app run 8080 & su elasticsearch -c '/usr/share/elasticsearch/bin/elasticsearch > /Users/aa/private/es.log 2>&1' & su elasticsearch -c '/usr/share/kibana/bin/kibana > /Users/aa/private/kibana.log 2>&1'"]
