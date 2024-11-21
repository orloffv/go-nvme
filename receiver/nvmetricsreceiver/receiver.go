package nvmetricsreceiver

import (
	"context"
	"fmt"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
	"go.opentelemetry.io/collector/model/pdata"
	"go.uber.org/zap"
	"github.com/dswarbrick/go-nvme/nvme"
)

type nvmeMetricsReceiver struct {
	logger   *zap.Logger
	config   *Config
	consumer consumer.MetricsConsumer
	cancel   context.CancelFunc
}

func newNVMeMetricsReceiver(logger *zap.Logger, config *Config, consumer consumer.MetricsConsumer) (component.MetricsReceiver, error) {
	return &nvmeMetricsReceiver{
		logger:   logger,
		config:   config,
		consumer: consumer,
	}, nil
}

func (r *nvmeMetricsReceiver) Start(ctx context.Context, host component.Host) error {
	ctx, cancel := context.WithCancel(ctx)
	r.cancel = cancel

	go func() {
		// Collect metrics periodically
		for {
			select {
			case <-ctx.Done():
				return
			default:
				r.collectMetrics()
			}
		}
	}()

	return nil
}

func (r *nvmeMetricsReceiver) Shutdown(ctx context.Context) error {
	if r.cancel != nil {
		r.cancel()
	}
	return nil
}

func (r *nvmeMetricsReceiver) collectMetrics() {
	// Example of collecting NVMe metrics
	device := nvme.NewNVMeDevice("/dev/nvme0")
	if err := device.Open(); err != nil {
		r.logger.Error("Failed to open NVMe device", zap.Error(err))
		return
	}
	defer device.Close()

	controller, err := device.IdentifyController(nil)
	if err != nil {
		r.logger.Error("Failed to identify NVMe controller", zap.Error(err))
		return
	}

	// Log the collected metrics
	r.logger.Info("NVMe Controller", zap.String("ModelNumber", controller.ModelNumber), zap.String("SerialNumber", controller.SerialNumber))

	// Convert the collected metrics to OpenTelemetry metrics and send to the consumer
	metrics := pdata.NewMetrics()
	rm := metrics.ResourceMetrics().AppendEmpty()
	ilms := rm.InstrumentationLibraryMetrics().AppendEmpty()
	metricsSlice := ilms.Metrics()

	// Add NVMe controller metrics
	addMetric(metricsSlice, "nvme.controller.vendor_id", pdata.MetricDataTypeIntGauge, int64(controller.VendorID))
	addMetric(metricsSlice, "nvme.controller.model_number", pdata.MetricDataTypeString, controller.ModelNumber)
	addMetric(metricsSlice, "nvme.controller.serial_number", pdata.MetricDataTypeString, controller.SerialNumber)
	addMetric(metricsSlice, "nvme.controller.firmware_version", pdata.MetricDataTypeString, controller.FirmwareVersion)
	addMetric(metricsSlice, "nvme.controller.oui", pdata.MetricDataTypeIntGauge, int64(controller.OUI))
	addMetric(metricsSlice, "nvme.controller.max_data_xfer_size", pdata.MetricDataTypeIntGauge, int64(controller.MaxDataXferSize))

	if err := r.consumer.ConsumeMetrics(context.Background(), metrics); err != nil {
		r.logger.Error("Failed to consume metrics", zap.Error(err))
	}
}

func addMetric(metrics pdata.MetricSlice, name string, dataType pdata.MetricDataType, value interface{}) {
	metric := metrics.AppendEmpty()
	metric.SetName(name)
	metric.SetDataType(dataType)

	switch dataType {
	case pdata.MetricDataTypeIntGauge:
		metric.IntGauge().DataPoints().AppendEmpty().SetValue(value.(int64))
	case pdata.MetricDataTypeString:
		metric.StringDataPoints().AppendEmpty().SetValue(value.(string))
	}
}
