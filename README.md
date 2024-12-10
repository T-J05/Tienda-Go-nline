# **Tienda Go-Online 🛒**  
¡Bienvenidos a la Tienda Go-Online! Una tienda en línea minimalista y eficiente que conecta a los usuarios con los productos que aman, mientras ofrece un poderoso panel administrativo para la gestión. 🐧✨  

---

## **💡 Características Principales**

### **Frontend (Go + HTML/Templates):**
- 🛍️ **Catálogo de productos**: Visualiza productos con imágenes, nombres y precios.
- 🛒 **Carrito de compras**: Agrega productos, revisa tu carrito y confirma pedidos con facilidad.
- 🏠 **Interfaz limpia y rápida**: Sin JavaScript en el frontend, todo el renderizado ocurre del lado del servidor.

### **Backend (Node.js + MongoDB):**
- 🛠️ **CRUD completo**: Crear, editar y eliminar productos desde el panel administrativo.
- 📦 **Gestión de pedidos**: Visualización y seguimiento de pedidos realizados.
- 🔐 **Autenticación segura**: Login mediante tokens JWT, diseñado para garantizar máxima seguridad.
- 🎨 **Vistas dinámicas**: Interfaz renderizada con Pug para mayor personalización.

---

## **⚙️ Requisitos**

Asegúrate de tener las siguientes herramientas instaladas en tu sistema antes de comenzar:  
1. **MongoDB**: Puede ser local o en la nube (cambia las URL de conexión según corresponda).  
2. **Go**: Instala la última versión desde su [sitio oficial](https://go.dev/).  
3. **Node.js**: También necesitas [NPM](https://nodejs.org/).  
4. **Nodemon** *(opcional)*: Para reinicio automático del servidor Node.js durante el desarrollo.

---

## **🚀 Pasos de Instalación**

1. Clona el repositorio:  
   ```bash
   git clone https://github.com/tuusuario/tienda-GO-online.git

2. **Instala dependencias del backend:**  
   Ve a la carpeta `backend` y ejecuta:  
   ```bash
   cd backend
   npm install
   ```

3. **Instala dependencias del frontend:**  
   Ve a la carpeta `frontend` y ejecuta:  
   ```bash
   cd ../frontend
   go mod tidy
   ```

4. **Configura las variables de entorno:**  
   Crea un archivo `.env` en el directorio `backend` con el siguiente contenido de ejemplo:  
   ```env
   PORT=3000
   SECRET_KEY=tu_secreto_super_seguro
   SALT_ROUNDS=10
   MONGO_URI=mongodb://localhost:27017/tienda
   ```

5. **Inicia el backend:**  
   Ve a la carpeta `backend` y ejecuta el servidor:  
   ```bash
   cd backend
   npm start
   ```

6. **Inicia el frontend:**  
   Ve a la carpeta `frontend` y ejecuta el servidor:  
   ```bash
   cd ../frontend
   go run main.go
   ```

## **Estructura del Proyecto**

### **Backend (Node.js - MongoDB - Pug)**

- **config:**  
  Conexión a la base de datos MongoDB.

- **controllers:**  
  Interacción con la base de datos (CRUD de productos y gestión de pedidos).

- **middlewares:**  
  Implementación de autenticación con tokens JWT.

- **models:**  
  Modelos para las colecciones de la base de datos.

- **public:**  
  Archivos estáticos como CSS e imágenes.

- **routes:**  
  Rutas para el login, CRUD de productos y visualización de pedidos.

- **views:**  
  Vistas renderizadas con Pug.

- **app.js:**  
  Archivo principal que inicializa MongoDB y la aplicación utilizando Express.

---

### **Frontend (Go - HTML/Templates - MongoDB)**

- **config:**  
  Configuración para la conexión a la base de datos MongoDB.

- **controllers:**  
  Controladores para manejar la interacción con la base de datos.

- **models:**  
  Definición de los modelos para las colecciones en MongoDB.

- **routes:**  
  Rutas para la funcionalidad de pedidos y carrito de compras.

- **static:**  
  Archivos CSS e imágenes necesarias para la interfaz.

- **templates:**  
  Archivos HTML que estructuran el frontend de la tienda online.

---

## **Flujo de Utilización**

### 1. **Inicio de la aplicación:**
   - Los usuarios pueden acceder a la tienda online desde su navegador.
   - Visualizan los productos disponibles, con detalles como la imagen, el nombre y el precio.

### 2. **Carrito de compras:**
   - Los usuarios pueden agregar productos al carrito de compras.
   - También pueden ver los productos en el carrito y confirmar su pedido para continuar.

### 3. **Panel de administrador:**
   - El acceso al panel de administrador requiere autenticación mediante un token JWT.
   - Una vez autenticado, el administrador puede:
     - **Gestionar productos:** Crear, editar y eliminar productos.
     - **Visualizar pedidos:** Consultar los pedidos realizados por los clientes.
