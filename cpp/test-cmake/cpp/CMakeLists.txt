cmake_minimum_required(VERSION 3.5)
project(testcmake)

set(CMAKE_CXX_STANDARD 14)

include_directories(
    /home/codeman/.itsoflife/git/ws/ws-misc-git/cpp/test-cmake/include
    )

add_executable(kaledioscope
        add.cc)

