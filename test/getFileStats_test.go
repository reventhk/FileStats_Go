package test

import (

	"testing"
	file_stats "github.com/reventhk/FileStats_Go"
	"strings"
	"fmt"
	 
	 
	 
)

func TestAddFile(t *testing.T){

	fs := file_stats.FileMetadata{"/Users/revanthkumar/Desktop",10, false}
	err := file_stats.AddFile(fs)
	if !strings.Contains(err.Error(),"Given Directory, provide file path"){
		t.Errorf("AddFile func failed for Dir path, Expected %v ","Given Directory, provide file path")
	}else{
		t.Logf("AddFile func is success, Expected %v got %v","Given Directory, provide file path", err.Error() )
	}

}


/* // TestFatal is used to do tests which are supposed to be fatal
func TestFatal(t *testing.T) {
    origLogFatalf := file_stats.LogFatalf

    // After this test, replace the original fatal function
    defer func() { file_stats.LogFatalf = origLogFatalf } ()
	
    errors := []string{}
    file_stats.LogFatalf = func(format string, args ...interface{}) {
        if len(args) > 0 {
            errors = append(errors, fmt.Sprintf(format, args))
        } else {
            errors = append(errors, format)
        }
	}
	file_stats.GetStats()
    if len(errors) != 1 {
        t.Errorf("excepted one error, actual %v", len(errors))
    }
} */

