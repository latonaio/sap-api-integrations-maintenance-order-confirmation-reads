package main

import (
	sap_api_caller "sap-api-integrations-maintenance-order-confirmation-reads/SAP_API_Caller"
	"sap-api-integrations-maintenance-order-confirmation-reads/sap_api_input_reader"

	"github.com/latonaio/golang-logging-library/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs/SDC_Maintenance_Order_Confirmation_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
	)
	
    accepter := inoutSDC.Accepter
	if len(accepter) == 0 || accepter[0] == "All" {
		accepter = []string{
			"Header",
		}
	}
	
	caller.AsyncGetMaintenanceOrderConfirmation(
		inoutSDC.MaintenanceOrderConfirmation.MaintOrderConf,
		accepter,
	)
}
