# speech-edge

Low-power speech-to-text edge router with local/cloud cascades,
 Kafka transport, OpenTelemetry metrics, and k3s deploys.

![Kafka Mode: ZooKeeper](https://img.shields.io/badge/kafka--mode-zookeeper-blue)
![CI: GitHub Actions](https://github.com/jacobluanjohnston/speech-edge/actions/workflows/test.yaml/badge.svg)
![Languages: Rust, Go, Python](https://img.shields.io/badge/languages-rust--go--python-yellowgreen)

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

## Other Notes

### Kafka Deployment Notes

This project uses Kafka with ZooKeeper mode for CI and local development to ensure
 maximum compatibility and reliability across environments.

KRaft (Kafka Raft Mode)

- Kafka is transitioning towards a ZooKeeper-less architecture known as KRaft mode,
 which offers simplified deployment, faster start-up, and a native quorum-based
 controller management
- This project does not currently use KRaft by default. Future branches may include
 KRaft support.
