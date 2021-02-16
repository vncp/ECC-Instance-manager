# ECC-remote-frontend
Front-end forms for the Engineering Computer Center's remote instances. Allowing users to request and manage their virtualized instances.

Authentication built using JWT, handled in the backend by checking authenticity of user credentials through PAM modules. This is handled through "msteinert/pam" Go-Pam wrapper.  
Database will be MySQL.  
Front-end made in NextJS/TypeScript.  
Back-end made with Go "gorilla/mux".  

# To-Do
### Frontend
- [x] Login Form
- [x] Request Form
- [x] JWT Authentication
- [x] Management Table / Actions
- [x] Request Table / Actions

### Backend
- [x] Base API
- [x] JWT Management
- [ ] PAM Authentication
- [ ] Set up MySQL Database
- [ ] Integrate Database with API
- [ ] Interface with Apache Guacamole Server

### Security
- [ ] TLS/HTTPS
- [ ] Add Rate Limiting
- [ ] Burp Proxy (SQL Injection, False Authentication, etc.)

# Sample
### Login Page
![Login Page](sample/login.png)
### Request Form
Allows users to submit a request which can be approved by staff  
![Request Form Page](sample/requestform.png)
### User Management
Allows users to independently restart their instances through server it's not responding.  
![User Management](sample/usermanagement.png)
### Request Management
Allows staff to respond to requests.  
![Request Management](sample/requestmanagement.png)
### Instance Management
Allows for management and info on instances.  
![Instance Management](sample/instancemanagement.png)
