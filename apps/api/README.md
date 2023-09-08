# LearnHub QL

```shell
atlas migrate diff migration_name \
  --dir "file://migrations" \
  --to "ent://src/schema" \
  --dev-url "postgresql://username:password@localhost:5432/temp?search_path=public&sslmode=disable"
```

```shell
atlas migrate apply \
  --dir "file://migrations" \
  --url "postgresql://username:password@localhost:5432/learnhub-ql?search_path=public&sslmode=disable"
```

```shell
atlas migrate status \
  --dir "file://migrations" \
  --url "postgresql://username:password@localhost:5432/learnhub-ql?search_path=public&sslmode=disable"
```
