# Ascii Art Output

## Description

The initial project, ascii-art, was a project that was used to print text data in an artistic format. The follow-up project provided the user with an option to pick a font file to be used to display their characters. This proect was known as ascii-art-fs. There were a few font files, which we referred to as banner files. This files contained graphic representations of printable ascii characters. When the user needed to print their data on the console, the program would instead print the graphic representation of the same. 

This project is an extension by the virtue that it offers the user the ability to use an optional flag when invoking the program. This flag is then used to decide where the artistic representation will be saved. The new concepts introduced is output files. The user now has the option to save the results within a .txt file.


## Features
- Converts strings into ASCII art representations.
- Supports multiple banner formats including `shadow`, `standard`, and `thinkertoy`.
- Support for saving the graphic representation inside a file
- if a file is not provided, the program defaults to printing the results in the console


#### Limitations
- The program is designed to work with printable ASCII characters only.
- Characters outside the range of space (` `) to tilde (`~`) are not supported and may not render correctly.

#### Prerequisites
- Go programming language
- ASCII banner files in the correct format

#### Banner files [BANNER]

The mentioned project, ascii-art-fs only accepted three banner files. We will stick to tha same banner files which include: 

  + [`shadow`](https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/shadow.txt)
  + [`standard`](https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/standard.txt)
  + [`thinkertoy`](https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/thinkertoy.txt)



#### --output optional argument [OPTION]

When a user runs a program while following the strict guidelines under the usage section. Suppose that they provided a --output=temp.txt, an output similar to this will be saved within the provided file.

```bash

   _              _   _          
  | |            | | | |         
  | |__     ___  | | | |   ___   
  |  _ \   / _ \ | | | |  / _ \  
  | | | | |  __/ | | | | | (_) | 
  |_| |_|  \___| |_| |_|  \___/  
                                                                                 
                                                                                 

```

It should not go without mentioning that one can not provide a filename that already exists within the system; specifically not the banner files as a file to save the output in.


## Usage

The format: (replace the square bracket with the specific requirement): 

go run . [OPTION] [STRING] [BANNER]

**OPTION** allows you to specify the file where your art will be saved

**STRING** is the actual string whose representation will be printed

**BANNER** is also an optional argument to switch between the banner files to be used

**NOTE:** While the OPTION and BANNER are totally optional, in order to get the artistic representation, you will need to provide the STRING argument: 

go run . [STRING]

The following is an example of how to do so:

##### installation

```bash
git clone https://learn.zone01kisumu.ke/git/shfana/ascii-art-output.git
cd ascii-art-output
clear
```
##### run the program

```bash
go run . --output=test.txt "Hello"
```

Additionally, to switch between banner files, you can provide an additional argument which has to be the last argument

```bash
go run . --output=test.txt "Hello" "shadow"
```


## Contributions

#### contibutors

- [sfana](https://learn.zone01kisumu.ke/git/shfana)
- [bnyatoro](https://learn.zone01kisumu.ke/git/bnyatoro)
- [anoduor](https://learn.zone01kisumu.ke/git/anoduor)

#### To contribute

Go to the repository at
[ascii-art-output](https://learn.zone01kisumu.ke/git/shfana/ascii-art-output.git) and fork the repository. Clone your fork locally and make any changes you need to be made. After completing all git processes and pusing to your fork, you can issue a pull request and we assure you that your contributions wiil be taken into consideration.


