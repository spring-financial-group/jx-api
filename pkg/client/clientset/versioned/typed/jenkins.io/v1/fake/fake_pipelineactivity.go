// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	jenkinsiov1 "github.com/jenkins-x/jx-api/pkg/apis/jenkins.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePipelineActivities implements PipelineActivityInterface
type FakePipelineActivities struct {
	Fake *FakeJenkinsV1
	ns   string
}

var pipelineactivitiesResource = schema.GroupVersionResource{Group: "jenkins.io", Version: "v1", Resource: "pipelineactivities"}

var pipelineactivitiesKind = schema.GroupVersionKind{Group: "jenkins.io", Version: "v1", Kind: "PipelineActivity"}

// Get takes name of the pipelineActivity, and returns the corresponding pipelineActivity object, and an error if there is any.
func (c *FakePipelineActivities) Get(ctx context.Context, name string, options v1.GetOptions) (result *jenkinsiov1.PipelineActivity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(pipelineactivitiesResource, c.ns, name), &jenkinsiov1.PipelineActivity{})

	if obj == nil {
		return nil, err
	}
	return obj.(*jenkinsiov1.PipelineActivity), err
}

// List takes label and field selectors, and returns the list of PipelineActivities that match those selectors.
func (c *FakePipelineActivities) List(ctx context.Context, opts v1.ListOptions) (result *jenkinsiov1.PipelineActivityList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(pipelineactivitiesResource, pipelineactivitiesKind, c.ns, opts), &jenkinsiov1.PipelineActivityList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &jenkinsiov1.PipelineActivityList{ListMeta: obj.(*jenkinsiov1.PipelineActivityList).ListMeta}
	for _, item := range obj.(*jenkinsiov1.PipelineActivityList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pipelineActivities.
func (c *FakePipelineActivities) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(pipelineactivitiesResource, c.ns, opts))

}

// Create takes the representation of a pipelineActivity and creates it.  Returns the server's representation of the pipelineActivity, and an error, if there is any.
func (c *FakePipelineActivities) Create(ctx context.Context, pipelineActivity *jenkinsiov1.PipelineActivity, opts v1.CreateOptions) (result *jenkinsiov1.PipelineActivity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(pipelineactivitiesResource, c.ns, pipelineActivity), &jenkinsiov1.PipelineActivity{})

	if obj == nil {
		return nil, err
	}
	return obj.(*jenkinsiov1.PipelineActivity), err
}

// Update takes the representation of a pipelineActivity and updates it. Returns the server's representation of the pipelineActivity, and an error, if there is any.
func (c *FakePipelineActivities) Update(ctx context.Context, pipelineActivity *jenkinsiov1.PipelineActivity, opts v1.UpdateOptions) (result *jenkinsiov1.PipelineActivity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(pipelineactivitiesResource, c.ns, pipelineActivity), &jenkinsiov1.PipelineActivity{})

	if obj == nil {
		return nil, err
	}
	return obj.(*jenkinsiov1.PipelineActivity), err
}

// Delete takes name of the pipelineActivity and deletes it. Returns an error if one occurs.
func (c *FakePipelineActivities) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(pipelineactivitiesResource, c.ns, name), &jenkinsiov1.PipelineActivity{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePipelineActivities) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(pipelineactivitiesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &jenkinsiov1.PipelineActivityList{})
	return err
}

// Patch applies the patch and returns the patched pipelineActivity.
func (c *FakePipelineActivities) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *jenkinsiov1.PipelineActivity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(pipelineactivitiesResource, c.ns, name, pt, data, subresources...), &jenkinsiov1.PipelineActivity{})

	if obj == nil {
		return nil, err
	}
	return obj.(*jenkinsiov1.PipelineActivity), err
}
