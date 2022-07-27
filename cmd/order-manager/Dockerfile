FROM uhub.service.ucloud.cn/entropypool/alpine:3.16

WORKDIR /app

COPY order-manager /app

ENTRYPOINT ["/app/order-manager"]

EXPOSE 50360 50361 50362 50363

CMD ["run"]
