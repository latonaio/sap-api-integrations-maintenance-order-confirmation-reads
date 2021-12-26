package sap_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema     string `json:"api_schema"`
	MaterialCode  string `json:"material_code"`
	Plant         string `json:"plant/supplier"`
	Stock         string `json:"stock"`
	DocumentType  string `json:"document_type"`
	DocumentNo    string `json:"document_no"`
	PlannedDate   string `json:"planned_date"`
	ValidatedDate string `json:"validated_date"`
	Deleted       bool   `json:"deleted"`
}

type SDC struct {
	ConnectionKey                string `json:"connection_key"`
	Result                       bool   `json:"result"`
	RedisKey                     string `json:"redis_key"`
	Filepath                     string `json:"filepath"`
	MaintenanceOrderConfirmation struct {
		MaintOrderConf                 string      `json:"MaintOrderConf"`
		MaintOrderConfCntrValue        string      `json:"MaintOrderConfCntrValue"`
		MaintenanceOrder               string      `json:"MaintenanceOrder"`
		MaintenanceOrderOperation      string      `json:"MaintenanceOrderOperation"`
		MaintenanceOrderSubOperation   string      `json:"MaintenanceOrderSubOperation"`
		PersonnelNumber                string      `json:"PersonnelNumber"`
		ActualWorkQuantity             string      `json:"ActualWorkQuantity"`
		ActualWorkQuantityUnit         string      `json:"ActualWorkQuantityUnit"`
		ActualDuration                 string      `json:"ActualDuration"`
		ActualDurationUnit             string      `json:"ActualDurationUnit"`
		OperationConfirmedStartDate    string      `json:"OperationConfirmedStartDate"`
		OperationConfirmedStartTime    string      `json:"OperationConfirmedStartTime"`
		OperationConfirmedEndDate      string      `json:"OperationConfirmedEndDate"`
		OperationConfirmedEndTime      string      `json:"OperationConfirmedEndTime"`
		IsFinalConfirmation            bool        `json:"IsFinalConfirmation"`
		NoFurtherWorkQuantityIsExpd    bool        `json:"NoFurtherWorkQuantityIsExpd"`
		RemainingWorkQuantity          string      `json:"RemainingWorkQuantity"`
		RemainingWorkQuantityUnit      string      `json:"RemainingWorkQuantityUnit"`
		PostingDate                    string      `json:"PostingDate"`
		ActivityType                   string      `json:"ActivityType"`
		OpenReservationsIsCleared      bool        `json:"OpenReservationsIsCleared"`
		ConfirmationText               string      `json:"ConfirmationText"`
		VarianceReasonCode             string      `json:"VarianceReasonCode"`
		CapacityInternalID             string      `json:"CapacityInternalID"`
		NmbrOfMaintTechnicianCapSplits int         `json:"NmbrOfMaintTechnicianCapSplits"`
		MaterialDocument               string      `json:"MaterialDocument"`
		AccountingIndicatorCode        string      `json:"AccountingIndicatorCode"`
		ActyConfFcstdEndDate           string      `json:"ActyConfFcstdEndDate"`
		ActyConfFcstdEndTime           string      `json:"ActyConfFcstdEndTime"`
	} `json:"MaintenanceOrderConfirmation"`
	APISchema                      string   `json:"api_schema"`
	Accepter                       []string `json:"accepter"`
	MaintenanceOrderConfirmationNo string   `json:"maintenance_order_confirmation"`
	Deleted                        bool     `json:"deleted"`
}
