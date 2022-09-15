package model

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -L./lib -lGoIntegrationC
#include "simple_lib.h"
*/
//import "C"

type inputs_T struct {
    target float64
    current float64
}

type outputs_T struct {
    command float64
}

type parameters_T struct {
    Kp float64
}

type simdata_T struct {
    inputs inputs_T
    outputs outputs_T
    parameters parameters_T
}

var simdata simdata_T

func Initialize() error {

    simdata.inputs.target  = 0.1
    simdata.inputs.current = 0.2

    simdata.outputs.command = 0.3

    simdata.parameters.Kp = 0.4

    return nil

}

func Step() error {
    simdata.outputs.command = (simdata.inputs.target - simdata.inputs.current) * simdata.parameters.Kp
    return nil
}

func GetOutput(name string) float64 {
    
    if name == "command" {
        return simdata.outputs.command
    } else {
        return 0.0
    }

}

func GetParameter(name string) float64 {
    if name == "Kp"{
        return simdata.parameters.Kp
    } else {
        return 0.0
    }
}

func SetInput(name string, value float64) error {

    if name == "target" {
        simdata.inputs.target = value
    } else if name == "current" {
        simdata.inputs.current = value
    }
    return nil

}

func SetParameter(name string, value float64) error {

    if name == "Kp" {
        simdata.parameters.Kp = value
    }
    return nil

}
