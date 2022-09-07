package structs

import "google.golang.org/api/pubsub/v1"

// Pubsubから受信するデータの構造体
type PushRequest struct {
	Message      pubsub.PubsubMessage `json:"message"`
	Subscription string               `json:"subscription"`
}