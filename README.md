# learnpixel
Some simple Go apps to learn pixel (https://github.com/faiface/pixel), a simple 2d game library written in Go.

## how to run the stuff
This examples are using Go modules. 

The OpenGL version used is **OpenGL 3.3**.

- On macOS, you need Xcode or Command Line Tools for Xcode (`xcode-select --install`) for required
  headers and libraries.
- On Ubuntu/Debian-like Linux distributions, you need `libgl1-mesa-dev` and `xorg-dev` packages.
- On CentOS/Fedora-like Linux distributions, you need `libX11-devel libXcursor-devel libXrandr-devel
  libXinerama-devel mesa-libGL-devel libXi-devel` packages.
- See [here](http://www.glfw.org/docs/latest/compile.html#compile_deps) for full details.

## qlines 
Bouncing lines example. Based on my old Atari ST demo, showcasing the Integer-Bresenham algorithm.

## asteroids
Work in progress. Also based on an old Atari ST demo. Some of my first Megamax C programs.

