# sabaticos_crud
:heavy_check_mark: Check: Repositorio API CRUD para el sistema de gestión de sabáticos.


## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones
* [Golang 1.25](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
* [BeeGo 1.12.3](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)
* [PostgreSQL](https://www.postgresql.org/)
* [Docker](https://docs.docker.com/engine/install/ubuntu/)


### Variables de Entorno
```shell
# Parámetros de la API
API_NAME=sabatico_crud
SABATICOS_CRUD_HTTPPORT=8080
SABATICOS_CRUD_RUNMODE=dev

#Database
SABATICOS_CRUD_DB_USER=postgres
SABATICOS_CRUD_DB_PASS=postgres
SABATICOS_CRUD_DB_URL=localhost
SABATICOS_CRUD_DB_NAME=database
SABATICOS_CRUD_DB_SCHEMA=schema
SABATICOS_CRUD_DB_PORT=port

```

### Ejecución del Proyecto
```shell
# 1. Clonar el repositorio
git clone -b develop https://github.com/udistrital/sabaticos_crud.git

# 2. Moverse a la carpeta del repositorio
cd sabaticos_crud

# 3. Moverse a la rama **develop**
git pull origin develop && git checkout develop

# 4. Configurar las variables de entorno
export SABATICOS_CRUD_HTTPPORT=8080

# 5. Ejecutar el proyecto
bee run
```

### Ejecución Dockerfile
```shell
# El Dockerfile está implementado para el despliegue mediante
# el sistema de integración continua (CI).

# 1. Compilar el proyecto para Linux
GOOS=linux GOARCH=amd64 go build -o main

# 2. Construir la imagen
docker build -t sabati .

# 3. Ejecutar el contenedor
docker run --name sabaticos_crud \
  --env-file custom.env \
  -p 8080:8080 \
  sabaticos_crud

# 4. Comprobar que el contenedor esté en ejecución
docker ps
```


### Ejecución docker-compose
```shell
# No implementado actualmente.
```

## Estado CI

| Develop | release/0.0.1 | Master | Sonar |
| -- | -- | -- | -- |
| [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/sabaticos_crud/status.svg?ref=refs/heads/develop)](https://hubci.portaloas.udistrital.edu.co/udistrital/sabaticos_crud)| [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/sabaticos_crud/status.svg?ref=refs/heads/release/0.0.1)](https://hubci.portaloas.udistrital.edu.co/udistrital/sabaticos_crud) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/sabaticos_crud/status.svg?ref=refs/heads/master)](https://hubci.portaloas.udistrital.edu.co/udistrital/sabaticos_crud) | [![Quality Gate Status](https://sonarqube.portaloas.udistrital.edu.co/api/project_badges/measure?project=sabaticos_crud&metric=alert_status)](https://sonar.portaloas.udistrital.edu.co/dashboard?id=sabaticos_crud) |


## Modelo de datos

[Modelo de datos](database/modelos.svg)



## Licencia

This file is part of sabaticos_crud.

sabaticos_crud is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

sabaticos_crud is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with sabaticos_crud. If not, see https://www.gnu.org/licenses/.
