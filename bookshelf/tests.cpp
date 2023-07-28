#include <catch2/catch_test_macros.hpp>
#include "tuple.h"



TEST_CASE("Colors are computed","[colors]"){
  color c1(0.9, 0.6, 0.75);
  color c2(0.7, 0.1, 0.25);
  REQUIRE((c1 * c2) == color(0.4, 0.6, 0.8)); 
}



