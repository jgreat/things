from golang:1.15.4-buster

# Share code as volume on /app
WORKDIR /app

COPY ./.dev/run.sh /

RUN apt-get update \
  && apt-get install inotify-tools -y \
  && rm -rf /var/lib/apt/lists/*

RUN  go get github.com/go-delve/delve/cmd/dlv \
  && addgroup --system --gid 1000 app \
  && adduser --system --ingroup app --uid 1000 app

USER app

CMD [ "/run.sh"]
