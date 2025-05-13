# Keepalived

Para usar `keepalived` es necesario tener al menos dos nodos (se puede usar con un unico nodo, pero no tendria mucho sentido). Se crean los contenedores de prueba:

```sh
docker build . -t keepalived-node

docker run -d --name node1 \
  --net keepalived-net \
  --ip 192.168.100.10 \
  --cap-add=NET_ADMIN \
  --cap-add=NET_RAW \
  keepalived-node

# Se modifica el status (BACKUP) y priority (< 100)
docker run -d --name node2 \
  --net keepalived-net \
  --ip 192.168.100.11 \
  --cap-add=NET_ADMIN \
  --cap-add=NET_RAW \
  keepalived-node
```

Se puede comprobar la red dentro de los contenedores:

```sh
docker exec -it node1 bash
ip a | grep eth0
```

Cuando el nodo MASTER (`node1`) falle, la IP virtual (`192.168.1.100`) pasará automáticamente al nodo BACKUP, lo que mantiene el servicio accesible sin intervención manual.

## Usos combinados

Keepalived se puede combinar con **nginx o HAProxy** para el balanceo de carga.

```text
[Cliente] → 192.168.1.100 (VIP)
                     ↓
           +---------+---------+
           |                   |
     [nginx #1]         [nginx #2]
        MASTER            BACKUP
         ↑  ↓              ↑  ↓
     (app1, app2...)   (app1, app2...)
```

> Los servidores deberan de tener configurado su firewall para que acepten IP virtuales.

`ping` para probar su funcionamiento:
```sh
docker run -it --rm --net keepalived-net busybox ping 192.168.100.100
docker stop node1
docker start node1
docker stop node2
docker stop node1
```
