package main
 
import (

	"fmt"
	file_stats "byndr/FileOpr"
)


func main(){
	
	 

	fs := file_stats.FileMetadata{"/Users/revanthkumar/Desktop/Visa_Questionnaire_Revanth.xlsx",10, false}

	//fs := file_stats.FileMetadata{"/Users/revanthkumar/Desktop/20663801_1644573678906979_5240984234551598314_n.jpg",10, false}

	//fs := file_stats.FileMetadata{"/Users/revanthkumar/Desktop/revanth_kumar_latest.docx",10, false}

	//fs := file_stats.FileMetadata{"/Users/revanthkumar/Desktop/revanth kumar_latest.pdf",10, false}

	//fs := file_stats.FileMetadata{"/Users/revanthkumar/Desktop/Downloads/go1.13.4.darwin-amd64.pkg",10, false}

	err := file_stats.AddFile(fs)
	
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(file_stats.GetStats())

}