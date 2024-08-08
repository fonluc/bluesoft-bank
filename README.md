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

![Captura de tela 2024-08-07 213043](https://github.com/user-attachments/assets/1ca7db52-86b3-41d2-b19d-d7ca433424f4)

### **Structure `Account`**

- **Campos:**
    - `ID int64` — Identificador único de la cuenta.
    - `Balance float64` — Saldo actual de la cuenta.
    - `Customer Customer` — Referencia al cliente que posee la cuenta.
- **Métodos:**
    - `GetBalance() float64` — Consulta el saldo de la cuenta.
    - `GetTransactions() []Transaction` — Consulta los movimientos más recientes de la cuenta.
    - `GenerateMonthlyStatement() []Transaction` — Genera el extracto mensual de la cuenta.

### **Structure `SavingsAccount`**

- **Campos y Métodos:** Hereda campos y métodos de `Account`. Representa una cuenta de ahorros para personas físicas.

### **Structure `CheckingAccount`**

- **Campos y Métodos:** Hereda campos y métodos de `Account`. Representa una cuenta corriente para empresas.

### **Structure `Transaction`**

- **Campos:**
    - `ID int64` — Identificador único del movimiento.
    - `Type string` — Tipo del movimiento (depósito o retiro).
    - `Amount float64` — Valor del movimiento.
    - `Date time.Time` — Fecha del movimiento.
- **Métodos:**
    - `PerformDeposit(amount float64) error` — Realiza un depósito en la cuenta.
    - `PerformWithdrawal(amount float64) error` — Realiza un retiro en la cuenta.

### **Structure `Customer`**

- **Campos:**
    - `ID int64` — Identificador único del cliente.
    - `Name string` — Nombre del cliente.
    - `City string` — Ciudad de origen del cliente.
- **Métodos:**
    - `GetTransactionReport(month time.Month, year int) []Transaction` — Obtiene el informe de transacciones del cliente para un mes específico.

¿Qué arquitectura y tecnologías usarías para resolver el caso Bluesoft Bank?


Sube la implementación del caso a un repositorio público y comparte el enlace.
