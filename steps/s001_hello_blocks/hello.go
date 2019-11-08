package main

import (
    "crypto/sha256"
    "encoding/binary"
    "encoding/hex"
    "bytes"
    "fmt"
)

type Block struct {
    PrevHash string
    Height uint64
}

/**
 * Create a cryptographic hash by converting all the fields of a Block to binary.
 * The result is a hexadecimal string, since the `sha256` function is used, the result
 * is always 64 characters long (32 bytes or 256 bits).
 *
 * To simplify the rest of the code, `HashBlock(nil)` is not an error.
 * Instead we return all zeros of the correct length.
 */
func HashBlock(block *Block) string {
    if block == nil {
        return "0000000000000000000000000000000000000000000000000000000000000000"
    }

    // This function converts the block to bytes by writing the fields into a Buffer,
    // then sending the Buffer contents to an sha256 object.  We do it this way so it
    // is easy to examine the bytes by printing the Buffer contents.

    buf := new(bytes.Buffer)

    // Write the PrevHash field
    binPrevBlockHash, err := hex.DecodeString(block.PrevHash)
    if err != nil { panic("Error decoding block.PrevHash") }
    buf.Write(binPrevBlockHash)

    // Write the Height field
    err = binary.Write(buf, binary.LittleEndian, block.Height)
    if err != nil { panic("Error writing block.Height") }

    // Done writing fields, get the Buffer contents
    blockBytes := buf.Bytes()

    // Uncomment one of these statements to print out the bytes
    // fmt.Printf("%s\n", hex.Dump(blockBytes))              // Pretty hex dump format
    // fmt.Printf("%s\n", hex.EncodeToString(blockBytes))    // Mashed-together characters format

    // Compute the hash of blockBytes using the sha256 cryptographic hash algorithm
    hasher := sha256.New()
    hasher.Write(blockBytes)
    hash := hex.EncodeToString(hasher.Sum(nil))

    // Uncomment this statement to print out the hash
    // fmt.Printf("The hash of these bytes is %s\n", hash)

    return hash
}



/**
 * Extend the blockchain by one block.
 *
 * prevBlockHash : The result of `HashBlock(prevBlock)`.
 * prevBlock : The most recent block, or `nil` if calling `ProduceBlock()` for the first time on a new blockchain.
 */
func ProduceBlock(prevBlockHash string, prevBlock *Block) *Block {
    // This function creates a new block.
    //
    // The implementation is fairly straightforward matter of creating a Block instance and filling in the fields.
    //
    // This function's API is slightly weird, it requires the caller to compute `prevBlockHash = HashBlock(prevBlock)`.
    // Why we don't simplify the API by computing `prevBlockHash` ourselves, reducing the number of arguments
    // from two to one?  Two reasons:
    //
    // - Immediately, it makes the placement of the `HashBlock()` debugging output less confusing.
    // - Eventually, we will have more data with the same data flow as `prevBlockHash`, so writing code to route this data
    // now will be useful later.

    newBlock := new(Block)

    if prevBlock == nil {
        newBlock.PrevHash = prevBlockHash
        newBlock.Height = 1
    } else {
        newBlock.PrevHash = prevBlockHash
        newBlock.Height = prevBlock.Height + 1
    }

    return newBlock
}

func main() {
    nilHash := HashBlock(nil)

    fmt.Printf("----------------------------------------------------------------\n")
    block1 := ProduceBlock(nilHash, nil)
    block1_hash := HashBlock(block1)
    fmt.Printf("Block 1:  %+v\n", *block1)

    fmt.Printf("----------------------------------------------------------------\n")
    block2 := ProduceBlock(block1_hash, block1)
    block2_hash := HashBlock(block2)
    fmt.Printf("Block 2:  %+v\n", *block2)

    fmt.Printf("----------------------------------------------------------------\n")
    block3 := ProduceBlock(block2_hash, block2)
    block3_hash := HashBlock(block3)
    fmt.Printf("Block 3:  %+v\n", *block3)

    fmt.Printf("----------------------------------------------------------------\n")
    block4 := ProduceBlock(block3_hash, block3)
    block4_hash := HashBlock(block4)
    fmt.Printf("Block 4:  %+v\n", *block4)

    fmt.Printf("----------------------------------------------------------------\n")
    block5 := ProduceBlock(block4_hash, block4)
    block5_hash := HashBlock(block5)
    fmt.Printf("Block 5:  %+v\n", *block5)

    _ = block5_hash                // Prevent unused variable error
}
