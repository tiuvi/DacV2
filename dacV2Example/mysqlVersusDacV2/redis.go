package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)


func testRedis() {

	// Crear cliente Redis
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Dirección de Redis
		DB:   0,                // Base de datos predeterminada
	})

	ctx := context.Background()

	// Insertar 1000 valores booleanos en Redis

	start := time.Now()

	for i := 0; i < totalOperation; i++ {
		key := fmt.Sprintf("bool_key_%d", i)
		value := i%2 == 0 // Alternar entre true y false
		err := client.Set(ctx, key, value, 0).Err()
		if err != nil {
			log.Fatal("❌ Error al insertar en Redis:", err)
		}
	}

	duration := time.Since(start)
	fmt.Printf("✅ Inserción completa en Redis. Tiempo transcurrido: %s\n", duration)

	// Leer los 1000 valores booleanos de Redis
	start = time.Now()

	for i := 0; i < totalOperation; i++ {
		key := fmt.Sprintf("bool_key_%d", i)
		val, err := client.Get(ctx, key).Bool()
		if err != nil {
			log.Fatal("❌ Error al leer de Redis:", err)
		}
		_ = val // No se usa, solo para simular la lectura
	}

	duration = time.Since(start)
	fmt.Printf("✅ Lectura completa en Redis. Tiempo transcurrido: %s\n", duration)

	keys, err := client.Keys(ctx, "bool_key_*").Result()
	if err != nil {
		log.Fatal("❌ Error obteniendo claves:", err)
	}

	// Verificar si hay claves para borrar
	if len(keys) > 0 {
		err = client.Del(ctx, keys...).Err()
		if err != nil {
			log.Fatal("❌ Error al borrar claves:", err)
		}
		fmt.Printf("✅ Se eliminaron %d claves correctamente\n", len(keys))
	} else {
		fmt.Println("✅ No hay claves para eliminar")
	}
	
	// Cerrar conexión con Redis
	client.Close()
}
