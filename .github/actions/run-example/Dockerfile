# define base image
FROM golang:1.12 AS build
# set the working directory
WORKDIR /src
# copy the sorce code into the /src directory
COPY /home/runner/work/go-module-example/go-module-example/module-one .
COPY /home/runner/work/go-module-example/go-module-example/module-two .
# build the moduletwo.go
RUN go build -o /out/moduletwo ./module-two/moduletwo.go
FROM scratch as bin
COPY --from=build /out/moduletwo /
RUN ./moduletwo

