# OS pipes processing with MAP/FOLD by Gomacro and Golang
## Why
- For replacement awk in my every day data processing;
- For learning Gomacro scripting
- For next using in my distribute system over NATS messaging
- As crossplatform tool for pipes datastream processing in Linux/Windows/Unix 

## pgs_maps 

*pgs_maps* read lines from stdin and apply gomacro function for every line,so it looks like programming with awk but with power of Golang scripting

Mapping function may be arbitrary but input and output is strings.  

Example of file strings.gmacro: 
 func call(inp string) string {return inp+inp}

> pgs_map.exe -code strings.gmacro -tracing  < input_file
>
> Output:
-  C:\go.w\FastHTTPD_GET_NATS\NATS_Pipes\pgs_awk>pgs_maps.exe -code strings.gmacro -tracing < file
- 2023/01/03 18:58:07 Code file:"strings.gmacro" args=[] mask="%#v" 
- 2023/01/03 18:58:07 Code:func call(inp string) string {return inp+inp}
- 2023/01/03 18:58:07 '1'[0] -> "11"
- "11"
- 2023/01/03 18:58:07 '2'[0] -> "22"
- "22"
- 2023/01/03 18:58:07 'a'[0] -> "aa"
- "aa"
- 2023/01/03 18:58:07 'b'[0] -> "bb"
- "bb" 

## pgs_folds

*pgs_folds* also read lines from stdin and apply gomacro function for every line with accumulator, like fold in FP languages

Folding function may be arbitrary but input and output is strings

Example for fold transformation (inside file strings_fold.gmacro):
 func call(inp string, acc string) string {return acc+inp}

> pgs_folds.exe -code strings_fold.gmacro -tracing < input_file
> 
> Output:
-  C:\go.w\FastHTTPD_GET_NATS\NATS_Pipes\pgs_awk>pgs_folds.exe -code strings_fold.gmacro -tracing  <1 
-  111211121112311121112111231112111211123a111211121112311121112111231112111211123a111211121112311121112111231112111211123ab

