/*
Copyright 2020.

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
	"strings"
	"time"

	bmh_v1alpha1 "github.com/metal3-io/baremetal-operator/pkg/apis/metal3/v1alpha1"
	adiiov1alpha1 "github.com/openshift/assisted-service/internal/controller/api/v1alpha1"
	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// BMHReconciler reconciles a Agent object
type BMHReconciler struct {
	client.Client
	Log    logrus.FieldLogger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=metal3.io,resources=baremetalhosts,verbs=get;list;watch;update;patch

func (r *BMHReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	bmh := &bmh_v1alpha1.BareMetalHost{}

	err := r.Get(ctx, req.NamespacedName, bmh)
	if err != nil {
		r.Log.WithError(err).Errorf("Failed to get resource %s", req.NamespacedName)
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	if bmh.Spec.Image != nil && bmh.Spec.Image.URL != "" {
		return ctrl.Result{}, nil
	}
	for _, label := range bmh.Labels {
		if strings.HasPrefix(label, "InstallEnv:") {
			installEnvName := strings.TrimPrefix(label, "InstallEnv:")
			// TODO how to get the Namespace? encoded in the string?
			// InstallEnv:ns/name
			installEnv := &adiiov1alpha1.InstallEnv{}
			if err := r.Get(ctx, types.NamespacedName{Name: installEnvName, Namespace: "??"}, installEnv); err != nil {
				return ctrl.Result{}, client.IgnoreNotFound(err)
			}
			if installEnv.Status.ISODownloadURL == "" {
				return ctrl.Result{RequeueAfter: time.Minute}, nil // the image has not been created yet, try later.
			}

			if bmh.Spec.Image == nil {
				bmh.Spec.Image = &bmh_v1alpha1.Image{}
			}
			bmh.Spec.Image.URL = installEnv.Status.ISODownloadURL
			return ctrl.Result{}, r.Client.Update(ctx, bmh)
		}
	}

	return ctrl.Result{}, nil
}

func (r *BMHReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&bmh_v1alpha1.BareMetalHost{}).
		Complete(r)
}
