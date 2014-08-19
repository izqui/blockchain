FROM google/golang

WORKDIR /gopath/src/blockchain
ADD . /gopath/src/blockchain
RUN go get blockchain

CMD []
EXPOSE 9119
ENTRYPOINT blockchain
