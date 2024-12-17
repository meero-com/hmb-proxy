# Guild Proxy

Abstract asynchronous processing and make it synchronous for your clients.

## Goal

The proxy is made to allow under the hood, loosely coupled processing with parallel processing in mind.

Instead of having a "fire & forget", callback supported asynchronous processing, customers can rely on simple HTTP calls while having all the benefits of async processing.

Using a timeout / retry / exponential back-off strategy, users can completly abstract the implementation complexity of asynchronous systems.
