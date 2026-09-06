package cola

const mensajePanic string = "La cola esta vacia"

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

type nodoCola[T any] struct {
	dato T
	prox *nodoCola[T]
}

/* Se crea una Cola nueva, del tipo T, inicializada vacía, y se retorna*/
func CrearColaEnlazada[T any]() Cola[T] {
	return &colaEnlazada[T]{nil, nil}
}

/* EstaVacia devuelve verdadero si la cola no tiene elementos encolados, false en caso contrario. */
func (cola *colaEnlazada[T]) EstaVacia() bool {
	return cola.primero == nil
}

/*VerPrimero obtiene el valor del primero de la cola. Si está vacía, entra en pánico con un mensaje "La cola esta vacia".*/
func (cola *colaEnlazada[T]) VerPrimero() T {
	if cola.primero == nil {
		panic(mensajePanic)
	}
	return cola.ultimo.dato
}

/* Encolar agrega un nuevo elemento a la cola, al final de la misma.*/
func (cola *colaEnlazada[T]) Encolar(dato_nuevo T) {
	nuevo_ultimo := &nodoCola[T]{dato_nuevo, nil}
	if cola.EstaVacia() {
		cola.ultimo = nuevo_ultimo
		cola.primero = nuevo_ultimo
	} else {
		cola.ultimo.prox = nuevo_ultimo
		cola.ultimo = nuevo_ultimo
	}
}

/* Desencolar saca el primer elemento de la cola. Si la cola tiene elementos, se quita el primero de la misma, y se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La cola esta vacia". */
func (cola *colaEnlazada[T]) Desencolar() T {
	if cola.EstaVacia() {
		panic(mensajePanic)
	}
	desencolado := cola.primero
	cola.primero = cola.primero.prox

	if cola.primero == nil {
		cola.ultimo = nil
	}
	return desencolado.dato
}
