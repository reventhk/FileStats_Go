package file_stats

import (

	"fmt"
	"os"
	"log"
	"path/filepath"
	"time"
	"sort"
	"os/exec"
	"strings"

)

type FileMetadata struct {
	Path string `json:"path"` // the file's absolute path
	Size int64 `json:"size"` // the file size in bytes
	IsBinary bool `json:"is_binary"` // whether the file is a binary file or a simple text file
}

type TimeStampData struct {
	Path string `json:"path"` // the file's absolute path
	TimeStamp int64 `json:"time_stamp"` // the file size in bytes
}

type LargestFileInfo struct {
	Path string `json:"path"`
	Size int64 `json:"size"`
}

type ExtInfo struct {
	Extension string `json:"extension"`
	NumOccurrences int64 `json:"num_occurrences"`
}

type FileStats struct {
	NumFiles int64 `json:"num_files"`
	LargestFile LargestFileInfo `json:"largest_file"`
	AverageFileSize float64 `json:"avg_file_size"`
	MostFrequentExt ExtInfo `json:"most_frequent_ext"`
	TextPercentage float32 `json:"text_percentage"`
	MostRecentPaths []string `json:"most_recent_paths"`
}

type MetaJson struct {
	FileList []FileMetadata `json:"file_list"`
	TimeStampData []TimeStampData `json:"timestamp_list"`
}

var tmpFile string = "./file_data.tmp"
var lastAddedFileLengh int = 10
var textFilesTypes = []string{"text/","wordprocessingml.document", "application/pdf", "spreadsheetml.sheet","vnd.ms-excel", "vnd.ms-powerpoint", "presentationml.presentation", "msword", "application/rtf" }

 

// Function to add file into system, Accepts 
//  FileMetadata  structuretype 
//	FileMetadata struct {
//	Path string `json:"path"` // the file's absolute path
//	Size int64 `json:"size"` // the file size in bytes
//	IsBinary bool `json:"is_binary"` // whether the file is a binary file or a simple text file
//}
// Provide default values for Size and IsBinary bool
func AddFile(metadata FileMetadata) error {
	if _, err := os.Stat(metadata.Path); err == nil {
		var storedData MetaJson
		 
		calcFileSize(&metadata)

		if _, err := os.Stat(tmpFile); err == nil {
			if err := Load(tmpFile, &storedData); err != nil {	  		 
				  log.Fatalln(err)
			}else {
			
			}
		}
		 
		checkAndUpdateJson(&storedData, metadata)

		if err := Save(tmpFile, storedData); err != nil {
			 
			log.Fatalln(err)
		  }
			// load it back
		
		  //fmt.Println(o2)

		  return nil
	  } else if os.IsNotExist(err) {
		 
		fmt.Println("File is not exists !!")

		return err
	  
	  } else {

		return err
	  }


}

func checkAndUpdateJson(tmpJson *MetaJson, metadata FileMetadata) {

	var is_file_aaded bool
	var timeStampData TimeStampData
	
	if len(tmpJson.FileList) > 0 {
	 for _ , metadata_item := range tmpJson.FileList { 
		 
		if metadata_item.Path == metadata.Path{
			is_file_aaded = true
		}
	 } 
	}else{
		is_file_aaded = false
	}

	metadata.IsBinary = isBinaryFile(metadata.Path)



	if !is_file_aaded{
		tmpJson.FileList = append(tmpJson.FileList, metadata)
		timeStampData.Path = metadata.Path
		timeStampData.TimeStamp = time.Now().Unix()

		tmpJson.TimeStampData = append(tmpJson.TimeStampData, timeStampData) 

	 }
	

}

//Returns all files stats till now added by user in below format 
//type FileStats struct {
//	NumFiles int64 `json:"num_files"`
//	LargestFile LargestFileInfo `json:"largest_file"`
//	AverageFileSize float64 `json:"avg_file_size"`
//	MostFrequentExt ExtInfo `json:"most_frequent_ext"`
//	TextPercentage float32 `json:"text_percentage"`
//	MostRecentPaths []string `json:"most_recent_paths"`
//}
// TextPercentage may not give correct value on windows os

func GetStats() FileStats{

	var storedData MetaJson
	var fileStats FileStats
	var largeFileInfo LargestFileInfo
	var tmpLargeSize int64
	var tmpLargeName string
	var totalFileSize float64
	var fileExtArr []string
	var textFilesListLen int


	if _, err := os.Stat(tmpFile); err == nil {
		if err := Load(tmpFile, &storedData); err != nil {   
			  log.Fatalln(err)
		}else {
		}
	}

	for _ , metadata_item := range storedData.FileList { 

		if tmpLargeSize < metadata_item.Size{
			 tmpLargeSize = metadata_item.Size
			 tmpLargeName = metadata_item.Path
		}

		totalFileSize = totalFileSize + float64(metadata_item.Size)

		fileExtArr = append(fileExtArr,  filepath.Ext(metadata_item.Path))
		if metadata_item.IsBinary == false {
			textFilesListLen = textFilesListLen + 1
		}
	}

	fileStats.NumFiles = int64(len(storedData.FileList))

	largeFileInfo.Size = tmpLargeSize
	largeFileInfo.Path = tmpLargeName

	fileStats.LargestFile = largeFileInfo
	
	fileStats.AverageFileSize = totalFileSize/float64(len(storedData.FileList))

	fileStats.MostFrequentExt = getMostFreqFileExt(fileExtArr)

	fileStats.MostRecentPaths =  getlastAddedFilesList(storedData.TimeStampData)

	fileStats.TextPercentage = float32((float32( textFilesListLen)/float32( len(storedData.FileList))) * 100)
 	 
	 return fileStats

}

// Funtion to get most frequest files(extentions) added by user
func getMostFreqFileExt(arr []string) ExtInfo{

	dict:= make(map[string]int)
	for _ , num :=  range arr {
	dict[num] = dict[num]+1
	}

	var tmpKey string
	var tmpVal int 
	var tmpExt ExtInfo
	for key, value := range dict {
		if tmpVal < value {
			tmpVal = value
			tmpKey = key
		}

	}

	tmpExt.Extension = tmpKey
	tmpExt.NumOccurrences = int64(tmpVal)

	return tmpExt
}

// funtion to get last 10 files added (sorted with time stamp)
func getlastAddedFilesList(timeStampData []TimeStampData) []string {

	if len(timeStampData) < lastAddedFileLengh {
		lastAddedFileLengh = len(timeStampData)
	}

	var data []string 
	data = make([]string, lastAddedFileLengh, lastAddedFileLengh)
	 // Sort by timestamp, keeping original order or equal elements.
	 sort.SliceStable(timeStampData, func(i, j int) bool {
    return timeStampData[i].TimeStamp < timeStampData[j].TimeStamp
	})
	
	timeStampData = timeStampData[0:lastAddedFileLengh]
	for index , item := range timeStampData {

		data[index] = item.Path
	} 
 
return data
}

// calculate given file size
func calcFileSize(metadata *FileMetadata) {

	fi, _ := os.Stat(metadata.Path);
	
	switch mode := fi.Mode(); {
    case mode.IsDir():
        // do directory stuff
		log.Fatalln("Given Directory, provide file path")
         
    }

	metadata.Size = fi.Size()
} 

// check weather file is binary or not.
// Here textFilesTypes array considered as text file types, 
// remaining any other file type considered as binary
// as we are using bash commands, this function may not work in windows
func isBinaryFile(path string) bool {
	
	out, err := exec.Command("file","--mime", path).Output()
	
	 isBinaryFileBool := true
	if err != nil {
		 log.Fatal(err)
	 }

	out_string := string(out)
 
	for _, item := range textFilesTypes {
 
	   if strings.Contains(out_string, item) {
 
			isBinaryFileBool = false
			break
	   }
	}
return isBinaryFileBool
}
