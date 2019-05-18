# cmake file for reference. by Anshuman Dhuliya

cmake_minimum_required(VERSION 3.10)

project(HelloWorld) # project name

# setting custom variables
set (MY_VAR "Hello")
set (OTHER_VAR "${MY_VAR}, World")
set (PROJECT_DIR "/home/codeman/.itsoflife/git/ws/ws-misc-git/cpp/cmake/hello-world")

# systematize the build dir
# so when you do: mkdir build; cd build; cmake ..; make; it creates
# the directory structure inside `build/`
set(CMAKE_ARCHIVE_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/lib)
set(CMAKE_LIBRARY_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/lib)
set(CMAKE_RUNTIME_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/bin)

# add include directories
# include_directories(${PROJECT_DIR}/include)
include_directories(${CMAKE_CURRENT_SOURCE_DIR}/include)

set (CMAKE_CXX_STANDARD 14) # enable C++14 standard
set (CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -Wall")

add_executable(helloworld
  src/main.cpp
  # more cpp files go here
  )

target_link_libraries(helloworld helloworldlib)

add_executable(iterVector
  spike/iterVector.cpp
  )

#add_library (my_library STATIC|SHARED|MODULE my_source.cpp)
add_library (helloworldlib STATIC
  src/helloworldlib.cpp
  )

# add_subdirectory (project1) # including project1 into the main project
# add_subdirectory (project2) # including project2 into the main project

# a phony myclean target : the `clean` target has some builtin behavior
add_custom_target(myclean
  COMMAND ${CMAKE_COMMAND} --build . --target clean --config $<CONFIG>
  COMMAND rm -Rf bin CMakeCache.txt  CMakeFiles  cmake_install.cmake  lib  Makefile
  )

