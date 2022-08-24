/* Alta3 Research | RZFeeser
   Map - hosts to IP:port resolution  */

package main

import (  
    "fmt"
)

func main() {  
    
    // create a map type, "languageExt"
    var languageExt = map[string]string{
        "Python": ".py",
        "Golang": ".go",
        "Jave": ".java",
        "Ansible": ".yml",
        "C++": ".cpp",
    }
    
    // Print the languageExt Map
    fmt.Println(languageExt)

}
