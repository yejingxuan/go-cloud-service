FROM alpine:latest

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk --no-cache add ca-certificates tzdata libc6-compat && cp -r -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN mkdir -p /mnt/glusterfs/rscb
COPY build/task /usr/bin/task
ENTRYPOINT ["/usr/bin/task"]