---

# 🚦 FullCycle Rate Limiter

Um serviço simples de **rate limiting** utilizando Redis, com suporte a limitação por IP ou por Token, configurável por variáveis de ambiente.

---

## 🔌 Portas Utilizadas

* **Servidor Web:** `8080`

---

## 📡 Endpoints Disponíveis

* `GET http://localhost:8080/` → Resposta: `Hello World`

---

## ⚙️ Configuração das Variáveis de Ambiente

| Variável                                     | Descrição                                                                         | Padrão |
| -------------------------------------------- | --------------------------------------------------------------------------------- | ------ |
| `WEB_SERVER_PORT`                            | Porta em que a aplicação será executada                                           | `8080` |
| `REDIS_URL_ADDRESS`                          | Endereço de conexão com o Redis                                                   | -      |
| `NR_MAXIMO_REQUEST_POR_SEGUNDO_IP`           | Número máximo de requisições por segundo de um mesmo IP                           | `10`   |
| `DURACAO_BLOQUEIO_IP`                        | Duração do bloqueio (em segundos) para IPs que excederem o limite                 | `10s`  |
| `NR_MAXIMO_REQUEST_POR_SEGUNDO_TOKEN_PADRAO` | Limite de requisições por segundo para tokens sem configuração específica         | `10`   |
| `DURACAO_BLOQUEIO_TOKEN_PADRAO`              | Duração do bloqueio para tokens sem configuração exclusiva que excederem o limite | `10s`  |

---

### 🔐 Variáveis Exclusivas por Token

Você pode definir regras específicas para um token usando as seguintes variáveis:

| Exemplo (Token = `ABCD123`)                   | Valor |
| --------------------------------------------- | ----- |
| `NR_MAXIMO_REQUEST_POR_SEGUNDO_TOKEN_ABCD123` | `100` |
| `DURACAO_BLOQUEIO_TOKEN_ABCD123`              | `30s` |

---

## ▶️ Como Rodar o Projeto

1. Execute o comando abaixo para iniciar o Redis e o servidor da aplicação:

```bash
docker-compose up
```

---

## ✅ Como Testar o Projeto

### 1. Testes Unitários

Execute os testes com os comandos:

```bash
go test ./internal/entity/
go test ./internal/usecase/
```

---

### 2. Testes com `curl`

#### Por IP (Cabeçalho `X-Real-IP`)

```bash
curl --parallel --parallel-immediate --parallel-max 20 \
  --header "X-Real-IP:10.10.10.1" \
  localhost:8080/ localhost:8080/ localhost:8080/ localhost:8080/ localhost:8080/ \
  localhost:8080/ localhost:8080/ localhost:8080/ localhost:8080/ localhost:8080/
```

#### Por Token (Cabeçalho `API_KEY`)

```bash
curl --parallel --parallel-immediate --parallel-max 20 \
  --header "API_KEY:ABCD1234" \
  localhost:8080/ localhost:8080/ localhost:8080/ localhost:8080/ localhost:8080/ \
  localhost:8080/ localhost:8080/ localhost:8080/ localhost:8080/ localhost:8080/
```

---

Se quiser, posso substituir o conteúdo do arquivo por essa versão revisada. Deseja que eu faça isso?
