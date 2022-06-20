-- name: CreateProduto :one
INSERT INTO produtos
(codigo_barras, nome, descricao, foto, valorpago, valorvenda, qtde, und_cod, cat_cod, scat_cod)
VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING *;

-- name: GetProduto :one
SELECT * FROM produtos
WHERE codigo_barras=$1 LIMIT 1;

-- name: ListProdutos :many
SELECT * FROM produtos
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateProduto :one
UPDATE produtos
SET  nome=$2, descricao=$3, foto=$4, valorpago=$5, valorvenda=$6, qtde=$7, und_cod=$8, cat_cod=$9, scat_cod=$10, data_atualizacao=now()
WHERE codigo_barras=$1
RETURNING *;

-- name: DeleteProduto :exec
DELETE FROM produtos
WHERE id=$1;


