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

### **Estructura `Cuenta`**

- **Campos:**
    - `ID int64` — Identificador único de la cuenta.
    - `Saldo float64` — Saldo actual de la cuenta.
    - `Cliente Cliente` — Referencia al cliente que posee la cuenta.
- **Métodos:**
    - `ConsultarSaldo() float64` — Consulta el saldo de la cuenta.
    - `ConsultarMovimientos() []Movimiento` — Consulta los movimientos más recientes de la cuenta.
    - `GenerarExtractoMensual() []Movimiento` — Genera el extracto mensual de la cuenta.

### **Estructura `CuentaAhorros`**

- **Campos y Métodos:** Hereda campos y métodos de `Cuenta`. Representa una cuenta de ahorros para personas físicas.

### **Estructura `CuentaCorriente`**

- **Campos y Métodos:** Hereda campos y métodos de `Cuenta`. Representa una cuenta corriente para empresas.

### **Estructura `Movimiento`**

- **Campos:**
    - `ID int64` — Identificador único del movimiento.
    - `Tipo string` — Tipo del movimiento (depósito o retiro).
    - `Valor float64` — Valor del movimiento.
    - `Fecha time.Time` — Fecha del movimiento.
- **Métodos:**
    - `RealizarDeposito(valor float64) error` — Realiza un depósito en la cuenta.
    - `RealizarRetiro(valor float64) error` — Realiza un retiro en la cuenta.

### **Estructura `Cliente`**

- **Campos:**
    - `ID int64` — Identificador único del cliente.
    - `Nombre string` — Nombre del cliente.
    - `Ciudad string` — Ciudad de origen del cliente.
- **Métodos:**
    - `ObtenerInformeTransacciones(mes time.Month, año int) []Movimiento` — Obtiene el informe de transacciones del cliente para un mes específico.



¿Qué arquitectura y tecnologías usarías para resolver el caso Bluesoft Bank?


Sube la implementación del caso a un repositorio público y comparte el enlace.
