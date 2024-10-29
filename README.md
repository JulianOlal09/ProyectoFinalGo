# ProyectoFinalGo
Proyecto Final Taller Go (Equipo 6)

A continuación se detallan las instrucciones y pasos a seguir para configurar y ejecutar el proyecto. 

1. Clonar el repositorio para
  
   git clone <https://github.com/JulianOlal09/ProyectoFinalGo>
   cd proyecto-final-go



2. Configurar las bases de datos MySQL

   Antes de ejecutar el código primero hay que verificar la conexión, asegurar que las credenciales proporcionadas en main.go coincidan con el usuario que esta definido en el archivo main.go de cada parte del proyecto.
   
   Estudiantes
   dsn := "root:toor@tcp(localhost:3306)/escuela"

   Materias
   dsn := "root:abigail30@tcp(127.0.0.1:3306)/escuela"

   Calificaciones
   dsn := "root:Anettevira26?@tcp(localhost:3307)/escuela"


4. Con la base de datos y la tabla correspondiente ya existentes, ejecute el servidor con el siguiente comando:
   
   Estudiantes y Calificaciones
   go run main.go

   Materias
   go run main.go subject.go
   

   
