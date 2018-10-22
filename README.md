Objetivo: verificar la conveniencia del uso de compresión de datos en las respuestas entregadas por los servicios

Caracteristicas del script de muestra:
    Es posible especifidar de forma estandar si se desea compresión de datos en la respuesta.


Equipo
    OS: Ubuntu 18.04.1 LTS
    OS Type: 64-bit
    Memory: 15,2 GiB
    Processor: Intel® Core™ i5-8250U CPU @ 1.60GHz × 8
    Disk: 113,0 GB

Muestra de datos: 400.000 requests

Compresión solicitada al servicio: gzip
    Comando ejecutado x 4: time seq 100000 | parallel -n0 ./throwRequests.sh gzip
    Tiempo real          : 267,395 s| 265,878 s| 265,702 s| 265,702 s
    Tiempo promedio      : 266 (266.17) s
    Tamaño respuesta     : 2.7 kb
    Tiempo promedio truncado de proceso de escritura : 226.39µs

Compresión solicitada al servicio: NO
    Comando ejecutado x 4: time seq 100000 | parallel -n0 ./throwRequests.sh none
    Tiempo real          : 247,835 s | 250,844 s | 250,630 s | 249,295 s
    Tiempo promedio      : 250 (249.651) s
    Tamaño respuesta     : 7.5 kb
    Tiempo promedio truncado de proceso de escritura: 45.89µs


Conclusión:
        Comprimir con gzip demora 0.0001805 fracciones de segundos más que sin comprimir
        Los datos comprimidos son 64% mas pequeños de los no comprimidos
        En 1 millon de requests son ahorrados: 4.5 GB








