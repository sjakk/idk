#include "canvas.h"
#include "tuple.h"

struct Projectile {
    point position;
    vector velocity;

    Projectile(const point& pos, const vector& vel) : position(pos), velocity(vel) {}
};


struct Environment {
    vector gravity;
    vector wind;

    Environment(const vector& g, const vector& w) : gravity(g), wind(w) {}
};



Projectile tick(const Environment& env, const Projectile& proj) {
    point position = proj.position + proj.velocity;
    vector velocity = proj.velocity + env.gravity + env.wind;
    return Projectile(position, velocity);
}


   
