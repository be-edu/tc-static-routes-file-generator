package main

import (
	"file-persistence-mod/fileinfo"
	"file-persistence-mod/xmlio"
	"fmt"
	"log"
	"std-io-mod/stdreader"
	"tc-util-mod/tcutil"
)

func main() {
	printHeaderInformation()
	printEmptyLine()

	baseIpAddress, err := promptUserForBaseIpAddress()
	checkAndHandleError(err)

	tcConfig := tcutil.GetTcConfigBasedOnBaseIpAddressWithFirstThreeSegments(baseIpAddress)

	err = xmlio.WriteDataToXmlFile(tcutil.XmlFileName, tcutil.XmlFileHeader, tcConfig)
	checkAndHandleError(err)

	err = printProcessingInformation()
	checkAndHandleError(err)

	printEmptyLine()
	promptUserToClose()
}

func printHeaderInformation() {
	printSeparationLine()
	fmt.Println("TwinCAT static routes file generator")
	printSeparationLine()
	fmt.Println("Description:")
	fmt.Println("The programm will generate the last numbers ([0-255]) for the provided first")
	fmt.Println("three numbers of the desired IPv4 range (e.g. 192.168.1).")
	fmt.Println("Based on that, the static routes file is generated.")
	printSeparationLine()
}

func printSeparationLine() {
	fmt.Println("----------------------------------------------------------------------------")
}

func printEmptyLine() {
	fmt.Println()
}

func promptUserForBaseIpAddress() (string, error) {
	inputPromptInformation := getInputPromptInformation()
	baseIpAddress, err := stdreader.GetDataFromUser(inputPromptInformation, '\n')

	if err != nil {
		return "", err
	}

	return baseIpAddress, nil
}

func getInputPromptInformation() string {
	firstPart := "Enter first three numbers of the desired IPv4 range (e.g. 192.168.1):"
	secondPart := "-> "
	inputPromptInformation := fmt.Sprintf("%s\n%s", firstPart, secondPart)

	return inputPromptInformation
}

func checkAndHandleError(err error) {
	if err != nil {
		handleError(err)
	}
}

func handleError(err error) {
	fmt.Println("The programm will be shortly terminated, because the following error occurred:")
	log.Fatal(err)
}

func printProcessingInformation() error {
	absoluteFileName, err := fileinfo.GetAbsoluteFileNameFromWorkingDirectory(tcutil.XmlFileName)

	if err != nil {
		return err
	}

	fmt.Printf("File successfully generated under %s\n", absoluteFileName)

	return nil
}

func promptUserToClose() {
	inputPromptInformation := "Press enter to close application\n"
	stdreader.GetDataFromUser(inputPromptInformation, '\n')
}
