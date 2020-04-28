#!/bin/bash

export GOOS=linux
export VOLANTMQ_WORK_DIR=/usr/lib/volantmq
export VOLANTMQ_BUILD_FLAGS="-i"
export VOLANTMQ_PLUGINS_DIR=/usr/lib/volantmq/plugins
export GO111MODULE=on
export GOPATH=/home/administrator/go

sudo mkdir -p $VOLANTMQ_WORK_DIR/bin
sudo mkdir -p $VOLANTMQ_WORK_DIR/conf
sudo mkdir -p $VOLANTMQ_PLUGINS_DIR

#Compilar servidor VOLANT

sudo cp $GOPATH/src/github.com/fatimafp95/volantmq_2/tools/print_version.sh /bin
export GO111MODULE=off
go get github.com/troian/govvv
cd cmd/volantmq/
sudo /usr/local/go/bin/go build $VOLANTMQ_BUILD_FLAGS -o $VOLANTMQ_WORK_DIR/bin/volantmq
sudo cp volantmq $VOLANTMQ_WORK_DIR/bin/volantmq

#Compilar debug plugins
go get gitlab.com/VolantMQ/vlplugin/debug
cd $GOPATH/src/gitlab.com/VolantMQ/vlplugin/debug
export GO111MODULE=on
go mod tidy
sudo /usr/local/go/bin/go build $VOLANTMQ_BUILD_FLAGS -buildmode=plugin -ldflags "-X main.version=$(print_version.sh)" -o $VOLANTMQ_WORK_DIR/plugins/debug.so

#Health plugin
export GO111MODULE=off
go get gitlab.com/VolantMQ/vlplugin/health
export GO111MODULE=on
sudo /usr/local/go/bin/go build $VOLANTMQ_BUILD_FLAGS -buildmode=plugin -ldflags "-X main.version=$(print_version.sh)" -o $VOLANTMQ_WORK_DIR/plugins/health.so

#build metrics plugins 
export GO111MODULE=off
go get gitlab.com/VolantMQ/vlplugin/monitoring/prometheus
cd $GOPATH/src/gitlab.com/VolantMQ/vlplugin/monitoring/prometheus
export GO111MODULE=on
go mod tidy
sudo /usr/local/go/bin/go build $VOLANTMQ_BUILD_FLAGS -buildmode=plugin -ldflags "-X main.version=$(print_version.sh)" -o $VOLANTMQ_WORK_DIR/plugins/prometheus.so

export GO111MODULE=off
go get gitlab.com/VolantMQ/vlplugin/monitoring/systree
cd $GOPATH/src/gitlab.com/VolantMQ/vlplugin/monitoring/systree
export GO111MODULE=on
go mod tidy
sudo /usr/local/go/bin/go build $VOLANTMQ_BUILD_FLAGS -buildmode=plugin -ldflags "-X main.version=$(print_version.sh)" -o $VOLANTMQ_WORK_DIR/plugins/systree.so

#build persistence plugins
export GO111MODULE=off
go get gitlab.com/VolantMQ/vlplugin/persistence/bbolt
cd $GOPATH/src/gitlab.com/VolantMQ/vlplugin/persistence/bbolt
export GO111MODULE=on
go mod tidy
cd plugin
sudo /usr/local/go/bin/go build $VOLANTMQ_BUILD_FLAGS -buildmode=plugin -ldflags "-X main.version=$(print_version.sh)" -o $VOLANTMQ_WORK_DIR/plugins/persistence_bbolt.so

#build auth plugins
export GO111MODULE=off
go get gitlab.com/VolantMQ/vlplugin/auth/http
cd $GOPATH/src/gitlab.com/VolantMQ/vlplugin/auth/http
GO111MODULE=on
go mod tidy
sudo /usr/local/go/bin/go build $VOLANTMQ_BUILD_FLAGS -buildmode=plugin -ldflags "-X main.version=$(print_version.sh)" -o $VOLANTMQ_WORK_DIR/plugins/auth_http.so

#Ejecución del broker con el fichero de configuración
cd $GOPATH/src/github.com/fatimafp95/volantmq_2
sudo cp -r /usr/lib/volantmq /var/lib
sudo /usr/lib/volantmq/bin/volantmq --config=examples/config.yaml
