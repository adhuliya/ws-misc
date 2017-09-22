A Simple Debian Package
============================

Source is [here][source1].

A simple debian package can be created using the following steps

1. Create and compile a `helloworld.c` program that prints `Hello, World!`

        // Contents of helloworld.c
        #include <stdio.h>
        int main() {
          printf("Hello, World!\n");
          return 0;
        }

2. Compile the above program and name it `helloworld`

        $ clang helloworld.c -o helloworld

3. Create a folder that will house the package and do the following,

        $ mkdir helloworld-1.0
        $ cd helloworld-1.0
        $ mkdir DEBIAN
        $ mkdir -p usr/bin

4. Now copy the `helloworld` executable as follows, 

        $ cp ../helloworld usr/bin
        # /usr/bin/helloworld will be the final location
        # after package installation.

5. Create a file `control` inside `DEBIAN` directory created earlier with the following content,

        Package: helloworld
        Version: 1.0
        Section: custom
        Priority: optional
        Architecture: all
        Essential: no
        Installed-Size: 1024
        Maintainer: helloworld.org
        Description: Print hello world on the screen

6. Now make the `deb` package after moving to the directory housing `helloworld-1.0` directory.

        $ sudo dpkg --build helloworld-1.0/

7. A file `helloworld-1.0.deb` must be created. Rename it as follows, (assuming 64 bit machine)

        $ mov helloworld-1.0.deb helloworld-1.0_amd64.deb
        # for 32 bit machine use _i386 instead of _amd64 suffix

8. Now install the package as follows,

        $ sudo dpkg -i helloworld-1.0_amd64.deb

9. Run the hello world program like other programs,

        $ helloworld
        Hello, World!


[source1]:https://linuxconfig.org/easy-way-to-create-a-debian-package-and-local-package-repository
