package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	mockdb "pdv/db/mock"
	db "pdv/db/sqlc"
	"pdv/util"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetProduto(t *testing.T) {
	produto := randomProduto()
	testCases := []struct {
		name             string
		produtoCodBarras int64
		buildStubs       func(store *mockdb.MockStore)
		expectStatus     int
	}{
		{
			name:             "OK",
			produtoCodBarras: produto.CodigoBarras,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetProduto(gomock.Any(), gomock.Eq(produto.CodigoBarras)).
					Return(produto, nil).
					Times(1)
			},
			expectStatus: http.StatusOK,
		},
		{
			name:             "NotFound",
			produtoCodBarras: produto.CodigoBarras,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetProduto(gomock.Any(), gomock.Eq(produto.CodigoBarras)).
					Return(db.Produto{}, sql.ErrNoRows).
					Times(1)
			},
			expectStatus: http.StatusNotFound,
		},
		{
			name:             "InternalError",
			produtoCodBarras: produto.CodigoBarras,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetProduto(gomock.Any(), gomock.Eq(produto.CodigoBarras)).
					Return(db.Produto{}, sql.ErrConnDone).
					Times(1)
			},
			expectStatus: http.StatusInternalServerError,
		},
		{
			name:             "InvalidID",
			produtoCodBarras: 0,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetProduto(gomock.Any(), gomock.Any()).
					Times(0)
			},
			expectStatus: http.StatusBadRequest,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/api/v1/products/%d", tc.produtoCodBarras)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			require.Equal(t, tc.expectStatus, recorder.Code)
		})
	}
}

func TestCreateProduto(t *testing.T) {
	produto := randomProduto()

	testCases := []struct {
		name         string
		body         gin.H
		buildStubs   func(store *mockdb.MockStore)
		expectStatus int
	}{
		{
			name: "OK",
			body: gin.H{
				"codigo_barras": produto.CodigoBarras,
				"nome":       produto.Nome,
				"descricao":  produto.Descricao,
				"foto":       produto.Foto,
				"valorpago":  produto.Valorpago,
				"valorvenda": produto.Valorvenda,
				"qtde":       produto.Qtde,
				"und_cod":    produto.UndCod,
				"cat_cod":    produto.CatCod,
				"scat_cod":   produto.ScatCod,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreateProdutoParams{
					CodigoBarras: 0,
					Nome:         produto.Nome,
					Descricao:    produto.Descricao,
					Foto:         produto.Foto,
					Valorpago:    produto.Valorpago,
					Valorvenda:   produto.Valorvenda,
					Qtde:         produto.Qtde,
					UndCod:       produto.UndCod,
					CatCod:       produto.CatCod,
					ScatCod:      produto.ScatCod,
				}
				store.EXPECT().CreateProduto(gomock.Any(), gomock.Eq(arg)).Times(1).Return(produto, nil)
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "InvalidCurrency",
			body: gin.H{
				"nome":          produto.Nome,
				"codigo_barras": "invalid",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateProduto(gomock.Any(), gomock.Any()).
					Times(0)
			},
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "InvalidOwner",
			body: gin.H{
				"codigo_barras": "",
				"nome":          produto.Nome,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateProduto(gomock.Any(), gomock.Any()).
					Times(0)
			},
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "InternalError",
			body: gin.H{
				"nome":          produto.Nome,
				"codigo_barras": produto.CodigoBarras,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateProduto(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Produto{}, sql.ErrConnDone)
			},
			expectStatus: http.StatusInternalServerError,
		},
	}
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := NewServer(store)
			recorder := httptest.NewRecorder()

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/api/v1/products"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			require.Equal(t, tc.expectStatus, recorder.Code)
		})
	}
}

func randomProduto() db.Produto {
	return db.Produto{
		CodigoBarras: util.RandomInt(0, 100),
		Nome:         util.RandomString(6),
		Descricao:    util.RandomString(6),
		Foto:         util.RandomString(6),
		Valorpago:    util.RandomFloat(0, 100),
		Valorvenda:   util.RandomFloat(0, 100),
		Qtde:         util.RandomInt(0, 100),
		UndCod:       util.RandomInt(0, 100),
		CatCod:       util.RandomInt(0, 100),
		ScatCod:      util.RandomInt(0, 100),
	}
}
