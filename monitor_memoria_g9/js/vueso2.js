const app = new Vue({
	el: "#app",
	data: {
		titulo: "Tareas con VUE js",
		usuarios: [
            {usuario : "root", pass : "root", nombre: "root", logueado:false},
            {usuario : "admin", pass : "admin", nombre: "admin", logueado:false},
            {usuario : "201114490", pass : "201114490", nombre: "Luis Ricardo", logueado:false},
            {usuario : "201114056", pass : "201114056", nombre: "Robinson Jonathan", logueado:false}
        ],
        mensaje: "",
        indice: -1,
        nombre: "",
        user:"",
        password:"",
        url:""

	},
	methods: {
        verificar() {
            this.indice = -1;
            this.nombre = "";
            for (var x = 0; x < this.usuarios.length; x++) {
                let usr = this.usuarios[x];
                if (usr.usuario == this.user  && usr.pass == this.password) {
                    this.indice = x;
                    this.usuarios[this.indice].logueado =  true;
                    localStorage.setItem("usuarios-vue", JSON.stringify(this.usuarios));
                    localStorage.setItem("indice-vue", JSON.stringify(this.indice));
				}
            }
            if ( this.indice == -1){
                this.mensaje = "error de autenticacion"
                localStorage.setItem("usuarios-vue", JSON.stringify(this.usuarios));
                localStorage.setItem("indice-vue", JSON.stringify(this.indice));
            }else {
                window.location.href = 'dashboard';
            }
        },
        cerrar(){
            this.nombre = "";
            this.indice = -1;
            localStorage.setItem("usuarios-vue", JSON.stringify(this.usuarios));
            localStorage.setItem("indice-vue", JSON.stringify(this.indice));
            window.location.href = '/';
        }
	},
	created: function () {
        let datosDBindice = JSON.parse(localStorage.getItem("indice-vue"));
        if (datosDBindice === null ){
            localStorage.setItem("indice-vue", JSON.stringify(this.indice));
            localStorage.setItem("usuarios-vue", JSON.stringify(this.usuarios));
        }

        datosDBindice = JSON.parse(localStorage.getItem("indice-vue"));
        let datosDB = JSON.parse(localStorage.getItem("usuarios-vue"));
		if (datosDB === null) {
            this.usuarios = [];
            this.indice = -1;
		} else {
            this.usuarios = datosDB;
            this.indice = datosDBindice;
            if (this.indice != -1) {
                this.nombre = this.usuarios[this.indice].nombre;
            }
        }
        this.url = window.location.pathname;
        if (this.indice == -1 && this.url != "/") {
            window.location.href = '/';
        }

	},
});