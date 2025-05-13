# nginx

Para un ambiente aislado se puede realizar todo dentro de un contenedor:

```sh
docker build . -t nginx-tmp
docker run -it --rm nginx-testing:latest
```

Despues, dentro del contenedor:
```sh
go run main.go & # por defecto puerto 4000
go run main.go --addr=":4001" &
go run main.go --addr=":4002" &

service nginx start # iniciar nginx service
nginx -s reload -t  # cargar configuración
```

Ahora con `curl` es posible ver que el cuerpo de la respuesta cambia:
```sh
curl -i localhost
```

