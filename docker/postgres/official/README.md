Resource:
<https://github.com/docker-library/postgres/tree/4c2e78c234a8e4293a5d6bb6a4d20421236d98d8/13>

Not using it since the following error popped up on building:

    + key=B97B0AFCAA1A47F044F244A07FCC7D46ACCC4CF8
    + mktemp -d
    + export GNUPGHOME=/tmp/tmp.sv46t565Hw
    + gpg --batch --keyserver ha.pool.sks-keyservers.net --recv-keys B97B0AFCAA1A47F044F244A07FCC7D46ACCC4CF8
    gpg: keybox '/tmp/tmp.sv46t565Hw/pubring.kbx' created
    gpg: keyserver receive failed: Connection timed out
