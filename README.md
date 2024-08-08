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

### Sube la implementación del caso a un repositorio público y comparte el enlace.

#### **1. Configuración del Entorno**

1. **Creación del Directorio del Proyecto**

   Se creó un nuevo directorio para el proyecto y se accedió a él desde el terminal.

2. **Inicialización del Repositorio Git**

   Se inicializó un repositorio Git en el directorio del proyecto.

3. **Inicialización del Módulo Go**

   Se inicializó un módulo Go con un nombre de módulo específico en el repositorio de Git.

#### **2. Estructura de Directorios y Archivos**

1. **Creación del Directorio `cmd` y Archivo `main.go`**

   Se creó un directorio `cmd` y se añadió el archivo `main.go` para configurar el servidor HTTP y las rutas de la API usando Gin.

2. **Creación del Directorio `internal` y Subdirectorios**

   Se creó una estructura de directorios `internal` con subdirectorios para `account`, `customer`, `transaction`, `persistence`, `messaging` y `services`. Dentro de estos directorios, se crearon archivos vacíos correspondientes.

3. **Contenido de los Archivos**

   - **`account.go`**: Define estructuras y métodos para cuentas, transacciones y tipos de cuentas.
   - **`transaction.go`**: Define funciones para realizar depósitos y retiros.
   - **`customer.go`**: Define la estructura de `Customer` y el método para obtener un informe de transacciones.
   - **`postgres.go`**: Define la función para conectar con una base de datos PostgreSQL.
   - **`kafka.go`**: Define la función para conectar con Kafka.
   - **`account_service.go`**: Define los manejadores de rutas para las operaciones de la API.

4. **Archivos Dockerfile y `docker-compose.yml`**

   - **`Dockerfile`**: Configura el entorno de ejecución de Docker para el aplicativo Go.
   - **`docker-compose.yml`**: Configura los servicios de Docker para el aplicativo, PostgreSQL y Kafka.

#### **3. Configuración del Entorno y Pruebas**

1. **Configuración de la Base de Datos PostgreSQL**

   - Instalación en Windows: Descargue el instalador, ejecute la instalación y configure el PostgreSQL.
   - Creación de la base de datos y tablas necesarias usando el pgAdmin.

2. **Configuración de Kafka**

   - Kafka se gestiona a través de Docker. Se inicia Kafka utilizando Docker Compose.

3. **Pruebas de Conexión**

   - **Base de Datos**: Se actualizó el archivo `postgres.go` para probar la conexión a la base de datos PostgreSQL y se creó un archivo de prueba `test_db_connection.go` para verificar la conexión.
   - **Kafka**: Se actualizó el archivo `kafka.go` y se creó `test_kafka_connection.go` para probar la conexión con Kafka.

4. **Pruebas Adicionales**

   - Se escribieron pruebas unitarias para las funciones de servicio en `internal/services/account_service_test.go`.
   - Se probaron las rutas de la API localmente utilizando herramientas como `curl`, `Postman`, o `httpie`.
