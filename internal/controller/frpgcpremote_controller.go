/*
Copyright 2023.

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

package controller

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	experimentalmilkshakescloudv1 "milkshakes.cloud/frp-gcp-remote-operator/api/v1"
)

// FrpGCPRemoteReconciler reconciles a FrpGCPRemote object
type FrpGCPRemoteReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

func rangeIn(low, hi int) int {
	return low + rand.Intn(hi-low)
}

func constructJobForFrpGCPProvisioning(fgr *experimentalmilkshakescloudv1.FrpGCPRemote) (*batchv1.Job, error) {
	name := fmt.Sprintf("frpsgcpd-%s-%s", fgr.Spec.DNSAName, fgr.Spec.DNSZone)

	labels := fgr.ObjectMeta.Labels
	annotations := fgr.ObjectMeta.Annotations

	provisioningJob := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Labels:            labels,
			Annotations:       annotations,
			Name:              name,
			CreationTimestamp: metav1.Time{Time: time.Now()},
			OwnerReferences:   []metav1.OwnerReference{*metav1.NewControllerRef(fgr, experimentalmilkshakescloudv1.SchemeBuilder.GroupVersion.WithKind(fgr.TypeMeta.Kind))},
		},
	}

	// TODO: Build job spec
	fgr.Spec.JobTemplate.Spec.DeepCopyInto(&provisioningJob.Spec)

	return provisioningJob, nil
}

//+kubebuilder:rbac:groups=experimental.milkshakes.cloud.milkshakes.cloud,resources=frpgcpremotes,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=experimental.milkshakes.cloud.milkshakes.cloud,resources=frpgcpremotes/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=experimental.milkshakes.cloud.milkshakes.cloud,resources=frpgcpremotes/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the FrpGCPRemote object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.16.0/pkg/reconcile
func (r *FrpGCPRemoteReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	var frpRemote experimentalmilkshakescloudv1.FrpGCPRemote
	if err := r.Get(ctx, req.NamespacedName, &frpRemote); err != nil {
		logger.Error(err, "Unable to fetch FrpGCPRemote resource")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if frpRemote.Status.ProvisionStatus != "Started" {
		provisioningJob, err := constructJobForFrpGCPProvisioning(&frpRemote.Spec)
		if err != nil {
			logger.Error(err, "Unable to create provisioning job object")
			return ctrl.Result{}, client.IgnoreNotFound(err)
		}
		if err := r.Create(ctx, provisioningJob); err != nil {
			logger.Error(err, "Unable to create provisioning job resource")
			return ctrl.Result{}, client.IgnoreNotFound(err)
		}

		frpRemote.Status.ProvisionStatus = "Started"
		frpRemote.Status.Active = "Provisioning"

		dnsName := fmt.Sprintf("%s.%s.%s", frpRemote.Spec.DNSAName, frpRemote.Spec.DNSZone, frpRemote.Spec.DNSBaseDomain)
		frpRemote.Status.RemoteDNSName = dnsName

		if err := r.Status().Update(ctx, &frpRemote); err != nil {
			logger.Error(err, "unable to update FrpGCPRemote status")
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *FrpGCPRemoteReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&experimentalmilkshakescloudv1.FrpGCPRemote{}).
		Complete(r)
}
