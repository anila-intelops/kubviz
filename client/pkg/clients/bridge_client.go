package clients

import (
	"encoding/json"
	"log"

	"github.com/intelops/kubviz/client/pkg/clickhouse"
	"github.com/nats-io/nats.go"
)

type GitBridge string

const (
	bridgeSubjects GitBridge = "GITMETRICS.*"
	bridgeSubject  GitBridge = "GITMETRICS.git"
	bridgeConsumer GitBridge = "Git-Consumer"
)

func (n *NATSContext) SubscribeGitBridgeNats(conn clickhouse.DBInterface) {
	n.stream.Subscribe(string(bridgeSubject), func(msg *nats.Msg) {
		type pubData struct {
			Metrics json.RawMessage `json:"metrics"`
			Repo    string          `json:"repo"`
			Event   string          `json:"event"`
		}
		msg.Ack()
		repo := msg.Header.Get("repo")
		event := msg.Header.Get("event")
		metrics := &pubData{
			Metrics: json.RawMessage(msg.Data),
			Repo:    repo,
			Event:   event,
		}
		// metrics := &models.Gitevent{}
		data, err := json.Marshal(metrics)
		if err != nil {
			log.Fatal(err)
		}
		conn.InsertGitEvent(string(data))
		log.Println("Inserted Git metrics:", string(data))
	}, nats.Durable(string(bridgeConsumer)), nats.ManualAck())
}
