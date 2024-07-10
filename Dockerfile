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

RUN echo "xpack.security.enabled: true" >> /usr/share/elasticsearch/config/elasticsearch.yml && echo "xpack.security.authc.api_key.enabled: true" >> /usr/share/elasticsearch/config/elasticsearch.yml && echo "xpack.security.authc.realms.native.native1.order: 0" >> /usr/share/elasticsearch/config/elasticsearch.yml && echo "xpack.security.authc.realms.native.native1.enabled: true" >> /usr/share/elasticsearch/config/elasticsearch.yml && echo "xpack.security.authc.realms.native.native1.type: native" >> /usr/share/elasticsearch/config/elasticsearch.yml && echo "xpack.security.authc.realms.native.native1.bind_dn: cn=aa" >> /usr/share/elasticsearch/config/elasticsearch.yml && echo "xpack.security.authc.realms.native.native1.bind_password: flyisgreat" >> /usr/share/elasticsearch/config/elasticsearch.yml && echo "xpack.security.authc.realms.native.native1.usernames: elastic" >> /usr/share/elasticsearch/config/elasticsearch.yml

FROM docker.elastic.co/kibana/kibana:7.17.3 AS kibana

FROM debian:bookworm
RUN apt-get update
RUN apt-get install -y vim jq wget curl openjdk-17-jre-headless
RUN rm -rf /var/lib/apt/lists/*
ENV discovery.type=single-node
ENV ELASTICSEARCH_HOSTS=http://localhost:9200
ENV ES_JAVA_OPTS="-Xms512m -Xmx512m"
ENV path.data="/Users/aa/elasticsearch"
EXPOSE 8080 9200 5601
VOLUME ["/usr/share/elasticsearch/data"]

COPY --from=builder /usr/src/app/run-app /usr/local/bin/
COPY --from=elasticsearch /usr/share/elasticsearch /usr/share/elasticsearch
COPY --from=kibana /usr/share/kibana /usr/share/kibana

RUN mkdir -p /Users/aa/elasticsearch

CMD ["sh", "-c", "/usr/local/bin/run-app run 8080 & /usr/share/elasticsearch/bin/elasticsearch & /usr/share/kibana/bin/kibana"]
