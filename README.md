A consensus-breaking change refers to any modification to the blockchain protocol that makes the new version incompatible with the old version. Consensus-breaking changes disrupt the existing agreed set of rules that maintain a unified ledger.

Consensus Breaking Change: Adding a new field, string BreakConsensus, to exchange.proto file
The exchange.proto file defines the structure of the Exchange datatype. The consensus-breaking change was adding a new field to the Exchange datatype. Transactions created with this new datatype will include the new field and the older software will not be able to recognise it, rejecting the transaction and resulting in a consensus failure.
