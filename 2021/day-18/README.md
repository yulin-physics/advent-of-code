## Day 18 Binary Tree

[[[[[9,8],1],2],3],4]
Wrong representation:

```
4
│
└─3
  │
  └─2
    │   ┌─8
    └─1─│
        └─9
```

Correct representation:

```
               Node
             /     \
            Node    4
          /    \
        Node    3
      /     \
     Node    2
   /      \
Node       1
/ \
8   9
```

after a single explode action, becomes [[[[0,9],2],3],4]

Do left first search, biasing the left, look for leaf node where there is a number

Each node needs

- a parent node
- indicate it is a left or right child
