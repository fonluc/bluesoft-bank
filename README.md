### Caso Bluesoft Bank

Bluesoft Bank es un banco tradicional que se encarga de guardar el dinero de sus ahorradores. Ofrece dos tipos de cuenta: ahorros para personas naturales y corrientes para empresas. Adicionalmente, para cada cuenta se pueden hacer consignaciones y retiros.

Además, debe soportar algunos requerimientos para sus ahorradores:

Consultar el saldo de la cuenta.
Consultar los movimientos más recientes.
Generar extractos mensuales.
Reglas de negocio:

Una cuenta no puede tener un saldo negativo.
El saldo de la cuenta siempre debe ser consistente frente a dos operaciones concurrentes (consignación, retiro).
También se deben generar reportes en tiempo real como:

Listado de clientes con el número de transacciones para un mes en particular, organizado de manera descendente (primero el cliente con mayor número de transacciones en el mes).
Clientes que retiran dinero fuera de la ciudad de origen de la cuenta con el valor total de los retiros realizados superior a $1.000.000.
En base a lo anterior, por favor responde los siguientes puntos:

Crea un diagrama de clases que modele el problema, identifica los elementos principales y sus relaciones.

![Captura de tela 2024-08-07 210024](https://github.com/user-attachments/assets/0a03b025-bdfc-48c8-a407-6d0202ae3d7e)

### **Estrutura `Conta`**

- **Campos:**
    - `ID int64` — Identificador único da conta.
    - `Saldo float64` — Saldo atual da conta
    - `Cliente Cliente` — Referência ao cliente que possui a conta.
- **Métodos:**
    - `ConsultarSaldo() float64` — Consulta o saldo da conta.
    - `ConsultarMovimentos() []Movimento` — Consulta os movimentos mais recentes da conta.
    - `GerarExtratoMensal() []Movimento` — Gera o extrato mensal da conta.

### **Estrutura `ContaAhorros`**

- **Campos e Métodos:** Herda campos e métodos de `Conta`. Representa uma conta de poupança para pessoas físicas.

### **Estrutura `ContaCorrente`**

- **Campos e Métodos:** Herda campos e métodos de `Conta`. Representa uma conta corrente para empresas.

### **Estrutura `Movimento`**

- **Campos:**
    - `ID int64` — Identificador único do movimento.
    - `Tipo string` — Tipo do movimento (depósito ou saque).
    - `Valor float64` — Valor do movimento.
    - `Data time.Time` — Data do movimento.
- **Métodos:**
    - `RealizarDeposito(valor float64) error` — Realiza um depósito na conta.
    - `RealizarSaque(valor float64) error` — Realiza um saque na conta.

### **Estrutura `Cliente`**

- **Campos:**
    - `ID int64` — Identificador único do cliente.
    - `Nome string` — Nome do cliente.
    - `Cidade string` — Cidade de origem do cliente.
- **Métodos:**
    - `ObterRelatorioTransacoes(mes time.Month, ano int) []Movimento` — Obtém o relatório de transações do cliente para um mês específico.

¿Qué arquitectura y tecnologías usarías para resolver el caso Bluesoft Bank?
Sube la implementación del caso a un repositorio público y comparte el enlace.
