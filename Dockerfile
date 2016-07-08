FROM alpine
COPY envexp /usr/bin/envexp
ENTRYPOINT ["/usr/bin/envexp"]
