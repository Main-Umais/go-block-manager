# go-block-manager

A simple Go package for managing blocks stored in a file.

# Functions

## DisplayAllBlocks

    func DisplayAllBlocks(filename string) error

### Description:

Reads all blocks from the specified file and displays them.

### Parameters:

filename (string): The name of the file containing blocks.

### Returns:

error: An error if any occurred during file reading.

## NewBlock

    func NewBlock(filename, block string) error

### Description:

Appends a new block to the specified file.

### Parameters:

filename (string): The name of the file to which the block will be appended.
block (string): The content of the new block to be added.

### Returns:

error: An error if any occurred during file writing.

# ModifyBlock

    func ModifyBlock(filename, oldBlock, newBlock string) error

### Description:

Modifies a specific block in the specified file.

### Parameters:

filename (string): The name of the file containing blocks.
oldBlock (string): The content of the block to be modified.
newBlock (string): The new content to replace the old block.

### Returns:

error: An error if any occurred during file reading or writing.

# Usage

    Import the blockmanager package in your Go project:
    import "github.com/Mian-Umais/go-block-manager/blockmanager"

Use the provided functions (DisplayAllBlocks, NewBlock, ModifyBlock) in your code to manage blocks stored in a file.
