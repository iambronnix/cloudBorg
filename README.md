# CloudBorg: Distributed Task Orchestrator (Prototype)

> **Status:** 🚧 Early-Stage Architectural Prototype. 
> The core interfaces, data structures, and concurrency models are in place. The networking and Docker execution layers are currently in development.

CloudBorg is a lightweight distributed task scheduling engine written in Go, inspired by Google's Borg and Kubernetes. This project is my deep dive into distributed systems architecture, focusing on how to safely manage state across multiple nodes.

## 🧠 System Architecture

The project is designed with a modular architecture to separate concerns:

*   **`manager`**: The master node. Holds the global state, manages the task queue, and handles scheduling logic.
*   **`worker`**: The agent that executes tasks. (Currently simulates task execution; Docker integration is on the roadmap).
*   **`scheduler`**: An extensible interface. The current implementation uses a simple `RoundRobin` strategy, but the interface allows for plug-and-play algorithms (e.g., Resource-Aware scheduling).
*   **`task` & `node`**: The core domain models representing workloads and compute resources.

## 🛡️ Engineering Highlights (Why this prototype matters)

Even in its early stages, this codebase implements production-grade Go patterns:

1.  **Concurrency Safety:** Distributed systems are inherently concurrent. I have embedded `sync.Mutex` into the `Manager` and `Worker` structs to prevent data races when multiple goroutines access `TaskDb` or `WorkerTaskMap`.
2.  **Interface Segregation:** The `Scheduler` interface allows the Manager to remain agnostic to *how* a node is picked. You can swap a Round-Robin scheduler for a complex scoring algorithm without changing the Manager's code.
3.  **State Machine:** Tasks move through explicit states (`Pending` -> `Scheduled` -> `Running` -> `Completed`), laying the foundation for a reliable event-driven system.

## 🗺️ Current Progress & Roadmap

*   [x] Define core domain models (`Task`, `TaskEvent`, `Node`).
*   [x] Implement thread-safe Manager and Worker state.
*   [x] Create `Scheduler` interface with a Round-Robin implementation.
*   [x] Simulate task assignment from Manager to Worker.
*   [ ] **Phase 2 (Next):** Replace in-memory queues with an HTTP/gRPC API so Manager and Worker can communicate over a network.
*   [ ] **Phase 3:** Integrate the Docker SDK to actually spin up containers inside the `Worker.RunTask()` method.
*   [ ] **Phase 4:** Implement heartbeat health checks to detect offline workers.

## 🚀 How to Run the Simulation

*(Note: Networking is currently mocked. The simulation runs in a single process).*

1. Clone the repo:
   ```bash
   git clone https://github.com/iamerick/cloudBorg.git
   cd cloudBorg
