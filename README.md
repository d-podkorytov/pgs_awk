# OS pipes processing with MAP/FOLD by Gomacro and Golang

## pgs_map 

pgs_map read lines from stdin and apply gomacro function for every line, like awk but with power of Golang

Mapping function may be arbitrary but input and output is strings.  

Example of file strings.gmacro: 
 func call(inp string) string {return inp+inp}

pgs_map.exe -code strings.gmacro -tracing  < input_file

Output:
 C:\go.w\FastHTTPD_GET_NATS\NATS_Pipes\pgs_awk>pgs_map.exe -code strings.gmacro -tracing < file
2023/01/03 18:58:07 Code file:"strings.gmacro" args=[] mask="%#v" 
2023/01/03 18:58:07 Code:func call(inp string) string {return inp+inp}
2023/01/03 18:58:07 '1'[0] -> "11"
"11"
2023/01/03 18:58:07 '2'[0] -> "22"
"22"
2023/01/03 18:58:07 'a'[0] -> "aa"
"aa"
2023/01/03 18:58:07 'b'[0] -> "bb"
"bb" 

## pgs_folds

pgs_map also read lines from stdin and apply gomacro function for every line with accumulator, like fold in FP languages

Folding function may be arbitrary but input and output is strings

Example for fold transformation (file strings_fold.gmacro):
 func call(inp string, acc string) string {return acc+inp}

pgs_folds.exe -code strings_fold.gmacro -tracing < input_file

Output:
 C:\go.w\FastHTTPD_GET_NATS\NATS_Pipes\pgs_awk>pgs_folds.exe -code strings_fold.gmacro -tracing  0<1 
 111211121112311121112111231112111211123a111211121112311121112111231112111211123a111211121112311121112111231112111211123ab

