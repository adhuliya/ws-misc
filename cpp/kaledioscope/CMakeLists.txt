cmake_minimum_required(VERSION 3.5)
project(kaledioscope)

set(CMAKE_CXX_STANDARD 14)

include_directories(
    /home/codeman/mydata/local/packages-live/clang-llvm6/llvm/include
    /home/codeman/mydata/local/packages-live/clang-llvm6/build/include
    )

add_executable(kaledioscope
        toy.cc)

