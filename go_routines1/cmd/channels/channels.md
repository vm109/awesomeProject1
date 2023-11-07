### Channels scenario

#### Scenario
- Lets say there are multiple different websites a document is published to. 
- for each website we need to calculate the total number of documents and aggregate.


#### How do we use channels?
- Divide the work between multiple goroutines. 
- total goroutines equal to number of websites. 
- once each goroutine get the total documents per website, we can aggregate the total documents for all websites.

