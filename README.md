# FileStats_Go
Library that performs statistics and aggregations for file metadata provided to it

## import go package by using below command
go get -v -u github.com/reventhk/FileStats_Go
and 
write in your main package file.
file_stats "github.com/reventhk/FileStats_Go"



## Input
 AddFile(metaData FileMetadata)
 Function to add file into system, Accepts 
  FileMetadata  structuretype 
  
  provide metadata as below structure
  <pre>
	FileMetadata struct {
	Path string `json:"path"` // the file's absolute path
	Size int64 `json:"size"` // the file size in bytes
	IsBinary bool `json:"is_binary"` // whether the file is a binary file or a simple text file		 
        }
  </pre>	
 Provide default values for Size and IsBinary bool
 
 
 ## output
 
 GetStats() FileStats 
 
 Returns all files stats till now added by user in below format 
 <pre>
type FileStats struct {
	NumFiles int64 `json:"num_files"`
	LargestFile LargestFileInfo `json:"largest_file"`
	AverageFileSize float64 `json:"avg_file_size"`
	MostFrequentExt ExtInfo `json:"most_frequent_ext"`
	TextPercentage float32 `json:"text_percentage"`
	MostRecentPaths []string `json:"most_recent_paths"`
}
</pre>
 TextPercentage may not give correct value on windows os
 
 


