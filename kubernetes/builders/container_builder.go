package builders

import (
	corev1 "k8s.io/api/core/v1"
)

// Container Struct for containers
type Container struct {
	Name      string
	Image     string
	Tag       string
	Port      int32
	Container *corev1.Container
}

// NewContainerBuilder return a container struct
func NewContainerBuilder() Container {
	return Container{}
}

// SetName Set container name
func (c *Container) SetName(name string) *Container {
	c.Name = name
	return c
}

// SetImage Set container image
func (c *Container) SetImage(image string) *Container {
	c.Image = image
	return c
}

// SetTag Set container tag
func (c *Container) SetTag(tag string) *Container {
	c.Tag = tag
	return c
}

// SetPort Set container port
func (c *Container) SetPort(port int32) *Container {
	c.Port = port
	return c
}

// Build build container
func (c *Container) Build() *corev1.Container {
	container := &corev1.Container{
		Name:  c.Name,
		Image: c.Image + ":" + c.Tag,
		Ports: []corev1.ContainerPort{
			{
				Name:          "http",
				Protocol:      corev1.ProtocolTCP,
				ContainerPort: c.Port,
			},
		},
	}
	c.Container = container
	return c.Container
}
