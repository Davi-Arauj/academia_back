package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	mockdb "pdv/db/mock"
	db "pdv/db/sqlc"
	"pdv/util"
	"testing"

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

			url := fmt.Sprintf("/api/v1/produtos/%d", tc.produtoCodBarras)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			require.Equal(t, tc.expectStatus, recorder.Code)
		})
	}
}

// func TestCreateProduto(t *testing.T) {
// 	produto := randomProduto()

// 	testCases := []struct{
// 		name             string
// 		body 			 gin.H
// 		buildStubs       func(store *mockdb.MockStore)
// 		expectStatus     int
// 	}{
// 		{
// 			name: "OK",
// 			body: gin.H{
// 				"cod_barras":produto.CodigoBarras,
// 			},
// 			buildStubs: func(store *mockdb.MockStore) {
// 				arg := db.CreateProdutoParams{
// 					CodigoBarras: produto.CodigoBarras,
// 					Nome: produto.Nome,
// 					Valorpago: produto.Valorpago,
// 				}
// 			},
// 		},
// 	}

// }

func randomProduto() db.Produto {
	return db.Produto{
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
}
