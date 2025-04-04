# 🏷️ FullCycle Auction API

Este é um clone do sistema de leilões desenvolvido em Go, utilizando MongoDB como base de dados e o framework Gin para exposição da API.
Fonte: https://github.com/devfullcycle/labs-auction-goexpert

O foco deste projeto é apresentar a funcionalidade de **fechamento automático de leilões**, após um tempo definido via variável de ambiente `AUCTION_INTERVAL`.

---

## 🚀 Como rodar o projeto

1. **Clonar o repositório**

```bash
git clone https://github.com/jeancarlosdanese/fullcycle-auction-go.git
cd fullcycle-auction-go
```

2. **Configurar variáveis no `.env`**

Crie o arquivo `cmd/auction/.env` com o seguinte conteúdo:

```env
MONGODB_URL=mongodb://mongodb:27017
MONGODB_DB=auction_db
AUCTION_INTERVAL=20s
```

> A variável `AUCTION_INTERVAL` define o tempo de duração de cada leilão.

3. **Subir com Docker Compose**

```bash
docker-compose up --build
```

> Isso inicia o serviço da API (`localhost:8080`) e o banco MongoDB (`localhost:27017`)

---

## 🎯 Objetivo do Teste

Verificar se um leilão criado é **automaticamente encerrado** após o tempo configurado (`AUCTION_INTERVAL`).

---

## 🧪 Testando com `curl`

### ✅ 1. Criar um novo leilão

```bash
curl -X POST http://localhost:8080/auction \
  -H "Content-Type: application/json" \
  -d '{
    "product_name": "Casa ABCD",
    "category": "Imóveis",
    "description": "Casa da Rua ABCD, 223",
    "condition": 0
  }'
```

> Espera-se: HTTP 201 Created (sem corpo de resposta)

---

### 📥 2. Consultar os leilões existentes

```bash
curl "http://localhost:8080/auction?status=0"
```

Você verá algo como:

```json
[
  {
    "id": "b1705c7d-adb9-4585-ac97-58a55a364659",
    "product_name": "Casa ABCD",
    "category": "Imóveis",
    "description": "Casa da Rua ABCD, 223",
    "condition": 0,
    "status": 0,
    "timestamp": "2025-04-04T20:33:00Z"
  }
]
```

**status:** 0, 1 (Active, Completed)

---

### ⏱️ 3. Aguardar o intervalo (ex: 30s)

Após o tempo definido em `AUCTION_INTERVAL`, execute novamente:

```bash
curl http://localhost:8080/auction
```

Agora o campo `status` do leilão estará como `1`, indicando **leilão encerrado automaticamente**:

```json
{
  "status": 1 // Completed
}
```

---

## 📚 Documentação de apoio

| Endpoint              | Método | Descrição                            |
| --------------------- | ------ | ------------------------------------ |
| `/auction`            | GET    | Lista todos os leilões               |
| `/auction/:auctionId` | GET    | Detalha um leilão específico         |
| `/auction`            | POST   | Cria um novo leilão                  |
| `/auction/winner/:id` | GET    | Busca o lance vencedor (caso exista) |
| `/bid`                | POST   | Realiza um novo lance                |
| `/bid/:auctionId`     | GET    | Lista todos os lances de um leilão   |
| `/user/:userId`       | GET    | Busca um usuário                     |

---
