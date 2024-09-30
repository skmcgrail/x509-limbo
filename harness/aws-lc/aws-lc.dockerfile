FROM amazonlinux:2023 AS aws_lc_build

RUN yum install -y ninja-build cmake golang gcc gcc-c++ git

COPY --from=aws_lc_src . /src

WORKDIR /src

RUN cmake -B build -S . -GNinja -DBUILD_SHARED_LIBS=1 -DCMAKE_BUILD_TYPE=Debug -DCMAKE_INSTALL_PREFIX=/opt/aws-lc && \
    cmake --build build && \
    cmake --install build

FROM amazonlinux:2023

RUN yum install -y gcc gcc-c++ make

COPY --from=aws_lc_build /opt/aws-lc /opt/aws-lc

WORKDIR /

ENV PKG_CONFIG_PATH=/opt/aws-lc/lib64/pkgconfig

COPY date.hpp json.hpp main.cpp Makefile /

RUN make

ENTRYPOINT ["./main"]
