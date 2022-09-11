
#include "simple_lib.h"

/* Global variables definition */
Inputs_T Model_U;
Outputs_T Model_Y;
Parameters_T Model_P;

void initialize(){
    /* Inputs */
    Model_U.current = 0.0;
    Model_U.target  = 0.0;
    /* Outputs */
    Model_Y.command = 0.0;
    /* Parameters */
    Model_P.Kp      = 0.8;
}

void step(){

    double error;

    error = Model_U.target - Model_U.current;

    Model_Y.command = error * Model_P.Kp;

}

void terminate(){
    
}