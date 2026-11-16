
# Sistema de Gestión de E-commerce

## Descripción
Aplicación de consola desarrollada en Go para administrar las operaciones básicas de una tienda virtual.

## Objetivo
Gestionar productos, clientes, inventario y pedidos mediante funciones que facilitan la organización de las operaciones de un comercio electrónico.

## Módulos
- **Productos:** registro, consulta, búsqueda, actualización y eliminación.
- **Clientes:** registro, consulta, búsqueda y actualización.
- **Inventario:** consulta y modificación de existencias.
- **Pedidos y ventas:** registro de pedidos, cálculo de totales y descuento automático del inventario.

## Tecnologías utilizadas
- Lenguaje Go.
- Visual Studio Code.
- Git y GitHub.

## Paquetes utilizados
Se utilizan paquetes de la biblioteca estándar de Go: fmt, bufio, os, strings, strconv y time.

## Alcance y limitaciones
Los datos se almacenan temporalmente en memoria y se pierden al cerrar el programa. El sistema no procesa pagos reales ni gestiona envíos. Cada pedido registra un solo producto.

## Ejecución
Desde la terminal, ejecutar:

`go run main.go productos.go clientes.go inventario.go pedidos.go`
