package helixsaga

import (
	"fmt"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog"

	helixsagav1 "github.com/Shanghai-Lunara/helixsaga-operator/pkg/apis/helixsaga/v1"
	helixsagaclientset "github.com/Shanghai-Lunara/helixsaga-operator/pkg/generated/helixsaga/clientset/versioned"
	helixsagascheme "github.com/Shanghai-Lunara/helixsaga-operator/pkg/generated/helixsaga/clientset/versioned/scheme"
	informersext "github.com/Shanghai-Lunara/helixsaga-operator/pkg/generated/helixsaga/informers/externalversions"
	informers "github.com/Shanghai-Lunara/helixsaga-operator/pkg/generated/helixsaga/informers/externalversions/helixsaga/v1"
	k8scorev1 "github.com/nevercase/k8s-controller-custom-resource/core/v1"
)

func NewController(
	controllerName string,
	kubeclientset kubernetes.Interface,
	sampleclientset helixsagaclientset.Interface,
	stopCh <-chan struct{}) k8scorev1.KubernetesControllerV1 {

	exampleInformerFactory := informersext.NewSharedInformerFactory(sampleclientset, time.Second*30)
	fooInformer := exampleInformerFactory.Nevercase().V1().HelixSagas()

	//roInformerFactory := informersv2.NewSharedInformerFactory(sampleclientset, time.Second*30)

	opt := k8scorev1.NewOption(&helixsagav1.HelixSaga{},
		controllerName,
		OperatorKindName,
		helixsagascheme.AddToScheme(scheme.Scheme),
		sampleclientset,
		fooInformer,
		fooInformer.Informer(),
		CompareResourceVersion,
		Get,
		Sync,
		SyncStatus)
	opts := k8scorev1.NewOptions()
	if err := opts.Add(opt); err != nil {
		klog.Fatal(err)
	}
	op := k8scorev1.NewKubernetesOperator(kubeclientset, stopCh, controllerName, opts)
	kc := k8scorev1.NewKubernetesController(op)
	//roInformerFactory.Start(stopCh)
	exampleInformerFactory.Start(stopCh)
	return kc
}

func NewOption(controllerName string, cfg *rest.Config, stopCh <-chan struct{}) k8scorev1.Option {
	c, err := helixsagaclientset.NewForConfig(cfg)
	if err != nil {
		klog.Fatal("Error building clientSet: %s", err.Error())
	}
	informerFactory := informersext.NewSharedInformerFactory(c, time.Second*30)
	fooInformer := informerFactory.Nevercase().V1().HelixSagas()
	opt := k8scorev1.NewOption(&helixsagav1.HelixSaga{},
		controllerName,
		OperatorKindName,
		helixsagascheme.AddToScheme(scheme.Scheme),
		c,
		fooInformer,
		fooInformer.Informer(),
		CompareResourceVersion,
		Get,
		Sync,
		SyncStatus)
	informerFactory.Start(stopCh)
	return opt
}

func CompareResourceVersion(old, new interface{}) bool {
	newResource := new.(*helixsagav1.HelixSaga)
	oldResource := old.(*helixsagav1.HelixSaga)
	if newResource.ResourceVersion == oldResource.ResourceVersion {
		// Periodic resync will send update events for all known Deployments.
		// Two different versions of the same Deployment will always have different RVs.
		return true
	}
	return false
}

func Get(foo interface{}, nameSpace, ownerRefName string) (obj interface{}, err error) {
	kc := foo.(informers.HelixSagaInformer)
	return kc.Lister().HelixSagas(nameSpace).Get(ownerRefName)
}

func Sync(obj interface{}, clientObj interface{}, ks k8scorev1.KubernetesResource, recorder record.EventRecorder) error {
	hs := obj.(*helixsagav1.HelixSaga)
	clientSet := clientObj.(helixsagaclientset.Interface)
	for _, v := range hs.Spec.Applications {
		klog.Info("v:", v)
		if err := NewStatefulSetAndService(ks, clientSet, hs, v.Spec); err != nil {
			klog.V(2).Info(err)
			return err
		}
	}
	recorder.Event(hs, corev1.EventTypeNormal, SuccessSynced, MessageResourceSynced)
	return nil
}

func SyncStatus(obj interface{}, clientObj interface{}, ks k8scorev1.KubernetesResource, recorder record.EventRecorder) error {
	clientSet := clientObj.(helixsagaclientset.Interface)
	ss := obj.(*appsv1.StatefulSet)
	var objName string
	if t, ok := ss.Labels[k8scorev1.LabelController]; ok {
		objName = t
	} else {
		return fmt.Errorf(ErrResourceNotMatch, "no controller")
	}
	hs, err := clientSet.NevercaseV1().HelixSagas(ss.Namespace).Get(objName, metav1.GetOptions{})
	if err != nil {
		return err
	}
	var appName string
	if t, ok := ss.Labels[k8scorev1.LabelName]; ok {
		appName = t
	} else {
		return fmt.Errorf(ErrResourceNotMatch, "no appName")
	}
	if err := updateStatus(hs, clientSet, ss, appName); err != nil {
		return err
	}
	recorder.Event(hs, corev1.EventTypeNormal, SuccessSynced, MessageResourceSynced)
	return nil
}
