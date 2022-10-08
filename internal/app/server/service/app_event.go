package service

import (
	"context"
	"encoding/json"
	"sort"
	"strconv"
	"time"

	"github.com/kiaedev/kiae/api/graph/model"
	"github.com/kiaedev/kiae/pkg/loki"
)

type AppEventService struct {
	loki *loki.Client
}

func NewAppEventService(loki *loki.Client) *AppEventService {
	return &AppEventService{loki: loki}
}

func (s *AppEventService) List(ctx context.Context, query string) ([]*model.Event, error) {
	results, err := s.loki.QueryRange(query, 2000, time.Now().Add(-24*15*time.Hour), time.Now(), "")
	if err != nil {
		return nil, err
	}

	events := make(Events, 0)
	for _, result := range results {
		events = append(events, s.formatEvents(result.Stream, result.Values)...)
	}

	sort.Sort(events)
	return events.Events(), nil
}

type EValue struct {
	Body       string `json:"body"`
	Severity   string `json:"severity"`
	Attributes struct {
		K8SEventAction    string `json:"k8s.event.action"`
		K8SEventCount     int    `json:"k8s.event.count"`
		K8SEventName      string `json:"k8s.event.name"`
		K8SEventReason    string `json:"k8s.event.reason"`
		K8SEventStartTime string `json:"k8s.event.start_time"`
		K8SEventUid       string `json:"k8s.event.uid"`
		K8SNamespaceName  string `json:"k8s.namespace.name"`
	} `json:"attributes"`
	Resources struct {
		K8SNodeName              string `json:"k8s.node.name"`
		K8SObjectApiVersion      string `json:"k8s.object.api_version"`
		K8SObjectFieldpath       string `json:"k8s.object.fieldpath"`
		K8SObjectKind            string `json:"k8s.object.kind"`
		K8SObjectName            string `json:"k8s.object.name"`
		K8SObjectResourceVersion string `json:"k8s.object.resource_version"`
		K8SObjectUid             string `json:"k8s.object.uid"`
	} `json:"resources"`
}

func (s *AppEventService) formatEvents(stream loki.Stream, values [][]string) Events {
	events := make(Events, 0)
	for _, value := range values {
		nanoUnix, _ := strconv.ParseInt(value[0], 10, 64)

		var ev EValue
		_ = json.Unmarshal([]byte(value[1]), &ev)
		events = append(events, eEvent{
			Event: &model.Event{
				UID:  ev.Attributes.K8SEventUid,
				Name: ev.Attributes.K8SEventName,
				InvolvedObject: &model.InvolvedObject{
					Kind:      ev.Resources.K8SObjectKind,
					Namespace: ev.Attributes.K8SNamespaceName,
					Name:      ev.Resources.K8SObjectName,
				},
				Reason:    ev.Attributes.K8SEventReason,
				Message:   ev.Body,
				Type:      ev.Severity,
				Count:     ev.Attributes.K8SEventCount,
				StartedAt: time.Unix(0, nanoUnix).Format(time.RFC3339),
			},
			LastTimestamp: time.Unix(0, nanoUnix),
		})
	}
	return events
}

type eEvent struct {
	*model.Event

	LastTimestamp time.Time
}

type Events []eEvent

func (es Events) Events() []*model.Event {
	events := make([]*model.Event, len(es))
	for idx, e := range es {
		events[idx] = e.Event
	}
	return events
}

func (es Events) Len() int {
	return len(es)
}

func (es Events) Less(i, j int) bool {
	return es[j].LastTimestamp.Before(es[i].LastTimestamp)
}

func (es Events) Swap(i, j int) {
	es[i], es[j] = es[j], es[i]
}
