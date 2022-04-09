package main

import "testing"

func TestMain(t *testing.T) {
	
	defer func() {
        if r := recover(); r != nil {
           t.Error("This func should not panic")
        }
    }()
	main()
}
