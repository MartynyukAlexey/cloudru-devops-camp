# two phase build (i've seen it in golang and frontend projects)
# we first build the binary in the heavy container with toolchain and depencencies.
# then we copy the binary in the lean image and start it there

FROM golang:1.24-alpine AS build-stage

WORKDIR /01-application

COPY go.mod go.sum ./ 

# I didn't use any dependencies, so no need for this step
# RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /build ./



FROM scratch AS release-stage

WORKDIR /01-application

COPY --from=build-stage /build /build

ENTRYPOINT ["/build"]