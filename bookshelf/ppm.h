#include <sstream>
#include "canvas.h"


std::string canvas_to_ppm(const Canvas& canvas) {
    std::stringstream ppm;
    ppm << "P3\n";
    ppm << canvas.get_width() << " " << canvas.get_height() << "\n";
    ppm << "255\n";

    for (int y = 0; y < canvas.get_height(); ++y) {
        for (int x = 0; x < canvas.get_width(); ++x) {
            const color& c = canvas.get_pixel(x, y);
            int r = static_cast<int>(std::round(c.x() * 255));
            int g = static_cast<int>(std::round(c.y() * 255));
            int b = static_cast<int>(std::round(c.z() * 255));
            ppm << r << " " << g << " " << b << " ";
        }
        ppm << "\n";
    }

    return ppm.str();
}
