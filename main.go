package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	lector := bufio.NewReader(os.Stdin)
	productos := []Producto{}
	clientes := []Cliente{}
	pedidos := []Pedido{}

	for {
		fmt.Println("\n================================")
		fmt.Println("     SISTEMA DE E-COMMERCE")
		fmt.Println("================================")
		fmt.Println("1. Gestionar productos")
		fmt.Println("2. Gestionar clientes")
		fmt.Println("3. Gestionar inventario")
		fmt.Println("4. Gestionar pedidos y ventas")
		fmt.Println("5. Salir")
		fmt.Println("================================")

		fmt.Print("Seleccione una opcion: ")
		opcion, err := lector.ReadString('\n')

		if err != nil {
			fmt.Println("Error al leer la opcion.")
			return
		}

		opcion = strings.TrimSpace(opcion)

		switch opcion {
		case "1":
			productos = gestionarProductos(productos, lector)
		case "2":
			clientes = gestionarClientes(clientes, lector)
		case "3":
			productos = gestionarInventario(productos, lector)
		case "4":
			pedidos, productos = gestionarPedidos(pedidos, productos, clientes, lector)
		case "5":
			fmt.Println("Gracias por utilizar nuestro sistema.")
			return
		default:
			fmt.Println("Opcion no valida.")
		}
	}
}
