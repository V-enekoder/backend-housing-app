
<p align="center"><h1 align="center">Backend TutorIA</h1></p>
<p align="center">

</p>
<p align="center">
<img src="https://img.shields.io/github/license/V-enekoder/backend-housing-app?style=default&logo=opensourceinitiative&logoColor=white&color=0080ff" alt="license">
<img src="https://img.shields.io/github/last-commit/V-enekoder/backend-housing-app?style=default&logo=git&logoColor=white&color=0080ff" alt="last-commit">
<img src="https://img.shields.io/github/languages/top/V-enekoder/backend-housing-app?style=default&color=0080ff" alt="repo-top-language">
<img src="https://img.shields.io/github/languages/count/V-enekoder/backend-housing-app?style=default&color=0080ff" alt="repo-language-count">
</p>
<p align="center"><!-- default option, no dependency badges. -->
</p>
<p align="center">
	<!-- default option, no dependency badges. -->
</p>
<br>

## 🚀 Tecnologías

* Go
* Gin
* GORM
* SQLite


### 🛠️ Makefile

El Makefile incluye comandos útiles para gestionar el proyecto. Por favor asegúrate de tener `make` disponible en tu PATH. Entre estos están:

* `env`: Crea un archivo .env con las variable de entorno listas para configurarse
* `seed`: Realiza un llenado de datos inicial

Modo de Uso
```sh
❯ make comando
```
---

## 🖥️ Problemas Para Usuarios Windows

Si estás utilizando **Windows**, es posible que encuentres problemas con dependencias que usen **cgo** (por ejemplo, go-sqlite3), debido a la falta de compiladores adecuados. Para resolver este problema, se deja este tutorial: https://www.youtube.com/watch?v=1FuOwvj94zA

---


### ⚙️ Instalación


1. Clonar el repositorio:
```sh
❯ git clone https://github.com/V-enekoder/backend-housing-app.git
```

2. Navegar al directorio del proyecto:
```sh
❯ cd backend-housing-app
```

3. Instalar las dependencias del proyecto:
```sh
  ❯ go mod tidy
```
4. Copiar .env y luego configurar variables de entorno:
```sh
  ❯ make env
```
6. Realizar seeding inicial:
```sh
  ❯ make seed
```


### 🤖 Uso &nbsp; [<img align="center" src="https://img.shields.io/badge/Go-00ADD8.svg?style={badge_style}&logo=go&logoColor=white" />](https://golang.org/)

Pa ejecutar el programa se puede utilizar 

```sh
❯ go run main.go
```
o
```sh
❯ make run
```
