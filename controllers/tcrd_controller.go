/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	netv1 "k8s.io/api/networking/v1"

	tapiv1 "github.com/Tasdidur/finalCRD/api/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// TCrdReconciler reconciles a TCrd object
type TCrdReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=tapi.tasdid,resources=tcrds,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=tapi.tasdid,resources=tcrds/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=tapi.tasdid,resources=tcrds/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the TCrd object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.8.3/pkg/reconcile
func (r *TCrdReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	Log := log.FromContext(ctx)

	// your logic here
	dep := &appsv1.Deployment{}
	if err := r.Get(ctx,req.NamespacedName,dep); err != nil{
		Log.Error(err,"unable to fetch deploymetns")
		return ctrl.Result{},client.IgnoreNotFound(err)
	} else{
		myCrd := &tapiv1.TCrd{}
		r.Get(ctx,req.NamespacedName,myCrd)
		if dep.Spec.Template.Name != myCrd.Spec.Name{
			fmt.Println("creating deployment...")
			newDep := newDeployment(myCrd)
			err = r.Create(ctx,newDep)
			if err !=nil{
				Log.Error(err,"unable to create deployment")
				return ctrl.Result{},client.IgnoreNotFound(err)
			}else {
				r.Status().Update(ctx,newDep)
			}
		}
	}


	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *TCrdReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&tapiv1.TCrd{}).
		Owns(&appsv1.Deployment{}).
		Owns(&corev1.Service{}).
		Owns(&netv1.Ingress{}).
		Complete(r)
}
