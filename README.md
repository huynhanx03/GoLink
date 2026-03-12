# GoLink - URL Shortener
GoLink is a URL shortener service built with Golang, ScyllaDB, Redis, and Envoy Proxy.

## Contents

I. [System Architecture](#i-system-architecture)

II. [Tech Stack](#ii-tech-stack)

III. [Capacity Estimates](#iii-capacity-estimates)

IV. [How to Run](#iv-how-to-run)

---

## I. System Architecture

### 1. Unique ID Generation: Snowflake Node
Utilizing **Twitter Snowflake** algorithm for unique ID generation **Optimal Configuration (Bit Breakdown)**: Comparison with 7-bit hard limit for Base62 encoding (7 characters).

-   **Total Bits**: `42` bits (Output 7 chars).
-   **Timestamp**: `30` bits (Epoch: 01/01/2026, Unit: Seconds) - System Lifespan: **34 years** (Until 2060).
-   **Node**: `2` bits - Max Scale: **4 nodes**.
-   **Step**: `10` bits - ID Generation Speed: **1,024 req/s/node** (Cluster Capacity: **~4,000 req/s**).

Zero-Latency Short Code Pool:
-   **Strategy**: **Pre-generation** utilizing **MPMC Lock-free Queue**.
-   **Capacity**: **120,000** codes, automatically refilled by background workers.

### 2. Database: ScyllaDB
-   **Performance**: Write-Heavy Optimization (LSM Tree) - Extremely high write throughput, suitable for continuous short link generation.
-   **Architecture**: Peer-to-Peer (Masterless) - No Single Point of Failure (SPOF).
-   **Scalability**: Linear Scalability - Easily add nodes to increase capacity without downtime.

### 3. Caching: Redis
-   **Strategy**: Cache-Aside Pattern - Automatically syncs cache when new data arrives from DB.
-   **Efficiency**: High Hit Ratio - Serves 90% of "Hot Traffic" directly from RAM (< 1ms latency).

### 4. API Gateway: Envoy Proxy
-   **Routing**: L7 Load Balancing - Precise routing to Generation or Redirection Service.
-   **Protection**: Rate Limiting & CORS - Anti-Spam/Abuse (100 req/s/IP).

### 5. Identity Service
Centralized authentication and authorization hub. See [Identity Service Documentation](Identity/README.md) for details.

### 6. Orchestrator Service
Coordinates distributed transactions using the Saga pattern to ensure system-wide consistency. See [Orchestrator Service Documentation](Orchestrator/README.MD) for details.

### 7. Notification Service
Enterprise-grade multi-channel notification engine (Email, Push, SSE, Webhooks). See [Notification Service Documentation](Notification/README.md) for details.

---

## II. Tech Stack

*   **Language**: Golang
*   **Communication**: gRPC, HTTP (Gin)
*   **Database**: ScyllaDB, PostgreSQL, MongoDB
*   **Cache**: Redis, TinyLFU (Local Cache)
*   **Message Queue**: Kafka
*   **Gateway**: Envoy Proxy
*   **Containerization**: Docker & Docker Compose

---

## III. Capacity Estimates

### Benchmark
-   **Write Volume**: 1,000,000 links/day (Avg ~12 TPS, **Peak ~60 TPS**).
-   **Read Volume**: 100,000,000 clicks/day (Avg ~1,200 RPS, **Peak ~6,000 RPS**).
-   **Data Size**: 0.5 KB average/record.

### ID Space
-   **Max Capacity**: ~3.5 Trillion IDs (Base62 7 chars).
-   **Exhaustion Time**: At 1 million links/day -> Takes **~10,000 years** to exhaust.

### Storage
-   **1 Year**: $1M \times 365 \times 0.5KB \approx \textbf{180 GB}$.

### Throughput & Scalability
-   **ScyllaDB Node** (Foundation):
    -   **Storage**: 2TB+ per node (Linear scale with node count).
    -   **Capacity**: ~20,000 Write-QPS / ~15,000 Read-QPS per node (Disk I/O).
-   **Redis** (Accelerator):
    -   **Performance**: >100,000 QPS per instance (Memory-bound).
    -   **Capacity**: Stores **~25 million Hot Keys** (per 16GB Node) -> Easily Scale-out to **Redis Cluster** for billions of keys.
    -   **Role**: Handles 90% Read Traffic, protects DB from Hot Keys.
-   **Total System**: Capable of **40k Write/s** and **100k Read/s**, hundreds of times higher than actual demand.

---

## IV. How to Run

### 1. Prerequisites
- **Docker & Docker Compose**
- **Go 1.22+** (for local development)

### 2. Start Infrastructure
Launch all required databases, message brokers, and proxies:
```bash
make docker-up
```

### 3. Initialize System
Wait until ScyllaDB and other services are ready, then run:
```bash
make init-all
```

### 4. Start Core Services
Run each service in a separate terminal:
```bash
make run-generation
make run-redirection
make run-identity
make run-orchestrator
```

---

## V. Roadmap & Contributing

We are constantly working to improve this system. Future plans include:
- [ ] **Analytics Service**: Real-time traffic analysis and click-through rate (CTR) tracking.
- [ ] **Project Website**: A modern dashboard for managing links and viewing analytics.
- [ ] **Monitoring & Observability**: Integrated tracing (Jaeger) and structured logging dashboards.

> **Interested in contributing?**
> We welcome all contributions! Whether it's fixing bugs, improving docs, or proposing new features, feel free to submit a Pull Request.

---

## VI. License

This project is licensed under the Apache License 2.0. See the [LICENSE](LICENSE) file for the full text.
