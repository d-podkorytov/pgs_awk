// All input from is strings 
package main

import (
   "fmt"
   "io/ioutil"
   "os"
   "bufio"
   "flag"
   "log"  
   "github.com/cosmos72/gomacro/fast"

)

var (
 	tracing      = flag.Bool("tracing", false, "Tracing")
 	output_mask  = flag.String("mask","%#v", "Output mask")
 	fold_mask    = flag.String("fold_mask","%v\ncall(%#v)", "Fold mask")

 	script       = flag.String("code","code.gmacro", "GoMacro code file")
)

func main() {

	//flag.Usage = usage
	flag.Parse()

inputFile := *script 
if *tracing {log.Printf("Code file:%#v args=%v mask=%#v \n",*script,flag.Args(),*output_mask)}

if len(os.Args)<2   {log.Printf("Run it:\n%v -code script_file.gomacro [-mask \"%v\"] [-tracing] < input_file\n",os.Args[0])
                     os.Exit(1) }

   // read the whole content of file and pass it to file variable, in case of error pass it to err variable
   code, err := ioutil.ReadFile(inputFile)
   if err != nil {
      log.Printf("Could not read the gomacro script file '%v' due to this %s error \n", inputFile,err)
      os.Exit(1)
   }
   codeContent := string(code)
   // print code file content
   if *tracing  {log.Println("Code:\n",codeContent)}

    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        fmt.Println(RunGomacroMaps(*tracing,codeContent,scanner.Text())) // Println will add back the final '\n'
    }
    if err := scanner.Err(); err != nil {
        log.Println("reading standard input:", err)
    }


}

/*
var (
 	tracing      = flag.Bool("tracing", false, "Tracing")
 	output_mask  = flag.String("mask","%#v", "Output mask")
 	fold_mask    = flag.String("fold_mask","%v\ncall(%#v)", "Fold mask")

 	script       = flag.String("code","code.gmacro", "GoMacro code file")
)

*/
func RunGomacroMaps(tracing bool,code string,toeval string) string {
    interp := fast.New()
    // slow call but is scripting
    //vals, _ := interp.Eval(code+"\n"+"call(\""+toeval+"\")\n")    
    vals, _ := interp.Eval(fmt.Sprintf(*fold_mask,code,toeval))    

    r:=""

for i:=0;i<len(vals);i++ {
 if tracing { log.Printf("'%v'[%v] -> %#v \n",toeval,i,vals[i].ReflectValue())}
            r+=fmt.Sprintf(*output_mask,vals[i].ReflectValue()) 
                         } 
 return r
}

func Gomacro_test() {
 interp  := fast.New()
 vals, err_vals := interp.Eval("func call(i int) int {\nreturn i}\ncall(1+2)\n")
 //if *tracing {log.Printf("\nresult %v %#v\n",vals[0].ReflectValue(),  err_vals)} 
 fmt.Printf("\nresult %#v \n",vals[0].ReflectValue()) 
 if err_vals != nil {log.Printf("\nerror %#v\n", err_vals)} 

// fmt.Printf("\nresult %v\n",RunGomacro("func call(i int) int {\n return i }\n","1+2")) 
}


