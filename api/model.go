package api

import "time"

type Req struct {
	CodigoBarras int64   `json:"codigo_barras" codinome:"codigo_barras"`
	Nome         string  `json:"nome" binding:"required,gte=1" minLength:"1" codinome:"nome"`
	Descricao    string  `json:"descricao" codinome:"descricao"`
	Foto         string  `json:"foto" codinome:"foto"`
	Valorpago    float64 `json:"valorpago" codinome:"valorpago"`
	Valorvenda   float64 `json:"valorvenda" codinome:"valorvenda"`
	Qtde         int64   `json:"qtde" minLength:"1" codinome:"qtde"`
	UndCod       int64   `json:"und_cod" codinome:"und_cod"`
	CatCod       int64   `json:"cat_cod" codinome:"cat_cod"`
	ScatCod      int64   `json:"scat_cod" codinome:"scat_cod"`
}

type Res struct {
	ID              int64     `json:"id,omitempty" codinome:"id"`
	DataCriacao     time.Time `json:"data_criacao,omitempty" codinome:"data_criacao"`
	DataAtualizacao time.Time `json:"data_atualizacao" codinome:"data_atualizacao"`
	CodigoBarras    int64     `json:"codigo_barras" codinome:"codigo_barras"`
	Nome            string    `json:"nome" binding:"required,gte=1" minLength:"1" codinome:"nome"`
	Descricao       string    `json:"descricao" codinome:"descricao"`
	Foto            string    `json:"foto" codinome:"foto"`
	Valorpago       float64   `json:"valorpago" codinome:"valorpago"`
	Valorvenda      float64   `json:"valorvenda" codinome:"valorvenda"`
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
	CodigoBarras int64 `uri:"codigo_barras" binding:"required,min=1"`
}

type listProdutoRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}
