package model

/*
#cgo CFLAGS: -g -Wall -I/deps
#cgo LDFLAGS: -L/deps -lGoIntegrationC
#include "simple_lib.h"
*/
import "C"
/*
type (
	inputs_T C.struct_Inputs_T
)
type (
	outputs_T C.struct_Outputs_T
)
type (
	parameters_T C.struct_Parameters_T
)
*/

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

	simdata.inputs.target = 0.1
	simdata.inputs.current = 0.2
	simdata.outputs.command = 0.3
	simdata.parameters.Kp = 0.4

    C.Model_U.target  = C.double(simdata.inputs.target)
    C.Model_U.current = C.double(simdata.inputs.current)

    C.Model_Y.command = C.double(simdata.outputs.command)

    C.Model_P.Kp = C.double(simdata.parameters.Kp)

    return nil

}

func Step() error {

	C.Model_U.target = C.double(simdata.inputs.target)
	C.Model_U.current = C.double(simdata.inputs.current)
	C.Model_P.Kp = C.double(simdata.parameters.Kp)

    C.Model_Y.command = (C.Model_U.target - C.Model_U.current) * C.Model_P.Kp

	simdata.outputs.command = float64(C.Model_Y.command)

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
        simdata.inputs.target = value
		C.Model_U.target = C.double(simdata.inputs.target)		
    } else if name == "current" {
		simdata.inputs.current = value
        C.Model_U.current = C.double(simdata.inputs.current)
    }
    return nil

}

func SetParameter(name string, value float64) error {

    if name == "Kp" {
		simdata.parameters.Kp = value
        C.Model_P.Kp = C.double(simdata.parameters.Kp)
    }
    return nil

}
