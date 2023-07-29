package builders

import (
	apiv1 "k8s.io/api/core/v1"
)

// Container Struct for containers
type Container struct {
	Name  string
	Image string
	Tag   string
	Port  int32
}

// NewContainerBuilder return a container struct
func NewContainerBuilder() Container {
	return Container{}
}

// SetName Set container name
func (c *Container) SetName(name string) {
	c.Name = name
}

// SetImage Set container image
func (c *Container) SetImage(image string) {
	c.Image = image
}

// SetTag Set container tag
func (c *Container) SetTag(tag string) {
	c.Tag = tag
}

// SetPort Set container port
func (c *Container) SetPort(port int32) {
	c.Port = port
}

// Build build container
func (c *Container) Build() apiv1.Container {
	container := apiv1.Container{
		Name:  c.Name,
		Image: c.Image + ":" + c.Tag,
		Ports: []apiv1.ContainerPort{
			{
				Name:          "http",
				Protocol:      apiv1.ProtocolTCP,
				ContainerPort: c.Port,
			},
		},
	}
	return container
}
