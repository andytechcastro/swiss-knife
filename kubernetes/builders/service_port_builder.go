package builders

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// Ports ports struct for service ports
type Ports struct {
	Name        string
	Port        int32
	Protocol    corev1.Protocol
	TargetPort  intstr.IntOrString
	NodePort    int32
	ServicePort *corev1.ServicePort
}

// NewServicePort get a new service port
func NewServicePort() *Ports {
	return &Ports{}
}

// SetName set name of the service port
func (p *Ports) SetName(name string) *Ports {
	p.Name = name
	return p
}

// SetPort set port of the service port
func (p *Ports) SetPort(port int32) *Ports {
	p.Port = port
	return p
}

// SetProtocol set protocol of the service port
func (p *Ports) SetProtocol(protocol corev1.Protocol) *Ports {
	p.Protocol = protocol
	return p
}

// SetTargetPort set target port for service port
func (p *Ports) SetTargetPort(targetPort int) *Ports {
	p.TargetPort = intstr.FromInt(targetPort)
	return p
}

// SetNodePort set Node port for service port
func (p *Ports) SetNodePort(nodePort int32) *Ports {
	p.NodePort = nodePort
	return p
}

// Build build the service port
func (p *Ports) Build() *corev1.ServicePort {
	p.ServicePort = &corev1.ServicePort{
		Name:       p.Name,
		Port:       p.Port,
		Protocol:   p.Protocol,
		TargetPort: p.TargetPort,
		NodePort:   p.NodePort,
	}
	return p.ServicePort
}
