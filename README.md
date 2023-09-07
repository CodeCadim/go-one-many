This is an example of managing a database one-to-many relationship in Go.

A one-to-many relation is described as the relation between two structures in which the first one may have multiple links to the other, but the other has only one link.

For example, an author may have several posts, but a post only have one author.

    +-------+             +-------+
    |Author |             |Post   |
    +-------+             +-------+
    |- ID   |     +-- * ->|- ID   |
    |- Name |     |       |- Title|
    |- Posts|- 1 -+       |- ...  |
    |- ...  |             +-------+  
    +-------+            

The issue introduced here is visible when requesting datas involving the two parts of the relation, ie retreiving author's and post's informations. Each result of the first structure will be present as many times it has links to the second structure. 
