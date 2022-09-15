package main

import (
    "testing"
)

func TestInitialize(t *testing.T) {

    // [ EXERCISE ]
    Initialize()

    // [ VERIFY ]
    if simdata.inputs.target != 0.1 {
        t.Errorf("simdata.inputs.target = %f (expected = %f)", simdata.inputs.target, 0.1)            
    }
    if simdata.inputs.current != 0.2 {
        t.Errorf("simdata.inputs.current = %f (expected = %f)", simdata.inputs.current, 0.2)            
    }
    if simdata.outputs.command != 0.3 {
        t.Errorf("simdata.outputs.command = %f (expected = %f)", simdata.outputs.command, 0.3)            
    }
    if simdata.parameters.Kp != 0.4 {
        t.Errorf("simdata.parameters.Kp = %f (expected = %f)", simdata.parameters.Kp, 0.4)            
    }

}

func TestGetOutputCommand(t *testing.T) {
    
    // [ SETUP ]
    name := "command"
    expected := 0.3

    // [ EXERCISE ]
    actual := GetOutput(name)

    // [ VERIFY ]
    if actual != expected {
        t.Errorf("%s output = %f, expected = %f", name, actual, expected)
    }

}

func TestGetParameterKp(t *testing.T) {
    
    // [ SETUP ]
    name := "Kp"
    expected := 0.4

    // [ EXERCISE ]
    actual := GetParameter(name)

    // [ VERIFY ]
    if actual != expected {
        t.Errorf("%s parameter = %f, expected = %f", name, actual, expected)
    }
    
}

func TestSetParameterKp(t *testing.T) {
    
    // [ SETUP ]
    name := "Kp"
    value := 5.2

    // [ EXERCISE ]
    SetParameter(name, value)

    // [ VERIFY ]
    if simdata.parameters.Kp != value {
        t.Errorf("simdata.parameters.Kp = %f (expected = %f)", simdata.parameters.Kp, value)
    }
    
}

func TestSetInputTarget(t *testing.T) {
    
    // [ SETUP ]
    name := "target"
    value := 2.0

    // [ EXERCISE ]
    SetInput(name,value)

    // [ VERIFY ]
    if simdata.inputs.target != 2.0 {
        t.Errorf("simdata.inputs.target = %f (expected = %f)", simdata.inputs.target, 2.0)
    }
    
}

func TestSetInputCurrent(t *testing.T) {
    
    // [ SETUP ]
    name := "current"
    value := 4.5

    // [ EXERCISE ]
    SetInput(name,value)

    // [ VERIFY ]
    if simdata.inputs.current != 4.5 {
        t.Errorf("simdata.inputs.current = %f (expected = %f)", simdata.inputs.current, 4.5)
    }
    
}

func TestStep(t *testing.T) {
    
    // [ SETUP ]
    simdata.inputs.target = 4.0
    simdata.inputs.current = 2.0
    simdata.parameters.Kp = 5

    // [ EXERCISE ]
    Step()

    // [ VERIFY ]
    if simdata.outputs.command != 10.0 {
        t.Errorf("simdata.outputs.command = %f (expected = %f)", simdata.outputs.command, 10.0)
    }
    
}