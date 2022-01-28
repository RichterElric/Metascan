package jsonWriter

import (
	"Metascan/main/log_templates/Log"
	"encoding/json"
	"os"
)

func convertToJSON(log Log.Log) []byte {
	print(log.Scan_date)
	b, _ := json.Marshal(log)

	return b
}

func writeInFile(json []byte) {
	_ = os.WriteFile("/opt/scan/metascan_results/result.json", json, 0644)

}

func WriteJSON(log Log.Log) bool {
	var resultAsJSON = convertToJSON(log)
	writeInFile(resultAsJSON)
	return true
}
