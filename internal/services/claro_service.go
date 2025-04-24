package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type ClaroRequest struct {
	Delivery struct {
		Distrito          string `json:"distrito"`
		Departamento      string `json:"departamento"`
		FlagAceptaEntrega string `json:"flagAceptaEntrega"`
		TipoDespacho      string `json:"tipoDespacho"`
		Provincia         string `json:"provincia"`
		Direccion         string `json:"direccion"`
	} `json:"delivery"`
	ProductoLista []struct {
		Familia   string `json:"familia"`
		Modalidad string `json:"modalidad"`
		Campania  string `json:"campania"`
	} `json:"producto_lista"`
	Cobertura struct {
		FlagCobertura int    `json:"flagCobertura"`
		CoordenadaX   string `json:"coordenadaX"`
		CoordenadaY   string `json:"coordenadaY"`
	} `json:"cobertura"`
	TipoLead string `json:"tipoLead"`
	Pago     struct {
		PurchaseNumber string      `json:"purchaseNumber"`
		FormTarjeta    string      `json:"formTarjeta"`
		FlagAceptaPago interface{} `json:"flagAceptaPago"`
		PrecioNiubiz   interface{} `json:"precioNiubiz"`
	} `json:"pago"`
	Addons                   []interface{} `json:"addons"`
	IdentificadorTransaccion string        `json:"identificadorTransaccion"`
	Bundles                  []interface{} `json:"bundles"`
	Notificacion             struct {
		CallCenter  string `json:"callCenter"`
		MailCliente string `json:"mailCliente"`
		Grupo       string `json:"grupo"`
	} `json:"notificacion"`
	Persona struct {
		FormNombre              string `json:"formNombre"`
		NumeroDocumento         string `json:"numeroDocumento"`
		Telefono                string `json:"telefono"`
		TipoDocumento           string `json:"tipoDocumento"`
		Email                   string `json:"email"`
		FlagTerminosCondiciones string `json:"flagTerminosCondiciones"`
		FlagProteccionDatos     string `json:"flagProteccionDatos"`
	} `json:"persona"`
	IdentificadorCanal string      `json:"identificadorCanal"`
	Bambulytics        interface{} `json:"bambulytics"`
}

func SubscribeToClaro(apiURL string, payload ClaroRequest) (string, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Go-Client")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("failed to subscribe to Claro: " + resp.Status)
	}

	return "Suscrito exitosamente a Claro ;)", nil
}
