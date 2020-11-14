# Client/Server TCP

# Objetivo

Hacer uso del paquete `net` de Go para implementar un cliente-servidor usando el protocolo `TCP`.

# Descripción

Desarrollar un servidor de procesos con las siguientes características:

- El servidor comenzará corriendo 5 procesos de manera concurrente y siempre estará mostrando los procesos que está corriendo.
- Cada proceso tendrá un contador que incrementará en 1 cada 500 milisegundos.
- Cuando se conecte un cliente al servidor, el servidor le asignará un proceso, esto es, se le envía un proceso al cliente para que el cliente siga corriendo ese proceso.
- Antes terminar un cliente, este debe de retornar el proceso al servidor con la intención de que el servidor siga corriendo el proceso que el cliente procesaba.
