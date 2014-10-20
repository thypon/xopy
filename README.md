XopY
====

XopY is the bytes swiss-army knife.

Given two files _x_ and _y_ xopy operates bytewise against them, giving a compound result.
The operations included are:

- **eq**, copies _x/y_ bytes if _x/y_ bytes are equal, replaces with _pattern_ otherwise
- **and**, applies bitwise and against _x_ and _y_
- **or**, applies bitwise or against _x_ and _y_
- **xor**, applies bitwiser xor against _x_ and _y_
- **cut**, applies a cut returning _y_ bytes after a given _threshold_, _x_ bytes otherwise

Examples
--------

    $ echo -n "dogey" > dogey.txt
    $ echo -n "doge" > doge.txt
    $ echo -n "zarro" > zarro.txt
    $ echo -n "lorem" > lorem.txt
    
### Eq

    $ xopy -x doge.txt -y dogey.txt | xxd
    0000000: 646f 6765 00                             doge.
    
### And

    $ xopy -x doge.txt -y dogey.txt -operation and | xxd
    0000000: 646f 6765 08                             doge.
    
### Or

    $ xopy -x doge.txt -y dogey.txt -operation or | xxd
    0000000: 646f 6765 7b                             doge{
    
### Xor

    $ xopy -x doge.txt -y dogey.txt -operation xor | xxd
    0000000: 0000 0000 73                             ....s
    
### Cut

    $ xopy -x doge.txt -y zarro.txt -operation cut -threshold 2 | xxd
    0000000: 646f 7272 6f                             dorro
    
#### Pipe

XopY is also easily pipable:

    $ xopy -x doge.txt -y zarro.txt -operation cut -threshold 2 | xopy -y lorem.txt | xxd
    0000000: 006f 7200 00                             .or..
    
Usage
-----

    Usage of xopy:
    -extend=false: Extend the file read to the longer one
    -operation="eq": Select the operation: eq, and, or, xor, cut
    -output="STDOUT": The Output File, STDOUT if not specified
    -pattern="\x00": Select the replacing pattern
    -threshold=8: selection threshold in cut mode
    -x="STDIN": The X-file
    -y="": The Y-file
