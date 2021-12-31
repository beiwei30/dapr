package handlers

import (
	"context"
	ctrl "sigs.k8s.io/controller-runtime"
)

// Handler is the interface for dealing with Dapr CRDs state changes.
type Handler interface {
	Init() error
	Reconcile(context.Context, ctrl.Request) (ctrl.Result, error)
}
