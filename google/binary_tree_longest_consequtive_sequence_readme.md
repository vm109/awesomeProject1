Problem description: 
Find the longest consecutive sequence in a binary tree.

What I understood: 
In a given binary tree find a sequence of child nodes where there are no sibling nodes
i.e in the below example tree child nodes of node `4` does not have siblings. 
so the consecutive sequence is thought to be [ 4, 6, 7, 8 ] and so the length of the 
consecutive sequence is 4.

``` 
   1
    \
     3
    /  \
   2     4
        /
       6
        \
         7
          \
           8
```

Reading the algorithm after writing my algorithm: 
After reading the algorithm in geeks-for-geeks turned out the longest consecutive sequence is about the numbers. 
that is in above tree the consecutive sequences are [3,4] [6,7,8] so the longest consecutive sequence is length `3`.