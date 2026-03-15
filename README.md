# 📐 Euclides CLI

Uma **interface de linha de comando (CLI)** desenvolvida em **Go** para realizar operações matemáticas, como:

* Raiz quadrada
* Exponenciação
* Cálculo de porcentagem
* Regra de três simples

O projeto utiliza a biblioteca Cobra para organizar os comandos da CLI de forma modular.

**Versão:** 0.1.0

---

# ⚙️ Tecnologias

* Go
* Cobra
* Bibliotecas padrão do Go:

  * `math`
  * `strconv`
  * `fmt`
  * `log`

---

# 📦 Instalação para teste

### 1️⃣ Clone o repositório

```bash
git clone https://github.com/seu-usuario/seu-repositorio.git
cd seu-repositorio
```

### 2️⃣ Instale as dependências

```bash
go mod tidy
```

### 3️⃣ Compile o projeto para seu sistema operacional

```bash
go build
```

---

# 🚀 Como usar

Os comandos seguem a seguinte estrutura:

```bash
# Dentro do Repositorio
./euc comando [argumentos]

# Exportado para o sistema
euc comando [argumentos]
```

---

# 🔢 Comandos

## 1️⃣ Raiz Quadrada

Calcula a raiz quadrada de um número.

### Comando

```
sqrt <numero>
```

### Exemplo

```bash
euc sqrt 25
```

### Saída

```
The square root of 25.00 is 5.0000
```

### Regras

* Aceita **1 argumento**
* O número **não pode ser negativo**

---

## 2️⃣ Exponenciação

Eleva um número a uma determinada potência.

### Comando

```
exp <base> <expoente>
```

### Exemplo

```bash
euc exp 2 4
```

### Saída

```
The result is 2 raised to the power of 4, which equals 16
The formula is: 2 ^ 4 = 16
```

### Regras

* Requer **2 argumentos**
* A base **não pode ser 0**
* O expoente **deve ser maior que 0**

---

## 3️⃣ Porcentagem

Calcula a porcentagem de um número.

### Comando

```
perc <numero> <porcentagem>
```

### Exemplo

```bash
euc perc 200 10
```

### Saída

```
10.00% of 200.00 equals 20.00
```

### Regras

* Requer **2 argumentos**
* Valor máximo permitido: **32.000.000.000**

---

## 4️⃣ Regra de Três Simples

Calcula um valor utilizando a **regra de três simples**.

### Comando

```
srot <a> <b> <c>
```

### Fórmula

```
X = (a * b) / c
```

### Exemplo

```bash
euc srot 10 20 5
```

### Saída

```
The value of X is 40
```
---

# 5️⃣ Juros Simples

Calcula o valor de **juros simples** com base no capital inicial, taxa de juros e tempo.

Este comando aplica a fórmula clássica de **juros simples**

Onde:

* **J** = Juros
* **C** = Capital inicial
* **i** = Taxa de juros (em decimal)
* **t** = Tempo

---

## Comando

```bash
sint <capital> <taxa_juros> <tempo>
```

---

## Exemplo

```bash
euc sint 1000 0.01 12
```

### Saída esperada

```
The result of interest and 120
```

Nesse exemplo:

* Capital = 1000
* Taxa de juros = 0.01 (1%)
* Tempo = 12

Resultado:

```
J = 1000 × 0.01 × 12 = 120
```

---

## Regras

* O comando espera **3 argumentos obrigatórios**:

  1. Capital inicial
  2. Taxa de juros
  3. Tempo

* A **taxa de juros deve ser informada em formato decimal**, por exemplo:

| Porcentagem | Forma Decimal |
| ----------- | ------------- |
| 1%          | 0.01          |
| 5%          | 0.05          |
| 10%         | 0.10          |

* O programa retorna erro se:

  * algum argumento **não for um número**
  * a **taxa de juros for maior que 0.1**

---

# 🛡 Tratamento de Erros

O programa valida os dados antes de realizar os cálculos, verificando:

* Números inválidos
* Valores negativos
* Argumentos ausentes
* Valores acima do limite permitido

Os erros são tratados utilizando o pacote `log` do Go.

---

# 📊 Exemplo de Uso

```bash
# raiz quadrada
euc sqrt 81

# exponenciação
euc exp 5 3

# porcentagem
euc perc 1000 15

# regra de três
euc srot 12 8 4

# juros simples
euc sint 200 0.05 12
```

---

### 📜 Licença

Este projeto pode ser distribuído sob a licença **GPL-3.0 license**.

---

**Assinado:** dvcDaniel