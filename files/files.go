package files


import (
	"fmt"
	"os"
)

func Fmt_write(filename string, name string) {

	// try {
	// 	file, err := os.Create(filename)
	// } catch () {
	// }


	//file, err := os.Create(filename)
	file, err := os.OpenFile(filename,os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("Error creating file:", err)
		//return
	}

	if file != nil {

		// Ensure the file gets closed when the function ends
		defer file.Close() 
		// Write content to the file
		_, err = fmt.Fprintln(file, "Hello " + name + " , this is some text in the file!")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
		// Print a success message
		fmt.Println("Text written to file successfully!")
	}
	
}