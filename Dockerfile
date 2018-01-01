FROM alpine:3.7

RUN mkdir -p /source
WORKDIR /source
COPY ./bin/namer ./
COPY names.yml ./

EXPOSE 8080

ENTRYPOINT ["./namer"]
CMD ["help"]
