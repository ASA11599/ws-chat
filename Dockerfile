FROM golang:1.21.5 AS build

WORKDIR /usr/src/ws-chat/
COPY go.* .

RUN go mod download

COPY internal internal/
COPY main.go .

RUN CGO_ENABLED=0 go build -o ws-chat main.go

FROM scratch

WORKDIR /

COPY --from=build /usr/src/ws-chat/ws-chat .

ENV HOST=0.0.0.0
ENV PORT=80

EXPOSE 80

CMD [ "/ws-chat" ]
