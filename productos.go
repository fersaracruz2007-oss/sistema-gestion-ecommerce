package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Producto struct {
	ID     int
	Nombre string
	Precio float64
	Stock  int
}

func leerTexto(lector *bufio.Reader, mensaje string) string {
	fmt.Print(mensaje)
	texto, _ := lector.ReadString('\n')
	return strings.TrimSpace(texto)
}

func registrarProducto(productos []Producto, lector *bufio.Reader) []Producto {
	nombre := leerTexto(lector, "Nombre del producto: ")
	precioTexto := leerTexto(lector, "Precio: ")
	stockTexto := leerTexto(lector, "Cantidad disponible: ")

	precio, err1 := strconv.ParseFloat(precioTexto, 64)
	stock, err2 := strconv.Atoi(stockTexto)

	if nombre == "" || err1 != nil || err2 != nil || precio <= 0 || stock < 0 {
		fmt.Println("Datos invalidos. No se registro el producto.")
		return productos
	}

	id := 1
	for _, p := range productos {
		if p.ID >= id {
			id = p.ID + 1
		}
	}

	nuevo := Producto{ID: id, Nombre: nombre, Precio: precio, Stock: stock}
	productos = append(productos, nuevo)

	fmt.Println("Producto registrado correctamente.")
	return productos
}

func mostrarProductos(productos []Producto) {
	if len(productos) == 0 {
		fmt.Println("No hay productos registrados.")
		return
	}

	for _, p := range productos {
		fmt.Printf("ID: %d | Nombre: %s | Precio: $%.2f | Stock: %d\n",
			p.ID, p.Nombre, p.Precio, p.Stock)
	}
}

func buscarProducto(productos []Producto, lector *bufio.Reader) {
	idTexto := leerTexto(lector, "Ingrese el ID del producto: ")
	id, err := strconv.Atoi(idTexto)

	if err != nil {
		fmt.Println("ID invalido.")
		return
	}

	for _, p := range productos {
		if p.ID == id {
			fmt.Printf("Producto: %s | Precio: $%.2f | Stock: %d\n",
				p.Nombre, p.Precio, p.Stock)
			return
		}
	}

	fmt.Println("Producto no encontrado.")
}

func actualizarProducto(productos []Producto, lector *bufio.Reader) []Producto {
	idTexto := leerTexto(lector, "ID del producto a actualizar: ")
	id, err := strconv.Atoi(idTexto)

	if err != nil {
		fmt.Println("ID invalido.")
		return productos
	}

	for i, p := range productos {
		if p.ID == id {
			nombre := leerTexto(lector, "Nuevo nombre: ")
			precioTexto := leerTexto(lector, "Nuevo precio: ")
			stockTexto := leerTexto(lector, "Nuevo stock: ")

			precio, err1 := strconv.ParseFloat(precioTexto, 64)
			stock, err2 := strconv.Atoi(stockTexto)

			if nombre == "" || err1 != nil || err2 != nil ||
				precio <= 0 || stock < 0 {
				fmt.Println("Datos invalidos.")
				return productos
			}

			productos[i] = Producto{
				ID: id, Nombre: nombre, Precio: precio, Stock: stock,
			}

			fmt.Println("Producto actualizado correctamente.")
			return productos
		}
	}

	fmt.Println("Producto no encontrado.")
	return productos
}

func eliminarProducto(productos []Producto, lector *bufio.Reader) []Producto {
	idTexto := leerTexto(lector, "ID del producto a eliminar: ")
	id, err := strconv.Atoi(idTexto)

	if err != nil {
		fmt.Println("ID invalido.")
		return productos
	}

	for i, p := range productos {
		if p.ID == id {
			productos = append(productos[:i], productos[i+1:]...)
			fmt.Println("Producto eliminado correctamente.")
			return productos
		}
	}

	fmt.Println("Producto no encontrado.")
	return productos
}

func gestionarProductos(productos []Producto, lector *bufio.Reader) []Producto {
	for {
		fmt.Println("\n--- GESTION DE PRODUCTOS ---")
		fmt.Println("1. Registrar producto")
		fmt.Println("2. Mostrar productos")
		fmt.Println("3. Buscar producto")
		fmt.Println("4. Actualizar producto")
		fmt.Println("5. Eliminar producto")
		fmt.Println("6. Volver al menu principal")

		opcion := leerTexto(lector, "Seleccione una opcion: ")

		switch opcion {
		case "1":
			productos = registrarProducto(productos, lector)
		case "2":
			mostrarProductos(productos)
		case "3":
			buscarProducto(productos, lector)
		case "4":
			productos = actualizarProducto(productos, lector)
		case "5":
			productos = eliminarProducto(productos, lector)
		case "6":
			return productos
		default:
			fmt.Println("Opcion no valida.")
		}
	}
}
