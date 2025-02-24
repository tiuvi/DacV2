package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func testMysql() {

	println("total de operacion msql de escritura y lectura: " , totalOperation * 2)

	// Conectar a MySQL con el usuario root (sin especificar base de datos)
	dsn := "miusuario:mipassword@tcp(127.0.0.1:3306)/"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("❌ Error al conectar a MySQL:", err)
	}

	// Crear la base de datos
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS midatabase")
	if err != nil {
		log.Fatal("❌ Error al crear la base de datos:", err)
	}

	err = db.Close()
	if err != nil {
		log.Fatal("❌ Error al crear la tabla:", err)
	}

	dsn = "miusuario:mipassword@tcp(127.0.0.1:3306)/midatabase"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("❌ Error al conectar a la base de datos:", err)
	}
	defer db.Close()

	// Crear tabla "boolean_table"
	createTableQuery := `CREATE TABLE IF NOT EXISTS boolean_table (
			id INT AUTO_INCREMENT PRIMARY KEY,
			value BOOLEAN
		)`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal("❌ Error al crear la tabla:", err)
	}

	start := time.Now()

	insertQuery := `INSERT INTO boolean_table (value) VALUES (?)`

	for i := 0; i < totalOperation; i++ {
		// Inserta un booleano aleatorio (true o false)
		value := i%2 == 0 // Alterna entre true y false
		_, err = db.Exec(insertQuery, value)
		if err != nil {
			log.Fatal("❌ Error al insertar datos:", err)
		}
	}

	duration := time.Since(start)
	fmt.Printf("✅ Inserción completa mysql. Tiempo transcurrido: %s\n", duration)

	
	query := "SELECT id, value FROM boolean_table LIMIT " + strconv.Itoa(totalOperation)
	start = time.Now()

	// Leer los primeros 1000 registros
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("❌ Error al leer datos:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var value bool
		err := rows.Scan(&id, &value)
		if err != nil {
			log.Fatal("❌ Error al escanear datos:", err)
		}
	}

	duration = time.Since(start)
	fmt.Printf("✅ Lectura completa mysql. Tiempo transcurrido: %s\n", duration)


	// Verifica si hubo algún error durante la iteración
	if err = rows.Err(); err != nil {
		log.Fatal("❌ Error durante la iteración de filas:", err)
	}
	
	deleteQuery := `DELETE FROM boolean_table`
	_, err = db.Exec(deleteQuery)
	if err != nil {
		log.Fatal("❌ Error al borrar datos:", err)
	}
	fmt.Println("✅ Todos los registros han sido borrados de la tabla.")
}