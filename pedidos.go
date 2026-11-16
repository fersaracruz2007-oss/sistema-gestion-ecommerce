package main

import (
	"bufio"
	"fmt"
	"strconv"
	"time"
)

type Pedido struct {
	ID             int
	ClienteNombre  string
	ProductoNombre string
	Cantidad       int
	Total          float64
	Fecha          string
}

func registrarPedido(pedidos []Pedido, productos []Producto, clientes []Cliente, lector *bufio.Reader) ([]Pedido, []Producto) {
	if len(clientes) == 0 || len(productos) == 0 {
		fmt.Println("Debe registrar al menos un cliente y un producto.")
		return pedidos, productos
	}

	fmt.Println("\n--- CLIENTES DISPONIBLES ---")
	mostrarClientes(clientes)

	clienteTexto := leerTexto(lector, "Ingrese el ID del cliente: ")
	clienteID, err := strconv.Atoi(clienteTexto)

	if err != nil {
		fmt.Println("ID de cliente invalido.")
		return pedidos, productos
	}

	var clienteEncontrado *Cliente
	for i := range clientes {
		if clientes[i].ID == clienteID {
			clienteEncontrado = &clientes[i]
			break
		}
	}

	if clienteEncontrado == nil {
		fmt.Println("Cliente no encontrado.")
		return pedidos, productos
	}

	fmt.Println("\n--- PRODUCTOS DISPONIBLES ---")
	mostrarProductos(productos)

	productoTexto := leerTexto(lector, "Ingrese el ID del producto: ")
	productoID, err := strconv.Atoi(productoTexto)

	if err != nil {
		fmt.Println("ID de producto invalido.")
		return pedidos, productos
	}

	indiceProducto := -1
	for i, p := range productos {
		if p.ID == productoID {
			indiceProducto = i
			break
		}
	}

	if indiceProducto == -1 {
		fmt.Println("Producto no encontrado.")
		return pedidos, productos
	}

	cantidadTexto := leerTexto(lector, "Cantidad que desea comprar: ")
	cantidad, err := strconv.Atoi(cantidadTexto)

	if err != nil || cantidad <= 0 {
		fmt.Println("La cantidad debe ser mayor que cero.")
		return pedidos, productos
	}

	producto := productos[indiceProducto]

	if cantidad > producto.Stock {
		fmt.Println("Stock insuficiente. No se registro el pedido.")
		return pedidos, productos
	}

	total := producto.Precio * float64(cantidad)

	id := 1
	for _, p := range pedidos {
		if p.ID >= id {
			id = p.ID + 1
		}
	}

	nuevoPedido := Pedido{
		ID:             id,
		ClienteNombre:  clienteEncontrado.Nombre,
		ProductoNombre: producto.Nombre,
		Cantidad:       cantidad,
		Total:          total,
		Fecha:          time.Now().Format("02/01/2006 15:04"),
	}

	pedidos = append(pedidos, nuevoPedido)
	productos[indiceProducto].Stock -= cantidad

	fmt.Println("\n¡Pedido registrado correctamente!")
	fmt.Printf("Cliente: %s\n", clienteEncontrado.Nombre)
	fmt.Printf("Producto: %s\n", producto.Nombre)
	fmt.Printf("Cantidad: %d\n", cantidad)
	fmt.Printf("Total a pagar: $%.2f\n", total)
	fmt.Printf("Stock restante: %d\n", productos[indiceProducto].Stock)

	return pedidos, productos
}

func mostrarPedidos(pedidos []Pedido) {
	fmt.Println("\n--- LISTADO DE PEDIDOS ---")

	if len(pedidos) == 0 {
		fmt.Println("No hay pedidos registrados.")
		return
	}

	for _, p := range pedidos {
		fmt.Printf(
			"Pedido: %d | Cliente: %s | Producto: %s | Cantidad: %d | Total: $%.2f | Fecha: %s\n",
			p.ID, p.ClienteNombre, p.ProductoNombre,
			p.Cantidad, p.Total, p.Fecha,
		)
	}
}

func gestionarPedidos(pedidos []Pedido, productos []Producto, clientes []Cliente, lector *bufio.Reader) ([]Pedido, []Producto) {
	for {
		fmt.Println("\n--- GESTION DE PEDIDOS Y VENTAS ---")
		fmt.Println("1. Registrar pedido")
		fmt.Println("2. Mostrar pedidos")
		fmt.Println("3. Volver al menu principal")

		opcion := leerTexto(lector, "Seleccione una opcion: ")

		switch opcion {
		case "1":
			pedidos, productos = registrarPedido(pedidos, productos, clientes, lector)
		case "2":
			mostrarPedidos(pedidos)
		case "3":
			return pedidos, productos
		default:
			fmt.Println("Opcion no valida.")
		}
	}
}
