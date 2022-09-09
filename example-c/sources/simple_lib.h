
#ifndef SIMPLELIB__
#define SIMPLELIB__

/* Structure definitions */
typedef struct {
    double target;
    double current;
} Inputs_T;

typedef struct {
    double command;
} Outputs_T;


typedef struct {
    double Kp;
} Parameters_T;

/* Interface variable declaration */
extern Inputs_T Model_U;
extern Outputs_T Model_Y;
extern Parameters_T Model_P;

/* Function declaration */
extern void initialize();
extern void step();
extern void terminate();

#endif