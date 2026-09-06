package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

const PanicPilaVacia = "La pila esta vacia"
const maxVolumen = 10000

type persona struct {
	nombre string
	edad   int
	empleo bool
}

/* PRUEBAS DE PILA INT*/
func TestPilaVaciaInt(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, PanicPilaVacia, func() { pila.Desapilar() })
	require.PanicsWithValue(t, PanicPilaVacia, func() { pila.VerTope() })
}

func TestApilarAlgoInt(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	require.False(t, pila.EstaVacia())
}

func TestCasoBordeInt(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(0)
	require.False(t, pila.EstaVacia())
	require.Equal(t, 0, pila.Desapilar())
	require.True(t, pila.EstaVacia())
}

func TestApilarUnoInt(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	require.False(t, pila.EstaVacia())
	require.Equal(t, 1, pila.VerTope())
}

func TestDesapilarAlgoInt(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	desapilado := pila.Desapilar()
	require.True(t, pila.EstaVacia())
	require.Equal(t, 1, desapilado)
}

func TestVolumenInt(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i < maxVolumen; i++ {
		pila.Apilar(i)
		require.Equal(t, i, pila.VerTope())
	}

	require.False(t, pila.EstaVacia())

	for i := maxVolumen - 1; i >= 0; i-- {
		require.Equal(t, i, pila.Desapilar())
	}

	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, PanicPilaVacia, func() { pila.Desapilar() })
	require.PanicsWithValue(t, PanicPilaVacia, func() { pila.VerTope() })
}

/*PRUEBAS DE PILA STRING*/
func TestPilaVaciaString(t *testing.T) {
	pilaString := TDAPila.CrearPilaDinamica[string]()
	require.True(t, pilaString.EstaVacia())
}

func TestCasoBordeString(t *testing.T) {
	pilaString := TDAPila.CrearPilaDinamica[string]()
	require.True(t, pilaString.EstaVacia())
	stringPrueba := ""
	pilaString.Apilar(stringPrueba)
	require.False(t, pilaString.EstaVacia())
	require.Equal(t, stringPrueba, pilaString.Desapilar())
	require.True(t, pilaString.EstaVacia())
}

func TestApilarAlgoString(t *testing.T) {
	pilaString := TDAPila.CrearPilaDinamica[string]()
	require.True(t, pilaString.EstaVacia())
	primerString := "A"
	pilaString.Apilar(primerString)
	require.Equal(t, primerString, pilaString.VerTope())
	require.False(t, pilaString.EstaVacia())
}
func TestLIFO(t *testing.T) {
	stringsPrueba := []string{"A", "B", "C", "D", "E", "F"}
	pilaString := TDAPila.CrearPilaDinamica[string]()
	for i := 0; i < len(stringsPrueba); i++ {
		pilaString.Apilar(stringsPrueba[i])
	}
	require.False(t, pilaString.EstaVacia())

	for i := len(stringsPrueba) - 1; i >= 0; i-- {
		require.Equal(t, stringsPrueba[i], pilaString.Desapilar())
	}

	require.True(t, pilaString.EstaVacia())
}

/*PRUEBAS CON STRUCT*/

func TestPilaPersona(t *testing.T) {
	pilaPersona := TDAPila.CrearPilaDinamica[persona]()
	require.True(t, pilaPersona.EstaVacia())
	require.PanicsWithValue(t, PanicPilaVacia, func() { pilaPersona.VerTope() })
	require.PanicsWithValue(t, PanicPilaVacia, func() { pilaPersona.Desapilar() })
}

func TestPilaPersonaApilar(t *testing.T) {
	pilaPersona := TDAPila.CrearPilaDinamica[persona]()
	require.True(t, pilaPersona.EstaVacia())
	Jorge := persona{nombre: "Jorge", edad: 50, empleo: true}
	pilaPersona.Apilar(Jorge)
	require.False(t, pilaPersona.EstaVacia())
	require.Equal(t, Jorge, pilaPersona.Desapilar())
	require.True(t, pilaPersona.EstaVacia())
}

/*PRUEBAS DE PILAS CON PUNTEROS*/

func TestPilaPunteros(t *testing.T) {
	pilaPunteros := TDAPila.CrearPilaDinamica[*int]()
	require.True(t, pilaPunteros.EstaVacia())
	pilaPunteros.Apilar(nil)
	require.False(t, pilaPunteros.EstaVacia())
	require.Nil(t, pilaPunteros.VerTope())
	require.Nil(t, pilaPunteros.Desapilar())
	require.True(t, pilaPunteros.EstaVacia())
}
