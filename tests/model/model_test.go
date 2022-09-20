package model

import (
    "testing"
)

func TestInitialize(t *testing.T) {

    // [ EXERCISE ]
    Initialize()

    // [ VERIFY ]
    if GetOutput("command") != 0.0 {
        t.Errorf("command = %f (expected = %f)", GetOutput("command") , 0.0)            
    }
    if GetParameter("Kp") != 0.8 {
        t.Errorf("Kp = %f (expected = %f)", GetParameter("Kp"), 0.8)            
    }

}

func TestGetOutputCommand(t *testing.T) {
    
    // [ SETUP ]
    name := "command"
    expected := 0.0

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
    expected := 0.8

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
    if GetParameter("Kp") != value {
        t.Errorf("Kp = %f (expected = %f)", GetParameter("Kp"), value)
    }
    
}

func TestSetInputTarget(t *testing.T) {
    
    // [ SETUP ]
    name := "target"
    value := 2.0

    // [ EXERCISE ]
    SetInput(name,value)
    
}

func TestSetInputCurrent(t *testing.T) {
    
    // [ SETUP ]
    name := "current"
    value := 4.5

    // [ EXERCISE ]
    SetInput(name,value)
    
}

func TestStep(t *testing.T) {
    
    // [ SETUP ]
    SetInput("target",4.0)
    SetInput("current",2.0)
    SetParameter("Kp",5)

    // [ EXERCISE ]
    Step()

    // [ VERIFY ]
    if GetOutput("command") != 10.04 {
        t.Errorf("command = %f (expected = %f)", GetOutput("command"), 10.0)
    }
    
}