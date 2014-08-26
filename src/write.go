package main

import (
  "bufio"
  "fmt"
  "os"
  "io/ioutil"
)

func getTempFile() string {
  f, _ := ioutil.TempFile(os.TempDir(), "docker_evolve_")

  defer os.Remove(f.Name())
  return f.Name()
}

func WriteLinesToTempFile(lines []string) (string, error) {
  path := getTempFile();
  file, err := os.Create(path)
  if err != nil {
    return path, err
  }
  defer file.Close()

  w := bufio.NewWriter(file)
  fmt.Fprintln(w, "#!/bin/bash")
  for _, line := range lines {
    fmt.Fprintln(w, line)
  }

  return path, w.Flush()
}