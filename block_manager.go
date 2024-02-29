package blockmanager

import (
    "fmt"
    "io/ioutil"
    "os"
)

// DisplayAllBlocks reads all blocks from the file and displays them.
func DisplayAllBlocks(filename string) error {
    blocks, err := readBlocksFromFile(filename)
    if err != nil {
        return err
    }
    fmt.Println("All Blocks:")
    for _, block := range blocks {
        fmt.Println(block)
    }
    return nil
}

// NewBlock appends a new block to the file.
func NewBlock(filename, block string) error {
    f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        return err
    }
    defer f.Close()

    _, err = fmt.Fprintf(f, "%s\n", block)
    if err != nil {
        return err
    }
    return nil
}

// ModifyBlock modifies a specific block in the file.
func ModifyBlock(filename, oldBlock, newBlock string) error {
    blocks, err := readBlocksFromFile(filename)
    if err != nil {
        return err
    }

    for i, block := range blocks {
        if block == oldBlock {
            blocks[i] = newBlock
            break
        }
    }

    err = writeBlocksToFile(filename, blocks)
    if err != nil {
        return err
    }
    return nil
}

func readBlocksFromFile(filename string) ([]string, error) {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    blocks := splitIntoBlocks(string(data))
    return blocks, nil
}

func writeBlocksToFile(filename string, blocks []string) error {
    data := []byte(joinBlocks(blocks))
    err := ioutil.WriteFile(filename, data, 0644)
    if err != nil {
        return err
    }
    return nil
}

func splitIntoBlocks(data string) []string {
    return strings.Split(data, "\n")
}

func joinBlocks(blocks []string) string {
    return strings.Join(blocks, "\n")
}
