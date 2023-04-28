FROM golang:alpine AS build
WORKDIR /src
COPY *.go go.* /src/
RUN go build -o /bin/redir

FROM scratch
WORKDIR /app
COPY --from=build /bin/redir ./redir
EXPOSE 3000
ENTRYPOINT ["/app/redir"]
