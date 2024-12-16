package dmarc

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

type policyEvaluated struct {
	Dispotion string `xml:"disposition"`
	Dkim      string `xml:"dkim"`
	Spf       string `xml:"spf"`
}

type row struct {
	SourceIp        string          `xml:"source_ip"`
	Count           int             `xml:"count"`
	PolicyEvaluated policyEvaluated `xml:"policy_evaluated"`
}

type identifiers struct {
	EnvelopeTo   string `xml:"envelope_to"`
	EnvelopeFrom string `xml:"envelope_from"`
	HeaderFrom   string `xml:"header_from"`
}

type dkim struct {
	Domain   string `xml:"domain"`
	Selector string `xml:"selector"`
	Result   string `xml:"result"`
}

type spf struct {
	Domain string `xml:"domain"`
	Scope  string `xml:"scope"`
	Result string `xml:"result"`
}

type record struct {
	Row         row         `xml:"row"`
	Identifiers identifiers `xml:"identifiers"`
	AuthResults []dkim      `xml:"auth_results>dkim"`
	Spf         spf         `xml:"auth_results>spf"`
}

type policyPublished struct {
	Domain string `xml:"domain"`
	Adkim  string `xml:"adkim"`
	Aspf   string `xml:"aspf"`
	P      string `xml:"p"`
	Sp     string `xml:"sp"`
	Pct    int    `xml:"pct"`
	Fo     string `xml:"fo"`
}

type dateRange struct {
	Begin int `xml:"begin"`
	End   int `xml:"end"`
}

type reportMetadata struct {
	OrgName   string    `xml:"org_name"`
	Email     string    `xml:"email"`
	ReportId  string    `xml:"report_id"`
	DateRange dateRange `xml:"date_range"`
}

type feedback struct {
	Version         string          `xml:"version"`
	ReportMetadata  reportMetadata  `xml:"report_metadata"`
	PolicyPublished policyPublished `xml:"policy_published"`
	Records         []record        `xml:"record"`
}

var ParserCmd = &cobra.Command{
	Use: "parse",
	Run: func(cmd *cobra.Command, args []string) {
		file, _ := cmd.Flags().GetString("file")
		parseXml(file)
	},
}

var Filename string

func parseXml(filePath string) {
	fmt.Println("Parsing XML...")

	// Open the XML file
	xmlFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()

	// Read the XML file
	byteValue, err := io.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Parse the XML data
	var feedback feedback
	err = xml.Unmarshal(byteValue, &feedback)
	if err != nil {
		fmt.Println("Error parsing XML:", err)
		return
	}

	fmt.Println(feedback)

}

func init() {
	ParserCmd.Flags().StringVarP(&Filename, "file", "f", "", "XML file to parse")
	ParserCmd.MarkFlagRequired("file")
}
