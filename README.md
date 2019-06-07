# Adaptive Huffman coding compressor
---
Course project of *Digital Media Technology*, Sun Yat-Sen University, Software Engineering 2017
### Getting Started
---
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.
#### Prequisites
This project was finished by the Go programming language
, therefore a [Go environment](https://golang.google.cn/) on your computer is neccessary.
#### Building & Installing
If you are using *Makefile* you can just enter this directory and type `make build`, then this program will be built in `./bin/huffman-compressor`, furthermore you can run the command `make install` to get it installed.  

If you wouldn't like to use *Makefile*, building this program by yourself is also alternative. The building command is just like this:
```bash
go build -o ./bin/huffmantree -v ./src
```
The built binary file can be excute directly, or you can copy it into your system path to add it into commands.
### Running the tests
---
There's only one test case, which tests basic compression on common text files. I thought we should have had more test cases, but there only one required in our task. Once you would like to add more test cases, just copy your file into `./test`.   

If you are using *Makefile*, it's easy to test. What you have to do is just copy your file into `./test` directory, then run the command `make test`. Scripts will automatically run this program to compress each file in the `./test` directory and extract them out. Then the scripts will compare the extracted file and original file, if they are the same, it will print `Test case ${filename of testcase} passed` on the screen, and it will print `Test case ${filename of testcase} failed` otherwise.  

If you wouldn't like to use *Makefile*, the only way to test this program is to [run it](#usage).

### <span id="usage">Usage</span>
---
You can execute this program without any parameter to get following help messages:
```
Usage:
  -dst destination file
        specify the destination file to be compressed or to be decompressed
  -ext
        extract file
  -src source file
        specify the source file to be compressed or to be decompressed
```
#### Examples  
**Compressing file ./story.txt into ./story.txt.huff**
```bash
huffman-compress -src ./story.txt -dst ./story.txt.huff
```
**Extracting file from ./story.txt.huff into ./story.ext.txt**
```bash
huffman-compress -src ./story.txt.huff -dst ./story.ext.txt --ext
```

### Author
---
* Skida12138, suqd@mail2.sysu.edu.cn

### License
---
This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details
