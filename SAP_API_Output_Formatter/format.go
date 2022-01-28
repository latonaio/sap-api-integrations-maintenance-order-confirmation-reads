package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-maintenance-order-confirmation-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) ([]Header, error) {
	pm := &responses.Header{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	header := make([]Header, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		header = append(header, Header{
			MaintOrderConf:                 data.MaintOrderConf,
			MaintOrderConfCntrValue:        data.MaintOrderConfCntrValue,
			MaintenanceOrder:               data.MaintenanceOrder,
			MaintenanceOrderOperation:      data.MaintenanceOrderOperation,
			MaintenanceOrderSubOperation:   data.MaintenanceOrderSubOperation,
			PersonnelNumber:                data.PersonnelNumber,
			ActualWorkQuantity:             data.ActualWorkQuantity,
			ActualWorkQuantityUnit:         data.ActualWorkQuantityUnit,
			ActualDuration:                 data.ActualDuration,
			ActualDurationUnit:             data.ActualDurationUnit,
			OperationConfirmedStartDate:    data.OperationConfirmedStartDate,
			OperationConfirmedStartTime:    data.OperationConfirmedStartTime,
			OperationConfirmedEndDate:      data.OperationConfirmedEndDate,
			OperationConfirmedEndTime:      data.OperationConfirmedEndTime,
			IsFinalConfirmation:            data.IsFinalConfirmation,
			NoFurtherWorkQuantityIsExpd:    data.NoFurtherWorkQuantityIsExpd,
			RemainingWorkQuantity:          data.RemainingWorkQuantity,
			RemainingWorkQuantityUnit:      data.RemainingWorkQuantityUnit,
			PostingDate:                    data.PostingDate,
			ActivityType:                   data.ActivityType,
			OpenReservationsIsCleared:      data.OpenReservationsIsCleared,
			ConfirmationText:               data.ConfirmationText,
			VarianceReasonCode:             data.VarianceReasonCode,
			CapacityInternalID:             data.CapacityInternalID,
			NmbrOfMaintTechnicianCapSplits: data.NmbrOfMaintTechnicianCapSplits,
			MaterialDocument:               data.MaterialDocument,
			AccountingIndicatorCode:        data.AccountingIndicatorCode,
			ActyConfFcstdEndDate:           data.ActyConfFcstdEndDate,
			ActyConfFcstdEndTime:           data.ActyConfFcstdEndTime,
		})
	}

	return header, nil
}
