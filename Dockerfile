FROM golang:1.16.4

ENV WORKPATH=/app/src/github.com/jiaqi-yin/go-verification-code
COPY . ${WORKPATH}
WORKDIR ${WORKPATH}/src

RUN go build -o go-verification-code-api .

EXPOSE 8080

CMD [ "./go-verification-code-api" ]