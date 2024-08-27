Difference between Channels:
1. Buffered Channels: These type of channels have a specific capacity. The send operation is not blocked until the buffer is full.
    Pros:
   1. We can handle more amount of data at once.
   2. The sender and receiver don't have to wait for each other.
    Cons:
   1. The buffer size should be known, otherwise small buffer size might fill up and too big buffer size might cause the waste of resources.
   2. When the buffer is full, the send operation is blocked.
2. Unbuffered Channels: These don't have any capacity to hold the values. The values are sent only when the sender and receiver are ready.
    Pros:
   1. It is simple to use, because the buffer capacity is not mentioned. We can create the according to our requirement.
   2. We can send the values once both the sides are ready.
    Cons:
   1. If any one, sender or receiver is not available, then we cannot the send the values.
   2. We have to wait until the other side is ready, which can cause waiting issues.