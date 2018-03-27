# diffset

Diff two files, treating each line as an element in a set.

## Usage

```
> diffset --help
Usage of diffset:
  -intersect
    	Only show intersections
  -new string
    	New file path
  -old string
    	Old file path
```

# Examples

```
[diffset]> echo "a
dquote> b
dquote> c" > 1

[diffset]> echo "c
dquote> d
dquote> e" > 2

[diffset]> diffset 1 2
- a
- b
+ d
+ e

[diffset]> diffset -intersect 1 2
c
```
