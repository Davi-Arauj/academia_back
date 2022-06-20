-- name: CreateCliente :one
INSERT INTO clientes
(nome,email,cpf,fone,sexo,foto,data_nascimento )
VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING *;

-- name: GetCliente :one
SELECT * FROM clientes
WHERE nome=$1 LIMIT 1;

-- name: ListClientes :many
SELECT * FROM clientes
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateCliente :one
UPDATE clientes
SET  nome=$2, email=$3, cpf=$4, fone=$5, sexo=$6, foto=$7, data_nascimento=$8, data_atualizacao=now()
WHERE id=$1
RETURNING *;

-- name: DeleteCliente :exec
DELETE FROM clientes
WHERE id=$1;
