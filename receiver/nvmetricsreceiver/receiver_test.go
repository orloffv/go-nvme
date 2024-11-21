package nvmetricsreceiver

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configmodels"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.uber.org/zap"
)

func TestNewFactory(t *testing.T) {
	factory := NewFactory()
	assert.NotNil(t, factory)
}

func TestCreateDefaultConfig(t *testing.T) {
	cfg := createDefaultConfig()
	assert.NotNil(t, cfg)
	assert.Equal(t, "nvmetricsreceiver", cfg.Type())
	assert.Equal(t, "nvmetricsreceiver", cfg.Name())
}

func TestCreateMetricsReceiver(t *testing.T) {
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig().(*Config)
	params := component.ReceiverCreateParams{Logger: zap.NewNop()}
	consumer := consumertest.NewNop()

	receiver, err := factory.CreateMetricsReceiver(context.Background(), params, cfg, consumer)
	assert.NoError(t, err)
	assert.NotNil(t, receiver)
}

func TestNVMeMetricsReceiver_Start(t *testing.T) {
	cfg := createDefaultConfig().(*Config)
	params := component.ReceiverCreateParams{Logger: zap.NewNop()}
	consumer := consumertest.NewNop()

	receiver, err := newNVMeMetricsReceiver(params.Logger, cfg, consumer)
	assert.NoError(t, err)
	assert.NotNil(t, receiver)

	err = receiver.Start(context.Background(), nil)
	assert.NoError(t, err)

	err = receiver.Shutdown(context.Background())
	assert.NoError(t, err)
}
