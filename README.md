# Bashista

[![Build Status](https://ci.mrcyjanek.net/badge/fe8e9f89?branch=master)](https://ci.mrcyjanek.net/repos/453)

Make bash.org.pl quotes annoying.

## How does it work?

By running `bashista` you will get a pop-up notification with random quote from bash.org.pl

You can get static binaries from here: https://static.mrcyjanek.net/abstruse/bashista/

Or if you are on a debian-like system you can use my debian repo to get updates whenever they are out:

```bash
wget 'https://static.mrcyjanek.net/abstruse/apt-repository/mrcyjanek-repo/mrcyjanek-repo_2.0-1_all.deb' -O cyjanrepo.deb && \
    apt install ./cyjanrepo.deb && \
    rm ./cyjanrepo.deb && \
    apt update && \
    apt install bashista -y
```