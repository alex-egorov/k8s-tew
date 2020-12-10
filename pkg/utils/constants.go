package utils

// Versions
const VersionConfig = "2.4.0"
const VersionK8s = "v1.20.0"
const VersionKubeAPIServer = "k8s.gcr.io/kube-apiserver:" + VersionK8s
const VersionKubeControllerManager = "k8s.gcr.io/kube-controller-manager:" + VersionK8s
const VersionKubeScheduler = "k8s.gcr.io/kube-scheduler:" + VersionK8s
const VersionKubeProxy = "k8s.gcr.io/kube-proxy:" + VersionK8s
const VersionEtcd = "quay.io/coreos/etcd:v3.4.14"
const VersionContainerd = "1.4.3"
const VersionRunc = "1.0.0-rc92"
const VersionCrictl = "1.19.0"
const VersionHelm = "3.4.2"
const VersionGobetween = "docker.io/yyyar/gobetween:0.8.0"
const VersionVirtualIP = "docker.io/darxkies/virtual-ip:1.0.1"
const VersionBusybox = "docker.io/library/busybox:1.32.0"
const VersionPause = "k8s.gcr.io/pause:3.3"
const VersionCoreDNS = "docker.io/coredns/coredns:1.8.0"
const VersionCalicoPod2Daemon = "quay.io/calico/pod2daemon-flexvol:v3.17.0"
const VersionCalicoTypha = "quay.io/calico/typha:v3.17.0"
const VersionCalicoNode = "quay.io/calico/node:v3.17.0"
const VersionCalicoCni = "quay.io/calico/cni:v3.17.0"
const VersionCalicoKubeControllers = "quay.io/calico/kube-controllers:v3.17.0"
const VersionCeph = "docker.io/ceph/ceph:v15.2.7"
const VersionCsiCephPlugin = "quay.io/cephcsi/cephcsi:v3.1.1"
const VersionCsiAttacher = "quay.io/k8scsi/csi-attacher:v2.1.1"
const VersionCsiProvisioner = "quay.io/k8scsi/csi-provisioner:v1.6.0"
const VersionCsiDriverRegistrar = "quay.io/k8scsi/csi-node-driver-registrar:v1.3.0"
const VersionCsiSnapshotter = "quay.io/k8scsi/csi-snapshotter:v2.1.1"
const VersionCsiSnapshotController = "quay.io/k8scsi/snapshot-controller:v2.0.1"
const VersionCsiResizer = "quay.io/k8scsi/csi-resizer:v0.5.0"
const VersionVelero = "docker.io/velero/velero:v1.5.2"
const VersionVeleroPluginAWS = "docker.io/velero/velero-plugin-for-aws:v1.1.0"
const VersionVeleroPluginCSI = "docker.io/velero/velero-plugin-for-csi:v0.1.2"
const VersionVeleroResticRestoreHelper = "docker.io/velero/velero-restic-restore-helper:v1.5.2"
const VersionMinioServer = "docker.io/minio/minio:RELEASE.2020-10-18T21-54-12Z"
const VersionMinioClient = "docker.io/minio/mc:RELEASE.2020-10-03T02-54-56Z"
const VersionMetalLBController = "docker.io/metallb/controller:v0.9.5"
const VersionMetalLBSpeaker = "docker.io/metallb/speaker:v0.9.5"
const VersionElasticsearch = "docker.elastic.co/elasticsearch/elasticsearch:7.9.2"
const VersionKibana = "docker.elastic.co/kibana/kibana:7.9.2"
const VersionCerebro = "docker.io/lmenezes/cerebro:0.9.2"
const VersionFluentBit = "docker.io/fluent/fluent-bit:1.6.1"
const VersionMetricsScraper = "docker.io/kubernetesui/metrics-scraper:v1.0.6"
const VersionKubernetesDashboard = "docker.io/kubernetesui/dashboard:v2.0.5"
const VersionCertManagerController = "quay.io/jetstack/cert-manager-controller:v1.0.3"
const VersionCertManagerCAInjector = "quay.io/jetstack/cert-manager-cainjector:v1.0.3"
const VersionCertManagerWebHook = "quay.io/jetstack/cert-manager-webhook:v1.0.3"
const VersionNginxIngressAdmissionWebhook = "docker.io/jettech/kube-webhook-certgen:v1.3.0"
const VersionNginxIngressController = "k8s.gcr.io/ingress-nginx/controller:v0.40.2"
const VersionMetricsServer = "k8s.gcr.io/metrics-server/metrics-server:v0.3.7"
const VersionKubeStateMetrics = "quay.io/coreos/kube-state-metrics:v1.9.7"
const VersionGrafana = "docker.io/grafana/grafana:7.2.2"
const VersionPrometheus = "quay.io/prometheus/prometheus:v2.22.0"
const VersionNodeExporter = "quay.io/prometheus/node-exporter:v1.0.1"
const VersionAlertManager = "quay.io/prometheus/alertmanager:v0.21.0"
const VersionMysql = "docker.io/library/mysql:8.0.22"
const VersionWordpress = "docker.io/library/wordpress:5.5.3-apache"

// Settings
const ProjectTitle = "Kubernetes - The Easier Way"
const Username = "k8s-tew"
const ClusterName = "k8s-tew"
const KeyUsername = "username"
const KeyPassword = "password"
const KeySecret = "secret"
const MaxPods = 110
const RsaSize = 2048
const CaValidityPeriod = 20
const ClientValidityPeriod = 15
const GrafanaSize = 2
const PrometheusSize = 2
const MinioSize = 2
const ElasticsearchCount = 1
const ElasticsearchSize = 10
const AlertManagerCount = 1
const AlertManagerSize = 2
const KubeStateMetricsCount = 1
const BaseDirectory = "assets"
const ClusterDomain = "cluster.local"
const ClusterIpRange = "10.32.0.0/24"
const CalicoTyphaIp = "10.32.0.5"
const ClusterDnsIp = "10.32.0.10"
const ClusterCidr = "10.200.0.0/16"
const CephClusterName = "ceph"
const CephPlacementGroups = 256
const CephExpectedNumberOfObjects = 1000000
const MetalLBAddresses = "192.168.0.16/28"
const ResolvConf = "/etc/resolv.conf"
const PublicNetwork = "192.168.100.0/24"
const HelmServiceAccount = "tiller"
const Email = "k8s-tew@gmail.com"
const DeploymentDirectory = "/"
const IngressDomain = "k8s-tew.net"
const IngressSubdomainWordpress = "wordpress"
const AdminUserName = "admin-user"
const AdminUserNamespace = "kube-system"
const KubernetesDashboardNamespace = "kube-system"

// Ports
const PortVipRaftController uint16 = 16277
const PortVipRaftWorker uint16 = 16728
const PortLoadBalancer uint16 = 16443
const PortKubernetesDashboard uint16 = 32443
const PortApiServer uint16 = 6443
const PortCephManager uint16 = 30700
const PortCephRadosGateway uint16 = 30750
const PortMinio uint16 = 30800
const PortGrafana uint16 = 30900
const PortKibana uint16 = 30980
const PortCerebro uint16 = 30990
const PortWordpress uint16 = 30100

// URLs
const K8sBaseName = "kubernetes-node-linux-amd64"
const K8sDownloadUrl = "https://storage.googleapis.com/kubernetes-release/release/{{.Versions.K8S}}/{{.Filename}}.tar.gz"
const EtcdBaseName = "etcd-{{.Versions.Etcd | image_tag}}-linux-amd64"
const EtcdDownloadUrl = "https://github.com/coreos/etcd/releases/download/{{.Versions.Etcd | image_tag}}/{{.Filename}}.tar.gz"
const CniBaseName = "cni-plugins-amd64-v{{.Versions.CNI}}"
const CniDownloadUrl = "https://github.com/containernetworking/plugins/releases/download/v{{.Versions.CNI}}/{{.Filename}}.tgz"
const ContainerdBaseName = "containerd-{{.Versions.Containerd}}-linux-amd64"
const ContainerdDownloadUrl = "https://github.com/containerd/containerd/releases/download/v{{.Versions.Containerd}}/{{.Filename}}.tar.gz"
const RuncDownloadUrl = "https://github.com/opencontainers/runc/releases/download/v{{.Versions.Runc}}/runc.amd64"
const CrictlBaseName = "crictl-v{{.Versions.CriCtl}}-linux-amd64"
const CrictlDownloadUrl = "https://github.com/kubernetes-incubator/cri-tools/releases/download/v{{.Versions.CriCtl}}/{{.Filename}}.tar.gz"
const GobetweenBaseName = "gobetween_{{.Versions.Gobetween}}_linux_amd64"
const GobetweenDownloadUrl = "https://github.com/yyyar/gobetween/releases/download/{{.Versions.Gobetween}}/{{.Filename}}.tar.gz"
const HelmBaseName = "helm-v{{.Versions.Helm}}-linux-amd64"
const HelmDownloadUrl = "https://get.helm.sh/{{.Filename}}.tar.gz"
const VeleroBaseName = "velero-{{.Versions.Velero | image_tag}}-linux-amd64"
const VeleroDownloadUrl = "https://github.com/heptio/velero/releases/download/{{.Versions.Velero | image_tag}}/{{.Filename}}.tar.gz"

// Config
const ConfigFilename = "config.yaml"

// Node Labels
const NodeBootstrapper = "bootstrapper"
const NodeController = "controller"
const NodeWorker = "worker"
const NodeStorage = "storage"

// Features
const FeatureStorage = "storage"
const FeatureMonitoring = "monitoring"
const FeatureLogging = "logging"
const FeatureBackup = "backup"
const FeatureShowcase = "showcase"
const FeatureIngress = "ingress"

// OS
const OsUbuntu = "ubuntu"
const OsUbuntu1804 = "ubuntu/18.04"
const OsCentos = "centos"
const OsCentos75 = "centos/7.5"

// Sub-Directories
const SubdirectoryTemporary = "tmp"
const SubdirectoryConfig = "etc"
const SubdirectorySystemd = "systemd"
const SubdirectorySystem = "system"
const SubdirectoryK8sTew = "k8s-tew"
const SubdirectoryCertificates = "ssl"
const SubdirectoryOptional = "opt"
const SubdirectoryVariable = "var"
const SubdirectoryLogging = "log"
const SubdirectoryLibrary = "lib"
const SubdirectoryRun = "run"
const SubdirectoryBinary = "bin"
const SubdirectoryK8s = "k8s"
const SubdirectoryEtcd = "etcd"
const SubdirectoryCri = "cri"
const SubdirectoryCni = "cni"
const SubdirectoryKubeconfig = "kubeconfig"
const SubdirectorySecurity = "security"
const SubdirectorySetup = "setup"
const SubdirectoryContainerd = "containerd"
const SubdirectoryImages = "images"
const SubdirectoryProfileD = "profile.d"
const SubdirectoryLoadBalancer = "lb"
const SubdirectoryHelm = "helm"
const SubdirectoryKubelet = "kubelet"
const SubdirectoryPods = "pods"
const SubdirectoryPluginsRegistry = "plugins_registry"
const SubdirectoryManifests = "manifests"
const SubdirectoryCeph = "ceph"
const SubdirectoryCephBootstrapMds = "bootstrap-mds"
const SubdirectoryCephBootstrapOsd = "bootstrap-osd"
const SubdirectoryCephBootstrapRbd = "bootstrap-rbd"
const SubdirectoryCephBootstrapRgw = "bootstrap-rgw"
const SubdirectoryVelero = "velero"
const SubdirectoryHost = "host"
const SubdirectoryPlugins = "plugins"
const SubdirectoryCsiCephfsPlugin = "csi-cephfsplugin"
const SubdirectoryCsiRbdPlugin = "csi-rbdplugin"

// Directories
const DirectoryConfig = "config"
const DirectoryCertificates = "certificates"
const DirectoryCniConfig = "cni-config"
const DirectoryCriConfig = "cri-config"
const DirectoryK8sSecurityConfig = "security-config"
const DirectoryK8sConfig = "k8s-config"
const DirectoryK8sKubeConfig = "kube-config"
const DirectoryK8sSetupConfig = "setup-config"
const DirectoryBinaries = "binaries"
const DirectoryK8sBinaries = "k8s-binaries"
const DirectoryEtcdBinaries = "etcd-binaries"
const DirectoryCniBinaries = "cni-binaries"
const DirectoryCriBinaries = "cri-binaries"
const DirectoryDynamicData = "dynamic-data"
const DirectoryEtcdData = "etcd-data"
const DirectoryContainerdData = "containerd-data"
const DirectoryImages = "images"
const DirectoryLogging = "logging"
const DirectoryService = "service"
const DirectoryContainerdState = "containerd-state"
const DirectoryAbsoluteContainerdState = "absolute-containerd-state"
const DirectoryProfile = "profile"
const DirectoryGobetweenBinaries = "gobetween-binaries"
const DirectoryGobetweenConfig = "gobetween-config"
const DirectoryHelmData = "helm-data"
const DirectoryKubeletData = "kubelet-data"
const DirectoryPodsData = "pods-data"
const DirectoryTemporary = "temporary"
const DirectoryK8sManifests = "kubelet-manifests"
const DirectoryCeph = "ceph"
const DirectoryCephConfig = "ceph-config"
const DirectoryCephData = "ceph-data"
const DirectoryCephBootstrapMds = "bootstrap-mds"
const DirectoryCephBootstrapOsd = "bootstrap-osd"
const DirectoryCephBootstrapRbd = "bootstrap-rbd"
const DirectoryCephBootstrapRgw = "bootstrap-rgw"
const DirectoryVeleroBinaries = "velero"
const DirectoryHostBinaries = "host-binaries"
const DirectoryKubeletPlugins = "kubelet-plugins"
const DirectoryKubeletPluginsRegistry = "kubelet-plugins-registry"
const DirectoryVarRun = "var-run"
const DirectoryRun = "run"

// Binaries
const BinaryK8sTew = "k8s-tew"
const BinaryHelm = "helm"
const BinaryContainerd = "containerd"
const BinaryContainerdShimRuncV2 = "containerd-shim-runc-v2"
const BinaryCtr = "ctr"
const BinaryRunc = "runc"
const BinaryCrictl = "crictl"
const BinaryEtcd = "etcd"
const BinaryEtcdctl = "etcdctl"
const BinaryKubectl = "kubectl"
const BinaryKubelet = "kubelet"
const BinaryGobetween = "gobetween"
const BinaryVelero = "velero"

// Certificates
const PemCa = "ca.pem"
const PemCaKey = "ca-key.pem"
const PemKubernetes = "kubernetes.pem"
const PemKubernetesKey = "kubernetes-key.pem"
const PemAdmin = "admin.pem"
const PemAdminKey = "admin-key.pem"
const PemProxy = "proxy.pem"
const PemProxyKey = "proxy-key.pem"
const PemControllerManager = "controller-manager.pem"
const PemControllerManagerKey = "controller-manager-key.pem"
const PemScheduler = "scheduler.pem"
const PemSchedulerKey = "scheduler-key.pem"
const PemKubelet = "kubelet-{{.Name}}.pem"
const PemKubeletKey = "kubelet-{{.Name}}-key.pem"
const PemServiceAccount = "service-account.pem"
const PemServiceAccountKey = "service-account-key.pem"
const PemVirtualIp = "virtual-ip.pem"
const PemVirtualIpKey = "virtual-ip-key.pem"
const PemAggregator = "aggregator.pem"
const PemAggregatorKey = "aggregator-key.pem"
const PemElasticsearch = "elasticsearch.pem"
const PemElasticsearchKey = "elasticsearch-key.pem"
const PemMinio = "minio.pem"
const PemMinioKey = "minio-key.pem"
const PemGrafana = "grafana.pem"
const PemGrafanaKey = "grafana-key.pem"
const PemCeph = "ceph.pem"
const PemCephKey = "ceph-key.pem"
const PemPrometheus = "prometheus.pem"
const PemPrometheusKey = "prometheus-key.pem"
const PemKubernetesDashboard = "kubernetes-dashboard.pem"
const PemKubernetesDashboardKey = "kubernetes-dashboard-key.pem"

// Kubeconfig
const KubeconfigAdmin = "admin.kubeconfig"
const KubeconfigControllerManager = "controller-manager-{{.Name}}.kubeconfig"
const KubeconfigScheduler = "scheduler-{{.Name}}.kubeconfig"
const KubeconfigProxy = "proxy-{{.Name}}.kubeconfig"
const KubeconfigKubelet = "kubelet-{{.Name}}.kubeconfig"

// Manifests
const ManifestEtcd = "etcd-{{.Name}}.yaml"
const ManifestKubeApiserver = "kube-apiserver-{{.Name}}.yaml"
const ManifestKubeControllerManager = "kube-controller-manager-{{.Name}}.yaml"
const ManifestKubeScheduler = "kube-scheduler-{{.Name}}.yaml"
const ManifestKubeProxy = "kube-proxy-{{.Name}}.yaml"
const ManifestGobetween = "gobetween-{{.Name}}.yaml"
const ManifestControllerVirtualIP = "controller-virtual-ip-{{.Name}}.yaml"
const ManifestWorkerVirtualIP = "worker-virtual-ip-{{.Name}}.yaml"

// Security
const EncryptionConfig = "encryption-config.yaml"

// Containerd
const ContainerdConfig = "config-{{.Name}}.toml"
const ContainerdSock = "containerd.sock"
const ContainerdKubernetesNamespace = "k8s.io"
const ContainerdServerName = "containerd"

// K8S Config
const K8sKubeletSetup = "kubelet-setup.yaml"
const K8sAdminUserSetup = "admin-user-setup.yaml"
const K8sHelmUserSetup = "helm-user-setup.yaml"
const K8sKubeProxyConfig = "kube-proxy-{{.Name}}-config.yaml"
const K8sKubeSchedulerConfig = "kube-scheduler-{{.Name}}-config.yaml"
const K8sKubeletConfig = "kubelet-{{.Name}}-config.yaml"
const K8sCorednsSetup = "coredns-setup.yaml"
const K8sCalicoSetup = "calico-setup.yaml"
const K8sMetalLBSetup = "metallb-setup.yaml"
const K8sEfkSetup = "efk-setup.yaml"
const K8sVeleroSetup = "velero-setup.yaml"
const K8sKubernetesDashboardSetup = "kubernetes-dashboard-setup.yaml"
const K8sKubernetesDashboardCertificates = "kubernetes-dashboard-certificates.yaml"
const K8sHelmSetup = "helm-setup.yaml"
const K8sCertManagerSetup = "cert-manager-setup.yaml"
const K8sNginxIngressSetup = "nginx-ingress-setup.yaml"
const K8sMetricsServerSetup = "metrics-server-setup.yaml"
const K8sPrometheusSetup = "prometheus-setup.yaml"
const K8sPrometheusRules = "prometheus-rules.yaml"
const K8sPrometheusAlerts = "prometheus-alerts.yaml"
const K8sPrometheusCertificates = "prometheus-certificates.yaml"
const K8sKubeStateMetricsSetup = "kube-state-metrics-setup.yaml"
const K8sNodeExporterSetup = "node-exporter-setup.yaml"
const K8sGrafanaSetup = "grafana.yaml"
const K8sGrafanaDashboards = "grafana-dashboards.yaml"
const K8sAlertManagerSetup = "alert-manager.yaml"
const K8sGrafanaCredentials = "grafana-credentials.yaml"
const K8sGrafanaCertificates = "grafana-certificates.yaml"
const K8sMinioCredentials = "minio-credentials.yaml"
const K8sMinioCertificates = "minio-certificates.yaml"
const K8sCerebroCredentials = "cerebro-credentials.yaml"
const K8sCephManagerCredentials = "ceph-manager-credentials.yaml"
const K8sCephCertificates = "ceph-certificates.yaml"
const K8sCephRadosGatewayCredentials = "ceph-rados-gateway-credentials.yaml"
const K8sElasticsearchCredentials = "elasticsearch-credentials.yaml"
const K8sElasticsearchCertificates = "elasticsearch-certificates.yaml"
const WordpressSetup = "wordpress-setup.yaml"

// Gobetween Config
const GobetweenConfig = "config.toml"

// Profile
const K8sTewProfile = "k8s-tew.sh"

// Logging
const AuditLog = "audit.log"

// Deployment
const DeploymentUser = "root"

// Service
const ServiceName = "k8s-tew"
const ServiceConfig = ServiceName + ".service"

// Ceph
const CephRbdPoolName = "cephrbd"
const CephFsPoolName = "cephfs"
const CephConfig = "ceph.conf"
const CephClientAdminKeyring = "ceph.client.admin.keyring"
const CephMonitorKeyring = "ceph.mon.keyring"
const CephKeyring = "ceph.keyring"
const CephBootstrapMdsKeyring = "ceph.bootstrap.mds.keyring"
const CephBootstrapOsdKeyring = "ceph.bootstrap.osd.keyring"
const CephBootstrapRbdKeyring = "ceph.bootstrap.rbd.keyring"
const CephBootstrapRgwKeyring = "ceph.bootstrap.rgw.keyring"
const CephSecrets = "ceph-secrets.yaml"
const CephSetup = "ceph-setup.yaml"
const CephCsi = "ceph-csi.yaml"

// Cluster Issuer
const LetsencryptClusterIssuer = "letsencrypt-cluster-issuer.yaml"

// Environment variables
const K8sTewBaseDirectory = "K8S_TEW_BASE_DIRECTORY"

// Virtual IP Manager
const ElectionNamespace = "/k8s-tew"
const ElectionController = "/controller-vip-manager"
const ElectionWorker = "/worker-vip-manager"

// Common Names
const CnAdmin = "admin"
const CnAggregator = "aggregator"
const CnSystemKubeControllerManager = "system:kube-controller-manager"
const CnSystemKubeScheduler = "system:kube-scheduler"
const CnSystemKubeProxy = "system:kube-proxy"
const CnSystemNodePrefix = "system:node:%s"
const CnElasticsearch = "elasticsearch"
const CnMinio = "minio"
const CnGrafana = "grafana"
const CnCeph = "ceph"
const CnPrometheus = "prometheus"
const CnKubernetesDashboard = "kubernetes-dashboard"

// Templates
const TemplateContainerdToml = "k8s/cri/containerd.toml"
const TemplateK8sTewService = "system/k8s-tew.service"
const TemplateK8sTewProfile = "system/k8s-tew.sh"
const TemplateEnvironment = "system/environment.sh"
const TemplateGobetweenToml = "k8s/lb/gobetween.toml"
const TemplateKubeProxyConfiguration = "k8s/kube-proxy-configuration.yaml"
const TemplateKubeSchedulerConfiguration = "k8s/kube-scheduler-configuration.yaml"
const TemplateKubeletConfiguration = "k8s/kubelet-configuration.yaml"
const TemplateEncryptionConfig = "k8s/encryption-config.yaml"
const TemplateKubeconfig = "k8s/kubeconfig.yaml"
const TemplateCredentials = "k8s/credentials.yaml"
const TemplateConfigMap = "k8s/config-map.yaml"
const TemplateServiceAccount = "k8s/service-account.yaml"
const TemplateKubeletSetup = "k8s/setup/kubelet-setup.yaml"
const TemplateCephClientKeyring = "ceph/client.keyring"
const TemplateCephClientAdminKeyring = "ceph/client-admin.keyring"
const TemplateCephMonitorKeyring = "ceph/monitor.keyring"
const TemplateCephConfig = "ceph/ceph.conf"
const TemplateCephSecrets = "k8s/setup/storage/ceph-secrets.yaml"
const TemplateCephSetup = "k8s/setup/storage/ceph-setup.yaml"
const TemplateCephCsi = "k8s/setup/storage/ceph-csi.yaml"
const TemplateLetsencryptClusterIssuerSetup = "k8s/setup/ingress/letsencrypt-cluster-issuer.yaml"
const TemplateCorednsSetup = "k8s/setup/dns/coredns.yaml"
const TemplateCalicoSetup = "k8s/setup/networking/calico.yaml"
const TemplateMetalLBSetup = "k8s/setup/networking/metallb.yaml"
const TemplateEfkSetup = "k8s/setup/logging/efk.yaml"
const TemplateVeleroSetup = "k8s/setup/backup/velero.yaml"
const TemplateKubernetesDashboardSetup = "k8s/setup/management/kubernetes-dashboard.yaml"
const TemplateHelmSetup = "k8s/setup/management/helm.yaml"
const TemplateCertManagerSetup = "k8s/setup/networking/cert-manager.yaml"
const TemplateNginxIngressSetup = "k8s/setup/networking/nginx-ingress.yaml"
const TemplateMetricsServerSetup = "k8s/setup/monitoring/metrics-server.yaml"
const TemplatePrometheusSetup = "k8s/setup/monitoring/prometheus.yaml"
const TemplatePrometheusRules = "k8s/setup/monitoring/prometheus-rules.yaml"
const TemplatePrometheusAlerts = "k8s/setup/monitoring/prometheus-alerts.yaml"
const TemplateKubeStateMetricsSetup = "k8s/setup/monitoring/kube-state-metrics.yaml"
const TemplateNodeExporterSetup = "k8s/setup/monitoring/node-exporter.yaml"
const TemplateGrafanaSetup = "k8s/setup/monitoring/grafana.yaml"
const TemplateGrafanaDashboards = "k8s/setup/monitoring/grafana-dashboards.yaml"
const TemplateAlertManagerSetup = "k8s/setup/monitoring/alert-manager.yaml"
const TemplateWordpressSetup = "k8s/setup/miscellaneous/wordpress.yaml"
const TemplateManifestEtcd = "k8s/manifests/etcd.yaml"
const TemplateManifestKubeApiserver = "k8s/manifests/kube-apiserver.yaml"
const TemplateManifestKubeControllerManager = "k8s/manifests/kube-controller-manager.yaml"
const TemplateManifestKubeScheduler = "k8s/manifests/kube-scheduler.yaml"
const TemplateManifestKubeProxy = "k8s/manifests/kube-proxy.yaml"
const TemplateManifestGobetween = "k8s/manifests/gobetween.yaml"
const TemplateManifestVirtualIP = "k8s/manifests/virtual-ip.yaml"

const NodeRoleController = "node-role.kubernetes.io/master"
const NodeRoleWorker = "node-role.kubernetes.io/worker"
const NodeRoleStorage = "node-role.kubernetes.io/storage"
const NodeNotReady = "node.kubernetes.io/not-ready"

const ConcurrentSshConnectionsLimit = 10

const GrafanaCredentials = "grafana-credentials"
const MinioCredentials = "minio-credentials"
const CerebroCredentials = "cerebro-credentials"
const CephManagerCredentials = "ceph-manager-credentials"
const CephRadosGatewayCredentials = "ceph-rados-gateway-credentials"
const ElasticsearchCredentials = "elasticsearch-credentials"

const ElasticsearchCertificates = "elasticsearch-certificates"
const MinioCertificates = "minio-certificates"
const GrafanaCertificates = "grafana-certificates"
const CephCertificates = "ceph-certificates"
const PrometheusCertificates = "prometheus-certificates"
const KubernetesDashboardCertificates = "kubernetes-dashboard-certificates"
