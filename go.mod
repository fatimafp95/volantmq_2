module github.com/VolantMQ/volantmq

go 1.13
//replace github.com/lucas-clemente/quic-go v0.15.1 => /home/administrator/go/src/github.com/fatimafp95/quic-go

require (
	github.com/VolantMQ/vlapi v0.5.4
//	github.com/blang/semver v3.5.1+incompatible // indirect
//	github.com/coreos/bbolt v1.3.3 // indirect
	github.com/gobwas/httphead v0.0.0-20180130184737-2c6c146eadee // indirect
	github.com/gobwas/pool v0.2.0 // indirect
	github.com/gobwas/ws v1.0.2
	github.com/gorilla/websocket v1.4.1 // indirect
//	github.com/lucas-clemente/quic-go v0.15.1
	github.com/fatimafp95/quic-go v0.0.0
	github.com/pkg/errors v0.8.1
	github.com/stretchr/testify v1.4.0
	github.com/troian/healthcheck v0.1.3
	github.com/vbauerster/mpb/v4 v4.9.4
	gitlab.com/VolantMQ/vlplugin/auth/http v0.0.2
	gitlab.com/VolantMQ/vlplugin/debug v0.0.8
	gitlab.com/VolantMQ/vlplugin/health v0.0.8
	gitlab.com/VolantMQ/vlplugin/monitoring/prometheus v0.0.5
	gitlab.com/VolantMQ/vlplugin/monitoring/systree v0.0.7
	gitlab.com/VolantMQ/vlplugin/persistence/bbolt v0.0.7
	gitlab.com/VolantMQ/vlplugin/persistence/mem v0.0.4
//	go.etcd.io/bbolt v1.3.3 // indirect
	go.uber.org/zap v1.12.0
	gopkg.in/yaml.v3 v3.0.0-20200121175148-a6ecf24a6d71
)


