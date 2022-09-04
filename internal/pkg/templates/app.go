package templates

import (
	"fmt"

	"github.com/kiaedev/kiae/api/app"
	"github.com/kiaedev/kiae/api/kiae"
	"github.com/kiaedev/kiae/api/project"
	"github.com/kiaedev/kiae/pkg/kiaeutil"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
	"github.com/saltbo/gopkg/strutil"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

type Application struct {
	Name        string
	Image       string
	Replicas    uint32
	Resources   v1.ResourceRequirements
	Ports       []*project.Port
	ConfigPaths []string
	Traits      []common.ApplicationTrait
	Middlewares []common.ApplicationComponent
}

func NewApplication(app *app.Application, proj *project.Project, traits []*kiae.Trait) (*v1beta1.Application, error) {
	appTmplModel := &Application{
		Name:        app.Name,
		Image:       app.Image,
		Ports:       proj.Ports,
		Replicas:    app.Replicas,
		Resources:   buildResources(app.Size, 0.5),
		ConfigPaths: buildMountPaths(kiaeutil.ConfigsMerge(proj.Configs, app.Configs)),
		Traits:      buildTraits(traits),
		Middlewares: buildMiddlewares(proj.Middlewares),
	}

	var oam v1beta1.Application
	err := New("app").Render(appTmplModel, &oam)
	if err != nil {
		return nil, err
	}

	return &oam, nil
}

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

func buildResources(size app.Size, oversoldRate float64) v1.ResourceRequirements {
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

func buildMountPaths(configs []*project.Configuration) []string {
	paths := make([]string, 0, len(configs))
	for _, cfg := range configs {
		if strutil.StrInSlice(cfg.MountPath, paths) {
			continue
		}

		paths = append(paths, cfg.MountPath)
	}

	return paths
}

func buildTraits(traits []*kiae.Trait) []common.ApplicationTrait {
	finalTraits := make([]common.ApplicationTrait, 0)
	for _, traitItem := range traits {
		finalTraits = append(finalTraits, common.ApplicationTrait{
			Type:       traitItem.Type,
			Properties: util.Object2RawExtension(traitItem.Properties),
		})

	}
	return finalTraits
}

func buildMiddlewares(middlewares []*project.Middleware) []common.ApplicationComponent {
	var res []common.ApplicationComponent
	// for _, middleware := range middlewares {
	// 	res = append(res, common.ApplicationComponent{
	// 		Name:       middleware.Name,
	// 		Type:       middleware.Type,
	// 		Properties: util.Object2RawExtension(middleware.Properties),
	// 	})
	// }
	return res
}
