#include <iostream>
#include <fstream>
#include <cmath>
#include "ppm.h"

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




int main() {

  // Define the initial projectile and environment
    point initialPosition(0, 1, 0);
    vector initialVelocity = Normalize(vector(1, 1.8, 0)) * 20;
    Projectile projectile(initialPosition, initialVelocity);

    vector gravity(0, -0.1, 0);
    vector wind(-0.01, 0, 0);
    Environment environment(gravity, wind);

    // Create a canvas
    Canvas canvas(1800, 1500);

    int ticks = 0;
    while (projectile.position.y() > 0) {
        projectile = tick(environment, projectile);

        int x = static_cast<int>(projectile.position.x());
        int y = canvas.get_height() - static_cast<int>(projectile.position.y());
        if (x >= 0 && x < canvas.get_width() && y >= 0 && y < canvas.get_height()) {
            canvas.write_pixel(x, y, color(1, 0, 0)); // Red color for the path of the projectile
        }

        ticks++;
    }

    std::string ppm_data = canvas_to_ppm(canvas);
    std::ofstream file("projectile.ppm");
    if (file.is_open()) {
        file << ppm_data;
        file.close();
        std::cout << "Canvas data written to 'projectile.ppm'." << std::endl;
    } else {
        std::cerr << "Failed to write to file." << std::endl;
    }

    return 0;


}




