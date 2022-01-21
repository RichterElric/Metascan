package htmlWriter

import "Metascan/main/log_templates/Log"

func convertToHtml(log Log.Log) string{
	print(log.Scan_date)
	var result = ""
	result += "<html>\n" +
		"<head></head>\n" +
		"<h1>Scan results</h1>\n" +
		"<p>Scan duration:" + log.Scan_date + "</p>\n" +
		"<p>Scan types:</p>\n"
	for _, s := range log.Scan_types {
		result += "- " + s
	}
	print(result)
	return "result"
}

func writeInFile(html string) {

}

func WriteHtml(log Log.Log) bool{
	var resultAsHtml = convertToHtml(log)
	writeInFile(resultAsHtml)
	return true
}
