package filehandler

import "os"
import "bufio"

func ReadVidoFile(file_path string) ([]byte, error) {
    f, err := os.Open(file_path)

    if err != nil {
        panic(err)
    }

    defer f.Close()

    stats, statsErr := f.Stat()
    if statsErr != nil {
        return nil, statsErr
    }

    var size int64 = stats.Size()
    bytes := make([]byte, size)

    bufr := bufio.NewReader(f)
    _,err = bufr.Read(bytes)

    return bytes, err
}

