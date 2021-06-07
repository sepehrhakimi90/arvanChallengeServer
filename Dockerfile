FROM golang
WORKDIR /app
COPY ./go.mod /app
RUN go mod download
COPY ./ /app
EXPOSE 8080
RUN go build -o /server
ENV MYSQL_PORT=3306
ENV REDIS_PORT=6379
ENV REDIS_CHANNEL="ruleChannel"
WORKDIR /
RUN mv /app/templates /templates
RUN rm -rf /app
CMD ["/server"]