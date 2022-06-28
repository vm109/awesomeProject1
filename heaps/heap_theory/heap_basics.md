Heap DataStructure: 

1. What is a Heap Data Structure? 
   Heap is a tree data structure. Where the Heap is a binary tree. 
   Heaps can be 2 types. 
   Max-Heap and Min-Heap
2. What is the difference between Max-Heap and Min-Heap?
   Max-Heap: Parent node is always greater than child nodes. 
   Min-Heap: Parent node is always smaller than child nodes. 
3. What is a Binary Heap? 
   1. A Binary heap is a complete tree.
   Complete tree in the sense - all levels of a tree are completely filled except may be the last level. 
   And in the last level the nodes should be as left as possible. 
   2. A binary heap must be either min-heap or max-heap
4. What is `heapify` in a tree?
   1. Process of reshaping a binary tree into a heap data structure 
   2. If a node's child nodes are heapified then only we can heapify that node.
   3. A heap should always be a complete binary tree. [ i.e all levels of tree should be filled except may be the last one]
      and the nodes should be as left as possible
   4. Once the heap is complete binary tree we can modify it ot become a max-heap by running a function called `heapify`
5. Heap sort algorithm for sorting in increasing order. 
   1. Build max heap from input data
   2. replace the last item of the heap with the root[ root is the largest in max heap ]
   3. heapify the root of the tree. 
   4. repeat the above steps while the size of the heap is greater than 1.