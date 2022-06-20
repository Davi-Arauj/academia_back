package api

import "time"

type Req struct {
	CodBarras  int64   `json:"codigo_barras" codinome:"cod_barras"`
	Nome       string  `json:"nome" binding:"required,gte=1" minLength:"1" codinome:"nome"`
	Descricao  string  `json:"descricao" codinome:"descricao"`
	Foto       string  `json:"foto" codinome:"foto"`
	ValorPago  float64 `json:"valorpago" codinome:"valor_pago"`
	ValorVenda float64 `json:"valorvenda" codinome:"valor_venda"`
	Qtde       float64 `json:"qtde" minLength:"1" codinome:"qtde"`
	UndCod     int64   `json:"und_cod" codinome:"und_cod"`
	CatCod     int64   `json:"cat_cod" codinome:"cat_cod"`
	ScatCod    int64   `json:"scat_cod" codinome:"scat_cod"`
}

type Res struct {
	ID              int64    `json:"id,omitempty" codinome:"id"`
	DataCriacao     time.Time `json:"data_criacao,omitempty" codinome:"data_criacao"`
	DataAtualizacao time.Time `json:"data_atualizacao" codinome:"data_atualizacao"`
	CodBarras       int64     `json:"cod_barras" codinome:"cod_barras"`
	Nome            string    `json:"nome" binding:"required,gte=1" minLength:"1" codinome:"nome"`
	Descricao       string    `json:"descricao" codinome:"descricao"`
	Foto            string    `json:"foto" codinome:"foto"`
	ValorPago       float64   `json:"valor_pago" codinome:"valor_pago"`
	ValorVenda      float64   `json:"valor_venda" codinome:"valor_venda"`
	Qtde            float64   `json:"qtde" minLength:"1" codinome:"qtde"`
	UndCod          int64     `json:"und_cod" codinome:"und_cod"`
	CatCod          int64     `json:"cat_cod" codinome:"cat_cod"`
	ScatCod         int64     `json:"scat_cod" codinome:"scat_cod"`
}

type ResPag struct {
	Dados []Res  `json:"dados,omitempty"`
	Prox  *bool  `json:"prox,omitempty"`
	Total *int64 `json:"total,omitempty"`
}

type GetProdutoReq struct {
	CodBarras int64 `uri:"cod_barras" binding:"required,min=1"`
}

type listProdutoRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}
