// example - rand.NewSource() and rand.New()

import (
    "fmt" // import the fmt package which allows you to use text
          //formatting, reading input & printing output functions

    "math/rand" //import the rand package which allows you to
                //generate random numbers
    
    "time" // import the package which will provide time functionality to measure time
)


source := rand.NewSource(time.Now().UnixNano())  // create a source from a timestamp
randomizer := rand.New(source)                   // we can draw from randomizer
fmt.Println(randomizer.Intn(11))                 // draw using our timestamp seed


source = rand.NewSource("qwerty")          // create a source from "qwerty"
randomizerTwo := rand.New(source)          // create a new non conflicting randomizer
fmt.Println(randomizerTwo.Intn(11))        // will be different from 'randomizer'

randomizerThree := rand.New(source)        // create a new non conflicting randomizer
fmt.Println(randomizerThree.Intn(11))      // will be the SAME result as 'randomizerTwo'

