# Comprobador de URLs

![Comprobador de URLs](https://img.shields.io/badge/Status-Completed-brightgreen)

Un proyecto desarrollado en **Go** con **Fiber** para comprobar el estado de una lista de URLs. Este proyecto permite ingresar múltiples URLs, verificar su estado HTTP y exportar los resultados en **JSON** o **PDF**.

## Características

- **Comprobación de URLs**:
  - Verifica el estado HTTP (`200`, `404`, etc.).
  - Muestra el tiempo de respuesta de cada URL.
- **Interfaz web**:
  - Diseñada con **HTML**, **CSS** y **Tailwind CSS**.
  - Fácil de usar.
- **Exportación**:
  - Guarda los resultados en formato **JSON** o **PDF**.
- **Íconos de estado**:
  - Muestra ✅ para URLs válidas y ❌ para URLs con errores.

---

## Tecnologías Utilizadas

### Backend
- **Golang**: Lenguaje principal del proyecto.
- **Fiber**: Framework de Go para crear aplicaciones web rápidas y eficientes.
- **Net/HTTP**: Para realizar solicitudes HTTP y medir tiempos de respuesta.

### Frontend
- **HTML5**: Para la estructura del proyecto.
- **CSS**: Utilizado en conjunto con **Tailwind CSS** para los estilos.
- **JavaScript**:
  - Gestión de eventos del formulario.
  - Exportación de resultados (JSON y PDF).
- **jsPDF**: Librería para generar archivos PDF dinámicamente desde el navegador.

---

## Requisitos Previos

1. Tener instalado **Go** (v1.20 o superior).
2. Tener instalado **Docker** (opcional, para contenerizar la aplicación).

---

## Instalación y Uso

1. **Clona este repositorio**:
   ```bash
   git clone https://github.com/tuusuario/url-checker.git
   cd url-checker
