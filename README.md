# Go - File to folder

[![Go Report Card](https://goreportcard.com/badge/github.com/dtroncy/go-file-to-folder)](https://goreportcard.com/report/github.com/dtroncy/go-file-to-folder)

In a folder, transfert each file in a specific folder with the same name.

From :
````
Root Folder
|
|-- file1.txt
|
|-- file2.txt
````
to :
````
Root Folder
|
|-- file1 (it's a folder with the name of the file)
   |
   |-- file1.txt
|
|-- file2 (it's a folder with the name of the file)
   |
   |-- file2.txt
````

To use it, copy script to root folder and launch it :

````
./go-file-to-folder
````
