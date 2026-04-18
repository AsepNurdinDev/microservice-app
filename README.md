# Microservice App 🌿

Aplikasi microservice berbasis Go dengan clean architecture, siap untuk production deployment di Ubuntu + Nginx, Kubernetes, monitoring Prometheus + Grafana, dan CI/CD via GitHub Actions.

## Arsitektur

```
Client
  │
  ▼
┌─────────────────────┐
│   Nginx (port 443)  │  ← TLS termination, rate limiting
└─────────┬───────────┘
          │
          ▼
┌─────────────────────┐
│  Gateway Service    │  ← JWT validation, reverse proxy (port 8080)
│     :8080           │
└────┬──────────┬─────┘
     │          │
     ▼          ▼
┌─────────┐ ┌──────────────┐
│  Auth   │ │   Article    │
│ Service │ │   Service    │
│  :8001  │ │    :8002     │
│         │ │              │
│PostgreSQL│ │   MongoDB   │
└─────────┘ └──────────────┘

Observability:
  Prometheus :9090 → scrapes /metrics dari semua service
  Grafana    :3000 → dashboard visualisasi
```

## Struktur Project

```
microservice-app/
├── services/
│   ├── auth-service/          # Autentikasi (Register/Login/JWT)
│   │   ├── cmd/main.go
│   │   ├── internal/
│   │   │   ├── domain/        # Entity + error domain
│   │   │   ├── repository/    # Interface repository
│   │   │   ├── usecase/       # Business logic
│   │   │   ├── delivery/http/ # HTTP handler
│   │   │   ├── infrastructure/# DB impl, JWT util
│   │   │   ├── middleware/    # Prometheus metrics
│   │   │   └── mocks/         # Testify mocks
│   │   ├── migrations/        # SQL migrations
│   │   └── test/
│   │
│   ├── article-service/       # CRUD artikel (MongoDB)
│   │   └── ... (struktur sama)
│   │
│   └── gateway-service/       # API gateway, JWT guard, reverse proxy
│       └── ...
│
├── deployments/
│   ├── kubernetes/            # K8s manifests (Deployment, Service, HPA, Ingress)
│   └── nginx/nginx.conf       # Nginx reverse proxy config
│
├── monitoring/
│   ├── prometheus/            # prometheus.yml scrape config
│   └── grafana/dashboards/    # Dashboard JSON + provisioning
│
├── .github/workflows/         # GitHub Actions CI/CD
├── scripts/
│   ├── deploy.sh              # Deploy ke Ubuntu server
│   └── k8s-apply.sh           # Apply K8s manifests
└── docker-compose.yml         # Local dev + staging
```

## Quick Start (Local Development)

### Prasyarat
- Docker Desktop / Docker Engine + Docker Compose v2
- Go 1.22+ (untuk development)

### 1. Setup environment

```bash
cp .env.example .env
# Edit .env — minimal ganti JWT_SECRET dengan string acak panjang
# Generate: openssl rand -hex 32
```

### 2. Jalankan semua service

```bash
docker compose up --build
```

Service akan tersedia di:
| Service      | URL                          |
|--------------|------------------------------|
| Gateway API  | http://localhost:8080        |
| Auth Service | http://localhost:8001        |
| Article Svc  | http://localhost:8002        |
| Prometheus   | http://localhost:9090        |
| Grafana      | http://localhost:3000        |

### 3. Test API

```bash
# Register
curl -X POST http://localhost:8080/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"password123"}'

# Login
TOKEN=$(curl -s -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"password123"}' \
  | jq -r '.access_token')

# Buat artikel
curl -X POST http://localhost:8080/articles \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"title":"Judul Artikel","content":"Konten artikel yang panjang..."}'

# List artikel
curl http://localhost:8080/articles?page=1&limit=10 \
  -H "Authorization: Bearer $TOKEN"
```

## Menjalankan Unit Tests

```bash
# Semua service sekaligus
for svc in auth-service article-service gateway-service; do
  echo "=== Testing $svc ==="
  cd services/$svc
  go test ./... -v -race -cover
  cd ../..
done

# Satu service saja
cd services/auth-service
go test ./... -v -cover
```

## Deploy ke Ubuntu Server + Nginx

### 1. Siapkan server

```bash
# Install Docker
curl -fsSL https://get.docker.com | sh
sudo usermod -aG docker $USER

# Install Nginx
sudo apt install -y nginx

# Copy nginx config
sudo cp deployments/nginx/nginx.conf /etc/nginx/sites-available/microservice-app
sudo ln -s /etc/nginx/sites-available/microservice-app /etc/nginx/sites-enabled/
sudo nginx -t && sudo systemctl reload nginx
```

### 2. Deploy aplikasi

```bash
# Clone project ke server
git clone <repo-url> /opt/microservice-app
cd /opt/microservice-app

# Setup environment
cp .env.example .env
# Edit .env dengan credential production

# Deploy
./scripts/deploy.sh production
```

### 3. Setup TLS dengan Let's Encrypt

```bash
sudo apt install -y certbot python3-certbot-nginx
sudo certbot --nginx -d your-domain.com
```

## Deploy ke Kubernetes

### 1. Edit image registry

Ganti `your-registry/` di file `deployments/kubernetes/03-*.yml` dengan registry Anda:
```bash
sed -i 's|your-registry|ghcr.io/your-org|g' deployments/kubernetes/*.yml
```

### 2. Buat secrets

```bash
# Generate JWT secret
JWT=$(openssl rand -hex 32)
kubectl create secret generic app-secrets \
  --from-literal=jwt-secret="$JWT" \
  --from-literal=db-password="strong-password" \
  --from-literal=db-user="user" \
  -n microservice-app
```

### 3. Apply manifests

```bash
./scripts/k8s-apply.sh apply
```

### 4. Verifikasi

```bash
kubectl get pods -n microservice-app
kubectl get hpa -n microservice-app
kubectl get ingress -n microservice-app
```

## CI/CD (GitHub Actions)

Pipeline berjalan otomatis pada:
- **Push ke `develop`**: Test + Build image
- **Push ke `main`**: Test + Build image + Deploy ke Kubernetes

### Setup Secrets di GitHub

Buka Settings → Secrets → Actions:
| Secret         | Keterangan                              |
|----------------|-----------------------------------------|
| `KUBECONFIG`   | Base64 dari kubeconfig file             |
| `CODECOV_TOKEN`| Token dari codecov.io (opsional)        |

```bash
# Encode kubeconfig
cat ~/.kube/config | base64 -w 0
```

## Monitoring

### Prometheus
- URL: http://localhost:9090
- Metrics tersedia di `/metrics` setiap service

### Grafana
- URL: http://localhost:3000
- Login: admin / admin (ganti di `.env`)
- Dashboard "Microservices Overview" sudah tersedia otomatis

### Metrics yang tersedia
| Metric                               | Keterangan                    |
|--------------------------------------|-------------------------------|
| `http_requests_total`                | Total request per service     |
| `http_request_duration_seconds`      | Latency histogram per service |
| `gateway_requests_total`             | Request masuk ke gateway      |
| `gateway_request_duration_seconds`   | Latency di gateway            |

## API Reference

### Auth Service

| Method | Path              | Auth | Keterangan            |
|--------|-------------------|------|-----------------------|
| POST   | `/auth/register`  | ❌   | Register user baru    |
| POST   | `/auth/login`     | ❌   | Login, dapat JWT      |
| GET    | `/health`         | ❌   | Health check          |
| GET    | `/metrics`        | ❌   | Prometheus metrics    |

### Article Service

| Method | Path              | Auth | Keterangan              |
|--------|-------------------|------|-------------------------|
| POST   | `/articles`       | ✅   | Buat artikel baru       |
| GET    | `/articles`       | ✅   | List artikel (paginasi) |
| GET    | `/articles/:id`   | ✅   | Ambil artikel by ID     |
| DELETE | `/articles/:id`   | ✅   | Hapus artikel           |
| GET    | `/health`         | ❌   | Health check            |
| GET    | `/metrics`        | ❌   | Prometheus metrics      |

## Security Checklist

- [x] JWT secret dibaca dari env var, tidak hardcoded
- [x] Password di-hash dengan bcrypt (cost=10)
- [x] Container berjalan sebagai non-root (UID 65534)
- [x] Docker image menggunakan `FROM scratch` (attack surface minimal)
- [x] `readOnlyRootFilesystem: true` di K8s
- [x] Nginx: security headers, rate limiting, TLS-only
- [x] `/metrics` endpoint diblokir dari akses publik via Nginx
- [ ] Ganti JWT_SECRET dengan nilai random di production
- [ ] Aktifkan TLS di Nginx dengan Let's Encrypt
- [ ] Gunakan Sealed Secrets atau Vault untuk K8s secrets