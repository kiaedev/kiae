package utils

import (
	"fmt"

	"github.com/kiaedev/kiae/api/app"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

type ResourceType struct {
	Name  app.Size `json:"name"`
	Label string   `json:"label"`

	LimitsCPU resource.Quantity `json:"-"`
	LimitsMem resource.Quantity `json:"-"`
}

func NewResourceType(name app.Size, cpu, mem string) ResourceType {
	return ResourceType{
		Name:  name,
		Label: fmt.Sprintf("%sC-%s", cpu, mem),

		LimitsCPU: resource.MustParse(cpu),
		LimitsMem: resource.MustParse(mem),
	}
}

func (rt *ResourceType) RequestsCPU(oversoldRate float64) resource.Quantity {
	newVal := float64(rt.LimitsCPU.MilliValue()) * oversoldRate
	return *resource.NewMilliQuantity(int64(newVal), rt.LimitsCPU.Format)
}

func (rt *ResourceType) RequestsMem(oversoldRate float64) resource.Quantity {
	newVal := float64(rt.LimitsMem.MilliValue()) * oversoldRate
	return *resource.NewMilliQuantity(int64(newVal), rt.LimitsMem.Format)
}

var AppResourceTypes = []ResourceType{
	NewResourceType(app.Size_SIZE_NANO, "0.1", "256Mi"),
	NewResourceType(app.Size_SIZE_MIRCO, "0.25", "512Mi"),
	NewResourceType(app.Size_SIZE_MINI, "0.5", "1Gi"),
	NewResourceType(app.Size_SIZE_SMALL, "1", "2Gi"),
	NewResourceType(app.Size_SIZE_MEDIUM, "2", "4Gi"),
	NewResourceType(app.Size_SIZE_LARGE, "4", "8Gi"),
	NewResourceType(app.Size_SIZE_XLARGE, "8", "16Gi"),
	// NewResourceType(app.Size_SIZE_XXLARGE, "16C", "256M"), // 暂不支持
}

func AppResourceTypeGet(name app.Size) ResourceType {
	for _, resourceType := range AppResourceTypes {
		if name == resourceType.Name {
			return resourceType
		}

	}

	return NewResourceType(app.Size_SIZE_MIRCO, "0.25C", "512Mi")
}

func BuildResources(size app.Size, oversoldRate float64) v1.ResourceRequirements {
	cr := AppResourceTypeGet(size)
	return v1.ResourceRequirements{
		Limits: v1.ResourceList{
			v1.ResourceCPU:              cr.LimitsCPU,
			v1.ResourceMemory:           cr.LimitsMem,
			v1.ResourceEphemeralStorage: resource.MustParse("10Gi"),
		},
		Requests: v1.ResourceList{
			v1.ResourceCPU:              cr.RequestsCPU(oversoldRate),
			v1.ResourceMemory:           cr.RequestsMem(oversoldRate),
			v1.ResourceEphemeralStorage: resource.MustParse("5Gi"),
		},
	}
}
