package api

import (
	"database/sql"
	"net/http"
	db "pdv/db/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (server *Server) createProduto(ctx *gin.Context) {
	var req Req
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateProdutoParams{
		CodigoBarras: req.CodBarras,
		Nome:         req.Nome,
		Descricao:    req.Descricao,
		Foto:         req.Foto,
		Valorpago:    int64(req.ValorPago),
		Valorvenda:   int64(req.ValorVenda),
		Qtde:         int64(req.Qtde),
		UndCod:       req.UndCod,
		CatCod:       req.CatCod,
		ScatCod:      req.ScatCod,
	}

	produto, err := server.store.CreateProduto(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}

		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, produto)
}

func (server *Server) getProduto(ctx *gin.Context) {
	var req GetProdutoReq
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	produto, err := server.store.GetProduto(ctx, req.CodBarras)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, produto)
}

func (server *Server) listProduto(ctx *gin.Context) {
	var req listProdutoRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.ListProdutosParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	produtos, err := server.store.ListProdutos(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, produtos)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
