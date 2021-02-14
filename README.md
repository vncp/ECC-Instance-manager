# ECC-remote-frontend
Front-end forms for the Engineering Computer Center's remote instances. Allowing users to request and manage their virtualized instances.

Authentication built using JWT, handled in the backend by checking authenticity of user credentials through PAM modules. This is handled through "msteinert/pam" Go-Pam wrapper.  
Database will be PostgreSQL.  
Front-end made in NextJS/TypeScript.  
Back-end made with Go "gorilla/mux".  
