A consensus-breaking change refers to any modification to the blockchain protocol that makes the new version incompatible with the old version. Consensus-breaking changes disrupt the existing agreed set of rules that maintain a unified ledger.

1) First Consensus Breaking Change: Adding a new field to exchange.proto file
The exchange.proto file defines the structure of the Exchange datatype. The consensus-breaking change was adding a new field to the Exchange datatype. Transactions created with this new datatype will include the new field and the older softwares will not be able to recognise it, rejecting the transaction and resulting in a consensus failure.

**Old exchange.proto:**
message Exchange {
  
  string date = 1; 
  string amount = 2; 
  string message = 3; 
  string reciever = 4; 
  string sender = 5;
  uint64 id = 6; 
}

**New exchange.proto:**
message Exchange {
  
  string date = 1; 
  string amount = 2; 
  string message = 3; 
  string reciever = 4; 
  string sender = 5;
  string BreakConsensus = 6; 
  uint64 id = 7; 
}
