package dpfm_api_caller

import (
	dpfm_api_input_reader "data-platform-api-planned-order-deletes-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-planned-order-deletes-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) HeaderRead(
	input *dpfm_api_input_reader.SDC,
	log *logger.Logger,
) *dpfm_api_output_formatter.Header {
	where := fmt.Sprintf("WHERE header.PlannedOrder = %d ", input.Header.PlannedOrder)
	rows, err := c.db.Query(
		`SELECT 
			header.PlannedOrder
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_planned_order_header_data as header 
		` + where + ` ;`)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ItemsRead(
	input *dpfm_api_input_reader.SDC,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Item {
	where := fmt.Sprintf("WHERE item.PlannedOrder IS NOT NULL\nAND header.PlannedOrder = %d", input.Header.PlannedOrder)
	rows, err := c.db.Query(
		`SELECT 
			item.PlannedOrder, item.PlannedOrderItem
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_planned_order_item_data as item
		INNER JOIN DataPlatformMastersAndTransactionsMysqlKube.data_platform_planned_order_header_data as header
		ON header.PlannedOrder = item.PlannedOrder ` + where + ` ;`)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItem(rows)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}

	return data
}
