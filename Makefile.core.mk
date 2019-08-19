# Copyright 2019 Istio Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

PROTOC = protoc

GOOGLEAPIS_SHA = 2912605a8d9a126bb24e418bc4b089ad0a2b62dc
GOOGLEAPIS_URL = https://raw.githubusercontent.com/googleapis/googleapis/$(GOOGLEAPIS_SHA)

CENSUS_SHA = e2601ef16f8a085a69d94ace5133f97438f8945f
CENSUS_URL = https://raw.githubusercontent.com/census-instrumentation/opencensus-proto/$(CENSUS_SHA)/src

PROMETHEUS_SHA = 6f3806018612930941127f2a7c6c453ba2c527d2
PROMETHEUS_URL = https://raw.githubusercontent.com/prometheus/client_model/$(PROMETHEUS_SHA)

GOGO_PROTO_PKG := github.com/gogo/protobuf/gogoproto
GOGO_TYPES := github.com/gogo/protobuf/types
GOGO_DESCRIPTOR := github.com/gogo/protobuf/protoc-gen-gogo/descriptor

importmaps := \
	gogoproto/gogo.proto=$(GOGO_PROTO_PKG) \
	google/protobuf/any.proto=$(GOGO_TYPES) \
	google/protobuf/descriptor.proto=$(GOGO_DESCRIPTOR) \
	google/protobuf/duration.proto=$(GOGO_TYPES) \
	google/protobuf/timestamp.proto=$(GOGO_TYPES) \
	google/protobuf/wrappers.proto=$(GOGO_TYPES) \
	google/protobuf/struct.proto=$(GOGO_TYPES) \

comma := ,
empty :=
space := $(empty) $(empty)
mapping_with_spaces := $(foreach map,$(importmaps),M$(map),)
MAPPING := $(subst $(space),$(empty),$(mapping_with_spaces))
GOGOSLICK_PLUGIN := --gogoslick_out=plugins=grpc,$(MAPPING)
GOGOFASTER_PLUGIN := --gogofaster_out=plugins=grpc,$(MAPPING)

googleapis_protos = \
	google/api/http.proto \
	google/api/annotations.proto \
	google/rpc/status.proto \
	google/rpc/code.proto \
	google/rpc/error_details.proto \
	google/type/color.proto \
	google/type/date.proto \
	google/type/dayofweek.proto \
	google/type/latlng.proto \
	google/type/money.proto \
	google/type/postal_address.proto \
	google/type/timeofday.proto \

googleapis_packages = \
	google/api \
	google/rpc \
	google/type \

census_protos = \
	opencensus/proto/stats/v1/stats.proto \
	opencensus/proto/trace/v1/trace.proto \
	opencensus/proto/trace/v1/trace_config.proto \

census_packages = \
	opencensus/proto/stats/v1 \
	opencensus/proto/trace/v1 \

all: build

$(googleapis_protos): %:
	@curl -sS $(GOOGLEAPIS_URL)/$@ -o googleapis/$@.tmp
	@sed -e '/^option go_package/d' googleapis/$@.tmp > googleapis/$@
	@rm googleapis/$@.tmp

$(googleapis_packages): %: $(googleapis_protos)
	@$(PROTOC) $(GOGOSLICK_PLUGIN):googleapis -I googleapis googleapis/$@/*.proto

$(census_protos): %:
	@curl -sS $(CENSUS_URL)/$@ -o $@
	@sed -i.tmp '/^option go_package/d' $@
	@rm $@.tmp

$(census_packages): %: $(census_protos)
	@$(PROTOC) $(GOGOFASTER_PLUGIN):. -I . $@/*.proto

prometheus/metrics.proto:
	@mkdir -p prometheus
	@curl -sS $(PROMETHEUS_URL)/metrics.proto -o prometheus/metrics.proto

prometheus/metrics.pb.go: prometheus/metrics.proto
	@$(PROTOC) $(GOGOFASTER_PLUGIN):. prometheus/metrics.proto

generate: $(googleapis_packages) $(census_packages) prometheus/metrics.pb.go

build: generate
	@go build ./...

clean:
	@rm -fr */*.pb.go */*/*.pb.go */*/*/*.pb.go

.PHONY: all format build $(googleapis_protos) $(googleapis_packages) clean

include Makefile.common.mk