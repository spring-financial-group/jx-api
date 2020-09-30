// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/jenkins-x/jx-api/pkg/apis/jenkins.io/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PipelineActivityLister helps list PipelineActivities.
// All objects returned here must be treated as read-only.
type PipelineActivityLister interface {
	// List lists all PipelineActivities in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.PipelineActivity, err error)
	// PipelineActivities returns an object that can list and get PipelineActivities.
	PipelineActivities(namespace string) PipelineActivityNamespaceLister
	PipelineActivityListerExpansion
}

// pipelineActivityLister implements the PipelineActivityLister interface.
type pipelineActivityLister struct {
	indexer cache.Indexer
}

// NewPipelineActivityLister returns a new PipelineActivityLister.
func NewPipelineActivityLister(indexer cache.Indexer) PipelineActivityLister {
	return &pipelineActivityLister{indexer: indexer}
}

// List lists all PipelineActivities in the indexer.
func (s *pipelineActivityLister) List(selector labels.Selector) (ret []*v1.PipelineActivity, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.PipelineActivity))
	})
	return ret, err
}

// PipelineActivities returns an object that can list and get PipelineActivities.
func (s *pipelineActivityLister) PipelineActivities(namespace string) PipelineActivityNamespaceLister {
	return pipelineActivityNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PipelineActivityNamespaceLister helps list and get PipelineActivities.
// All objects returned here must be treated as read-only.
type PipelineActivityNamespaceLister interface {
	// List lists all PipelineActivities in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.PipelineActivity, err error)
	// Get retrieves the PipelineActivity from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.PipelineActivity, error)
	PipelineActivityNamespaceListerExpansion
}

// pipelineActivityNamespaceLister implements the PipelineActivityNamespaceLister
// interface.
type pipelineActivityNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PipelineActivities in the indexer for a given namespace.
func (s pipelineActivityNamespaceLister) List(selector labels.Selector) (ret []*v1.PipelineActivity, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.PipelineActivity))
	})
	return ret, err
}

// Get retrieves the PipelineActivity from the indexer for a given namespace and name.
func (s pipelineActivityNamespaceLister) Get(name string) (*v1.PipelineActivity, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("pipelineactivity"), name)
	}
	return obj.(*v1.PipelineActivity), nil
}
