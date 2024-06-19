# VORTO CHALLENGE

I did not like my original version of the vorto challenge.
So I started from scratch last night.
In the end, I cheated, just to get the Output to pass the Python code.

In my study of the data, I created graphics to get visuals.




Based on [this Google Drive folder](https://drive.google.com/drive/folders/1Jb7FmR5Ftrg0jwgIJ-n_oKwOjyJ4gDHI)


## Objective
develope system that will compute lowest cost Routes from a Collection of LOADs.

## Logic

### Graphics
Observe ./output/graphic/v230
I was not sure about distribution of tags.





### 12HR
No driver is permitted to drive beyond the 12 hours
in a single route.


### Problem File
A collection of Load.txt Files.
Delivered as part of Challenge

### Load Problem 
A LOAD is combination of 
A collection of LatLong coordintate
2 points per LOAD
- PickUp
- DropOff

### DEPOT
A starting and End Location at
(Lat:0.0, Long:0.0)
All Routes must begin and End at Depot with out
violating 12HR



#### S100_Load 
A LOAD is combination of 
- P1_startPoint
- P2_PickUp
- P3_DropOff
- Pickedup // True if the Load has already been picked up.



### Route
A route is a collection of Loads performed by One Driver
The driver can not violate 12HR ( 12 hour rule )
There are X

