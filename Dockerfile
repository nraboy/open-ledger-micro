FROM debian:latest

RUN apt-get update
RUN apt-get install wget git --yes --quiet

RUN wget -O go.tgz "https://dl.google.com/go/go1.10.linux-amd64.tar.gz"
RUN wget -O nodejs.tgz "https://nodejs.org/dist/v9.9.0/node-v9.9.0-linux-x64.tar.gz"
RUN tar -C /usr/local/ -xzf go.tgz
RUN tar -C /usr/local/ -xzf nodejs.tgz && mv /usr/local/node-v9.9.0-linux-x64 /usr/local/nodejs
RUN rm go.tgz
RUN rm nodejs.tgz

EXPOSE 80 443 12345

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
ENV PATH /usr/local/nodejs/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
RUN mkdir -p "$GOPATH/src/github.com/nraboy/open-ledger-micro"

WORKDIR $GOPATH

COPY . $GOPATH/src/github.com/nraboy/open-ledger-micro/

RUN go get github.com/GeertJohan/go.rice
RUN go get github.com/GeertJohan/go.rice/rice
RUN go get github.com/gorilla/mux
RUN go get github.com/btcsuite/btcd
RUN go get github.com/btcsuite/btcutil
RUN go get github.com/gorilla/handlers

RUN cd $GOPATH/src/github.com/nraboy/open-ledger-micro/ui && npm install && npm run build
RUN cd $GOPATH/src/github.com/nraboy/open-ledger-micro && rice embed-go && go build && go install

CMD ["open-ledger-micro"]
