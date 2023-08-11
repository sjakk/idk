#ifndef __MATRIX__
#define __MATRIX__
#include <vector>
#include <iostream>

template<class T,size_t Rows,size_t Cols>

class Matrix{
	private:
		std::vector<std::vector<T> > data;

	
	public:
		Matrix(): data(Rows,std::vector<T>(Cols,T())){}

		Matrix(const Matrix&) = default;
		Matrix(Matrix&&) = default;
		
		T& operator()(size_t row,size_t col){
			return data[row][col];
		}
		const T& operator()(size_t row,size_t col) const{
			return data[row][col];
		};

		void print() const{
			 for (size_t i = 0; i < Rows; ++i) {
            		for (size_t j = 0; j < Cols; ++j) {
                		std::cout << data[i][j] << "\t";
            }
            std::cout << std::endl;
        }
	}


 friend const Matrix& operator*(const Matrix& m){
};
	


};




#endif
