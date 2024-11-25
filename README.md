# Key Value Store Cache Golang Implementation

To run the project:

```shell
go run cmd/main.go
```

## Available Command

**1. get**<br/>

Get values by key

```shell
get [key]
```

example:

```shell
get sde_bootcamp
```

output:

```shell
title: SDE-Bootcamp, price: 30000.00, enrolled: false, estimated_time: 30
```

Print "No entry found for [keys]" if get returns null.

**2. put**<br/>

Put attributes to key value store.

```shell
put [key] [attribute_key1] [attribute_value1] [attribute_key2] [attribute_value2]
```

example:

```shell
put sde_kickstart title SDE-Kickstart price 4000.00 enrolled true estimated_time 8
```

Do not print anything. Print "Data Type Error" if attribute has data type other than previous set.

**3. delete**<br/>

Delete key from key value store

```shell
delete [key]
```

example:

```shell
delete sde_bootcamp
```

**4. search**<br/>

Search keys contains attributes key value

```shell
search [attribute_key] [attribute_value]
```

Example:

```shell
search enrollment true
```

Output:

```shell
sde_bootcamp,sde_kickstart
```

Print in sorted order.

**5. keys**<br/>

Print all keys

```shell
keys
```

Output:

```shell
sde_bootcamp,sde_kickstart
```

Print in sorted order.
