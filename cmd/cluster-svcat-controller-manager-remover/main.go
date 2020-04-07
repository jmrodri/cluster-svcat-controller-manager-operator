package main

import (
	"os"

	operatorapiv1 "github.com/openshift/api/operator/v1"
	configclient "github.com/openshift/client-go/config/clientset/versioned"
	operatorclient "github.com/openshift/client-go/operator/clientset/versioned"
	operatorv1 "github.com/openshift/client-go/operator/clientset/versioned/typed/operator/v1"
	log "github.com/sirupsen/logrus"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var targetNamespaceName = "openshift-service-catalog-controller-manager-operator"

func createClientConfigFromFile(configPath string) (*rest.Config, error) {
	clientConfig, err := clientcmd.LoadFromFile(configPath)
	if err != nil {
		return nil, err
	}

	config, err := clientcmd.NewDefaultClientConfig(*clientConfig, &clientcmd.ConfigOverrides{}).ClientConfig()
	if err != nil {
		return nil, err
	}
	return config, nil
}

func deleteTargetNamespace(kubeClient *kubernetes.Clientset, target string) {
	log.Infof("Removing target namespace %s", target)
	if err := kubeClient.CoreV1().Namespaces().Delete(target, nil); err != nil && !apierrors.IsNotFound(err) {
		log.Errorf("problem removing target namespace [%s] :  %v", target, err)
	}
}

func deleteCustomResource(client operatorv1.OperatorV1Interface) {
	// DUMMY COMMIT
	log.Info("Removing the ServiceCatalogControllerManager CR")
	err := client.ServiceCatalogControllerManagers().Delete("cluster", &metav1.DeleteOptions{})
	if err != nil {
		log.Errorf("ServiceCatalogControllerManager cr deletion failed: %v", err)
	} else {
		log.Info("ServiceCatalogControllerManager cr removed successfully.")
	}
}

func deleteClusterOperator(clientConfig *rest.Config) {
	configClient, err := configclient.NewForConfig(clientConfig)
	if err != nil {
		log.Errorf("problem getting config client, error %v", err)
	}

	log.Info("Removing the service-catalog-controller-manager clusteroperator")
	err = configClient.ConfigV1().ClusterOperators().Delete("service-catalog-controller-manager", &metav1.DeleteOptions{})
	if err != nil && !apierrors.IsNotFound(err) {
		log.Errorf("problem removing cluster operator [service-catalog-controller-manager] :  %v", err)
	}
}

func deleteClusterRolesAndBindings(kubeClient *kubernetes.Clientset) {
	log.Info("Removing ClusterRoleBinding: openshift-service-catalog-controller-manager-operator")
	err := kubeClient.RbacV1().ClusterRoleBindings().Delete("openshift-service-catalog-controller-manager-operator", &metav1.DeleteOptions{})
	if err != nil && !apierrors.IsNotFound(err) {
		log.Errorf("problem removing cluster role binding [openshift-service-catalog-controller-manager-operator] :  %v", err)
	}

	log.Info("Removing ClusterRole: openshift-service-catalog-controller-manager-operator")
	err = kubeClient.RbacV1().ClusterRoles().Delete("openshift-service-catalog-controller-manager-operator", &metav1.DeleteOptions{})
	if err != nil && !apierrors.IsNotFound(err) {
		log.Errorf("problem removing cluster role [openshift-service-catalog-controller-manager-operator] :  %v", err)
	}
}

func main() {
	log.Info("Starting openshift-service-catalog-controller-manager-remover job")

	clientConfig, err := rest.InClusterConfig()
	if err != nil {
		clientConfig, err = createClientConfigFromFile(homedir.HomeDir() + "/.kube/config")
		if err != nil {
			log.Error("Failed to create LocalClientSet")
			panic(err.Error())
		}
	}

	kubeClient, err := kubernetes.NewForConfig(clientConfig)
	if err != nil {
		panic(err.Error())
	}

	operatorClient, err := operatorclient.NewForConfig(clientConfig)
	if err != nil {
		log.Errorf("problem getting operator client, error %v", err)
	}
	operatorConfigClient := operatorClient.OperatorV1()
	operatorConfig, err := operatorConfigClient.ServiceCatalogControllerManagers().Get("cluster", metav1.GetOptions{})
	if apierrors.IsNotFound(err) {
		log.Info("ServiceCatalogControllerManager cr has already been removed.")
		deleteTargetNamespace(kubeClient, targetNamespaceName)
		deleteClusterOperator(clientConfig)
		deleteClusterRolesAndBindings(kubeClient)
		os.Exit(0)
	} else if err != nil {
		log.Errorf("problem getting ServiceCatalogControllerManage CR, error %v", err)
	}

	// Handle the various ManagementStates
	switch operatorConfig.Spec.ManagementState {
	case operatorapiv1.Managed:
		log.Warning("We found a cluster-svcat-controller-manager-operator in Managed state. Aborting")
	case operatorapiv1.Unmanaged:
		log.Info("ServiceCatalogControllerManager managementState is 'Unmanaged'")
		deleteTargetNamespace(kubeClient, targetNamespaceName)
		deleteCustomResource(operatorConfigClient)
		deleteClusterOperator(clientConfig)
		deleteClusterRolesAndBindings(kubeClient)
	case operatorapiv1.Removed:
		log.Info("ServiceCatalogControllerManager managementState is 'Removed'")
		deleteTargetNamespace(kubeClient, targetNamespaceName)
		deleteCustomResource(operatorConfigClient)
		deleteClusterOperator(clientConfig)
		deleteClusterRolesAndBindings(kubeClient)
	default:
		log.Error("Unknown managementState")
	}
	log.Info("The openshift-service-catalog-controller-manater-remover job has finished.")
}
