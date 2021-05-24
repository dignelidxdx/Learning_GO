22/05/2021

go build src/main.go

./main

go run src/main.go

Go can declarate variable with a initial value

Documentation

https://golang.org/pkg/
https://pkg.go.dev/

mejores features de Go: concurrencia, lo hace muy optimo,

Paralelismo
Hilos de ejecucion, no es mas una tarea que va a ejecutar el procesador -> va directo al procesador

Goroutines: inicia el proceso de creacion del kerner de linux, el se queda alli esperando que el kerner estuviera listo, en el caso de la concurrencia el espera que esten listas, Gofers va a la tarea donde esta el linux creado para empezar a anunciar y todo esto lo hace de forma concurrente, es decir no espera que uno de los puntos de esas tareas este listas para avanzar hacia adelante o hacia atras con respecto a la tareas en su secuencia de ejecucion y una vez que este listo simplemente va a anunciar que la tarea esta iniciada y se va.