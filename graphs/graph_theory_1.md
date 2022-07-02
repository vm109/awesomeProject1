What is a Graph 
 - Graph has vertices also called as nodes.
 - a finite set of ordered pair (u,v) called edge 
 - where u is a node and v is a node. 

What are the common representation of graphs?
 - Adjacency Matrix 
 - Adjacency List 
 - Undirected Graphs Adjacency matrix are symmetric 
 - Adjacency matrix can also be used to represent weighted graphs 

What is adjacency matrix in Graph Representation?
 - Adjacency matrix representation is a way of representing Graph as a matrix of 0 and 1
 - The boolean values if there is a edge between the given two vertices.
 - Advantages:
   - Basics operations on edges are efficient with constant time. 
     - That is adding, deleting an edge, checking if an edge is present between two vertices are efficient in adjacency matrix.
   - If graph is dense and has many edges then adjacency matrix should be our first choice. 
 - Disadvantages:
   - The total space needed for the adjacency matrix is V*V 
   - To those graphs which are sparse in edges will use up memory for the few edges it has 
   - In that case adjacency lists are better choice than adjacency matrix 