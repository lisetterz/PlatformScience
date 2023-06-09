# SDE Code Excercise - Platform Science

This code intents to solve a practical tests given by Platform Science.

# Requirements

## Inputs

In oder for this program to work, 2 files are needed which are alredy provided within the [data](/data) repository.

**Drivers file**
<br>
 [This](/data/10-list-drivers.txt) document provides a list of n elements that are names of the available drivers to deliver the shipments separated by line. You can edit the same or create another file and provide the location of the file in the [addressFile](https://github.com/lisetterz/PlatformScience/blob/CleanCode/main.go#LL13C30-L13C30) constant.

**Addresses file**
<br>
[This](/data/10-list-addresses.txt) file contains a list of addresses of the shipments that need to be delivered. You can also edit this file respectine the line separator or create a new file and provide it as the input in the [driversFile](https://github.com/lisetterz/PlatformScience/blob/CleanCode/main.go#L14).

> Notes: The number of drivers must match the number of addresses in order to work.

## Required tools
- You need to have Go 1.17+ installed in your machine to run the code. [Here](https://go.dev/doc/install) is a guide on how to install int.
- Fork this repository.

# How to run
Open the terminal in the location of the code and run the below command:
```
go run main.go
```
The program will print a list of the relations and the total SS.
