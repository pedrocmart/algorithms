# Run-length encoding algorithm (RLE) 

## RLE works by taking the occurrence of each repeating character and outputting that number along with a single character of the repeating sequence.

You need to implement an algorithm that applies the RLE to a given string.

### Example

- For 
```
inputString = "abbaaaac"
```
, the output should be solution(inputString):
``` 
"1a2b4a1c"
```

### Input/Output

[execution time limit] 4 seconds (go)

- [input] string inputString

A string consisting of lowercase English letters only.

Guaranteed constraints:
1 ≤ inputString.length ≤ 1000.

- [output] string