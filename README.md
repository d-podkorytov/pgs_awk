# OS pipes processing with MAP/FOLD by Gomacro and Golang
## Why
- For replacement pure portable awk in my every day data processing;
- For learning Gomacro scripting and using it in future projects
- For next using in my distribute system over NATS messaging
- As crossplatform tool for pipes datastream processing in Linux/Windows/Unix/MacOS

## Requirements
- OS Linux/Windows and probably FreeBSD and MacOS
- Golang v1.19
- Gomacro from "https://github.com/cosmos72/gomacro"

## For more about Gomacro see 
- Docs: https://github.com/cosmos72/gomacro/tree/master/doc
- Examples: https://github.com/cosmos72/gomacro/tree/master/_example

## Compilation
> go build pgs_maps.go

> go build pgs_folds.go

## pgs_maps 

*pgs_maps* read lines from stdin and apply gomacro function for every line,so it looks like programming with awk but with power of Golang scripting

Mapping function may be arbitrary but input and output is strings.  

Example of file *strings.gmacro*: 
> func call(inp string) string {return inp+inp}

> C:\go.w\FastHTTPD_GET_NATS\NATS_Pipes\pgs_awk>pgs_map.exe -code strings.gmacro -tracing  < input_file

 Output:

> 2023/01/03 18:58:07 Code file:"strings.gmacro" args=[] mask="%#v" 

> 2023/01/03 18:58:07 Code:func call(inp string) string {return inp+inp}

> 2023/01/03 18:58:07 '1'[0] -> "11"

> "11"

> 2023/01/03 18:58:07 '2'[0] -> "22"

> "22"

> 2023/01/03 18:58:07 'a'[0] -> "aa"

> "aa"

> 2023/01/03 18:58:07 'b'[0] -> "bb"

> "bb" 

## pgs_folds

*pgs_folds* also read lines from stdin and apply gomacro function for every line with accumulator, like fold in FP languages

Folding function may be arbitrary but input and output is strings

Example for fold transformation (inside file *strings_fold.gmacro*):
> func call(inp string, acc string) string {return acc+inp}

> pgs_folds.exe -code strings_fold.gmacro -tracing < input_file
 
 Output:
>  C:\go.w\FastHTTPD_GET_NATS\NATS_Pipes\pgs_awk>pgs_folds.exe -code strings_fold.gmacro  <1 

>  111211121112311121112111231112111211123a111211121112311121112111231112111211123a111211121112311121112111231112111211123ab

## More examples

See *ls2hmtl.sh* and *ls2html.gomacro* files as an example for convertation directory to HTML text
