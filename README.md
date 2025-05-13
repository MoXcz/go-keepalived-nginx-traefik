# Load Balancer

## Uso

```sh
# Clonar repo
git clone https://github.com/MoXcz/go-keepalived-nginx-traefik.git && cd go-keepalived-nginx-traefik
```

### Traefik

```sh
docker compose up
```

### Nginx

```sh
cd nginx
docker build . -t nginx-tmp && docker run -it --rm nginx-testing:latest
# Dentro del contenedor:
go run main.go &
go run main.go --addr=":4001" &
go run main.go --addr=":4002" &

service nginx start
nginx -s reload -t
```

### Keepalived

```sh
cd keepalived
docker build . -t keepalived-node

docker run -d --name node1 \
  --net keepalived-net \
  --ip 192.168.100.10 \
  --cap-add=NET_ADMIN \
  --cap-add=NET_RAW \
  keepalived-node

docker run -d --name node2 \
  --net keepalived-net \
  --ip 192.168.100.11 \
  --cap-add=NET_ADMIN \
  --cap-add=NET_RAW \
  keepalived-node

docker exec -it node2 bash # Se modifica el status (BACKUP) y priority (< 100)

docker run -it --rm --net keepalived-net busybox ping 192.168.100.100
```

Un *load balancer* o *balanceador de carga* se encarga de distribuir la comunicación de diferentes conexiones entre múltiples servidores o instancias de una aplicación.

Para utilizar un balanceador de carga es primero necesario tener las *cargas* que balancear, las cuales consisten en la comunicación que existe entre diferentes dispositivos (comúnmente servidores).

A continuación se presentan distintas formas de utilizar un balanceador de carga dentro de Linux, Windows y dentro de aplicaciones.

## Nginx

Nginx es un servidor web que también puede funcionar como _reverse proxy_ o _load balancer_. En este caso, nos interesa configurar el bloque `upstream` para balanceo de carga.

### Configuración

Instalar [nginx](https://nginx.org/en/docs/install.html):

```sh
sudo apt install nginx # Debian / Ubuntu
```

Dependiendo del sistema utilizado, los archivos de configuración se encontrarán en:

* `/usr/local/nginx/conf`
* `/etc/nginx` (Debian / Ubuntu)
* `/usr/local/etc/nginx`

Existen otros [posibles archivos](https://stackoverflow.com/questions/41303885/nginx-do-i-really-need-sites-available-and-sites-enabled-folders) como `sites-available` y `sites-enabled`, pero no se utilizarán en este caso.

El archivo creado vendrá con una configuración por defecto que no se utilizará.

Nginx viene con una configuración predeterminada que no se utilizará. Se utiliza una configuración personalizada que defina un servicio HTTP a través de un bloque `upstream` que utilice los puertos `4000`, `4001` y `4002`. La configuración redirigirá todas las conexiones entrantes en el puerto `80` a estos puertos (utilizando un método de balanceo de carga sencillo de round robin).

## Traefik

Traefik es un _reverse proxy_ diseñado especialmente para trabajar con contenedores (como Docker y Kubernetes) y microservicios.

### Configuración

El archivo de configuración será `docker-compose.yml`, para poder utilizar Docker Compose y crear múltiples instancias de la misma aplicación de manera sencilla.

Dentro del archivo `docker-compose.yml`, se utiliza la etiqueta `"traefik.http.services.myapp.loadbalancer.server.port=4000"` para enrutar todas las conexiones hacia `myapp`, la cual está escuchando en el puerto 4000.

## Keepalived

Keepalived es una herramienta utilizada para proporcionar alta disponibilidad en sistemas Linux, particularmente útil para configurar balanceadores de carga redundantes. Su propósito principal es monitorear servicios y usar VRRP (Virtual Router Redundancy Protocol) para mantener una IP virtual (VIP) que siempre apunte al nodo activo.

Esto permite que, si un nodo (servidor) falla, otro tome su lugar de forma automática, sin interrupción del servicio.

> Keepalived no realiza el balanceo de carga por sí solo, sino que trabaja en conjunto con herramientas como nginx anteriormente descrita, proporcionando tolerancia a fallos a nivel de red.

### Configuración

Instalar `keepalived`:
```sh
sudo apt install keepalived # Debian / Ubuntu
```

La configuración de Keepalived se encuentra en `/etc/keepalived/keepalived.conf`, y el archivo de ejemplo predeterminado es `/etc/keepalived/keepalived.conf.sample`.

