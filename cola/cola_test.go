package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

const mensajePanic string = "La cola esta vacia"

func TestColaVaciaInt(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, mensajePanic, func() { cola.VerPrimero() })
	require.PanicsWithValue(t, mensajePanic, func() { cola.Desencolar() })
}

func TestColaEncolarInt(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(0)
	require.False(t, cola.EstaVacia())
	require.Equal(t, 0, cola.VerPrimero())
	require.Equal(t, 0, cola.Desencolar())
	/*require.True(t, cola.EstaVacia()) */
}

/*Prueba del invariante FIFO*/
func TestColaVolumen(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 0; i < 1001; i++ {
		cola.Encolar(i)
		require.Equal(t, i, cola.VerPrimero())
	}

	for i := 0; i < 1001; i++ {
		require.Equal(t, i, cola.Desencolar())
	}

	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, mensajePanic, func() { cola.VerPrimero() })
	require.PanicsWithValue(t, mensajePanic, func() { cola.Desencolar() })
}

func TestColaValoresBorde(t *testing.T) {
	cola_string := TDACola.CrearColaEnlazada[string]()
	cola_string.Encolar("")
	require.False(t, cola_string.EstaVacia())
	require.Equal(t, "", cola_string.VerPrimero())
	require.Equal(t, "", cola_string.Desencolar())
}
