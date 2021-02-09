# define base image
FROM golang:1.12 AS build
# set the working directory
# WORKDIR /src
# maybe don't set the working directory
# copy the sorce code into the /src directory
COPY ./module-one .
COPY ./module-two .
# build the moduletwo.go
RUN go build -o /out/moduletwo ./module-two/moduletwo.go
FROM scratch as bin
COPY --from=build /out/moduletwo /
RUN ./moduletwo

