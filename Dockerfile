FROM google/golang

WORKDIR /gopath/src/blockchain
ADD . /gopath/src/blockchain
RUN go get blockchain

CMD []
ENTRYPOINT ["/gopath/bin/blockchain"]