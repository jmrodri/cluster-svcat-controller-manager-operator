// Package v311_00_assets Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// bindata/v3.11.0/openshift-svcat-controller-manager/cm.yaml
// bindata/v3.11.0/openshift-svcat-controller-manager/cr-catalog-controller.yaml
// bindata/v3.11.0/openshift-svcat-controller-manager/crb-catalog-controller.yaml
// bindata/v3.11.0/openshift-svcat-controller-manager/crb-controller-namespace-viewer-binding.yaml
// bindata/v3.11.0/openshift-svcat-controller-manager/defaultconfig.yaml
// bindata/v3.11.0/openshift-svcat-controller-manager/ds.yaml
// bindata/v3.11.0/openshift-svcat-controller-manager/ns.yaml
// bindata/v3.11.0/openshift-svcat-controller-manager/role-cluster-info-configmap.yaml
// bindata/v3.11.0/openshift-svcat-controller-manager/role-configmap-accessor.yaml
// bindata/v3.11.0/openshift-svcat-controller-manager/rolebinding-cluster-info-configmap.yaml
// bindata/v3.11.0/openshift-svcat-controller-manager/rolebinding-configmap-accessor.yaml
// bindata/v3.11.0/openshift-svcat-controller-manager/sa.yaml
// bindata/v3.11.0/openshift-svcat-controller-manager/servicemonitor-role.yaml
// bindata/v3.11.0/openshift-svcat-controller-manager/servicemonitor-rolebinding.yaml
// bindata/v3.11.0/openshift-svcat-controller-manager/servicemonitor.yaml
// bindata/v3.11.0/openshift-svcat-controller-manager/svc.yaml
package v311_00_assets

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _v3110OpenshiftSvcatControllerManagerCmYaml = []byte(`apiVersion: v1
kind: ConfigMap
metadata:
  namespace: openshift-service-catalog-controller-manager
  name: config
data:
  config.yaml:
`)

func v3110OpenshiftSvcatControllerManagerCmYamlBytes() ([]byte, error) {
	return _v3110OpenshiftSvcatControllerManagerCmYaml, nil
}

func v3110OpenshiftSvcatControllerManagerCmYaml() (*asset, error) {
	bytes, err := v3110OpenshiftSvcatControllerManagerCmYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/openshift-svcat-controller-manager/cm.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftSvcatControllerManagerCrCatalogControllerYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: service-catalog-controller
rules:
- apiGroups:
  - ""
  resources:
  - secrets
  verbs:
  - create
  - update
  - patch
  - delete
  - get
  - list
  - watch
- apiGroups:
  - servicecatalog.k8s.io
  resources:
  - clusterservicebrokers/status
  - clusterserviceclasses/status
  - clusterserviceplans/status
  - serviceinstances/status
  - servicebindings/status
  - servicebindings/finalizers
  - serviceinstances/reference
  - servicebrokers/status
  - serviceclasses/status
  - serviceplans/status
  verbs:
  - update
- apiGroups:
  - servicecatalog.k8s.io
  resources:
  - clusterservicebrokers
  - serviceinstances
  - servicebindings
  - servicebrokers
  verbs:
  - list
  - get
  - watch
- apiGroups:
  - ""
  resources:
  - events
  verbs:
  - patch
  - create
- apiGroups:
  - servicecatalog.k8s.io
  resources:
  - clusterserviceclasses
  - clusterserviceplans
  - serviceclasses
  - serviceplans
  verbs:
  - create
  - delete
  - update
  - patch
  - get
  - list
  - watch
- apiGroups:
  - settings.k8s.io
  resources:
  - podpresets
  verbs:
  - create
  - update
  - delete
  - get
  - list
  - watch
`)

func v3110OpenshiftSvcatControllerManagerCrCatalogControllerYamlBytes() ([]byte, error) {
	return _v3110OpenshiftSvcatControllerManagerCrCatalogControllerYaml, nil
}

func v3110OpenshiftSvcatControllerManagerCrCatalogControllerYaml() (*asset, error) {
	bytes, err := v3110OpenshiftSvcatControllerManagerCrCatalogControllerYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/openshift-svcat-controller-manager/cr-catalog-controller.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftSvcatControllerManagerCrbCatalogControllerYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: service-catalog-controller-binding
roleRef:
  kind: ClusterRole
  name: service-catalog-controller
subjects:
- kind: ServiceAccount
  name: service-catalog-controller
  namespace: openshift-service-catalog-controller-manager
`)

func v3110OpenshiftSvcatControllerManagerCrbCatalogControllerYamlBytes() ([]byte, error) {
	return _v3110OpenshiftSvcatControllerManagerCrbCatalogControllerYaml, nil
}

func v3110OpenshiftSvcatControllerManagerCrbCatalogControllerYaml() (*asset, error) {
	bytes, err := v3110OpenshiftSvcatControllerManagerCrbCatalogControllerYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/openshift-svcat-controller-manager/crb-catalog-controller.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftSvcatControllerManagerCrbControllerNamespaceViewerBindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: service-catalog-controller-namespace-viewer-binding
roleRef:
  kind: ClusterRole
  name: namespace-viewer
subjects:
- kind: ServiceAccount
  name: service-catalog-controller
  namespace: openshift-service-catalog-controller-manager
`)

func v3110OpenshiftSvcatControllerManagerCrbControllerNamespaceViewerBindingYamlBytes() ([]byte, error) {
	return _v3110OpenshiftSvcatControllerManagerCrbControllerNamespaceViewerBindingYaml, nil
}

func v3110OpenshiftSvcatControllerManagerCrbControllerNamespaceViewerBindingYaml() (*asset, error) {
	bytes, err := v3110OpenshiftSvcatControllerManagerCrbControllerNamespaceViewerBindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/openshift-svcat-controller-manager/crb-controller-namespace-viewer-binding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftSvcatControllerManagerDefaultconfigYaml = []byte(`apiVersion: openshiftcontrolplane.config.openshift.io/v1
kind: ServiceCatalogControllerManagerConfig
`)

func v3110OpenshiftSvcatControllerManagerDefaultconfigYamlBytes() ([]byte, error) {
	return _v3110OpenshiftSvcatControllerManagerDefaultconfigYaml, nil
}

func v3110OpenshiftSvcatControllerManagerDefaultconfigYaml() (*asset, error) {
	bytes, err := v3110OpenshiftSvcatControllerManagerDefaultconfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/openshift-svcat-controller-manager/defaultconfig.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftSvcatControllerManagerDsYaml = []byte(`apiVersion: apps/v1
kind: DaemonSet
metadata:
  namespace: openshift-service-catalog-controller-manager
  name: controller-manager
  labels:
    app: svcat-controller-manager
    controller-manager: "true"
spec:
  updateStrategy:
    type: RollingUpdate
  selector:
    matchLabels:
      app: svcat-controller-manager
      controller-manager: "true"
  template:
    metadata:
      name: svcat-controller-manager
      labels:
        app: svcat-controller-manager
        controller-manager: "true"
    spec:
      serviceAccountName: service-catalog-controller
      containers:
      - name: controller-manager
        terminationMessagePolicy: FallbackToLogsOnError
        image: ${IMAGE}
        imagePullPolicy: IfNotPresent
        command: ["/usr/bin/service-catalog"]
        args:
        - controller-manager
        - --secure-port
        - "6443"
        - --leader-election-namespace
        - openshift-service-catalog-controller-manager
        - --leader-elect-resource-lock
        - configmaps
        - --cluster-id-configmap-namespace=openshift-service-catalog-controller-manager
        - --broker-relist-interval
        - "5m"
        - --feature-gates
        - OriginatingIdentity=true
        - --feature-gates
        - AsyncBindingOperations=true
        - --feature-gates
        - NamespacedServiceBroker=true
        resources:
          requests:
            memory: 100Mi
            cpu: 100m
        ports:
        - containerPort: 6443
        volumeMounts:
        - mountPath: /var/run/kubernetes-service-catalog
          name: apiserver-ssl
          readOnly: true
        - mountPath: /var/run/configmaps/config
          name: config
        - mountPath: /var/run/configmaps/client-ca
          name: client-ca
        - mountPath: /var/run/secrets/serving-cert
          name: serving-cert
        readinessProbe:
          httpGet:
            port: 6443
            path: /healthz/ready
            scheme: HTTPS
          failureThreshold: 3
          initialDelaySeconds: 20
          periodSeconds: 10
          successThreshold: 1
          timeoutSeconds: 10
        livenessProbe:
          httpGet:
            port: 6443
            path: /healthz
            scheme: HTTPS
          failureThreshold: 3
          initialDelaySeconds: 20
          periodSeconds: 10
          successThreshold: 1
          timeoutSeconds: 10
      volumes:
      - name: apiserver-ssl
        secret:
          defaultMode: 420
          secretName: serving-cert
          items:
          - key: tls.crt
            path: apiserver.crt
          - key: tls.key
            path: apiserver.key
      - name: config
        configMap:
          name: config
      - name: client-ca
        configMap:
          name: client-ca
      - name: serving-cert
        secret:
          secretName: serving-cert
      nodeSelector:
        node-role.kubernetes.io/master: ""
      priorityClassName: "system-cluster-critical"
      tolerations:
      - operator: Exists
`)

func v3110OpenshiftSvcatControllerManagerDsYamlBytes() ([]byte, error) {
	return _v3110OpenshiftSvcatControllerManagerDsYaml, nil
}

func v3110OpenshiftSvcatControllerManagerDsYaml() (*asset, error) {
	bytes, err := v3110OpenshiftSvcatControllerManagerDsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/openshift-svcat-controller-manager/ds.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftSvcatControllerManagerNsYaml = []byte(`apiVersion: v1
kind: Namespace
metadata:
  name: openshift-service-catalog-controller-manager
  labels:
    openshift.io/run-level: "1"
    openshift.io/cluster-monitoring: "true"
  annotations:
    openshift.io/node-selector: ""
`)

func v3110OpenshiftSvcatControllerManagerNsYamlBytes() ([]byte, error) {
	return _v3110OpenshiftSvcatControllerManagerNsYaml, nil
}

func v3110OpenshiftSvcatControllerManagerNsYaml() (*asset, error) {
	bytes, err := v3110OpenshiftSvcatControllerManagerNsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/openshift-svcat-controller-manager/ns.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftSvcatControllerManagerRoleClusterInfoConfigmapYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: cluster-info-configmap
  namespace: openshift-service-catalog-controller-manager
rules:
- apiGroups:     [""]
  resources:     ["configmaps"]
  resourceNames: ["cluster-info"]
  verbs:         ["get","create","list","watch","update"]
`)

func v3110OpenshiftSvcatControllerManagerRoleClusterInfoConfigmapYamlBytes() ([]byte, error) {
	return _v3110OpenshiftSvcatControllerManagerRoleClusterInfoConfigmapYaml, nil
}

func v3110OpenshiftSvcatControllerManagerRoleClusterInfoConfigmapYaml() (*asset, error) {
	bytes, err := v3110OpenshiftSvcatControllerManagerRoleClusterInfoConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/openshift-svcat-controller-manager/role-cluster-info-configmap.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftSvcatControllerManagerRoleConfigmapAccessorYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: configmap-accessor
  namespace: openshift-service-catalog-controller-manager
rules:
- apiGroups:
  - ""
  resources:
  - configmaps
  verbs:
  - list
  - watch
  - get
  - create
  - update
`)

func v3110OpenshiftSvcatControllerManagerRoleConfigmapAccessorYamlBytes() ([]byte, error) {
	return _v3110OpenshiftSvcatControllerManagerRoleConfigmapAccessorYaml, nil
}

func v3110OpenshiftSvcatControllerManagerRoleConfigmapAccessorYaml() (*asset, error) {
	bytes, err := v3110OpenshiftSvcatControllerManagerRoleConfigmapAccessorYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/openshift-svcat-controller-manager/role-configmap-accessor.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftSvcatControllerManagerRolebindingClusterInfoConfigmapYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: cluster-info-configmap-binding
  namespace: openshift-service-catalog-controller-manager
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: cluster-info-configmap
subjects:
- kind: ServiceAccount
  namespace: openshift-service-catalog-controller-manager
  name: service-catalog-controller
`)

func v3110OpenshiftSvcatControllerManagerRolebindingClusterInfoConfigmapYamlBytes() ([]byte, error) {
	return _v3110OpenshiftSvcatControllerManagerRolebindingClusterInfoConfigmapYaml, nil
}

func v3110OpenshiftSvcatControllerManagerRolebindingClusterInfoConfigmapYaml() (*asset, error) {
	bytes, err := v3110OpenshiftSvcatControllerManagerRolebindingClusterInfoConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/openshift-svcat-controller-manager/rolebinding-cluster-info-configmap.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftSvcatControllerManagerRolebindingConfigmapAccessorYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: configmap-accessor-binding
  namespace: openshift-service-catalog-controller-manager
roleRef:
  kind: Role
  name: configmap-accessor
subjects:
- kind: ServiceAccount
  namespace: openshift-service-catalog-controller-manager
  name: service-catalog-controller
`)

func v3110OpenshiftSvcatControllerManagerRolebindingConfigmapAccessorYamlBytes() ([]byte, error) {
	return _v3110OpenshiftSvcatControllerManagerRolebindingConfigmapAccessorYaml, nil
}

func v3110OpenshiftSvcatControllerManagerRolebindingConfigmapAccessorYaml() (*asset, error) {
	bytes, err := v3110OpenshiftSvcatControllerManagerRolebindingConfigmapAccessorYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/openshift-svcat-controller-manager/rolebinding-configmap-accessor.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftSvcatControllerManagerSaYaml = []byte(`kind: ServiceAccount
apiVersion: v1
metadata:
  name: service-catalog-controller
  namespace: openshift-service-catalog-controller-manager
`)

func v3110OpenshiftSvcatControllerManagerSaYamlBytes() ([]byte, error) {
	return _v3110OpenshiftSvcatControllerManagerSaYaml, nil
}

func v3110OpenshiftSvcatControllerManagerSaYaml() (*asset, error) {
	bytes, err := v3110OpenshiftSvcatControllerManagerSaYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/openshift-svcat-controller-manager/sa.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftSvcatControllerManagerServicemonitorRoleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: prometheus-k8s
  namespace: openshift-service-catalog-controller-manager
rules:
- apiGroups:
  - ""
  resources:
  - services
  - endpoints
  - pods
  verbs:
  - get
  - list
  - watch
`)

func v3110OpenshiftSvcatControllerManagerServicemonitorRoleYamlBytes() ([]byte, error) {
	return _v3110OpenshiftSvcatControllerManagerServicemonitorRoleYaml, nil
}

func v3110OpenshiftSvcatControllerManagerServicemonitorRoleYaml() (*asset, error) {
	bytes, err := v3110OpenshiftSvcatControllerManagerServicemonitorRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/openshift-svcat-controller-manager/servicemonitor-role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftSvcatControllerManagerServicemonitorRolebindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: prometheus-k8s
  namespace: openshift-service-catalog-controller-manager
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: prometheus-k8s
subjects:
- kind: ServiceAccount
  name: prometheus-k8s
  namespace: openshift-monitoring
`)

func v3110OpenshiftSvcatControllerManagerServicemonitorRolebindingYamlBytes() ([]byte, error) {
	return _v3110OpenshiftSvcatControllerManagerServicemonitorRolebindingYaml, nil
}

func v3110OpenshiftSvcatControllerManagerServicemonitorRolebindingYaml() (*asset, error) {
	bytes, err := v3110OpenshiftSvcatControllerManagerServicemonitorRolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/openshift-svcat-controller-manager/servicemonitor-rolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftSvcatControllerManagerServicemonitorYaml = []byte(`apiVersion: monitoring.coreos.com/v1
kind: ServiceMonitor
metadata:
  name: servicemonitor
  namespace: openshift-service-catalog-controller-manager
spec:
  endpoints:
  - bearerTokenFile: /var/run/secrets/kubernetes.io/serviceaccount/token
    interval: 30s
    port: https
    scheme: https
    tlsConfig:
      caFile: /etc/prometheus/configmaps/serving-certs-ca-bundle/service-ca.crt
      serverName: controller-manager.openshift-service-catalog-controller-manager.svc
  namespaceSelector:
    matchNames:
    - openshift-service-catalog-controller-manager
  selector: {}
`)

func v3110OpenshiftSvcatControllerManagerServicemonitorYamlBytes() ([]byte, error) {
	return _v3110OpenshiftSvcatControllerManagerServicemonitorYaml, nil
}

func v3110OpenshiftSvcatControllerManagerServicemonitorYaml() (*asset, error) {
	bytes, err := v3110OpenshiftSvcatControllerManagerServicemonitorYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/openshift-svcat-controller-manager/servicemonitor.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftSvcatControllerManagerSvcYaml = []byte(`apiVersion: v1
kind: Service
metadata:
  namespace: openshift-service-catalog-controller-manager
  name: controller-manager
  annotations:
    service.alpha.openshift.io/serving-cert-secret-name: serving-cert
spec:
  selector:
    controller-manager: "true"
  ports:
  - name: https
    port: 443
    targetPort: 6443
`)

func v3110OpenshiftSvcatControllerManagerSvcYamlBytes() ([]byte, error) {
	return _v3110OpenshiftSvcatControllerManagerSvcYaml, nil
}

func v3110OpenshiftSvcatControllerManagerSvcYaml() (*asset, error) {
	bytes, err := v3110OpenshiftSvcatControllerManagerSvcYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/openshift-svcat-controller-manager/svc.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"v3.11.0/openshift-svcat-controller-manager/cm.yaml":                                      v3110OpenshiftSvcatControllerManagerCmYaml,
	"v3.11.0/openshift-svcat-controller-manager/cr-catalog-controller.yaml":                   v3110OpenshiftSvcatControllerManagerCrCatalogControllerYaml,
	"v3.11.0/openshift-svcat-controller-manager/crb-catalog-controller.yaml":                  v3110OpenshiftSvcatControllerManagerCrbCatalogControllerYaml,
	"v3.11.0/openshift-svcat-controller-manager/crb-controller-namespace-viewer-binding.yaml": v3110OpenshiftSvcatControllerManagerCrbControllerNamespaceViewerBindingYaml,
	"v3.11.0/openshift-svcat-controller-manager/defaultconfig.yaml":                           v3110OpenshiftSvcatControllerManagerDefaultconfigYaml,
	"v3.11.0/openshift-svcat-controller-manager/ds.yaml":                                      v3110OpenshiftSvcatControllerManagerDsYaml,
	"v3.11.0/openshift-svcat-controller-manager/ns.yaml":                                      v3110OpenshiftSvcatControllerManagerNsYaml,
	"v3.11.0/openshift-svcat-controller-manager/role-cluster-info-configmap.yaml":             v3110OpenshiftSvcatControllerManagerRoleClusterInfoConfigmapYaml,
	"v3.11.0/openshift-svcat-controller-manager/role-configmap-accessor.yaml":                 v3110OpenshiftSvcatControllerManagerRoleConfigmapAccessorYaml,
	"v3.11.0/openshift-svcat-controller-manager/rolebinding-cluster-info-configmap.yaml":      v3110OpenshiftSvcatControllerManagerRolebindingClusterInfoConfigmapYaml,
	"v3.11.0/openshift-svcat-controller-manager/rolebinding-configmap-accessor.yaml":          v3110OpenshiftSvcatControllerManagerRolebindingConfigmapAccessorYaml,
	"v3.11.0/openshift-svcat-controller-manager/sa.yaml":                                      v3110OpenshiftSvcatControllerManagerSaYaml,
	"v3.11.0/openshift-svcat-controller-manager/servicemonitor-role.yaml":                     v3110OpenshiftSvcatControllerManagerServicemonitorRoleYaml,
	"v3.11.0/openshift-svcat-controller-manager/servicemonitor-rolebinding.yaml":              v3110OpenshiftSvcatControllerManagerServicemonitorRolebindingYaml,
	"v3.11.0/openshift-svcat-controller-manager/servicemonitor.yaml":                          v3110OpenshiftSvcatControllerManagerServicemonitorYaml,
	"v3.11.0/openshift-svcat-controller-manager/svc.yaml":                                     v3110OpenshiftSvcatControllerManagerSvcYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"v3.11.0": {nil, map[string]*bintree{
		"openshift-svcat-controller-manager": {nil, map[string]*bintree{
			"cm.yaml":                                      {v3110OpenshiftSvcatControllerManagerCmYaml, map[string]*bintree{}},
			"cr-catalog-controller.yaml":                   {v3110OpenshiftSvcatControllerManagerCrCatalogControllerYaml, map[string]*bintree{}},
			"crb-catalog-controller.yaml":                  {v3110OpenshiftSvcatControllerManagerCrbCatalogControllerYaml, map[string]*bintree{}},
			"crb-controller-namespace-viewer-binding.yaml": {v3110OpenshiftSvcatControllerManagerCrbControllerNamespaceViewerBindingYaml, map[string]*bintree{}},
			"defaultconfig.yaml":                           {v3110OpenshiftSvcatControllerManagerDefaultconfigYaml, map[string]*bintree{}},
			"ds.yaml":                                      {v3110OpenshiftSvcatControllerManagerDsYaml, map[string]*bintree{}},
			"ns.yaml":                                      {v3110OpenshiftSvcatControllerManagerNsYaml, map[string]*bintree{}},
			"role-cluster-info-configmap.yaml":             {v3110OpenshiftSvcatControllerManagerRoleClusterInfoConfigmapYaml, map[string]*bintree{}},
			"role-configmap-accessor.yaml":                 {v3110OpenshiftSvcatControllerManagerRoleConfigmapAccessorYaml, map[string]*bintree{}},
			"rolebinding-cluster-info-configmap.yaml":      {v3110OpenshiftSvcatControllerManagerRolebindingClusterInfoConfigmapYaml, map[string]*bintree{}},
			"rolebinding-configmap-accessor.yaml":          {v3110OpenshiftSvcatControllerManagerRolebindingConfigmapAccessorYaml, map[string]*bintree{}},
			"sa.yaml":                                      {v3110OpenshiftSvcatControllerManagerSaYaml, map[string]*bintree{}},
			"servicemonitor-role.yaml":                     {v3110OpenshiftSvcatControllerManagerServicemonitorRoleYaml, map[string]*bintree{}},
			"servicemonitor-rolebinding.yaml":              {v3110OpenshiftSvcatControllerManagerServicemonitorRolebindingYaml, map[string]*bintree{}},
			"servicemonitor.yaml":                          {v3110OpenshiftSvcatControllerManagerServicemonitorYaml, map[string]*bintree{}},
			"svc.yaml":                                     {v3110OpenshiftSvcatControllerManagerSvcYaml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
