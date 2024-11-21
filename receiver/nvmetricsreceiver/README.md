# NVMe Metrics Receiver

This receiver collects metrics from NVMe devices and integrates with the OpenTelemetry Collector.

## Overview

The NVMe Metrics Receiver is designed to collect metrics from NVMe devices and export them to the OpenTelemetry Collector. This receiver is based on the template from the nopreceiver in the OpenTelemetry Collector.

## Configuration

To use the NVMe Metrics Receiver, add the following configuration to your OpenTelemetry Collector configuration file:

```yaml
receivers:
  nvmetricsreceiver:
    # Configuration options for the NVMe Metrics Receiver
    # Add any necessary configuration options here

service:
  pipelines:
    metrics:
      receivers: [nvmetricsreceiver]
      exporters: [your_exporter]
```

## Metrics

The NVMe Metrics Receiver collects the following metrics from NVMe devices:

- NVMe controller information
- NVMe namespace information
- NVMe SMART attributes

## Usage

1. Clone the repository and navigate to the `receiver/nvmetricsreceiver` directory.
2. Build the receiver using the Go toolchain.
3. Add the NVMe Metrics Receiver to your OpenTelemetry Collector configuration file as shown in the Configuration section.
4. Run the OpenTelemetry Collector with the updated configuration.

## License

This project is licensed under the Apache License 2.0. See the [LICENSE](../../LICENSE) file for details.
