package main

import "fmt"

type Signal struct { Project string; Owner string; Profile string }

func main() {
  signal := Signal{Project: "weather-pulse-hkla", Owner: "JenniferWa19845", Profile: "0032"}
  fmt.Printf("%s is ready for %s (profile %s)", signal.Project, signal.Owner, signal.Profile)
}
