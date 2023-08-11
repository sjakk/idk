#include <cmath>
#include <iostream>

#ifndef __TUPLE__
#define __TUPLE__
class Tuple {
public:
  Tuple(): e{0,0,0}{};
  Tuple(double e0,double e1,double e2,double e3): e{e0,e1,e2,e3}{};
  Tuple(double e0,double e1,double e2): e{e0,e1,e2}{};
  Tuple(Tuple &&) = default;
  Tuple(const Tuple &) = default;
  Tuple &operator=(Tuple &&) = default;
  Tuple &operator=(const Tuple &) = default;
  ~Tuple(){};
        
      double operator[](int i)const{return e[i];}
      double& operator[](int i){ return e[i];}
      Tuple operator-()const{ return  Tuple(-e[0],-e[1],e[2],-e[3]);}
        
        double length(){return std::sqrt(length_squared());}
        double length_squared(){return e[0]*e[0]+e[1]*e[1]+e[2]*e[2]+e[3]*e[3];}
        

            


        Tuple& operator*=(const double t){
          e[0] *= t;
          e[1] *= t;
          e[2] *= t;
          e[3] *= t;

      return *this;
  }

        double x() const { return e[0]; }
        double y() const { return e[1]; }
        double z() const { return e[2]; }
        double w() const { return e[3]; }





public:
  double e[4];
};

inline std::ostream& operator<<(std::ostream &out, const Tuple &v) {
    return out << v.e[0] << ' ' << v.e[1] << ' ' << v.e[2] << ' ';
}

inline bool operator==(const Tuple &u,const Tuple &v){
  if(u.e[0] == v.e[0] && u.e[1] == v.e[1] && u.e[2] == v.e[2] && u.e[3] == v.e[3]) return true;
  return false;
}




inline Tuple operator+(const Tuple &u, const Tuple &v){
 return Tuple(u.e[0]+v.e[0],
        u.e[1]+v.e[1],
        u.e[2]+v.e[2],
        u.e[3]+v.e[3]);
}

inline Tuple operator-(const Tuple& u, const Tuple& v){
  return Tuple(
    u.e[0] - v.e[0],
    u.e[1] - v.e[1],
    u.e[2] - v.e[2],
    u.e[3] - v.e[3]
  );
}


inline Tuple operator*(const Tuple &u, const Tuple &v){
      return Tuple(u.e[0]*v.e[0],
            u.e[1]*v.e[1],
            u.e[2]*v.e[2],
            u.e[3]*v.e[3]);   
  }

inline Tuple operator*(double t,const Tuple& v){
    return Tuple(v.e[0]*t,t*v.e[1],t*v.e[2],t*v.e[3]);
  }

inline Tuple operator*(const Tuple &v,double t){
  return t * v;
}

inline Tuple operator/(Tuple v,double t){return (1/t) *v;}
inline Tuple Normalize(Tuple v){return v / v.length();}


inline double dot(const Tuple& u,const Tuple& v){
  return u.e[0] * v.e[0]
  + u.e[1] * v.e[1]
  + u.e[2] * v.e[2]
  + u.e[3] * v.e[3];
}

inline Tuple cross(const Tuple& u, const Tuple& v){
  return Tuple(
    u.e[1] * v.e[2] - u.e[2] * v.e[1],
    u.e[2] * v.e[0] - u.e[0] * v.e[2],
    u.e[0] * v.e[1] - u.e[1] * v.e[0]);
}


using point = Tuple;
using vector = Tuple;
using color = Tuple;
#endif
