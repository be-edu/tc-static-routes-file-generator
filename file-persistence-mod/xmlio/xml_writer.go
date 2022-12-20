// Package xmlio provides functionality for xml file input and output operations.
package xmlio

import (
	"encoding/xml"
	"fmt"
	"os"
)

// WriteDataToXmlFile writes the data to the xml file.
func WriteDataToXmlFile(fileName string, header string, data interface{}) error {
	dataMarshalled, err := xml.MarshalIndent(data, "", "	")

	if err != nil {
		return err
	}

	dataToWrite := ""
	if header != "" {
		// Put together header and marshalled bytes
		dataToWrite = fmt.Sprintf("%s\n%s", header, string(dataMarshalled))
	} else {
		dataToWrite = string(dataMarshalled)
	}

	dataMarshalled = []byte(dataToWrite)

	if err != nil {
		return err
	}

	var fileMode os.FileMode = 0644 // the 0644 is octal representation of the file mode
	err = os.WriteFile(fileName, dataMarshalled, fileMode)

	if err != nil {
		return err
	}

	return nil
}
