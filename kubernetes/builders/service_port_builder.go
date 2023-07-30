package builders

import (
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// Ports ports struct for service ports
type Ports struct {
	Name        string
	Port        int32
	Protocol    apiv1.Protocol
	TargetPort  intstr.IntOrString
	NodePort    int32
	ServicePort *apiv1.ServicePort
}

// SetName set name of the service port
func (p *Ports) SetName(name string) {
	p.Name = name
}

// SetPort set port of the service port
func (p *Ports) SetPort(port int32) {
	p.Port = port
}

// SetProtocol set protocol of the service port
func (p *Ports) SetProtocol(protocol apiv1.Protocol) {
	p.Protocol = protocol
}

// SetTargetPort set target port for service port
func (p *Ports) SetTargetPort(targetPort int) {
	p.TargetPort = intstr.FromInt(targetPort)
}

// SetNodePort set Node port for service port
func (p *Ports) SetNodePort(nodePort int32) {
	p.NodePort = nodePort
}

// Build build the service port
func (p *Ports) Build() *apiv1.ServicePort {
	p.ServicePort = &apiv1.ServicePort{
		Name:       p.Name,
		Port:       p.Port,
		Protocol:   p.Protocol,
		TargetPort: p.TargetPort,
		NodePort:   p.NodePort,
	}
	return p.ServicePort
}
