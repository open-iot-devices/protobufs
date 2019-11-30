FROM ubuntu:18.04

ENV PROTO_VERSION 3.11.0

# Install essentials
RUN apt-get update && apt-get install -y curl gcc g++ make golang git python python-pip && \
    pip install protobuf

# Install protobuf compiler
RUN curl -L https://github.com/protocolbuffers/protobuf/releases/download/v${PROTO_VERSION}/protobuf-all-${PROTO_VERSION}.tar.gz -o protobuf-${PROTO_VERSION}.tar.gz && \
    tar xzf protobuf-${PROTO_VERSION}.tar.gz && \
    cd protobuf-${PROTO_VERSION} && \
    ./configure && make install -j 4 && ldconfig && \
    cd .. && rm -rf protobuf*

# Install protobufs for golang compiler
RUN go get -u github.com/golang/protobuf/protoc-gen-go

# Install protobufs for c (nanobp)
RUN git clone https://github.com/nanopb/nanopb.git

COPY generate.sh /
CMD ["/generate.sh"]
