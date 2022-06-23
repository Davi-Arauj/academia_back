package db

import (
	"context"
	"database/sql"
	"pdv/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomProduto(t *testing.T) Produto {
	arg := CreateProdutoParams{
		CodigoBarras: util.RandomInt(0, 100),
		Nome:         util.RandomString(6),
		Descricao:    util.RandomString(6),
		Foto:         util.RandomString(6),
		Valorpago:    util.RandomInt(0, 100),
		Valorvenda:   util.RandomInt(0, 100),
		Qtde:         util.RandomInt(0, 100),
		UndCod:       util.RandomInt(0, 100),
		CatCod:       util.RandomInt(0, 100),
		ScatCod:      util.RandomInt(0, 100),
	}

	produto, err := testQueries.CreateProduto(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, produto)

	require.Equal(t, arg.CodigoBarras, produto.CodigoBarras)
	require.Equal(t, arg.Nome, produto.Nome)
	require.Equal(t, arg.Descricao, produto.Descricao)
	require.Equal(t, arg.Foto, produto.Foto)
	require.Equal(t, arg.Valorpago, produto.Valorpago)
	require.Equal(t, arg.Valorvenda, produto.Valorvenda)
	require.Equal(t, arg.Qtde, produto.Qtde)
	require.Equal(t, arg.UndCod, produto.UndCod)
	require.Equal(t, arg.CatCod, produto.CatCod)
	require.Equal(t, arg.ScatCod, produto.ScatCod)

	require.NotZero(t, produto.ID)
	require.NotZero(t, produto.DataCriacao)
	return produto
}

func TestCreateProduto(t *testing.T) {
	createRandomProduto(t)
}

func TestUpdateProduto(t *testing.T) {
	produto1 := createRandomProduto(t)

	arg := UpdateProdutoParams{
		CodigoBarras: produto1.CodigoBarras,
		Nome:         util.RandomString(6),
	}

	produto2, err := testQueries.UpdateProduto(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, produto2)

	require.Equal(t, produto1.CodigoBarras, produto2.CodigoBarras)
	require.Equal(t, arg.Nome, produto2.Nome)
	require.Equal(t, arg.Descricao, produto2.Descricao)
	require.Equal(t, arg.Foto, produto2.Foto)
	require.Equal(t, arg.Valorpago, produto2.Valorpago)
	require.Equal(t, arg.Valorvenda, produto2.Valorvenda)
	require.Equal(t, arg.Qtde, produto2.Qtde)
	require.Equal(t, arg.UndCod, produto2.UndCod)
	require.Equal(t, arg.CatCod, produto2.CatCod)
	require.Equal(t, arg.ScatCod, produto2.ScatCod)
}

func TestGetProduto(t *testing.T) {
	produto1 := createRandomProduto(t)
	produto2, err := testQueries.GetProduto(context.Background(), produto1.CodigoBarras)
	require.NoError(t, err)
	require.NotEmpty(t, produto2)

	require.Equal(t, produto1.CodigoBarras, produto2.CodigoBarras)
	require.Equal(t, produto1.Nome, produto2.Nome)
	require.Equal(t, produto1.Descricao, produto2.Descricao)
	require.Equal(t, produto1.Foto, produto2.Foto)
	require.Equal(t, produto1.Valorpago, produto2.Valorpago)
	require.Equal(t, produto1.Valorvenda, produto2.Valorvenda)
	require.Equal(t, produto1.Qtde, produto2.Qtde)
	require.Equal(t, produto1.UndCod, produto2.UndCod)
	require.Equal(t, produto1.CatCod, produto2.CatCod)
	require.Equal(t, produto1.ScatCod, produto2.ScatCod)
	require.WithinDuration(t, produto1.DataCriacao, produto2.DataCriacao, time.Second)
}

func TestDeleteProduto(t *testing.T) {
	produto1 := createRandomProduto(t)

	err := testQueries.DeleteProduto(context.Background(), produto1.ID)
	require.NoError(t, err)

	produto2, err := testQueries.GetProduto(context.Background(), produto1.CodigoBarras)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, produto2)
}

func TestListProduto(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomProduto(t)
	}

	arg := ListProdutosParams{
		Limit:  5,
		Offset: 5,
	}

	produtos, err := testQueries.ListProdutos(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, produtos, 5)

	for _, produto := range produtos {
		require.NotEmpty(t, produto)
	}
}
