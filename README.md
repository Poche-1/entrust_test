## Código fuente test CRUD (Entrust)

Para ejecutarlo se deberan realizar los siguientes pasos

1. Clonar el repositorio

2. Situado en la carpeta raiz  ejecutar el siguiente comando

```bash
$ docker-compose up
```

3. Una vez esten levantados los dos contenedores validar la ip del servidor, con el comando:

```bash
$ docker inspect entrustt | grep "IPAddress"
```
*Nota*: Realizar este paso con el fin de ajustar el valor de las constantes en el script de K6

4. Para ingresar al contenedor mariadb ingresar el siguiente comando

```bash
$ mysql -h 172.23.0.x -P 3306 --protocol=TCP -u root -p
```
*Nota: validar la ip del contenedor de mariadb*

5. Utilizar la contraseña que se encuentra descrita en el fichero docker-compose.yml

6. Ejecutar el siguiente comando para entrar al esquema 

```bash
$ use test_go
```

7. En la carpeta "script_database" ejecutar los script de base de datos para crear las tablas que usaremos para estas pruebas

###### Test con K6 

1. Ir a la carpeta K6
2. Ajustar en caso de ser necesario la ip del servidor para comenzar con las pruebas
3. Ejecutar el siguiente comando
```bash
$ k6 run script.js
```
