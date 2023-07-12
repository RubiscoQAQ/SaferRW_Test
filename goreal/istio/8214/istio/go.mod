module istio.io/istio

go 1.19

require (
	cloud.google.com/go v0.23.0
	code.cloudfoundry.org/copilot v0.0.0-20180808174356-6bade2a0677a
	contrib.go.opencensus.io/exporter/stackdriver v0.5.0
	github.com/DataDog/datadog-go v0.0.0-20180129105149-9487d3a9d3be
	github.com/alicebob/miniredis v0.0.0-20180201100744-9d52b1fc8da9
	github.com/aws/aws-sdk-go v1.13.24
	github.com/cactus/go-statsd-client v3.1.1+incompatible
	github.com/cenkalti/backoff v2.0.0+incompatible
	github.com/circonus-labs/circonus-gometrics v2.1.1+incompatible
	github.com/coreos/go-oidc v0.0.0-20180117170138-065b426bd416
	github.com/davecgh/go-spew v1.1.0
	github.com/emicklei/go-restful v2.6.0+incompatible
	github.com/envoyproxy/go-control-plane v0.5.0
	github.com/evanphx/json-patch v3.0.0+incompatible
	github.com/fluent/fluent-logger-golang v1.3.0
	github.com/fsnotify/fsnotify v1.4.7
	github.com/ghodss/yaml v1.0.0
	github.com/go-redis/redis v6.10.2+incompatible
	github.com/gogo/googleapis v0.0.0-20180223154316-0cd9801be74a
	github.com/gogo/protobuf v1.1.1
	github.com/gogo/status v1.0.3
	github.com/golang/protobuf v1.1.0
	github.com/golang/sync v0.0.0-20180314180146-1d60e4601c6f
	github.com/google/go-github v15.0.0+incompatible
	github.com/google/uuid v0.0.0-20161128191214-064e2069ce9c
	github.com/googleapis/gax-go v2.0.0+incompatible
	github.com/gorilla/mux v1.6.1
	github.com/gorilla/websocket v1.2.0
	github.com/grpc-ecosystem/go-grpc-middleware v0.0.0-20180323085839-aed189ae50cf
	github.com/grpc-ecosystem/go-grpc-prometheus v0.0.0-20160910222444-6b7015e65d36
	github.com/grpc-ecosystem/grpc-opentracing v0.0.0-20171214222146-0e7658f8ee99
	github.com/hashicorp/consul v1.0.6
	github.com/hashicorp/go-multierror v0.0.0-20171204182908-b7773ae21874
	github.com/hashicorp/vault v0.10.0
	github.com/howeyc/fsnotify v0.9.0
	github.com/natefinch/lumberjack v0.0.0-20170531160350-a96e63847dc3
	github.com/onsi/gomega v1.3.0
	github.com/open-policy-agent/opa v0.8.2
	github.com/openshift/api v0.0.0-20180301222200-dca24d1902af
	github.com/opentracing/opentracing-go v1.0.2
	github.com/pborman/uuid v0.0.0-20170612153648-e790cca94e6c
	github.com/pkg/errors v0.8.0
	github.com/pmezard/go-difflib v1.0.0
	github.com/prometheus/client_golang v0.9.0-pre1
	github.com/prometheus/client_model v0.0.0-20171117100541-99fa1f4be8e5
	github.com/prometheus/common v0.0.0-20180326160409-38c53a9f4bfc
	github.com/prometheus/prom2json v0.0.0-20180219155529-daa2bca1c13f
	github.com/satori/go.uuid v1.2.0
	github.com/signalfx/com_signalfx_metrics_protobuf v0.0.0-20170330202426-93e507b42f43
	github.com/signalfx/golib v1.1.6
	github.com/spf13/cobra v0.0.2
	github.com/spf13/pflag v1.0.0
	github.com/stretchr/testify v1.2.1
	github.com/uber/jaeger-client-go v2.12.0+incompatible
	go.opencensus.io v0.14.0
	go.uber.org/atomic v1.3.1
	go.uber.org/multierr v1.1.0
	go.uber.org/zap v1.7.1
	golang.org/x/net v0.0.0-20180320002117-6078986fec03
	golang.org/x/oauth2 v0.0.0-20180314180239-fdc9e635145a
	golang.org/x/time v0.0.0-20180314180208-26559e0f760e
	golang.org/x/tools v0.0.0-20180324185418-77106db15f68
	google.golang.org/genproto v0.0.0-20180323190852-ab0870e398d5
	google.golang.org/grpc v1.14.0
	gopkg.in/d4l3k/messagediff.v1 v1.2.1
	gopkg.in/validator.v2 v2.0.0-20180205153750-59c90c7046f6
	gopkg.in/yaml.v2 v2.2.1
	istio.io/api v0.0.0-20180817192745-214c7598afb7
	k8s.io/api v0.0.0-20180601181742-8b7507fac302
	k8s.io/apiextensions-apiserver v0.0.0-20180601203502-8e7f43002fec
	k8s.io/apimachinery v0.0.0-20180601181227-17529ec7eadb
	k8s.io/client-go v0.0.0-20180601184321-26a26f55b28a
	k8s.io/cluster-registry v0.0.6
	k8s.io/helm v2.9.1+incompatible
	k8s.io/ingress v0.0.0-20170803151325-fe19ebb09ee2
	k8s.io/kubernetes v1.10.4
)

require (
	github.com/Azure/go-autorest v9.10.0+incompatible // indirect
	github.com/BurntSushi/toml v0.3.0 // indirect
	github.com/Masterminds/semver v1.4.0 // indirect
	github.com/Masterminds/sprig v2.14.1+incompatible // indirect
	github.com/alecthomas/template v0.0.0-20160405071501-a0175ee3bccc // indirect
	github.com/alecthomas/units v0.0.0-20151022065526-2efee857e7cf // indirect
	github.com/alicebob/gopher-json v0.0.0-20180125190556-5a6b3ba71ee6 // indirect
	github.com/aokoli/goutils v1.0.1 // indirect
	github.com/apache/thrift v0.0.0-20161221203622-b2a4d4ae21c7 // indirect
	github.com/beorn7/perks v0.0.0-20180321164747-3a771d992973 // indirect
	github.com/circonus-labs/circonusllhist v0.0.0-20180104205821-1e65893c4458 // indirect
	github.com/codahale/hdrhistogram v0.9.0 // indirect
	github.com/cpuguy83/go-md2man v1.0.8 // indirect
	github.com/dchest/siphash v1.1.0 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/docker/distribution v2.6.0-rc.1.0.20180327202408-83389a148052+incompatible // indirect
	github.com/docker/spdystream v0.0.0-20170912183627-bc6354cbbc29 // indirect
	github.com/garyburd/redigo v1.6.0 // indirect
	github.com/go-ini/ini v1.33.0 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/golang/groupcache v0.0.0-20180203143532-66deaeb636df // indirect
	github.com/golang/snappy v0.0.0-20170215233205-553a64147049 // indirect
	github.com/google/go-querystring v0.0.0-20170111101155-53e6ce116135 // indirect
	github.com/google/gofuzz v0.0.0-20170612174753-24818f796faf // indirect
	github.com/googleapis/gnostic v0.1.0 // indirect
	github.com/gophercloud/gophercloud v0.0.0-20180327194212-2daf3049f2a9 // indirect
	github.com/gorilla/context v0.0.0-20160226214623-1ea25387ff6f // indirect
	github.com/hashicorp/errwrap v0.0.0-20141028054710-7554cd9344ce // indirect
	github.com/hashicorp/go-cleanhttp v0.0.0-20171218145408-d5fe4b57a186 // indirect
	github.com/hashicorp/go-retryablehttp v0.0.0-20170824180859-794af36148bf // indirect
	github.com/hashicorp/go-rootcerts v0.0.0-20160503143440-6bb64b370b90 // indirect
	github.com/hashicorp/golang-lru v0.0.0-20180201235237-0fb14efe8c47 // indirect
	github.com/hashicorp/hcl v0.0.0-20180404174102-ef8a98b0bbce // indirect
	github.com/hashicorp/serf v0.8.1 // indirect
	github.com/howeyc/gopass v0.0.0-20170109162249-bf9dde6d0d2c // indirect
	github.com/huandu/xstrings v1.0.0 // indirect
	github.com/imdario/mergo v0.0.0-20141206190957-6633656539c1 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jmespath/go-jmespath v0.0.0-20160202185014-0b12d6b521d8 // indirect
	github.com/json-iterator/go v0.0.0-20180315132816-ca39e5af3ece // indirect
	github.com/kr/logfmt v0.0.0-20140226030751-b84e30acd515 // indirect
	github.com/lyft/protoc-gen-validate v0.0.6 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.0 // indirect
	github.com/mitchellh/go-homedir v0.0.0-20161203194507-b8bc1bf76747 // indirect
	github.com/mitchellh/mapstructure v0.0.0-20180220230111-00c29f56e238 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v0.0.0-20180228065516-1df9eeb2bb81 // indirect
	github.com/onsi/ginkgo v1.6.0 // indirect
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/philhofer/fwd v1.0.0 // indirect
	github.com/pquerna/cachecontrol v0.0.0-20180306154005-525d0eb5f91d // indirect
	github.com/prometheus/procfs v0.0.0-20180321230812-780932d4fbbe // indirect
	github.com/rcrowley/go-metrics v0.0.0-20180125231941-8732c616f529 // indirect
	github.com/russross/blackfriday v1.5.1 // indirect
	github.com/ryanuber/go-glob v0.0.0-20160226084822-572520ed46db // indirect
	github.com/sethgrid/pester v0.0.0-20180227223404-ed9870dad317 // indirect
	github.com/signalfx/gohistogram v0.0.0-20160107210732-1ccfd2ff5083 // indirect
	github.com/sirupsen/logrus v1.0.5 // indirect
	github.com/square/certstrap v1.1.1 // indirect
	github.com/tinylib/msgp v1.0.2 // indirect
	github.com/tv42/httpunix v0.0.0-20150427012821-b75d8614f926 // indirect
	github.com/uber/jaeger-lib v1.4.0 // indirect
	github.com/yashtewari/glob-intersection v0.0.0-20180206001645-7af743e8ec84 // indirect
	github.com/yuin/gopher-lua v0.0.0-20180316054350-84ea3a3c79b3 // indirect
	golang.org/x/crypto v0.0.0-20180322175230-88942b9c40a4 // indirect
	golang.org/x/sync v0.0.0-20180314180146-1d60e4601c6f // indirect
	golang.org/x/sys v0.0.0-20180329131831-378d26f46672 // indirect
	golang.org/x/text v0.3.0 // indirect
	google.golang.org/appengine v1.0.0 // indirect
	gopkg.in/alecthomas/kingpin.v2 v2.2.6 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/logfmt.v0 v0.3.0 // indirect
	gopkg.in/square/go-jose.v2 v2.1.6 // indirect
	gopkg.in/stack.v1 v1.7.0 // indirect
	k8s.io/apiserver v0.0.0-20180327065226-f4a9d3132586 // indirect
	k8s.io/kube-openapi v0.0.0-20180216212618-50ae88d24ede // indirect
)
