#include <vector>
#include "tuple.h"

#ifndef __CANVAS__
#define __CANVAS__
class Canvas {
public:
    Canvas(int width, int height) : width(width), height(height), pixels(width * height, color(0, 0, 0)) {}
Canvas() = default;
    int get_width() const { return width; }
    int get_height() const { return height; }

    const color& get_pixel(int x, int y) const {
        return pixels[y * width + x];
    }

    void write_pixel(int x, int y, const color& c) {
        pixels[y * width + x] = c;
    }

private:
    int width;
    int height;
    std::vector<color> pixels;
};


inline std::string canvasToPpm(){
	return "";
}


#endif
