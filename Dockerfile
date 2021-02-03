FROM golang:alpine as builder

WORKDIR /build

# ENV GOPROXY=https://goproxy.cn

# RUN apk add --no-cache tzdata

COPY ./go.mod ./

COPY ./go.sum ./

RUN go mod download

COPY . .

# RUN go build -ldflags "-s -w" -o pma .
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o gemini


FROM scratch as runner

# COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

COPY --from=builder /build/gemini /opt/app/

# COPY config.json /opt/app/

EXPOSE 8080

ENTRYPOINT ["/opt/app/gemini"]