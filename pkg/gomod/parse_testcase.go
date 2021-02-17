package gomod

import "github.com/aquasecurity/go-dep-parser/pkg/types"

var (
	// docker run --name gomod --rm -it golang:1.15 bash
	// export USER=gomod
	// mkdir repo
	// cd repo
	// go mod init github.com/org/repo
	// go get github.com/stretchr/repoify
	// go list -m all | awk 'NR>1 {sub(/^v/, "", $2); sub("-.*", "", $2); printf("{\""$1"\", \""$2"\"},\n")}'
	GoModTestify = []types.Library{
		{"github.com/davecgh/go-spew", "1.1.0"},
		{"github.com/pmezard/go-difflib", "1.0.0"},
		{"github.com/stretchr/objx", "0.1.0"},
		{"github.com/stretchr/testify", "1.7.0"},
		{"gopkg.in/check.v1", "0.0.0"},
		{"gopkg.in/yaml.v3", "3.0.0"},
	}

	// docker run --name gomod --rm -it golang:1.15 bash
	// export USER=gomod
	// mkdir repo
	// cd repo
	// go mod init github.com/org/repo
	// go get github.com/urfave/cli
	// go list -m all | awk 'NR>1 {sub(/^v/, "", $2); sub("-.*", "", $2); printf("{\""$1"\", \""$2"\"},\n")}'
	GoModCLI = []types.Library{
		{"github.com/BurntSushi/toml", "0.3.1"},
		{"github.com/cpuguy83/go-md2man/v2", "2.0.0"},
		{"github.com/pmezard/go-difflib", "1.0.0"},
		{"github.com/russross/blackfriday/v2", "2.0.1"},
		{"github.com/shurcooL/sanitized_anchor_name", "1.0.0"},
		{"github.com/urfave/cli", "1.22.5"},
		{"gopkg.in/check.v1", "0.0.0"},
		{"gopkg.in/yaml.v2", "2.2.2"},
	}

	// docker run --name gomod --rm -it golang:1.15 bash
	// export USER=gomod
	// mkdir repo
	// cd repo
	// go mod init github.com/org/repo
	// go get github.com/aquasecurity/trivy
	// go list -m all | awk 'NR>1 {sub(/^v/, "", $2); sub("-.*", "", $2); printf("{\""$1"\", \""$2"\"},\n")}'
	GoModTrivy = []types.Library{
		{"cloud.google.com/go", "0.65.0"},
		{"cloud.google.com/go/bigquery", "1.8.0"},
		{"cloud.google.com/go/datastore", "1.1.0"},
		{"cloud.google.com/go/pubsub", "1.3.1"},
		{"cloud.google.com/go/storage", "1.10.0"},
		{"dmitri.shuralyov.com/gpu/mtl", "0.0.0"},
		{"github.com/Azure/azure-sdk-for-go", "38.0.0"},
		{"github.com/Azure/go-ansiterm", "0.0.0"},
		{"github.com/Azure/go-autorest/autorest", "0.9.3"},
		{"github.com/Azure/go-autorest/autorest/adal", "0.8.1"},
		{"github.com/Azure/go-autorest/autorest/date", "0.2.0"},
		{"github.com/Azure/go-autorest/autorest/mocks", "0.3.0"},
		{"github.com/Azure/go-autorest/autorest/to", "0.3.0"},
		{"github.com/Azure/go-autorest/autorest/validation", "0.1.0"},
		{"github.com/Azure/go-autorest/logger", "0.1.0"},
		{"github.com/Azure/go-autorest/tracing", "0.5.0"},
		{"github.com/BurntSushi/toml", "0.3.1"},
		{"github.com/BurntSushi/xgb", "0.0.0"},
		{"github.com/GoogleCloudPlatform/docker-credential-gcr", "1.5.0"},
		{"github.com/GoogleCloudPlatform/k8s-cloud-provider", "0.0.0"},
		{"github.com/Microsoft/go-winio", "0.4.15"},
		{"github.com/Microsoft/hcsshim", "0.8.6"},
		{"github.com/NYTimes/gziphandler", "0.0.0"},
		{"github.com/OneOfOne/xxhash", "1.2.7"},
		{"github.com/PuerkitoBio/purell", "1.1.1"},
		{"github.com/PuerkitoBio/urlesc", "0.0.0"},
		{"github.com/VividCortex/ewma", "1.1.1"},
		{"github.com/alcortesm/tgz", "0.0.0"},
		{"github.com/alecthomas/template", "0.0.0"},
		{"github.com/alecthomas/units", "0.0.0"},
		{"github.com/alicebob/gopher-json", "0.0.0"},
		{"github.com/alicebob/miniredis/v2", "2.14.1"},
		{"github.com/anmitsu/go-shlex", "0.0.0"},
		{"github.com/aquasecurity/bolt-fixtures", "0.0.0"},
		{"github.com/aquasecurity/fanal", "0.0.0"},
		{"github.com/aquasecurity/go-dep-parser", "0.0.0"},
		{"github.com/aquasecurity/go-gem-version", "0.0.0"},
		{"github.com/aquasecurity/go-npm-version", "0.0.0"},
		{"github.com/aquasecurity/go-pep440-version", "0.0.0"},
		{"github.com/aquasecurity/go-version", "0.0.0"},
		{"github.com/aquasecurity/testdocker", "0.0.0"},
		{"github.com/aquasecurity/trivy", "0.16.0"},
		{"github.com/aquasecurity/trivy-db", "0.0.0"},
		{"github.com/aquasecurity/vuln-list-update", "0.0.0"},
		{"github.com/araddon/dateparse", "0.0.0"},
		{"github.com/armon/consul-api", "0.0.0"},
		{"github.com/armon/go-socks5", "0.0.0"},
		{"github.com/aws/aws-sdk-go", "1.27.1"},
		{"github.com/beorn7/perks", "1.0.0"},
		{"github.com/bgentry/speakeasy", "0.1.0"},
		{"github.com/blang/semver", "3.5.0"},
		{"github.com/briandowns/spinner", "1.12.0"},
		{"github.com/caarlos0/env/v6", "6.0.0"},
		{"github.com/cenkalti/backoff", "2.2.1"},
		{"github.com/census-instrumentation/opencensus-proto", "0.2.1"},
		{"github.com/cespare/xxhash/v2", "2.1.1"},
		{"github.com/cheggaaa/pb/v3", "3.0.3"},
		{"github.com/chzyer/logex", "1.1.10"},
		{"github.com/chzyer/readline", "0.0.0"},
		{"github.com/chzyer/test", "0.0.0"},
		{"github.com/client9/misspell", "0.3.4"},
		{"github.com/cncf/udpa/go", "0.0.0"},
		{"github.com/cockroachdb/datadriven", "0.0.0"},
		{"github.com/containerd/containerd", "1.3.3"},
		{"github.com/containerd/continuity", "0.0.0"},
		{"github.com/coreos/etcd", "3.3.10"},
		{"github.com/coreos/go-etcd", "2.0.0"},
		{"github.com/coreos/go-oidc", "2.1.0"},
		{"github.com/coreos/go-semver", "0.3.0"},
		{"github.com/coreos/go-systemd", "0.0.0"},
		{"github.com/coreos/pkg", "0.0.0"},
		{"github.com/cpuguy83/go-md2man", "1.0.10"},
		{"github.com/cpuguy83/go-md2man/v2", "2.0.0"},
		{"github.com/creack/pty", "1.1.9"},
		{"github.com/davecgh/go-spew", "1.1.1"},
		{"github.com/deckarep/golang-set", "1.7.1"},
		{"github.com/dgrijalva/jwt-go", "3.2.0"},
		{"github.com/dgryski/go-rendezvous", "0.0.0"},
		{"github.com/dnaeon/go-vcr", "1.0.1"},
		{"github.com/docker/cli", "0.0.0"},
		{"github.com/docker/distribution", "2.7.1"},
		{"github.com/docker/docker", "1.4.2"},
		{"github.com/docker/docker-credential-helpers", "0.6.3"},
		{"github.com/docker/go-connections", "0.4.0"},
		{"github.com/docker/go-units", "0.4.0"},
		{"github.com/docker/spdystream", "0.0.0"},
		{"github.com/dustin/go-humanize", "1.0.0"},
		{"github.com/elazarl/goproxy", "0.0.0"},
		{"github.com/elazarl/goproxy/ext", "0.0.0"},
		{"github.com/emicklei/go-restful", "2.9.5"},
		{"github.com/emirpasic/gods", "1.12.0"},
		{"github.com/envoyproxy/go-control-plane", "0.9.4"},
		{"github.com/envoyproxy/protoc-gen-validate", "0.1.0"},
		{"github.com/evanphx/json-patch", "4.2.0"},
		{"github.com/fatih/color", "1.10.0"},
		{"github.com/flynn/go-shlex", "0.0.0"},
		{"github.com/fsnotify/fsnotify", "1.4.9"},
		{"github.com/ghodss/yaml", "1.0.0"},
		{"github.com/gin-contrib/sse", "0.1.0"},
		{"github.com/gin-gonic/gin", "1.5.0"},
		{"github.com/gliderlabs/ssh", "0.2.2"},
		{"github.com/go-git/gcfg", "1.5.0"},
		{"github.com/go-git/go-billy/v5", "5.0.0"},
		{"github.com/go-git/go-git-fixtures/v4", "4.0.1"},
		{"github.com/go-git/go-git/v5", "5.0.0"},
		{"github.com/go-gl/glfw", "0.0.0"},
		{"github.com/go-gl/glfw/v3.3/glfw", "0.0.0"},
		{"github.com/go-kit/kit", "0.8.0"},
		{"github.com/go-logfmt/logfmt", "0.3.0"},
		{"github.com/go-logr/logr", "0.1.0"},
		{"github.com/go-openapi/jsonpointer", "0.19.3"},
		{"github.com/go-openapi/jsonreference", "0.19.3"},
		{"github.com/go-openapi/spec", "0.19.3"},
		{"github.com/go-openapi/swag", "0.19.5"},
		{"github.com/go-playground/locales", "0.13.0"},
		{"github.com/go-playground/universal-translator", "0.17.0"},
		{"github.com/go-redis/redis", "6.15.7"},
		{"github.com/go-redis/redis/v8", "8.4.0"},
		{"github.com/go-restruct/restruct", "0.0.0"},
		{"github.com/go-sql-driver/mysql", "1.5.0"},
		{"github.com/go-stack/stack", "1.8.0"},
		{"github.com/gobwas/glob", "0.2.3"},
		{"github.com/goccy/go-yaml", "1.8.2"},
		{"github.com/gogo/protobuf", "1.3.1"},
		{"github.com/golang/glog", "0.0.0"},
		{"github.com/golang/groupcache", "0.0.0"},
		{"github.com/golang/mock", "1.4.4"},
		{"github.com/golang/protobuf", "1.4.2"},
		{"github.com/google/btree", "1.0.0"},
		{"github.com/google/go-cmp", "0.5.3"},
		{"github.com/google/go-containerregistry", "0.0.0"},
		{"github.com/google/go-github/v28", "28.1.1"},
		{"github.com/google/go-querystring", "1.0.0"},
		{"github.com/google/gofuzz", "1.0.0"},
		{"github.com/google/martian", "2.1.0"},
		{"github.com/google/martian/v3", "3.0.0"},
		{"github.com/google/pprof", "0.0.0"},
		{"github.com/google/renameio", "0.1.0"},
		{"github.com/google/subcommands", "1.0.1"},
		{"github.com/google/uuid", "1.1.1"},
		{"github.com/google/wire", "0.3.0"},
		{"github.com/googleapis/gax-go/v2", "2.0.5"},
		{"github.com/googleapis/gnostic", "0.2.2"},
		{"github.com/gophercloud/gophercloud", "0.1.0"},
		{"github.com/gopherjs/gopherjs", "0.0.0"},
		{"github.com/gorilla/context", "1.1.1"},
		{"github.com/gorilla/mux", "1.7.4"},
		{"github.com/gorilla/websocket", "1.4.0"},
		{"github.com/gregjones/httpcache", "0.0.0"},
		{"github.com/grpc-ecosystem/go-grpc-middleware", "1.0.1"},
		{"github.com/grpc-ecosystem/go-grpc-prometheus", "1.2.0"},
		{"github.com/grpc-ecosystem/grpc-gateway", "1.9.5"},
		{"github.com/hashicorp/errwrap", "1.0.0"},
		{"github.com/hashicorp/go-multierror", "1.1.0"},
		{"github.com/hashicorp/go-version", "1.2.1"},
		{"github.com/hashicorp/golang-lru", "0.5.3"},
		{"github.com/hashicorp/hcl", "1.0.0"},
		{"github.com/hpcloud/tail", "1.0.0"},
		{"github.com/ianlancetaylor/demangle", "0.0.0"},
		{"github.com/imdario/mergo", "0.3.5"},
		{"github.com/inconshreveable/mousetrap", "1.0.0"},
		{"github.com/jbenet/go-context", "0.0.0"},
		{"github.com/jessevdk/go-flags", "1.4.0"},
		{"github.com/jmespath/go-jmespath", "0.0.0"},
		{"github.com/joefitzgerald/rainbow-reporter", "0.1.0"},
		{"github.com/jonboulle/clockwork", "0.1.0"},
		{"github.com/json-iterator/go", "1.1.8"},
		{"github.com/jstemmer/go-junit-report", "0.9.1"},
		{"github.com/jtolds/gls", "4.20.0"},
		{"github.com/julienschmidt/httprouter", "1.2.0"},
		{"github.com/kevinburke/ssh_config", "0.0.0"},
		{"github.com/kisielk/errcheck", "1.2.0"},
		{"github.com/kisielk/gotool", "1.0.0"},
		{"github.com/knqyf263/go-apk-version", "0.0.0"},
		{"github.com/knqyf263/go-deb-version", "0.0.0"},
		{"github.com/knqyf263/go-rpm-version", "0.0.0"},
		{"github.com/knqyf263/go-rpmdb", "0.0.0"},
		{"github.com/knqyf263/nested", "0.0.1"},
		{"github.com/konsorten/go-windows-terminal-sequences", "1.0.2"},
		{"github.com/kr/logfmt", "0.0.0"},
		{"github.com/kr/pretty", "0.1.0"},
		{"github.com/kr/pty", "1.1.5"},
		{"github.com/kr/text", "0.2.0"},
		{"github.com/kylelemons/godebug", "1.1.0"},
		{"github.com/leodido/go-urn", "1.2.0"},
		{"github.com/magiconair/properties", "1.8.0"},
		{"github.com/mailru/easyjson", "0.7.0"},
		{"github.com/mattn/go-colorable", "0.1.8"},
		{"github.com/mattn/go-isatty", "0.0.12"},
		{"github.com/mattn/go-jsonpointer", "0.0.0"},
		{"github.com/mattn/go-runewidth", "0.0.9"},
		{"github.com/matttproud/golang_protobuf_extensions", "1.0.1"},
		{"github.com/maxbrunsfeld/counterfeiter/v6", "6.2.2"},
		{"github.com/mitchellh/go-homedir", "1.1.0"},
		{"github.com/mitchellh/mapstructure", "1.1.2"},
		{"github.com/modern-go/concurrent", "0.0.0"},
		{"github.com/modern-go/reflect2", "1.0.1"},
		{"github.com/morikuni/aec", "1.0.0"},
		{"github.com/munnerz/goautoneg", "0.0.0"},
		{"github.com/mwitkow/go-conntrack", "0.0.0"},
		{"github.com/mxk/go-flowrate", "0.0.0"},
		{"github.com/niemeyer/pretty", "0.0.0"},
		{"github.com/nxadm/tail", "1.4.4"},
		{"github.com/olekukonko/tablewriter", "0.0.2"},
		{"github.com/onsi/ginkgo", "1.14.2"},
		{"github.com/onsi/gomega", "1.10.3"},
		{"github.com/open-policy-agent/opa", "0.21.1"},
		{"github.com/opencontainers/go-digest", "1.0.0"},
		{"github.com/opencontainers/image-spec", "1.0.2"},
		{"github.com/opencontainers/runc", "0.1.1"},
		{"github.com/parnurzeal/gorequest", "0.2.16"},
		{"github.com/pelletier/go-toml", "1.2.0"},
		{"github.com/peterbourgon/diskv", "2.0.1"},
		{"github.com/peterh/liner", "0.0.0"},
		{"github.com/pkg/errors", "0.9.1"},
		{"github.com/pmezard/go-difflib", "1.0.0"},
		{"github.com/pquerna/cachecontrol", "0.0.0"},
		{"github.com/prometheus/client_golang", "1.0.0"},
		{"github.com/prometheus/client_model", "0.0.0"},
		{"github.com/prometheus/common", "0.4.1"},
		{"github.com/prometheus/procfs", "0.0.2"},
		{"github.com/rcrowley/go-metrics", "0.0.0"},
		{"github.com/remyoudompheng/bigfft", "0.0.0"},
		{"github.com/rogpeppe/fastuuid", "0.0.0"},
		{"github.com/rogpeppe/go-charset", "0.0.0"},
		{"github.com/rogpeppe/go-internal", "1.3.0"},
		{"github.com/rubiojr/go-vhd", "0.0.0"},
		{"github.com/russross/blackfriday", "1.5.2"},
		{"github.com/russross/blackfriday/v2", "2.0.1"},
		{"github.com/saracen/walker", "0.0.0"},
		{"github.com/satori/go.uuid", "1.2.0"},
		{"github.com/sclevine/spec", "1.2.0"},
		{"github.com/sergi/go-diff", "1.1.0"},
		{"github.com/shurcooL/sanitized_anchor_name", "1.0.0"},
		{"github.com/simplereach/timeutils", "1.2.0"},
		{"github.com/sirupsen/logrus", "1.5.0"},
		{"github.com/smartystreets/assertions", "1.2.0"},
		{"github.com/smartystreets/goconvey", "1.6.4"},
		{"github.com/soheilhy/cmux", "0.1.4"},
		{"github.com/sosedoff/gitkit", "0.2.0"},
		{"github.com/spf13/afero", "1.2.2"},
		{"github.com/spf13/cast", "1.3.0"},
		{"github.com/spf13/cobra", "0.0.5"},
		{"github.com/spf13/jwalterweatherman", "1.0.0"},
		{"github.com/spf13/pflag", "1.0.5"},
		{"github.com/spf13/viper", "1.3.2"},
		{"github.com/stretchr/objx", "0.3.0"},
		{"github.com/stretchr/testify", "1.6.1"},
		{"github.com/testcontainers/testcontainers-go", "0.3.1"},
		{"github.com/tmc/grpc-websocket-proxy", "0.0.0"},
		{"github.com/twitchtv/twirp", "5.10.1"},
		{"github.com/ugorji/go", "1.1.7"},
		{"github.com/ugorji/go/codec", "1.1.7"},
		{"github.com/urfave/cli", "1.22.5"},
		{"github.com/urfave/cli/v2", "2.3.0"},
		{"github.com/vdemeester/k8s-pkg-credentialprovider", "1.17.4"},
		{"github.com/vmware/govmomi", "0.20.3"},
		{"github.com/xanzy/ssh-agent", "0.2.1"},
		{"github.com/xiang90/probing", "0.0.0"},
		{"github.com/xordataexchange/crypt", "0.0.3"},
		{"github.com/yashtewari/glob-intersection", "0.0.0"},
		{"github.com/yuin/goldmark", "1.1.32"},
		{"github.com/yuin/gopher-lua", "0.0.0"},
		{"go.etcd.io/bbolt", "1.3.5"},
		{"go.etcd.io/etcd", "0.0.0"},
		{"go.opencensus.io", "0.22.4"},
		{"go.opentelemetry.io/otel", "0.14.0"},
		{"go.uber.org/atomic", "1.5.1"},
		{"go.uber.org/multierr", "1.4.0"},
		{"go.uber.org/tools", "0.0.0"},
		{"go.uber.org/zap", "1.13.0"},
		{"golang.org/x/crypto", "0.0.0"},
		{"golang.org/x/exp", "0.0.0"},
		{"golang.org/x/image", "0.0.0"},
		{"golang.org/x/lint", "0.0.0"},
		{"golang.org/x/mobile", "0.0.0"},
		{"golang.org/x/mod", "0.3.0"},
		{"golang.org/x/net", "0.0.0"},
		{"golang.org/x/oauth2", "0.0.0"},
		{"golang.org/x/sync", "0.0.0"},
		{"golang.org/x/sys", "0.0.0"},
		{"golang.org/x/text", "0.3.3"},
		{"golang.org/x/time", "0.0.0"},
		{"golang.org/x/tools", "0.0.0"},
		{"golang.org/x/xerrors", "0.0.0"},
		{"gonum.org/v1/gonum", "0.0.0"},
		{"gonum.org/v1/netlib", "0.0.0"},
		{"google.golang.org/api", "0.30.0"},
		{"google.golang.org/appengine", "1.6.6"},
		{"google.golang.org/genproto", "0.0.0"},
		{"google.golang.org/grpc", "1.31.0"},
		{"google.golang.org/protobuf", "1.25.0"},
		{"gopkg.in/alecthomas/kingpin.v2", "2.2.6"},
		{"gopkg.in/check.v1", "1.0.0"},
		{"gopkg.in/cheggaaa/pb.v1", "1.0.28"},
		{"gopkg.in/errgo.v2", "2.1.0"},
		{"gopkg.in/fsnotify.v1", "1.4.7"},
		{"gopkg.in/gcfg.v1", "1.2.0"},
		{"gopkg.in/go-playground/assert.v1", "1.2.1"},
		{"gopkg.in/go-playground/validator.v9", "9.31.0"},
		{"gopkg.in/inf.v0", "0.9.1"},
		{"gopkg.in/mgo.v2", "2.0.0"},
		{"gopkg.in/natefinch/lumberjack.v2", "2.0.0"},
		{"gopkg.in/resty.v1", "1.12.0"},
		{"gopkg.in/square/go-jose.v2", "2.2.2"},
		{"gopkg.in/tomb.v1", "1.0.0"},
		{"gopkg.in/warnings.v0", "0.1.2"},
		{"gopkg.in/yaml.v2", "2.4.0"},
		{"gopkg.in/yaml.v3", "3.0.0"},
		{"gotest.tools", "2.2.0"},
		{"honnef.co/go/tools", "0.0.1"},
		{"k8s.io/api", "0.17.4"},
		{"k8s.io/apimachinery", "0.17.4"},
		{"k8s.io/apiserver", "0.17.4"},
		{"k8s.io/client-go", "0.17.4"},
		{"k8s.io/cloud-provider", "0.17.4"},
		{"k8s.io/code-generator", "0.17.2"},
		{"k8s.io/component-base", "0.17.4"},
		{"k8s.io/csi-translation-lib", "0.17.4"},
		{"k8s.io/gengo", "0.0.0"},
		{"k8s.io/klog", "1.0.0"},
		{"k8s.io/klog/v2", "2.0.0"},
		{"k8s.io/kube-openapi", "0.0.0"},
		{"k8s.io/legacy-cloud-providers", "0.17.4"},
		{"k8s.io/utils", "0.0.0"},
		{"modernc.org/cc", "1.0.0"},
		{"modernc.org/golex", "1.0.0"},
		{"modernc.org/mathutil", "1.0.0"},
		{"modernc.org/strutil", "1.0.0"},
		{"modernc.org/xc", "1.0.0"},
		{"moul.io/http2curl", "1.0.0"},
		{"rsc.io/binaryregexp", "0.2.0"},
		{"rsc.io/quote/v3", "3.1.0"},
		{"rsc.io/sampler", "1.3.0"},
		{"sigs.k8s.io/structured-merge-diff", "1.0.1"},
		{"sigs.k8s.io/yaml", "1.1.0"},
	}
)