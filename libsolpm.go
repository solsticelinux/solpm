package libsolpm

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strings"
)

func ParseBashConfig(filepath string) (map[string]string, error) {
    config := make(map[string]string)
    file, err := os.Open(filepath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    re := regexp.MustCompile(`^([A-Za-z_][A-Za-z0-9_]*)=(.*)$`)

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())

        if strings.HasPrefix(line, "#") || line == "" {
            continue
        }

        matches := re.FindStringSubmatch(line)
        if len(matches) == 3 {
            key := matches[1]
            value := strings.Trim(matches[2], `"'`)
            config[key] = value
        }
    }

    return config, scanner.Err()
}

func main() {
    fmt.Println("Test libsolpm")

    config, err := ParseBashConfig("config.sh")
    if err != nil {
        fmt.Printf("Error reading config: %v\n", err)
        return
    }

    fmt.Println(config["VARIABLE_NAME"])
}

// this is probably incomplete, it's currently just a lib for solpm itself, kind of like how pacman depends on libalpm
// as of right now it should be able to support binary and source packages.

