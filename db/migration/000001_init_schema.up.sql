CREATE TABLE "produtos" (
	"id" bigserial PRIMARY KEY,
	"codigo_barras" bigint NOT NULL,
	"nome" varchar NOT NULL,
	"descricao" varchar NOT NULL,
	"foto" varchar NOT NULL,
	"valorpago" bigint NOT NULL,
	"valorvenda" bigint NOT NULL,
	"qtde" bigint NOT NULL,
	"und_cod" bigint NOT NULL,
	"cat_cod" bigint NOT NULL,
	"scat_cod" bigint NOT NULL,
	"data_criacao" timestamptz NOT NULL DEFAULT (now()),
	"data_atualizacao" timestamptz NULL
);

CREATE TABLE "clientes" (
  "id" varchar PRIMARY KEY NOT NULL,
  "nome" varchar NOT NULL,
  "email" varchar NOT NULL,
  "cpf" varchar,
  "fone" int,
  "foto" varchar,
  "sexo" varchar,
  "data_nascimento" timestamp,
  "data_criacao" timestamptz NOT NULL DEFAULT (now()),
  "data_atualizacao" timestamptz NULL
);