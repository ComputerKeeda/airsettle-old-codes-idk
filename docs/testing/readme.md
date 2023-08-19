this is command which can be used to run the testing test.go file

    go run ./cmd/airsettled/main.go


    package client_helper

    import (
        "encoding/json"
        "fmt"
        "os"

        handlers "airsettle/client_helper/handlers"
        tx "airsettle/client_helper/transaction"

        randomData "github.com/Pallinder/go-randomdata"
    )

    /*
    The Client_Helper function is a helper function that can be used to test the client_helper package.
    It creates a new account, stores it in the storage_account file, and then retrieves the account from
    the storage_account file.
    */
    func Client_Helper() {
        randomName := randomData.FirstName(randomData.RandomGender)
        data := handlers.CreateAccount("./client_helper/storage_account", randomName)
        // PrettyPrint(data)
        _ = data

        data2 := handlers.GetAccount("./client_helper/storage_account", randomName)
        fmt.Println(data2.Address)
        // PrettyPrint(data2.Address)

        tx.AddExecutionLayer(data2.Address)
    }

    // The PrettyPrint function takes in any data and prints it out in a formatted and indented JSON
    // string.
    func PrettyPrint(data interface{}) {
        prettyJSON, err := json.MarshalIndent(data, "", "   ")
        if err != nil {
            fmt.Printf("Failed to generate JSON: %v", err)
            return
        }

        fmt.Println(string(prettyJSON))
    }

    // The function `InputFileToTextConverter` takes a file path as input, reads the contents of the file,
    // and returns the contents as a slice of strings.

    func InputFileToTextConverter(filePath string) ([]string, error) {
        file, err := os.Open(filePath)
        if err != nil {
            return nil, err
        }
        defer file.Close()

        var inputs []string
        err = json.NewDecoder(file).Decode(&inputs)
        if err != nil {
            return nil, err
        }

        return inputs, nil
    }
