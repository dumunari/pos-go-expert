package main

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	s3Client *s3.S3
	s3Bucket string
	wg       sync.WaitGroup
)

func init() {
	// Initialize AWS S3 client
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String("us-east-1"),
			Credentials: credentials.NewStaticCredentials(
				"---",
				"---",
				"",
			),
		},
	)
	if err != nil {
		panic(err)
	}
	// Create S3 client
	s3Client = s3.New(sess)
	s3Bucket = "---"
}

func main() {
	// Create tmp directory if it doesn't exist
	if _, err := os.Stat("./tmp"); os.IsNotExist(err) {
		if err := os.Mkdir("./tmp", 0755); err != nil {
			panic(err)
		}
	}
	// Open the tmp directory
	dir, err := os.Open("./tmp")
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	// Create channels for controlling uploads and handling errors

	// Control channel to limit concurrent uploads
	// 100 is the maximum number of concurrent uploads
	uploadControl := make(chan struct{}, 100)

	// Channel to handle files that failed to upload
	// 10 is the maximum number of files that can be retried
	errorFileUpload := make(chan string, 10)

	// Start a goroutine to handle files that failed to upload
	// This goroutine will listen for files that failed to upload
	// and will retry uploading them
	// It will also limit the number of concurrent uploads
	// by using the uploadControl channel
	go func() {
		// for {
		// 	select {
		// 	case filename := <-errorFileUpload:
		// 		uploadControl <- struct{}{}
		// 		wg.Add(1)
		// 		go uploadFile(filename, uploadControl, errorFileUpload)
		// 	}
		// }
		for filename := range errorFileUpload {
			// When a file fails to upload, it will be sent to the errorFileUpload channel
			// This will trigger the goroutine to retry uploading the file
			// The uploadControl channel will be used to limit the number of concurrent uploads
			uploadControl <- struct{}{}
			// Add a wait group to wait for the upload to finish
			// This will ensure that the main goroutine waits for the upload to finish
			// before continuing to the next file
			// This will also ensure that the uploadControl channel is emptied
			// before the next file is uploaded
			// This will prevent the uploadControl channel from being full
			// and will allow the goroutine to continue uploading files
			wg.Add(1)
			// Call the uploadFile function to upload the file
			go uploadFile(filename, uploadControl, errorFileUpload)
		}
	}()

	// Read files from the tmp directory and upload them to S3
	// Use a loop to read files one by one
	for {
		files, err := dir.ReadDir(1)
		if err != nil {
			// If we reach the end of the directory, break the loop
			// If the error is io.EOF, it means we reached the end of the directory
			// and we can break the loop
			if err == io.EOF {
				break
			}

			// If an error occurs while reading the directory, print the error and continue
			// This will not stop the program, but will log the error
			// and continue to the next iteration
			fmt.Printf("Error reading directory: %s\n", err)
			continue
		}
		// If no files are found, break the loop
		if len(files) == 0 {
			break
		}
		// If files are found, upload the first file
		// Increment the wait group to wait for the upload to finish
		wg.Add(1)
		// Add the file to the uploadControl channel to limit concurrent uploads
		uploadControl <- struct{}{}
		// Call the uploadFile function to upload the file
		go uploadFile(files[0].Name(), uploadControl, errorFileUpload)
	}
	wg.Wait()
}

// uploadFile uploads a file to S3
// It takes the file name, upload control channel, and error channel as parameters
// It uses the S3 client to upload the file to the specified bucket
// If an error occurs during upload, it sends the file name to the error channel
// It also empties the upload control channel to allow the next file to be uploaded
func uploadFile(filename string, uploadControl <-chan struct{}, errorFileUpload chan<- string) {
	defer wg.Done()

	completeFileName := fmt.Sprintf("./tmp/%s", filename)
	fmt.Printf("Uploading file %s to bucket %s\n", completeFileName, s3Bucket)
	f, err := os.Open(completeFileName)

	if err != nil {
		fmt.Printf("Error opening file %s\n", completeFileName)
		<-uploadControl                     // Empty the upload control channel
		errorFileUpload <- completeFileName // Send the file name to the error channel
		return
	}
	defer f.Close()

	// Upload the file to S3
	// Use PutObject to upload the file
	// Use the bucket name and the file name as the key
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(filename),
		Body:   f,
	})

	// Check for errors during upload
	// If an error occurs, send the file name to the error channel
	if err != nil {
		fmt.Printf("Error uploading file %s\n", completeFileName)
		<-uploadControl // esvazia o canal
		errorFileUpload <- completeFileName
		return
	}
	fmt.Printf("File %s uploaded successfully\n", completeFileName)
	<-uploadControl // Empty the upload control channel
}
