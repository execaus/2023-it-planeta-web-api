CREATE TABLE "Account" (
	"id" serial NOT NULL,
	"first_name" character varying NOT NULL,
	"last_name" character varying NOT NULL,
	"email" character varying NOT NULL,
	"password" character varying NOT NULL,
	CONSTRAINT "Account_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);