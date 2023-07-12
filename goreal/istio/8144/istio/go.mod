module istio.io/istio

go 1.19

require (
	cloud.google.com/go v0.34.0
	code.cloudfoundry.org/copilot v0.0.0-20180522185031-1e7e324f8366
	contrib.go.opencensus.io/exporter/stackdriver v0.2.0
	github.com/DataDog/datadog-go v3.2.0+incompatible
	github.com/alicebob/miniredis v0.0.0-20180201100744-9d52b1fc8da9
	github.com/aws/aws-sdk-go v1.40.45
	github.com/bradfitz/slice v0.0.0-20160823171949-d9036e2120b5
	github.com/cactus/go-statsd-client v3.1.1+incompatible
	github.com/cenkalti/backoff v2.0.0+incompatible
	github.com/circonus-labs/circonus-gometrics v2.3.1+incompatible
	github.com/coreos/go-oidc v0.0.0-20180117170138-065b426bd416
	github.com/davecgh/go-spew v1.1.1
	github.com/emicklei/go-restful v2.6.0+incompatible
	github.com/envoyproxy/go-control-plane v0.11.1
	github.com/evanphx/json-patch v3.0.0+incompatible
	github.com/fluent/fluent-logger-golang v1.3.0
	github.com/ghodss/yaml v1.0.0
	github.com/go-redis/redis v6.10.2+incompatible
	github.com/gogo/googleapis v0.0.0-20180223154316-0cd9801be74a
	github.com/gogo/protobuf v1.3.2
	github.com/gogo/status v1.0.3
	github.com/golang/glog v1.1.1
	github.com/golang/protobuf v1.5.3
	github.com/golang/sync v0.0.0-20180314180146-1d60e4601c6f
	github.com/google/go-github v15.0.0+incompatible
	github.com/google/uuid v1.3.0
	github.com/googleapis/gax-go v2.0.0+incompatible
	github.com/gorilla/mux v1.6.1
	github.com/gorilla/websocket v1.2.0
	github.com/grpc-ecosystem/go-grpc-middleware v0.0.0-20180323085839-aed189ae50cf
	github.com/grpc-ecosystem/go-grpc-prometheus v0.0.0-20160910222444-6b7015e65d36
	github.com/grpc-ecosystem/grpc-opentracing v0.0.0-20171214222146-0e7658f8ee99
	github.com/hashicorp/consul v1.0.6
	github.com/hashicorp/go-multierror v1.1.0
	github.com/hashicorp/vault v0.10.0
	github.com/howeyc/fsnotify v0.9.0
	github.com/natefinch/lumberjack v0.0.0-20170531160350-a96e63847dc3
	github.com/onsi/gomega v1.3.0
	github.com/open-policy-agent/opa v0.8.2
	github.com/openshift/api v0.0.0-20180301222200-dca24d1902af
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pborman/uuid v0.0.0-20170612153648-e790cca94e6c
	github.com/pkg/errors v0.9.1
	github.com/pmezard/go-difflib v1.0.0
	github.com/prometheus/client_golang v1.16.0
	github.com/prometheus/client_model v0.4.0
	github.com/prometheus/common v0.44.0
	github.com/prometheus/prom2json v0.0.0-20180219155529-daa2bca1c13f
	github.com/satori/go.uuid v1.2.0
	github.com/signalfx/com_signalfx_metrics_protobuf v0.0.0-20170330202426-93e507b42f43
	github.com/signalfx/golib v1.1.6
	github.com/spf13/cobra v0.0.2
	github.com/spf13/pflag v1.0.0
	github.com/stretchr/testify v1.8.3
	github.com/uber/jaeger-client-go v2.12.0+incompatible
	go.opencensus.io v0.24.0
	go.uber.org/atomic v1.9.0
	go.uber.org/multierr v1.7.0
	go.uber.org/zap v1.19.1
	golang.org/x/net v0.11.0
	golang.org/x/oauth2 v0.9.0
	golang.org/x/time v0.0.0-20210723032227-1f47c861a9ac
	golang.org/x/tools v0.8.0
	google.golang.org/api v0.129.0
	google.golang.org/genproto v0.0.0-20230530153820-e85fd2cbaebc
	google.golang.org/grpc v1.56.1
	gopkg.in/d4l3k/messagediff.v1 v1.2.1
	gopkg.in/validator.v2 v2.0.0-20180205153750-59c90c7046f6
	gopkg.in/yaml.v2 v2.4.0
	istio.io/api v0.0.0-20180730173215-123b0a79a4db
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
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/Masterminds/semver v1.4.0 // indirect
	github.com/Masterminds/sprig v2.14.1+incompatible // indirect
	github.com/SAP/go-hdb v1.3.9 // indirect
	github.com/SermoDigital/jose v0.9.1 // indirect
	github.com/alecthomas/template v0.0.0-20160405071501-a0175ee3bccc // indirect
	github.com/alecthomas/units v0.0.0-20151022065526-2efee857e7cf // indirect
	github.com/alicebob/gopher-json v0.0.0-20180125190556-5a6b3ba71ee6 // indirect
	github.com/aokoli/goutils v1.0.1 // indirect
	github.com/apache/thrift v0.0.0-20161221203622-b2a4d4ae21c7 // indirect
	github.com/armon/go-metrics v0.3.9 // indirect
	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
	github.com/census-instrumentation/opencensus-proto v0.4.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/circonus-labs/circonusllhist v0.1.3 // indirect
	github.com/cncf/xds/go v0.0.0-20230607035331-e9ce68804cb4 // indirect
	github.com/codahale/hdrhistogram v0.9.0 // indirect
	github.com/cpuguy83/go-md2man v1.0.8 // indirect
	github.com/d4l3k/messagediff v1.2.1 // indirect
	github.com/dchest/siphash v1.1.0 // indirect
	github.com/denisenkom/go-mssqldb v0.12.3 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/docker/distribution v2.6.0-rc.1.0.20180327202408-83389a148052+incompatible // indirect
	github.com/docker/spdystream v0.0.0-20170912183627-bc6354cbbc29 // indirect
	github.com/dropbox/godropbox v0.0.0-20230623171840-436d2007a9fd // indirect
	github.com/duosecurity/duo_api_golang v0.0.0-20230418202038-096d3306c029 // indirect
	github.com/elazarl/go-bindata-assetfs v1.0.1 // indirect
	github.com/elazarl/goproxy v0.0.0-20221015165544-a0805db90819 // indirect
	github.com/envoyproxy/protoc-gen-validate v1.0.1 // indirect
	github.com/facebookgo/stack v0.0.0-20160209184415-751773369052 // indirect
	github.com/facebookgo/stackerr v0.0.0-20150612192056-c2fcf88613f4 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/garyburd/redigo v1.6.0 // indirect
	github.com/go-kit/kit v0.12.0 // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/gocql/gocql v1.5.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/snappy v0.0.3 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/go-querystring v0.0.0-20170111101155-53e6ce116135 // indirect
	github.com/google/gofuzz v1.0.0 // indirect
	github.com/google/s2a-go v0.1.4 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.2.5 // indirect
	github.com/googleapis/gax-go/v2 v2.11.0 // indirect
	github.com/googleapis/gnostic v0.1.0 // indirect
	github.com/gophercloud/gophercloud v0.0.0-20180327194212-2daf3049f2a9 // indirect
	github.com/gorilla/context v0.0.0-20160226214623-1ea25387ff6f // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.5.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-memdb v1.3.4 // indirect
	github.com/hashicorp/go-msgpack v0.5.5 // indirect
	github.com/hashicorp/go-plugin v1.4.10 // indirect
	github.com/hashicorp/go-retryablehttp v0.5.3 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/hcl v0.0.0-20180404174102-ef8a98b0bbce // indirect
	github.com/hashicorp/memberlist v0.5.0 // indirect
	github.com/hashicorp/serf v0.9.5 // indirect
	github.com/howeyc/gopass v0.0.0-20170109162249-bf9dde6d0d2c // indirect
	github.com/huandu/xstrings v1.0.0 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jefferai/jsonx v1.0.1 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/juju/errors v1.0.0 // indirect
	github.com/keybase/go-crypto v0.0.0-20200123153347-de78d2cb44f4 // indirect
	github.com/kr/logfmt v0.0.0-20140226030751-b84e30acd515 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/mitchellh/mapstructure v1.4.2 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/onsi/ginkgo v1.4.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/ory/dockertest/v3 v3.0.0-00010101000000-000000000000 // indirect
	github.com/pascaldekloe/goe v0.1.1 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/philhofer/fwd v1.0.0 // indirect
	github.com/pquerna/cachecontrol v0.0.0-20180306154005-525d0eb5f91d // indirect
	github.com/prometheus/procfs v0.11.0 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20180125231941-8732c616f529 // indirect
	github.com/rogpeppe/go-internal v1.10.0 // indirect
	github.com/russross/blackfriday v1.5.1 // indirect
	github.com/ryanuber/go-glob v0.0.0-20160226084822-572520ed46db // indirect
	github.com/sethgrid/pester v0.0.0-20180227223404-ed9870dad317 // indirect
	github.com/signalfx/gohistogram v0.0.0-20160107210732-1ccfd2ff5083 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/smartystreets/goconvey v1.8.0 // indirect
	github.com/square/certstrap v1.1.1 // indirect
	github.com/tinylib/msgp v1.0.2 // indirect
	github.com/tv42/httpunix v0.0.0-20150427012821-b75d8614f926 // indirect
	github.com/uber/jaeger-lib v1.4.0 // indirect
	github.com/yashtewari/glob-intersection v0.0.0-20180206001645-7af743e8ec84 // indirect
	github.com/yuin/gopher-lua v0.0.0-20180316054350-84ea3a3c79b3 // indirect
	go4.org v0.0.0-20180417224846-9599cf28b011 // indirect
	golang.org/x/crypto v0.10.0 // indirect
	golang.org/x/mod v0.10.0 // indirect
	golang.org/x/sync v0.3.0 // indirect
	golang.org/x/sys v0.9.0 // indirect
	golang.org/x/term v0.9.0 // indirect
	golang.org/x/text v0.10.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20230530153820-e85fd2cbaebc // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/alecthomas/kingpin.v2 v2.2.6 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/logfmt.v0 v0.3.0 // indirect
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
	gopkg.in/ory-am/dockertest.v3 v3.0.0-00010101000000-000000000000 // indirect
	gopkg.in/square/go-jose.v2 v2.1.6 // indirect
	gopkg.in/stack.v1 v1.7.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/apiserver v0.0.0-20180327065226-f4a9d3132586 // indirect
	k8s.io/kube-openapi v0.0.0-20180216212618-50ae88d24ede // indirect
)

replace (
	cloud.google.com/go => cloud.google.com/go v0.26.0
	github.com/armon/go-metrics => github.com/hashicorp/go-metrics v0.5.1
	github.com/ory/dockertest/v3 => gopkg.in/ory-am/dockertest.v3 v3.10.0
	github.com/prometheus/common => github.com/prometheus/common v0.0.0-20180326160409-38c53a9f4bfc
	google.golang.org/genproto => google.golang.org/genproto v0.0.0-20180323190852-ab0870e398d5
	gopkg.in/ory-am/dockertest.v3 => github.com/ory/dockertest/v3 v3.10.0
	k8s.io/kubernetes => k8s.io/kubernetes v1.10.4
)
