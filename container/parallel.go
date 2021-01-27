package main

import (
    "fmt"
    "log"
    "sync"
    "time"
    "os"
    "os/exec"
  "strings"
  "bufio"
  "math/rand"
  "syscall"

  "encoding/json"

  "net/http"
  "github.com/gorilla/mux"
  "github.com/rs/cors"
)

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}

func main() {
  rand.Seed(time.Now().UnixNano())
  fmt.Println("Start")
  parallelization := 6
  // list := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "l"}

  scmd := buildCmdArray("test.txt")
	for _, eachline := range scmd {
		// fmt.Println(strings.Join(eachline[4:len(eachline)], "|"))
		fmt.Println(strings.Join(eachline, "|"))
	}

  go func() {
    c := make(chan []string)

    var wg sync.WaitGroup
    wg.Add(parallelization)
    for ii := 0; ii < parallelization; ii++ {
      go func(c chan []string) {
          for {
              v, more := <-c
              if more == false {
                  wg.Done()
                  return
              }

              ret, err := do(v)
              if err != nil {
                  log.Printf("nop")
              }
              fmt.Println(ret)
          }
      }(c)
    }
    for {
      for _, a := range scmd {
        c <- a
      }
    }
    close(c)
    wg.Wait()
    fmt.Println("End")
  }()
  //
  // Mux router part
  //
  router := mux.NewRouter()

  router.HandleFunc("/", showResponse).Methods("GET")

  handler := cors.Default().Handler(router)

  port := fmt.Sprintf(":%s", getEnv("PORT", "8080"))

  log.Printf("Listening on %s...", port)
  log.Fatal(http.ListenAndServe(port, handler))
}

type Message struct {
	Text string
}

func showResponse(w http.ResponseWriter, r *http.Request) {
  m := Message{"Hello World"}

  w.Header().Set("Content-Type", "application/json")

  json.NewEncoder(w).Encode(&m)
}

func do(s []string) (string, error) {
  // time.Sleep(1 * time.Second)
  cmd := exec.Command(s[0], s[1:]...)
  cmd.Stdout = os.Stdout

  cmd.SysProcAttr = &syscall.SysProcAttr{
    Setpgid:   true,
    // Pdeathsig: syscall.SIGTERM,
  }

  if err := cmd.Start(); err != nil {
    log.Print("Command finished with error: %v", err)
    return "", err
  }

  // Wait for the process to finish or kill it after a timeout (whichever happens first):
  done := make(chan error, 1)
  go func() {
    done <- cmd.Wait()
  }()
  select {
  case <-time.After(5 * time.Second):
    if err := cmd.Process.Kill(); err != nil {
        log.Fatal("failed to kill process: ", err)
    }
    log.Println("process killed as timeout reached")
  case err := <-done:
    if err != nil {
        log.Print("process finished with error = %v", err)
        return "", err
    }
    log.Print("process finished successfully")
  }

  // log.Printf("Running command and waiting for it to finish...")

  n := rand.Intn(15)
  time.Sleep(time.Duration(n)*time.Second)

  fmt.Printf("Sleeped %d seconds...\n", n)
  return fmt.Sprintf("%s-%d", strings.Join(s, " "), time.Now().UnixNano()), nil
}

func buildCmdArray(filename string) [][]string {
	file, err := os.Open(filename)
	defer file.Close()
  
 
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
 
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
  var commands [][]string
 
	for scanner.Scan() {
    var txtline = scanner.Text()

		commands = append(commands, strings.Fields(txtline))
	}
 
 

  return commands
}
