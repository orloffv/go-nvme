package nvmetricsreceiver

import (
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/config/configmodels"
)

// Config defines configuration for NVMe metrics receiver.
type Config struct {
	configmodels.ReceiverSettings `mapstructure:",squash"`
	// Add any necessary configuration options here
}
