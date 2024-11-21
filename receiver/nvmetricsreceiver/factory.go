package nvmetricsreceiver

import (
	"context"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configmodels"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
)

// NewFactory creates a factory for NVMe metrics receiver.
func NewFactory() component.ReceiverFactory {
	return receiverhelper.NewFactory(
		"nvmetricsreceiver",
		createDefaultConfig,
		receiverhelper.WithMetrics(createMetricsReceiver),
	)
}

// createDefaultConfig creates the default configuration for the NVMe metrics receiver.
func createDefaultConfig() configmodels.Receiver {
	return &Config{
		ReceiverSettings: configmodels.ReceiverSettings{
			TypeVal: "nvmetricsreceiver",
			NameVal: "nvmetricsreceiver",
		},
	}
}

// createMetricsReceiver creates a metrics receiver based on provided config.
func createMetricsReceiver(
	_ context.Context,
	params component.ReceiverCreateParams,
	cfg configmodels.Receiver,
	consumer consumer.MetricsConsumer,
) (component.MetricsReceiver, error) {
	rCfg := cfg.(*Config)
	return newNVMeMetricsReceiver(params.Logger, rCfg, consumer)
}
