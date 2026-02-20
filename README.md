# Bulsara – Sitio Hugo

A continuación se detallan los pasos para clonar el repositorio y correr el sitio en local.

1. Clonar el repositorio y la rama correcta

  Usa el siguiente comando para clonar el proyecto y asegurarte de traer también los submódulos necesarios:
  
  git clone --recurse-submodules https://github.com/gatogaso/gatogaso.github.io Bulsara
  
⸻
2. Entrar a la carpeta del proyecto

  Una vez clonado el repositorio, ingresa al directorio del proyecto con:
  
  cd Bulsara

⸻
3. Para lanzar el sitio en local, ejecutá:

  hugo server

  Esto mostrará una salida como esta:

  Built in 152 ms
  Environment: "development"
  Serving pages from disk
  Running in Fast Render Mode. For full rebuilds on change: hugo server --disableFastRender
  Web Server is available at http://localhost:1313/ (bind address 127.0.0.1)
  Press Ctrl+C to stop

  Ahora podés abrir tu navegador e ir a: http://localhost:1313 para ver el sitio funcionando en vivo.

⸻
Para detener el servidor

  Presioná Ctrl + C en la terminal para detener el servidor.

⸻
