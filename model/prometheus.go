package model

import "time"

type Alert struct {
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
	StartsAt    time.Time         `json:"startsAt"`
	EndsAt      time.Time         `json:"endsAt"`
}

type Notification struct {
	Version           string            `json:"version"`
	// key identifying the group of alerts (e.g. to deduplicate)
	GroupKey          string            `json:"groupKey"`
	// <resolved|firing>
	Status            string            `json:"status"`
	Receiver          string            `json:"receiver"`
	GroupLabels       map[string]string `json:"groupLabels"`
	CommonLabels      map[string]string `json:"commonLabels"`
	CommonAnnotations map[string]string `json:"commonAnnotations"`
	// backlink to the Alertmanager.
	ExternalURL       string            `json:"externalURL"`
	Alerts            []Alert           `json:"alerts"`
}
