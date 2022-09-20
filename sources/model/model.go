package model

/*
#cgo CFLAGS: -g -Wall -I/deps
#cgo LDFLAGS: -L/deps -lGoIntegrationC
#include "simple_lib.h"
*/
import "C"

func Initialize() error {

    C.initialize()

    return nil

}

func Step() error {

    C.step()
    
    return nil
    
}

func GetOutput(name string) float64 {
    
    if name == "command" {
        return float64(C.Model_Y.command)
    } else {
        return 0.0
    }

}

func GetParameter(name string) float64 {
    
	if name == "Kp"{
        return float64(C.Model_P.Kp)
    } else {
        return 0.0
    }

}

func SetInput(name string, value float64) error {

    if name == "target" {
		C.Model_U.target = C.double(value)		
    } else if name == "current" {
        C.Model_U.current = C.double(value)
    }
    return nil

}

func SetParameter(name string, value float64) error {

    if name == "Kp" {
        C.Model_P.Kp = C.double(value)
    }
    return nil

}
