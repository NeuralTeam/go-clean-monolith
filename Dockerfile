FROM golang:1.21-alpine AS build

WORKDIR /tmp/backend
COPY . /tmp/backend

RUN go mod tidy
RUN go build -o /bin/backend ./cmd/main.go


FROM scratch

COPY --from=build /bin/backend /bin/backend

# You can uncomment this
#COPY --from=build /tmp/backend/.env /bin/.env

CMD ["/bin/backend", "start"]