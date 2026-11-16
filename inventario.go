package main

import (
	"bufio"
	"fmt"
	"strconv"
)

func consultarInventario(productos []Producto) {
	fmt.Println("\n--- INVENTARIO GENERAL ---")

	if len(productos) == 0 {
		fmt.Println("No hay productos registrados.")
		return
	}

	for _, p := range productos {
		fmt.Printf("ID: %d | Producto: %s | Existencias: %d\n",
			p.ID, p.Nombre, p.Stock)
	}
}

func consultarExistencias(productos []Producto, lector *bufio.Reader) {
	idTexto := leerTexto(lector, "Ingrese el ID del producto: ")
	id, err := strconv.Atoi(idTexto)

	if err != nil {
		fmt.Println("ID invalido.")
		return
	}

	for _, p := range productos {
		if p.ID == id {
			fmt.Printf("Producto: %s | Stock disponible: %d\n",
				p.Nombre, p.Stock)
			return
		}
	}

	fmt.Println("Producto no encontrado.")
}

func aumentarStock(productos []Producto, lector *bufio.Reader) []Producto {
	idTexto := leerTexto(lector, "ID del producto: ")
	id, err := strconv.Atoi(idTexto)

	if err != nil {
		fmt.Println("ID invalido.")
		return productos
	}

	cantidadTexto := leerTexto(lector, "Cantidad que desea agregar: ")
	cantidad, err := strconv.Atoi(cantidadTexto)

	if err != nil || cantidad <= 0 {
		fmt.Println("La cantidad debe ser un numero mayor que cero.")
		return productos
	}

	for i, p := range productos {
		if p.ID == id {
			productos[i].Stock += cantidad
			fmt.Println("Inventario actualizado correctamente.")
			fmt.Printf("Nuevo stock: %d\n", productos[i].Stock)
			return productos
		}
	}

	fmt.Println("Producto no encontrado.")
	return productos
}

func reducirStock(productos []Producto, lector *bufio.Reader) []Producto {
	idTexto := leerTexto(lector, "ID del producto: ")
	id, err := strconv.Atoi(idTexto)

	if err != nil {
		fmt.Println("ID invalido.")
		return productos
	}

	cantidadTexto := leerTexto(lector, "Cantidad que desea retirar: ")
	cantidad, err := strconv.Atoi(cantidadTexto)

	if err != nil || cantidad <= 0 {
		fmt.Println("La cantidad debe ser mayor que cero.")
		return productos
	}

	for i, p := range productos {
		if p.ID == id {
			if cantidad > p.Stock {
				fmt.Println("Stock insuficiente. No se realizo el cambio.")
				return productos
			}

			productos[i].Stock -= cantidad
			fmt.Println("Inventario actualizado correctamente.")
			fmt.Printf("Stock restante: %d\n", productos[i].Stock)
			return productos
		}
	}

	fmt.Println("Producto no encontrado.")
	return productos
}

func gestionarInventario(productos []Producto, lector *bufio.Reader) []Producto {
	for {
		fmt.Println("\n--- GESTION DE INVENTARIO ---")
		fmt.Println("1. Consultar inventario general")
		fmt.Println("2. Consultar existencias")
		fmt.Println("3. Aumentar stock")
		fmt.Println("4. Reducir stock")
		fmt.Println("5. Volver al menu principal")

		opcion := leerTexto(lector, "Seleccione una opcion: ")

		switch opcion {
		case "1":
			consultarInventario(productos)
		case "2":
			consultarExistencias(productos, lector)
		case "3":
			productos = aumentarStock(productos, lector)
		case "4":
			productos = reducirStock(productos, lector)
		case "5":
			return productos
		default:
			fmt.Println("Opcion no valida.")
		}
	}
}
