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

    **NOTE** that any directory path not present in the host system and is in the deb pacakge, will be automatically created in the host.

4. Now copy the `helloworld` executable as follows, 

        $ cp ../helloworld usr/bin
        # /usr/bin/helloworld will be the final location
        # after package installation.

5. Optinally (recommended) add a man page say `helloworld.1` (a text file with `troff` formatting commands (`groff` in GNU universe)). Since it is a command it will be added to section `1`. Creat a copy it in the directory given below:

        $ mkdir usr/local/share/man/man1
        $ cp ../helloworld.1 usr/local/share/man/man1

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

Further notes:

1. The `DEBIAN` directory can contain other (pre-defined) files and directories as well.

        DEBIAN/
            control (required)
            templates (optional)
            preinst (optional, chmod 0755)
            postinst (optional, chmod 0755)
            prerm (optional, chmod 0755)
            postrm (optional, chmod 0755)

    The `preinst` `postinst` `prerm` `postrm` are scripts that can run. `templates` is a directory that can house question files used to ask from the user during installation.

[source1]:https://linuxconfig.org/easy-way-to-create-a-debian-package-and-local-package-repository
