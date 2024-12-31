# Fun With Binary Trees

This was sparked by a discussion on [X/Twitter](https://x.com/vikhyatk/status/1873033432705712304). I wanted to see how difficult it would be to actually do this in Go. Turns out it's pretty simple

## High Level Explanation
There is a very basic implementation of a binary tree in the `data` package. It basically defines two structs; `Node` and `Btree`. A `Node` contains a `Value` and pointers to other `Node`s as `Left` and `Right`. The `Btree` contains a pointer to a `Root` `Node`. 

The instace of the binary tree is setup in the `main` package. It's simply hard coded because the `BTree` struct has no functions to add/remove/find `Node`s. If I were going to spend a lot of time on this they would be a good thing to add. For the purposes of this, I don't think it is necessary.

In order to facilitate traversing the tree I created a simple `Queue` struct also in the `data` package. I utilized generics here so that I could re-use the code later should I need a `Queue` implementation in the future. The `Queue` utilizes a slice under the hood and provides all of the necessary operations; `Enqueue`, `Dequeue`, `IsEmpty`, `Reset`, `Dump`, and `Peek`. I also created a generic constructor function that will initialize the `Queue`'s slice so that the calling code doesn't have to worry about manually instantiating it.

The `main` package instantiates an instance of a `Btree` with the required data. Then creates two instances of the `Queue` struct. One called `curQueue` and the other called `nextQueue`. It adds the root `Node` of the tree to the curQueue and enters a loop. The loop will exit when both queues are empty. Then while the curQueue is not empty the first Node in the Queue is pulled, and its Left and Right nodes are added to the nextQueue, and the Value of the current Node is printed. After the curQueue is emptied the next and cur Queues are switched, a new line is printed, and the loop continues.

In the end I'm not completely sure this is the most optimal solution as it requires two queues to keep track of everything, but I think it's a pretty good solution.
