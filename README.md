# speech-edge

Low-power speech-to-text edge router with local/cloud cascades,
 Kafka transport, OpenTelemetry metrics, and k3s deploys.

Written in **Rust**, **Go**, and **Python**.

Personal long-term side project to explore real-time observability,
power tradeoffs, and production-grade infra.

---

## Overview

- Rust (`edge-router`) - low-latency GPIO handling async Kafka producer
- Go (`stt-wrapper`) - simple cgo bindings to `whisper.cpp`
- Python (`power-sensor`) - fast prototyping for INA219 power sensor & metrics
- Kafka + k3s - back-pressure support, self-contained cluster demos
- OpenTelemetry + Grafana - end-to-end traces + p95/p99 latency tracking

---

## Prerequisites

Install core toolchains via Homebrew:

```bash
brew install go python@3.11 rustup-init kafkacat docker docker-compose
```

Then launch Docker:

```bash
open -a Docker
docker run --rm hello-world
docker compose version
```

---

## Language Setup

### Rust

Install via rustup:

```bash
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y
source ~/.cargo/env
rustup component add clippy
```

Quick sanity check:

```bash
cargo new hello && cd hello && cargo clippy
```

### Python

Use `pyenv` for version control:

```bash
pyenv install 3.11.9
pyenv global 3.11.9
python --version
```

### Go

No extra setup needed beyond `brew install go`:

```bash
go version
```

---

## Code Hygiene with pre-commit

Create a `.pre-commit-config.yaml` file with the necessary hooks.
Then run:

```bash
pre-commit install
pre-commit run --all-files
```

This installs and runs the following linters:

- **Python** - `black`, `ruff`, `mypy`
- **Rust** - `clippy`, `rustfmt`
- **Go** - `go-fmt`, `go-vet`, `go-test-mod`, `go-mod-tidy`

---

## Kafka End-to-End Test

To verify your Kafka + Zookeeper setup:
Install kafka for CLI

```bash
brew install kafka
```

Then run:

```bash
docker exec -t kafka kafka-topics --create \
--topic test-topic \
--bootstrap-server localhost:9092 \
--replication-factor 1 \
--partitions 1
```

Start a consumer:

```bash
kafka-console-consumer --bootstrap-server localhost:9092 \
--topic test-topic --from-beginning

```

Then, in another terminal, start a producer:

```bash
kafka-console-producer --bootstrap-server localhost:9092 \
--topic test-topic```

Type something in the producer terminal and make sure
 it appears on the consumer terminal.

## Project Status

This project is still in its early stages. It serves as a personal sandbox to understand:

- Adaptive local/cloud cascades for ML inference
- Distributed systems & observability
- Power/performance tradeoffs on edge devices
- Multi-language CI/CD and pre-commit pipelines

---

## Project Structure (planned)

```

speech-edge/
--- edge-router/      # Rust: local audio handling + Kafka producer
--- stt-wrapper/      # Go: Whisper ASR wrapper + stream bridge
--- power-sensor/     # Python: INA219 polling + metrics exporter
--- docker-compose.yml
--- README.md```
