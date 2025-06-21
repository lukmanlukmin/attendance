// Package event ...
package event

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// BuildKafkaPayload ...
func BuildKafkaPayload(data interface{}, topic string) (string, error) {
	hostName, err := os.Hostname()
	if err != nil {
		return "", fmt.Errorf("failed to get hostname: %w", err)
	}

	rawJSON, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal payload: %w", err)
	}

	hash, err := hashPayload(rawJSON)
	if err != nil {
		return "", fmt.Errorf("failed to hash payload: %w", err)
	}

	now := time.Now()

	payload := MessageFormat{
		Data: json.RawMessage(rawJSON),
		Metadata: Metadata{
			EmitHost:  hostName,
			EmitTime:  now.Unix(),
			Event:     topic,
			Hash:      hash,
			Timestamp: now,
		},
	}

	result, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal kafka message: %w", err)
	}

	return string(result), nil
}

func hashPayload(data []byte) (string, error) {
	hash := sha256.Sum256(data)
	return base64.StdEncoding.EncodeToString(hash[:]), nil
}
