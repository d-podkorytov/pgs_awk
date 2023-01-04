// All input from is strings 
// use channels
// r is accumulated state
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
 	output_mask  = flag.String("output_mask","%v", "Output mask %v or %#v ")
 	fold_mask    = flag.String("fold_mask","%v\ncall(%#v,%#v)", "Fold mask")
 	result_mask  = flag.String("result_mask","%v", "Result mask")

 	script       = flag.String("code","code.gmacro", "GoMacro code file")
)

func main() {

	//flag.Usage = usage
	flag.Parse()

inputFile := *script 
if *tracing  {log.Printf("Code file: %#v args=%#v mask=%#v \n",*script,flag.Args(),*output_mask)}

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
    r:=""
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        r+=RunGomacro2(*tracing,codeContent,r,scanner.Text()) // Println will add back the final '\n'
    }
    if err := scanner.Err(); err != nil {
        log.Println("Reading standard input:", err)
    }

    fmt.Println(*result_mask,r)
}

func RunGomacro2(tracing bool,code string,r,toeval string) string {
    interp := fast.New()
    // slow call but is scripting
    //vals, _ := interp.Eval(code+"\n"+"call(\""+toeval+"\")\n")
    e:=fmt.Sprintf(*fold_mask,code,toeval,r)
     if tracing { log.Printf("Eval(%#v) \n",e)}
    vals, vals_err := interp.Eval(e)    

if vals_err!=nil {log.Printf("Eval error %#v for code %#v\n",vals_err,e)
                  //os.Exit(1)
                 }

for i:=0;i<len(vals);i++ {
 if tracing { log.Printf("'%v'[%v] -> %#v \n",toeval,i,vals[i].ReflectValue())}
            r+=fmt.Sprintf(*output_mask,vals[i].ReflectValue()) 
                         } 
 return r
}

func Gomacro_test() {
 interp  := fast.New()
 vals, err_vals := interp.Eval("func call(i int) int {\nreturn i}\ncall(1+2)\n")
 fmt.Printf("\nresult %v %#v\n",vals[0].ReflectValue(),  err_vals) 

}


