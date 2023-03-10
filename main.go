package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

/*
----------------------------------------------------------------------------------------------

	type BastileResponse struct {
		Target  string
		Release string
	}

	func BastilleStart(response http.ResponseWriter, router *http.Request) {
		// bastille start TARGET
		vars := mux.Vars(router)
		target := vars["target"]
		out, err := exec.Command("bastille", "start", target).Output()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Executing: bastille start %s\n", target)
		fmt.Println(string(out))

		response.WriteHeader(http.StatusOK)
		response.Write([]byte(out))
	}

----------------------------------------------------------------------------------------------
*/
func BastilleStart(c *gin.Context) {
	target := c.Param("target")
	fmt.Printf("Target is: %s\n", target)
	out, err := exec.Command("bastille", "start", target).Output()

	if err != nil {
		fmt.Println("There was an error with the bastille start command")
		log.Fatal(err)
	}

	fmt.Printf("Executing: bastille start %s\n", target)
	fmt.Println(string(out))

	c.JSON(http.StatusOK, string(out))

}

func BastilleStop(c *gin.Context) {
	target := c.Param("target")
	out, err := exec.Command("bastille", "stop", target).Output()

	if err != nil {
		fmt.Println("There was an error with the bastille stop command")
		log.Fatal(err)
	}

	fmt.Printf("Executing: bastille stop %s\n", target)
	fmt.Println(string(out))

	c.JSON(http.StatusOK, string(out))
}

func BastilleRestart(c *gin.Context) {
	target := c.Param("target")
	out, err := exec.Command("bastille", "restart", target).Output()

	if err != nil {
		fmt.Println("There was an error with the bastille restart command")
		log.Fatal(err)
	}

	fmt.Printf("Executing: bastille restart %s\n", target)
	fmt.Println(string(out))

	c.JSON(http.StatusOK, string(out))

}

/*
----------------------------------------------------------------------------------------------

	func BastillePkg(response http.ResponseWriter, router *http.Request) {
		// bastille pkg TARGET args
		vars := mux.Vars(router)
		target := vars["target"]
		args := vars["args"]
		out, err := exec.Command("bastille", "pkg", target, args).Output()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Executing: bastille pkg %s %s\n", target, args)
		fmt.Println(string(out))

		response.WriteHeader(http.StatusOK)
		response.Write([]byte(out))
	}

	func BastillePkgInstall(response http.ResponseWriter, router *http.Request) {
		// bastille pkg TARGET install args -y
		vars := mux.Vars(router)
		target := vars["target"]
		args := vars["args"]
		out, err := exec.Command("bastille", "pkg", target, "install", args, "-y").Output()

		if err != nil {
			fmt.Printf("DEBUG: error encountered")
		}

		fmt.Printf("Executing: bastille pkg %s install %s -y\n", target, args)
		fmt.Println(string(out))

		response.WriteHeader(http.StatusOK)
		response.Write([]byte(out))
	}

	func BastillePkgUpgrade(response http.ResponseWriter, router *http.Request) {
		// bastille pkg TARGET upgrade -y
		vars := mux.Vars(router)
		target := vars["target"]
		out, err := exec.Command("bastille", "pkg", target, "upgrade", "-y").Output()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Executing: bastille pkg %s upgrade -y\n", target)
		fmt.Println(string(out))

		response.WriteHeader(http.StatusOK)
		response.Write([]byte(out))
	}

----------------------------------------------------------------------------------------------
*/
func BastilleList(c *gin.Context) {
	out, err := exec.Command("bastille", "list", "-a").Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Executing: bastille list -a\n")
	fmt.Println(string(out))
	c.IndentedJSON(http.StatusOK, out)
}

func BastilleUpdate(c *gin.Context) {
	release := c.Param("release")
	out, err := exec.Command("bastille", "update", release).Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Executing: bastille update %q\n", release)
	fmt.Println(string(out))

	c.IndentedJSON(http.StatusOK, out)

}

func BastilleCmd(c *gin.Context) {
	target := c.Param("target")
	args := c.Param("args")
	out, err := exec.Command("bastille", "cmd", target, args).Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Executing: bastille cmd %s %s\n", target, args)
	fmt.Println(string(out))
	c.IndentedJSON(http.StatusOK, out)
}

/*
----------------------------------------------------------------------------------------------

	func BastilleService(response http.ResponseWriter, router *http.Request) {
	    // bastille service TARGET args
	    vars := mux.Vars(router)
	    target := vars["target"]
	    service := vars["service"]
	    args := vars["args"]
	    out, err := exec.Command("bastille", "service", target, service, args).Output()

	    if err != nil {
	        log.Fatal(err)
	    }

	    fmt.Printf("Executing: bastille service %s %s %s\n", target, service, args)
	    fmt.Println(string(out))

	    response.WriteHeader(http.StatusOK)
	    response.Write([]byte(out))
	}

	func BastilleSysrc(response http.ResponseWriter, router *http.Request) {
	    // bastille sysrc TARGET args
	    vars := mux.Vars(router)
	    target := vars["target"]
	    args := vars["args"]
	    out, err := exec.Command("bastille", "sysrc", target, args).Output()

	    if err != nil {
	        log.Fatal(err)
	    }

	    fmt.Printf("Executing: bastille sysrc %s %s\n", target, args)
	    fmt.Println(string(out))

	    response.WriteHeader(http.StatusOK)
	    response.Write([]byte(out))
	}

--------------------------------------------------------------------------------------------------------------------
*/
func BastilleBootstrap(c *gin.Context) {
	release := c.Param("release")
	out, err := exec.Command("bastille", "bootstrap", release).Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Executing: bastille bootstrap %s\n", release)
	fmt.Println(string(out))

	c.IndentedJSON(http.StatusOK, out)
}

/*
--------------------------------------------------------------------------------------------------------------------

	func BastilleConfig(response http.ResponseWriter, router *http.Request) {
	    // bastille config TARGET args
	    vars := mux.Vars(router)
	    target := vars["target"]
	    args := vars["args"]
	    out, err := exec.Command("bastille", "config", target, args).Output()

	    if err != nil {
	        log.Fatal(err)
	    }

	    fmt.Printf("Executing: bastille config %s %s\n", target, args)
	    fmt.Println(string(out))

	    response.WriteHeader(http.StatusOK)
	    response.Write([]byte(out))
	}

--------------------------------------------------------------------------------------------------------------------
*/
func BastilleClone(c *gin.Context) {
	target := c.Param("target")
	newname := c.Param("newname")
	IP := c.Param("IP")
	out, err := exec.Command("bastille", "clone", target, newname, IP).Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Executing: bastille clone %s %s %s\n", target, newname, IP)
	fmt.Println(string(out))
	c.IndentedJSON(http.StatusOK, out)
}

func BastilleCreate(c *gin.Context) {
	name := c.Param("name")
	release := c.Param("release")
	IP := c.Param("IP")
	out, err := exec.Command("bastille", "create", name, release, IP).Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Executing: bastille create %s %s %s\n", name, release, IP)
	fmt.Println(string(out))

	c.IndentedJSON(http.StatusCreated, out)
}

func BastilleDestroy(c *gin.Context) {
	target := c.Param("target")
	out, err := exec.Command("bastille", "destroy", target).Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Executing: bastille destroy %s\n", target)
	fmt.Println(string(out))
}

/*----------------------------------------------------------------------------------------------------------------------------
func BastilleRdr(response http.ResponseWriter, router *http.Request) {
    // bastille rdr TARGET proto hostport jailport
    vars := mux.Vars(router)
    target := vars["target"]
    proto := vars["proto"]
    hostport := vars["hostport"]
    jailport := vars["jailport"]
    out, err := exec.Command("bastille", "rdr", target, proto, hostport, jailport).Output()

    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Executing: bastille rdr %s %s %d %d\n", target, proto, hostport, jailport)
    fmt.Println(string(out))

    response.WriteHeader(http.StatusOK)
    response.Write([]byte(out))
}

func BastilleTemplate(response http.ResponseWriter, router *http.Request) {
    // bastille template TARGET project/repo
    vars := mux.Vars(router)
    target := vars["target"]
    project := vars["project"]
    repo := vars["repo"]
    template := project + "/" + repo
    out, err := exec.Command("bastille", "template", target, template).Output()

    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Executing: bastille template %s %s\n", target, template)
    fmt.Println(string(out))

    response.WriteHeader(http.StatusOK)
    response.Write([]byte(out))
}*/

func main() {
	fmt.Println("Starting the Bastille API...")

	router := gin.Default()

	router.GET("/v1/bastille/list", BastilleList)
	router.POST("/v1/bastille/bootstrap/:release", BastilleBootstrap)
	/*----------------------------------------------------------------------------------------------------------------------------
	router.HandleFunc("/v1/bastille/clone/{target}/{newname}/{IP}", BastilleClone).Methods("POST")
	----------------------------------------------------------------------------------------------------------------------------*/
	router.POST("/v1/bastille/clone/:target/:newname/:IP", BastilleClone)
	router.POST("v1/bastille/cmd/:target/:args", BastilleCmd)
	/*----------------------------------------------------------------------------------------------------------------------------
	router.HandleFunc("/v1/bastille/config/{target}/{args}", BastilleConfig).Methods("POST")
	----------------------------------------------------------------------------------------------------------------------------*/
	router.POST("/v1/bastille/create/:name/:release/:IP", BastilleCreate)
	/*----------------------------------------------------------------------------------------------------------------------------
	router.HandleFunc("/v1/bastille/destroy/{target}", BastilleDestroy).Methods("POST")
	router.HandleFunc("/v1/bastille/pkg/{target}/install/{args}", BastillePkgInstall).Methods("POST")
	router.HandleFunc("/v1/bastille/pkg/{target}/upgrade", BastillePkgUpgrade).Methods("POST")
	router.HandleFunc("/v1/bastille/pkg/{target}/{args}", BastillePkg).Methods("POST")
	router.HandleFunc("/v1/bastille/rdr/{target}/{proto}/{hostport}/{jailport}", BastilleRdr).Methods("POST")
	router.HandleFunc("/v1/bastille/restart/{target}", BastilleRestart).Methods("POST")
	----------------------------------------------------------------------------------------------------------------------------*/
	router.POST("/v1/bastille/restart/:target", BastilleRestart)
	/*----------------------------------------------------------------------------------------------------------------------------
	router.HandleFunc("/v1/bastille/service/{target}/{service}/{args}", BastilleService).Methods("POST")
	----------------------------------------------------------------------------------------------------------------------------*/
	router.POST("/v1/bastille/start/:target", BastilleStart)
	router.POST("/v1/bastille/stop/:target", BastilleStop)
	/*----------------------------------------------------------------------------------------------------------------------------
	router.HandleFunc("/v1/bastille/sysrc/{target}/{args}", BastilleSysrc).Methods("POST")
	router.HandleFunc("/v1/bastille/template/{target}/{project}/{repo}", BastilleTemplate).Methods("POST")
	----------------------------------------------------------------------------------------------------------------------------*/
	router.POST("/v1/bastille/update/:release", BastilleUpdate)
	/*----------------------------------------------------------------------------------------------------------------------------
	router.HandleFunc("/v1/bastille/sysrc/{target}/{args}", BastilleSysrc).Methods("POST")
	----------------------------------------------------------------------------------------------------------------------------*/
	fmt.Println("Listening on port 12345")

	// http.ListenAndServe(":12345", router)
	router.Run(":12345")
}
