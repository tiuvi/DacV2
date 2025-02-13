package dacv2

import (
	"time"
)


func ScheduleTaskAfterExpiration(timeExpiration time.Time, task func()(newExpiration time.Time)) {

	// Calculamos la espera
	timeNow := time.Now()

	timeWait := timeExpiration.Sub(timeNow)

	// Ejecutamos la tarea después del tiempo de espera en una sola goroutine
	go func() {

		for {

			// Esperar hasta que transcurra el tiempo de espera
			select {

			case <-time.After(timeWait):
				// Ejecutar la tarea
				newExpiration := task() // Ejecutar la tarea

				// Recalcular el tiempo de espera para la próxima ejecución
				timeNow = time.Now()
				
				timeWait = newExpiration.Sub(timeNow)
			}
		}
	}()
}
