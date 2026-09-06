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
	require.True(t, cola.EstaVacia())
}
