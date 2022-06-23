package db

import (
	"context"
	"database/sql"
	"pdv/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomCliente(t *testing.T) Cliente {

	arg := CreateClienteParams{

		Nome:           util.RandomString(6),
		Email:          util.RandomString(6),
		Cpf:            util.RandomString(11),
		Fone:           util.RandomInt(0, 11),
		Foto:           util.RandomString(6),
		Sexo:           util.RandomString(6),
		DataNascimento: time.Now(),
	}

	cliente, err := testQueries.CreateCliente(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, cliente)

	require.Equal(t, arg.Nome, cliente.Nome)
	require.Equal(t, arg.Email, cliente.Email)
	require.Equal(t, arg.Cpf, cliente.Cpf)
	require.Equal(t, arg.Fone, cliente.Fone)
	require.Equal(t, arg.Foto, cliente.Foto)
	require.Equal(t, arg.Sexo, cliente.Sexo)

	require.NotZero(t, cliente.ID)
	require.NotZero(t, cliente.DataCriacao)
	return cliente
}

func TestCreateCliente(t *testing.T) {
	createRandomCliente(t)
}

func TestUpdateCliente(t *testing.T) {

	cliente1 := createRandomCliente(t)

	arg := UpdateClienteParams{
		ID:   cliente1.ID,
		Nome: util.RandomString(6),
	}

	cliente2, err := testQueries.UpdateCliente(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, cliente2)

	require.Equal(t, cliente1.ID, cliente2.ID)
	require.Equal(t, arg.Nome, cliente2.Nome)
	require.Equal(t, arg.Email, cliente2.Email)
	require.Equal(t, arg.Cpf, cliente2.Cpf)
	require.Equal(t, arg.Fone, cliente2.Fone)
	require.Equal(t, arg.Foto, cliente2.Foto)
	require.Equal(t, arg.Sexo, cliente2.Sexo)
}

func TestGetCliente(t *testing.T) {
	cliente1 := createRandomCliente(t)
	cliente2, err := testQueries.GetCliente(context.Background(), cliente1.Nome)
	require.NoError(t, err)
	require.NotEmpty(t, cliente2)
	require.Equal(t, cliente1, cliente2)

	require.WithinDuration(t, cliente1.DataCriacao, cliente2.DataCriacao, time.Second)
}

func TestDeleteCliente(t *testing.T) {
	cliente1 := createRandomCliente(t)

	err := testQueries.DeleteCliente(context.Background(), cliente1.ID)
	require.NoError(t, err)

	cliente2, err := testQueries.GetCliente(context.Background(), cliente1.Nome)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, cliente2)
}

func TestListClientes(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomCliente(t)
	}

	arg := ListClientesParams{
		Limit:  5,
		Offset: 5,
	}

	clientes, err := testQueries.ListClientes(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, clientes, 5)

	for _, cliente := range clientes {
		require.NotEmpty(t, cliente)
	}
}
