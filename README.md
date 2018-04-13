# zergling-rush

ref: [zergling-rush](https://www.codingame.com/training/community/zergling-rush)

Your task is to visualize what the base will look like after the zerglings arrive.

Note 1: zerglings enter from all sides (top, left, right, bottom).
<br>
Note 2: if no building can be reached or if there are no buildings at all, the zerglings will not stay (no zerglings should be included in the output).

For example, a base like:
```
...#####
...#...#
..B..B.#
...#...#
...#####
```
<br>
will become
<br>
```
...#####
.zz#...#
.zB..B.#
.zz#...#
...#####
```
Zerglings surround the building to the left, but cannot access the building to the right because their path is blocked.
