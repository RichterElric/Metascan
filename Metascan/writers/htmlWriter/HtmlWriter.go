package htmlWriter

import (
	"Metascan/main/log_templates/Log"
	"log"
	"os"
	"strconv"
)

func convertToHtml(log Log.Log) string{
	var result = ""
	result += "<html>\n" +
		"\t<head>" +
		"\t</head>\n" +
		"\t<body>" +
		"\t\t<h1>Scan results</h1>\n" +
		"\t\t<p>Scan duration:" + log.Scan_date + "</p>\n" +
		"\t\t<p>Scan types:</p>\n" +
		"\t\t<ul>\n"
	for _, s := range log.Scan_types {
		result += "\t\t\t<li>" + s + "</li>\n"
	}
	result += "\t\t</ul>\n" +
		"\t\t<p>Severity counters:</p>\n" +
		"\t\t<ul>\n" +
		"\t\t\t<li>High: " + strconv.Itoa(log.Severity_counters.High) + "</li>\n" +
		"\t\t\t<li>Medium: " + strconv.Itoa(log.Severity_counters.Medium) + "</li>\n" +
		"\t\t\t<li>Low: " + strconv.Itoa(log.Severity_counters.Low) + "</li>\n" +
		"\t\t\t<li>Info: " + strconv.Itoa(log.Severity_counters.Info) + "</li>\n" +
		"\t\t\t</ul>\n"
	result += "\t\t<table>\n" +
		"\t\t\t<thead>\n" +
		"\t\t\t\t<tr>\n" +
		"\t\t\t\t\t<th>File name</th>\n" +
		"\t\t\t\t\t<th>Issue</th>\n" +
		"\t\t\t\t\t<th>Severity</th>\n" +
		"\t\t\t\t\t<th>CVE/CWE</th>\n" +
		"\t\t\t\t\t<th>Description</th>\n" +
		"\t\t\t\t\t<th>Possible fix</th>\n" +
		"\t\t\t\t</tr>\n" +
		"\t\t\t</thead>\n" +
		"\t\t\t<tbody>\n"
	for _, entry := range log.Entries {
		result += "\t\t\t\t<tr>\n" +
			"\t\t\t\t\t<td>" + entry.Filename + "</td>\n" +
			"\t\t\t\t\t<td>" + entry.Issue_name + "</td>\n" +
			"\t\t\t\t\t<td>" + entry.Severity + "</td>\n" +
			"\t\t\t\t\t<td>" + entry.CVE + entry.CWE + "</td>\n" +
			"\t\t\t\t\t<td>" + entry.Description + "</td>\n" +
			"\t\t\t\t\t<td>" + entry.Fix + "</td>\n"
	}
	result += "\t\t\t\t</tr>" +
		"\t\t\t</tbody>\n" +
		"\t\t</table>\n" +
		"</body>"
	return result
}

func writeInFile(html string) {
	f, err := os.Create("/opt/scan/metascan_results/result.html")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(html)

	if err2 != nil {
		log.Fatal(err2)
	}
}

func WriteHtml(log Log.Log) bool{
	var resultAsHtml = convertToHtml(log)
	writeInFile(resultAsHtml)
	return true
}
