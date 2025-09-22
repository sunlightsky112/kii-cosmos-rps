# 🎓 Cosmos SDK Student Module — RPS

This project extends a basic Cosmos SDK blockchain with a custom module called `rps`, designed to manage student records on-chain. It implements full CRUD functionality, CLI integration, and gRPC support using idiomatic Cosmos SDK patterns.

## 📦 Module Overview

The `rps` module allows users to:

- Create student records
- Delete student records
- Query individual students
- Query all students with pagination

Each student is stored in the KVStore and includes metadata such as the block height at which they were created.

## 🧱 Data Model

### `Student` Message

Defined in `proto/lb/rps/v1/types.proto`:

```proto
message Student {
  string id = 1;           // Unique identifier (derived from signer address)
  string name = 2;         // Student's name
  uint64 age = 3;          // Student's age
  int64 created_at = 4;    // Block height when created
}
```

## 🔁 Transactions

### `CreateStudent`

- Stores a new student in KVStore
- `id` is derived from the signer address
- `created_at` is set using `ctx.BlockHeight()`
- Emits `create_student` event with:
  - `student_id`, `name`, `age`, `creator`, `created_at`

### `DeleteStudent`

- Removes student by `id`
- Emits `delete_student` event with:
  - `student_id`, `creator`

## 🔍 Queries

### `GetStudent`

- Retrieves a student by `id`
- Returns full `Student` object

### `GetStudents`

- Returns all students
- Supports pagination via `FilteredPaginate`


## 🧪 CLI Integration

AutoCLI is enabled via `x/rps/autocli.go`:

```go
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
  return &autocliv1.ModuleOptions{
    Tx: &autocliv1.ServiceCommandDescriptor{
      Service: types.MsgServiceDesc.ServiceName,
    },
    Query: &autocliv1.ServiceCommandDescriptor{
      Service: types.QueryServiceDesc.ServiceName,
    },
  }
}
```

### Available Commands

```bash
# Create a student
rpsd tx rps create-student --name "Alice" --age 21 --from validator

# Delete a student
rpsd tx rps delete-student --id <student-id> --from validator

# Query a student
rpsd query rps get-student --id <student-id>

# Query all students
rpsd query rps get-students
```

## ⚙️ Build & Setup

### Makefile Commands

| Command           | Description                                      |
|------------------|--------------------------------------------------|
| `make proto-gen` | Generate Go files from `.proto` definitions      |
| `make install`   | Build the blockchain binary `rpsd`               |
| `make init`      | Initialize the chain and genesis configuration   |
| `make proto-all` | Run linting and generation for all proto files   |

## 🧪 Testing Instructions

### 1. Initialize the chain

```bash
make init
```

### 2. Start the node

```bash
rpsd start
```

### 3. Create a student

```bash
rpsd tx rps create-student --name "Bob" --age 22 --from validator
```

### 4. Query a student

```bash
rpsd query rps get-student --id <student-id>
```

### 5. Query all students

```bash
rpsd query rps get-students
```

### 6. Delete a student

```bash
rpsd tx rps delete-student --id <student-id> --from validator
```
