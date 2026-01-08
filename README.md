# Kubernetes GitOps & Observability Demo

This project demonstrates a full-scale cloud-native environment, focusing on a **GitOps-based** CI/CD pipeline (ArgoCD) and a complete **Observability** stack (Prometheus, Grafana). It also includes the development and monitoring of a custom Go application.

## Technologies
* **Kubernetes** (Orchestration)
* **ArgoCD** (GitOps / Continuous Delivery)
* **Prometheus** (Monitoring & Alerting)
* **Grafana** (Visualization)
* **Docker & Go** (Application Development)

---

## Key Features

### 1. Prometheus & ArgoCD Integration
Monitoring the internal metrics of the CD pipeline.
* **Scrape Job:** Configured a new scrape job in Prometheus.
* **Service Discovery:** Utilized Kubernetes internal DNS resolution for cross-namespace communication.
    * Target: `argocd-server-metrics.argocd.svc.cluster.local:8083`
* **Result:** The ArgoCD target successfully reached **UP** state in the Prometheus Target list.

### 2. Grafana Dashboard & Data Source
Visualization of the collected metrics.
* **GitOps Deployment:** Added Grafana manifests to the repository and deployed via ArgoCD.
* **Auto-Provisioning:** Grafana is configured to automatically discover and connect to Prometheus as its Data Source upon startup, eliminating manual configuration.

### 3. Custom Go Application (Custom Exporter)
Development of a custom Go web server that exposes Prometheus-compatible metrics.
* **Instrumentation:** Implemented a specific HTTP handler for metric exposition:
    ```go
    http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) { ... }
    ```
    This endpoint returns data in the standard *Prometheus Exposition Format* (e.g., `go_app_status`, `go_app_random_value`).
* **Build & Deploy:**
    * Built and pushed the Docker image (v2).
    * Created a **Kubernetes Service** to provide stable network access instead of relying on ephemeral Pod IPs.
* **Monitoring:** Updated the Prometheus configuration with a `go-exporter` job using `static_configs` to scrape the application.

---

## Project Structure

Files are organized by function to ensure maintainability:

* **`/app`**: Go source code (`main.go`) and `Dockerfile`.
* **`/k8s`**: Kubernetes manifests (Deployments, Services, ConfigMaps).

---

## Usage

1.  **ArgoCD Sync:** All changes are applied via the ArgoCD interface.
2.  **Port Forwarding:** Access services locally:
    * *Grafana:* `kubectl port-forward svc/grafana 3000:3000`
    * *Prometheus:* `kubectl port-forward svc/prometheus 9090:9090`
    * *Go App:* `kubectl port-forward svc/kezdo-app-v2 8081:8080`
