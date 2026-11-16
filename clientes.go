package main

import (
	"bufio"
	"fmt"
	"strconv"
)

type Cliente struct {
	ID       int
	Nombre   string
	Correo   string
	Telefono string
}

func registrarCliente(clientes []Cliente, lector *bufio.Reader) []Cliente {
	nombre := leerTexto(lector, "Nombre del cliente: ")
	correo := leerTexto(lector, "Correo electronico: ")
	telefono := leerTexto(lector, "Telefono: ")

	if nombre == "" || correo == "" || telefono == "" {
		fmt.Println("Todos los campos son obligatorios.")
		return clientes
	}

	id := 1
	for _, c := range clientes {
		if c.ID >= id {
			id = c.ID + 1
		}
	}

	nuevo := Cliente{
		ID:       id,
		Nombre:   nombre,
		Correo:   correo,
		Telefono: telefono,
	}

	clientes = append(clientes, nuevo)
	fmt.Println("Cliente registrado correctamente.")

	return clientes
}

func mostrarClientes(clientes []Cliente) {
	if len(clientes) == 0 {
		fmt.Println("No hay clientes registrados.")
		return
	}

	for _, c := range clientes {
		fmt.Printf("ID: %d | Nombre: %s | Correo: %s | Telefono: %s\n",
			c.ID, c.Nombre, c.Correo, c.Telefono)
	}
}

func buscarCliente(clientes []Cliente, lector *bufio.Reader) {
	idTexto := leerTexto(lector, "Ingrese el ID del cliente: ")
	id, err := strconv.Atoi(idTexto)

	if err != nil {
		fmt.Println("ID invalido.")
		return
	}

	for _, c := range clientes {
		if c.ID == id {
			fmt.Printf("Cliente: %s | Correo: %s | Telefono: %s\n",
				c.Nombre, c.Correo, c.Telefono)
			return
		}
	}

	fmt.Println("Cliente no encontrado.")
}

func actualizarCliente(clientes []Cliente, lector *bufio.Reader) []Cliente {
	idTexto := leerTexto(lector, "ID del cliente a actualizar: ")
	id, err := strconv.Atoi(idTexto)

	if err != nil {
		fmt.Println("ID invalido.")
		return clientes
	}

	for i, c := range clientes {
		if c.ID == id {
			nombre := leerTexto(lector, "Nuevo nombre: ")
			correo := leerTexto(lector, "Nuevo correo: ")
			telefono := leerTexto(lector, "Nuevo telefono: ")

			if nombre == "" || correo == "" || telefono == "" {
				fmt.Println("Todos los campos son obligatorios.")
				return clientes
			}

			clientes[i] = Cliente{
				ID:       id,
				Nombre:   nombre,
				Correo:   correo,
				Telefono: telefono,
			}

			fmt.Println("Cliente actualizado correctamente.")
			return clientes
		}
	}

	fmt.Println("Cliente no encontrado.")
	return clientes
}

func gestionarClientes(clientes []Cliente, lector *bufio.Reader) []Cliente {
	for {
		fmt.Println("\n--- GESTION DE CLIENTES ---")
		fmt.Println("1. Registrar cliente")
		fmt.Println("2. Mostrar clientes")
		fmt.Println("3. Buscar cliente")
		fmt.Println("4. Actualizar cliente")
		fmt.Println("5. Volver al menu principal")

		opcion := leerTexto(lector, "Seleccione una opcion: ")

		switch opcion {
		case "1":
			clientes = registrarCliente(clientes, lector)
		case "2":
			mostrarClientes(clientes)
		case "3":
			buscarCliente(clientes, lector)
		case "4":
			clientes = actualizarCliente(clientes, lector)
		case "5":
			return clientes
		default:
			fmt.Println("Opcion no valida.")
		}
	}
}
