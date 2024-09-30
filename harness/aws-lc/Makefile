AWS_LC_SRC_DIR ?=
LIBRPATH = $(shell env PKG_CONFIG_PATH=${PKG_CONFIG_PATH} pkg-config --variable=libdir libcrypto)
CXXFLAGS = -std=c++17 -Wl,-rpath=${LIBRPATH}
CPPFLAGS = $(shell env PKG_CONFIG_PATH=${PKG_CONFIG_PATH} pkg-config --cflags libcrypto)
LDLIBS = $(shell env PKG_CONFIG_PATH=${PKG_CONFIG_PATH} pkg-config --libs libcrypto)

.PHONY: all
all: main

.PHONY: debug
debug: CXXFLAGS += -g -fsanitize=address,undefined
debug: main

.PHONY: clean
clean:
	rm -f main

.PHONY: verify-aws-lc-source
verify-aws-lc-source:
	@if [[ -z "${AWS_LC_SRC_DIR}" ]]; then \
		echo "AWS_LC_SRC_DIR environment variable is missing and is required to specify AWS-LC source location" && false; \
	fi

.PHONY: build-image
build-image: verify-aws-lc-source aws-lc.dockerfile
	docker build --platform linux/amd64 -f aws-lc.dockerfile --build-context aws_lc_src=${AWS_LC_SRC_DIR} . -t x509-limbo-aws-lc --load
