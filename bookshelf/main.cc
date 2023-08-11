#include <iostream>
#include <ctime>

#include "matrix.h"

int main() {

Matrix<int,4,4> m1;

for (size_t i = 0; i < 3; ++i) {
        for (size_t j = 0; j < 3; ++j) {
            m1(i, j) = i * 3 + j;
        }
    }

int delay = 3*CLOCKS_PER_SEC;

Matrix a = std::move(m1);

a.print();

clock_t now = clock();
while(clock() - now < delay);

m1.print();
 }




