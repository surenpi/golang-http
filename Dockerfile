FROM alpine:3.8

COPY bin/sample /usr/bin/sample

EXPOSE 9898

CMD ["sample"]
