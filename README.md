Base58 encoding
==============

Base58 is a lesser known encoding, but it is popular when encoded data is meant to be handled by human users.
The encoded data contains only alphanumerical characters, and avoids the easily confused characters 0, i, l and O
(zero, india, lima, capital oscar).

Documentation for this package can be found at http://godoc.org/github.com/sycamoreone/base58.

A warning
---------

Please note that base58 is actually a set of encodings, using slightly different alphabets.
If you need to be compatible with  another base58 library, then better check if the same alphabets are used.

This library uses the alphabet

    alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

which is also used by Bitcoin, but is different from the alphabet used for Ripple addresses and Flickr URLs.
