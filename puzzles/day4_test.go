  package puzzles
  
  import "testing"

  func startfunc() int {
    data, err := internal.ReadFile(filename)
    if err != nil {
      internal.LogOutput("error reading file: %s\n", err)
    }
}
