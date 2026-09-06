package pila

/*LAS CONSTANTES*/
const capacidadInicial = 10
const factorRedimension = 2
const denominadorMaximo = 4
const panicPilaVacia = "La pila esta vacia"

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func (p *pilaDinamica[T]) redimensionar(nuevo_tam int) {
	nuevos_datos := make([]T, nuevo_tam)
	copy(nuevos_datos, p.datos[:p.cantidad])
	p.datos = nuevos_datos
}

func CrearPilaDinamica[T any]() Pila[T] {
	return &pilaDinamica[T]{make([]T, capacidadInicial), 0}
}

func (p *pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

func (p *pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic(panicPilaVacia)
	}

	return p.datos[p.cantidad-1]
}

func (p *pilaDinamica[T]) Apilar(t T) {
	if p.cantidad == len(p.datos) {
		p.redimensionar(p.cantidad * factorRedimension)
	}
	p.datos[p.cantidad] = t
	p.cantidad++
}

func (p *pilaDinamica[T]) Desapilar() T {
	if p.EstaVacia() {
		panic(panicPilaVacia)
	}

	if (p.cantidad*denominadorMaximo) <= len(p.datos) && (p.cantidad/factorRedimension) > capacidadInicial {
		p.redimensionar(len(p.datos) / factorRedimension)
	}
	p.cantidad--
	return p.datos[p.cantidad]
}
