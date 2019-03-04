# build stage
FROM golang:alpine AS build-env
ADD . /src
RUN cd /src && go build -o app get_url_one_param.go

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/app /app/
ENTRYPOINT ./app