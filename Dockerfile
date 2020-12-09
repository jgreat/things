from golang:1.15.4-buster as builder

WORKDIR /app
COPY . /app

ENV CGO_ENABLED 0

RUN  go mod vendor \
  && go build -o api

FROM debian:buster-slim

RUN  addgroup --system --gid 1000 app \
  && adduser --system --ingroup app --uid 1000 app \
  && apt-get update \
  && apt-get upgrade -y \
  && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/api /

USER app

EXPOSE 8080

CMD [ "/api" ]
