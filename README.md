## Caso Bluesoft Bank

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

### Crea un diagrama de clases que modele el problema, identifica los elementos principales y sus relaciones.

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

### ¿Qué arquitectura y tecnologías usarías para resolver el caso Bluesoft Bank?

### **1. Arquitectura de Microservicios:**

- **Descripción:** La aplicación se divide en servicios independientes, como gestión de cuentas, movimientos e informes. Cada microservicio es responsable de una parte específica del sistema, permitiendo escalabilidad y mantenimiento independientes.

### **2. Capas de la Arquitectura:**

- **Capa de Presentación:** La interfaz de usuario interactúa con la API.
- **Capa de API:** Implementada con Gin, gestiona las solicitudes y respuestas.
- **Capa de Aplicación:** Implementa la lógica de negocios.
- **Capa de Persistencia:** Gestionada por PostgreSQL para almacenar datos financieros.
- **Capa de Mensajería:** Usando Kafka para procesar eventos en tiempo real.
- **Capa de Orquestación y Contenerización:** Docker para contenedores y Kubernetes para orquestación.

### **Tecnologías**

- **Lenguaje:** **Go (Golang)**
    - **Justificación:** Ideal para alto rendimiento y operaciones simultáneas con goroutines, manteniendo el código limpio y eficiente.
- **Framework:** **Gin**
    - **Justificación:** Framework rápido y ligero para crear APIs RESTful, ofreciendo un rendimiento superior y facilidad de uso.
- **Persistencia:** **PostgreSQL**
    - **Justificación:** Base de datos relacional confiable con soporte para transacciones ACID, capaz de manejar consultas complejas y escalabilidad.
- **Mensajería:** **Kafka**
    - **Justificación:** Plataforma de streaming que procesa y transmite grandes volúmenes de datos en tiempo real, esencial para informes y eventos.
- **Contenerización:** **Docker**
    - **Justificación:** Simplifica la creación y gestión de contenedores, garantizando consistencia entre entornos.
- **Orquestación:** **Kubernetes**
    - **Justificación:** Gestiona contenedores a gran escala, automatizando la implementación y gestión, garantizando escalabilidad y disponibilidad.
- **Plataforma en la Nube:** **Google Cloud**
    - **Justificación:** Ofrece servicios escalables y seguros para computación, almacenamiento y bases de datos, ideal para alojar el sistema bancario.

Sube la implementación del caso a un repositorio público y comparte el enlace.
