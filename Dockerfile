FROM golang:alpine AS build
WORKDIR /src
COPY *.go go.* /src
RUN go build -o /bin/redir
RUN touch /bin/_redirects

FROM scratch
WORKDIR /app
COPY --from=build /bin/redir ./redir
COPY --from=build /bin/_redirects ./_redirects
EXPOSE 3000
ENTRYPOINT ["/app/redir"]
