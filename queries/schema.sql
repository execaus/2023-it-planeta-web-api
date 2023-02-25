CREATE TABLE "Account" (
    "id" bigserial NOT NULL,
    "first_name" character varying NOT NULL,
    "last_name" character varying NOT NULL,
    "email" character varying NOT NULL,
    "password" character varying NOT NULL,
    "deleted" boolean NOT NULL,
    CONSTRAINT "Account_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);

CREATE TABLE "Animal" (
    "id" bigserial NOT NULL,
    "chipping_location" bigint NOT NULL,
    "weight" double precision NOT NULL,
    "length" double precision NOT NULL,
    "height" double precision NOT NULL,
    "gender" character varying NOT NULL,
    "life_status" character varying NOT NULL,
    "chipping_date" TIMESTAMP NOT NULL,
    "chipper" bigint NOT NULL,
    "death_date" TIMESTAMP,
    "deleted" boolean NOT NULL,
    CONSTRAINT "Animal_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);

CREATE TABLE "AnimalType" (
    "id" bigserial NOT NULL,
    "value" character varying NOT NULL,
    "deleted" boolean NOT NULL,
    CONSTRAINT "AnimalType_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);

CREATE TABLE "LocationPoint" (
    "id" bigserial NOT NULL,
    "latitude" double precision NOT NULL,
    "longitude" double precision NOT NULL,
    "deleted" boolean NOT NULL,
    CONSTRAINT "LocationPoint_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);

CREATE TABLE "AnimalToType" (
    "animal" bigint NOT NULL,
    "animal_type" bigint NOT NULL,
    CONSTRAINT "AnimalToType_pk" PRIMARY KEY ("animal", "type")
) WITH (
  OIDS=FALSE
);

CREATE TABLE "AnimalGender" (
    "value" character varying NOT NULL,
    CONSTRAINT "AnimalGender_pk" PRIMARY KEY ("value")
) WITH (
  OIDS=FALSE
);

CREATE TABLE "AnimalLifeStatus" (
    "value" character varying NOT NULL,
    CONSTRAINT "AnimalLifeStatus_pk" PRIMARY KEY ("value")
) WITH (
  OIDS=FALSE
);

CREATE TABLE "AnimalVisitedLocation" (
    "id" bigserial NOT NULL,
    "location" bigint NOT NULL,
    "animal" bigint NOT NULL,
    "date" TIMESTAMP NOT NULL,
    "deleted" boolean NOT NULL,
    CONSTRAINT "AnimalVisitedLocation_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);

ALTER TABLE "Animal" ADD CONSTRAINT "Animal_fk0" FOREIGN KEY ("chipping_location") REFERENCES "LocationPoint"("id");
ALTER TABLE "Animal" ADD CONSTRAINT "Animal_fk1" FOREIGN KEY ("gender") REFERENCES "AnimalGender"("value");
ALTER TABLE "Animal" ADD CONSTRAINT "Animal_fk2" FOREIGN KEY ("life_status") REFERENCES "AnimalLifeStatus"("value");
ALTER TABLE "Animal" ADD CONSTRAINT "Animal_fk3" FOREIGN KEY ("chipper") REFERENCES "Account"("id");

ALTER TABLE "AnimalToType" ADD CONSTRAINT "AnimalToType_fk0" FOREIGN KEY ("animal") REFERENCES "Animal"("id");
ALTER TABLE "AnimalToType" ADD CONSTRAINT "AnimalToType_fk1" FOREIGN KEY ("type") REFERENCES "AnimalType"("id");

ALTER TABLE "AnimalVisitedLocation" ADD CONSTRAINT "AnimalVisitedLocation_fk0" FOREIGN KEY ("location") REFERENCES "LocationPoint"("id");
ALTER TABLE "AnimalVisitedLocation" ADD CONSTRAINT "AnimalVisitedLocation_fk1" FOREIGN KEY ("animal") REFERENCES "Animal"("id");
