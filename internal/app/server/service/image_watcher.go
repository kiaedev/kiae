package service

import (
	"context"
	"log"
	"strings"

	"github.com/kiaedev/kiae/api/image"
	"github.com/kiaedev/kiae/internal/pkg/klient"
	"github.com/pivotal/kpack/pkg/apis/build/v1alpha2"
	kpack "github.com/pivotal/kpack/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ImageWatcher struct {
	imgSvc *ProjectImageSvc

	kpackCs *kpack.Clientset
}

func NewImageWatcher(imgSvc *ProjectImageSvc, kClients *klient.LocalClients) *ImageWatcher {
	iw := &ImageWatcher{
		imgSvc: imgSvc,

		kpackCs: kClients.KpackCs,
	}
	// go iw.checkNotDoneStatus(context.Background())
	return iw
}

func (i *ImageWatcher) OnAdd(obj interface{}) {
	newImg := obj.(*v1alpha2.Image)
	ctx := context.Background()
	i.syncImageStatus(ctx, newImg.Name, newImg.Status.LatestBuildRef)
}

func (i *ImageWatcher) OnUpdate(oldObj, newObj interface{}) {
	newImg := newObj.(*v1alpha2.Image)
	ctx := context.Background()
	i.syncImageStatus(ctx, newImg.Name, newImg.Status.LatestBuildRef)
}

func (i *ImageWatcher) OnDelete(obj interface{}) {

}

func (i *ImageWatcher) syncImageStatus(ctx context.Context, imageName, buildName string) {
	build, err := i.kpackCs.KpackV1alpha2().Builds("kiae-builder").Get(ctx, buildName, metav1.GetOptions{})
	if err != nil {
		return
	}

	var done int
	var failed bool
	for _, state := range build.Status.StepStates {
		if state.Terminated != nil && state.Terminated.ExitCode == 0 {
			done++
		}
		if state.Terminated != nil && state.Terminated.ExitCode > 0 {
			failed = true
			break
		}
	}

	status := image.Image_BUILDING
	if done > 0 && done == len(build.Status.StepStates) {
		status = image.Image_PUBLISHED
	} else if failed {
		status = image.Image_FAILED
	}

	if _, err := i.imgSvc.UpdateStatus(ctx, strings.TrimPrefix(imageName, "kiae-image-"), status); err != nil {
		log.Printf("failed to update status: %v", err)
	}
}
