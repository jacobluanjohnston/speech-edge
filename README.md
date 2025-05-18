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

Notes to self:
- https://training.confluent.io/
- https://kafka.apache.org/documentation
- https://hub.docker.com/r/confluentinc/cp-zookeeper
- virtualenv, added `source /home/admin/ina219-env/bin/activate` into .bashrc
