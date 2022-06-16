package main

import "fmt"

type inputErr struct {
    message      string
    missingField string
}

func (i *inputErr) Error() string {
    return i.message
}

func (i *inputErr) getMissingField() string {
    return i.missingField
}

func main() {
    err := validate("", "")
    if err != nil {
        if err, ok := err.(*inputErr); ok {
            fmt.Println(err)
            fmt.Printf("Missing Field is %s\n", err.getMissingField())
        }
    }
}

func validate(name, gender string) error {
    if name == "" {
        return &inputErr{message: "Name is mandatory", missingField: "name"}
    }
    if gender == "" {
        return &inputErr{message: "Gender is mandatory", missingField: "gender"}
    }
    return nil
}
