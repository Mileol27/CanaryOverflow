package main

import (
	"flag"
	"fmt"
	"log"

	"CanaryOverflow/internal/services"
)

func main() {
	// Define command-line flags
	phone := flag.String("phone", "", "Número de teléfono")
	name := flag.String("name", "", "Nombre del usuario")
	email := flag.String("email", "", "Correo electrónico del usuario")

	// Parse the flags
	flag.Parse()

	// Validate required flags
	if *phone == "" || *name == "" || *email == "" {
		log.Fatalf("Error: Debes proporcionar el número de teléfono, nombre y correo electrónico usando los flags --phone, --name y --email")
	}

	apiURL := "https://claro-middleware-apigw-brjb7ubo.uk.gateway.dev/api/frontoffice/vtex/createPreOrderCatologo"

	payload := services.ClaroRequest{
		Persona: struct {
			FormNombre              string `json:"formNombre"`
			NumeroDocumento         string `json:"numeroDocumento"`
			Telefono                string `json:"telefono"`
			TipoDocumento           string `json:"tipoDocumento"`
			Email                   string `json:"email"`
			FlagTerminosCondiciones string `json:"flagTerminosCondiciones"`
			FlagProteccionDatos     string `json:"flagProteccionDatos"`
		}{
			FormNombre:              *name,
			NumeroDocumento:         "65777335",
			Telefono:                *phone,
			TipoDocumento:           "dni",
			Email:                   *email,
			FlagTerminosCondiciones: "Si",
			FlagProteccionDatos:     "Si",
		},
	}

	response, err := services.SubscribeToClaro(apiURL, payload)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println(response)
}
