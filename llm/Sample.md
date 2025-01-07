# System Design Notes on Sharding and Database Optimization

## Key Concepts

### Query Optimization
- Use an **SQL optimizer** for querying a database.
- Indexing tables can improve performance but may not be sufficient for large datasets.

### Sharding
- **Definition**: Sharding is a method of partitioning data across multiple database servers.
- **Analogy**: Similar to sharing a pizza by slicing it for friends, where each server handles a "slice" of the database. 

### Horizontal vs. Vertical Partitioning
- **Horizontal Partitioning**: Involves breaking data into chunks based on a specific attribute (e.g., user ID).
- **Vertical Partitioning**: Involves breaking data based on columns.

## Sharding Mechanism
- **Partitioning by Key**: Use an attribute, such as user ID or location, to determine how to partition data.
- **Database Servers**: Focus on database servers that handle data consistently, unlike application servers (which are stateless).

## Important Principles
1. **Consistency**: Ensuring that data remains the same across all reads after a write operation.
2. **Availability**: Ensuring the database remains operational and doesn't crash, though consistency is prioritized over availability.

## Sharding Considerations
- When deciding what to shard on, consider the access patterns and data distribution (e.g., location-based sharding).
- Performance: Sharding can improve read and write performance by keeping data localized.

### Challenges with Sharding
- **Joins Across Shards**: Joins can cause expensive network calls if data is spread across multiple shards.
- **Inflexibility**: Once shards are created, it's challenging to adjust their size.

## Solutions for Sharding Challenges
- **Consistent Hashing**: A technique that allows shards to adapt dynamically, though its implementation can be complex.
- **Hierarchical Sharding**: Allows for dynamically splitting an overloaded shard into smaller pieces, improving flexibility.

### Indexing in Shards
- It's beneficial to create indexes on shards based on different attributes for efficient querying (e.g., finding users in a specific age range and location).

## Master-Slave Architecture
- **Master-Slave Setup**: A common architecture where:
  - Writes go to the master database.
  - Reads can be distributed across multiple slave databases.
  - If the master fails, a slave can be promoted to master, providing fault tolerance.

## Recommendations
- Start with simpler solutions like **indexing** or **NoSQL databases** before implementing sharding to avoid complexity.
- Focus on understanding consistent data management and scalability principles when designing systems.

---