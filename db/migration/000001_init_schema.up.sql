CREATE TABLE "produtos" (
	"id" bigserial PRIMARY KEY,
	"codigo_barras" bigint NULL,
	"nome" varchar NOT NULL,
	"descricao" varchar NULL,
	"foto" varchar NULL,
	"valorpago" bigint NOT NULL,
	"valorvenda" bigint NOT NULL,
	"qtde" bigint NOT NULL,
	"und_cod" bigint NULL,
	"cat_cod" bigint NULL,
	"scat_cod" bigint NULL,
	"data_criacao" timestamptz NOT NULL DEFAULT (now()),
	"data_atualizacao" timestamptz NULL
);