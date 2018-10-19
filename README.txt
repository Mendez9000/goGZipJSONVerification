time seq 10000 | parallel -n0 ./command_SIN_GZIP.sh

3 resultados promediados: 55.5 s


**********************************************************
57, 57, 58                57.3 s
time seq 10000 | parallel -n0 ./command_CON_GZIP.sh
PESO RESPUESTA 2.7 KB


PESO RESPUESTA 7.5 KB

El comprimido entrega a la red un archivo 3 veces menor en su tamaño.

Conclusion: entregar texto plano es 3% más rápido que entregar JSON comprimido.
            El JSON comprimido volcado a la red pesa el 30% del JSON original.






