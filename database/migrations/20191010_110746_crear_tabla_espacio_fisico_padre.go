package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaEspacioFisicoPadre_20191010_110746 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaEspacioFisicoPadre_20191010_110746{}
	m.Created = "20191010_110746"

	migration.Register("CrearTablaEspacioFisicoPadre_20191010_110746", m)
}

// Run the migrations
func (m *CrearTablaEspacioFisicoPadre_20191010_110746) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS oikos.espacio_fisico_padre ( id serial NOT NULL, padre_id integer NOT NULL, hijo_id integer NOT NULL, fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, CONSTRAINT pk_espacio_fisico_padre PRIMARY KEY (id), CONSTRAINT fk_espacio_fisico_espacio_fisico_padre FOREIGN KEY (padre_id) REFERENCES oikos.espacio_fisico(id), CONSTRAINT fk_espacio_fisico_espacio_fisico_hijo FOREIGN KEY (hijo_id) REFERENCES oikos.espacio_fisico(id), CONSTRAINT uq_hijo_id_espacio_fisico_padre UNIQUE (hijo_id), CONSTRAINT uq_hijo_id_padre_id_espacio_fisico_padre UNIQUE (padre_id, hijo_id) );")
	m.SQL("ALTER TABLE oikos.espacio_fisico_padre OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE oikos.espacio_fisico_padre IS 'Contiene las relaciones de los espacios fisicos.';")
	m.SQL("COMMENT ON COLUMN oikos.espacio_fisico_padre.id IS 'Identificador de la tabla.';")
	m.SQL("COMMENT ON COLUMN oikos.espacio_fisico_padre.padre_id IS 'Identificador del espacio fisico padre.';")
	m.SQL("COMMENT ON COLUMN oikos.espacio_fisico_padre.hijo_id IS 'Identificador del espacio fisico hijo.';")
	m.SQL("COMMENT ON COLUMN oikos.espacio_fisico_padre.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN oikos.espacio_fisico_padre.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")
	m.SQL("COMMENT ON CONSTRAINT uq_hijo_id_espacio_fisico_padre ON oikos.espacio_fisico_padre IS 'Restringe que el arbol no se vuelva un grafo.';")
	m.SQL("COMMENT ON CONSTRAINT uq_hijo_id_padre_id_espacio_fisico_padre ON oikos.espacio_fisico_padre    IS 'Restriccion para que solo pueda existir una unica relacion entre un padre y un hijo.';")
	
}

// Reverse the migrations
func (m *CrearTablaEspacioFisicoPadre_20191010_110746) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS oikos.espacio_fisico_padre")
}
