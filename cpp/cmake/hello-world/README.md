Sample CPP Project Structure
============================

Author: Anshuman Dhuliya

This is a sample cpp project structure, inspired from [here](https://hiltmon.com/blog/2013/07/03/a-simple-c-plus-plus-project-structure/).

### What is the folder structure?
* build: This folder contains all object files, and is removed on a clean. It also contains the final executables and libraries in sub-folder `bin/` and `lib/` respectively.

* doc: Any notes, like my assembly notes and configuration files, are here. I decided to create the development and production config files in here instead of in a separate config folder as they “document” the configuration.

* include: All project header files. All necessary third-party header files that do not exist under /usr/local/include are also placed here.

* lib: Any libs that get compiled by the project, third party or any needed in development. Prior to deployment, third party libraries get moved to /usr/local/lib where they belong, leaving the project clean enough to compile on our Linux deployment servers. I really use this to test different library versions than the standard.

* spike: I often write smaller classes or files to test technologies or ideas, and keep them around for future reference. They go here, where they do not dilute the real application’s files, but can still be found later. (I really liked this idea)

* src: The application and only the application’s source files.

* test: All test code files. You do write tests, no?

* misc: All non-code/non-doc resources go here (images, helpful content, etc..)


### What goes in the `.gitignore` file?
The minimal content could be the following,

    # Ignore the build and lib dirs
    build
    lib/*

