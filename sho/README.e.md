
# {{.Name}}/sho


## {{toc 5}}

# sho -- SimHash Oracle

The SimHash Oracle (sho) code is copied from Yahoo Inc's [github.com/yahoo/gryffin/html-distance](https://github.com/yahoo/gryffin/tree/master/html-distance). It uses BK Tree (Burkhard and Keller) for storing and verifying if a fingerprint is closed to a set of fingerprint within a defined proximity distance.

Distance is the hamming distance of the fingerprints. 

# API

#### > {{cat "example_test.go" | color "go"}}

All patches welcome.
