package main

import (
  "bufio"
  "fmt"
  "os"
  "regexp"
  "sync"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func telephoneNumbersInFile(path string) int {
  file, err := os.Open(path)
  check(err)
  defer file.Close()

  var telephone = regexp.MustCompile(`\(\d+\)\s\d+-\d+`)

  // do I need buffered channels here?
  jobs := make(chan string, 10)
  results := make(chan int, 10)

  // I think we need a wait group, not sure.
  wg := new(sync.WaitGroup)

  // start up some workers that will block and wait?
  for w := 1; w <= 3; w++ {
    wg.Add(1)
    go matchTelephoneNumbers(jobs, results, wg, telephone)
  }

  // go over a file line by line and queue up a ton of work
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    // Later I want to create a buffer of lines, not just line-by-line here ...
    jobs <- scanner.Text()

  }

  close(jobs)
  wg.Wait()
  close(results)

  // Add up the results from the results channel.
  // The rest of this isn't even working so ignore for now.
  counts := 0
  for v := range results {
    counts += v
  }

  return counts
}

func matchTelephoneNumbers(jobs <-chan string, results chan<- int, wg *sync.WaitGroup, telephone *regexp.Regexp) {
  // Decreasing internal counter for wait-group as soon as goroutine finishes
  defer wg.Done()

  // eventually I want to have a []string channel to work on a chunk of lines not just one line of text
  for j := range jobs {
    if telephone.MatchString(j) {
      results <- 1
    }
  }
}

func main() {
  // An artificial input source.  Normally this is a file passed on the command line.
  numberOfTelephoneNumbers := telephoneNumbersInFile(os.Args[1])
  fmt.Println(numberOfTelephoneNumbers)
}
