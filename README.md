# MathComputation

To Install Dependencies:
sh install.sh

To Build:
go build -o MathComp.exe

To Run:
MathComp.exe

To Test:
Hosting Two Rest API's:
1. AddTwoNumber :
    http://localhost:9001/api/v1/math/add?first=2&second=4

2. To Check API Usage:
    http://localhost:9001/api/v1/math/usage?endpoint=add
