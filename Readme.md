# pbar

Show a simple progress bar in the terminal window with basic options.

## Usage

A sample example is given in the example directory.
With 2 simple steps and an optional step a simple progressbar can be shown.

**Create a simple Bar object**

```go
b := pbar.NewBar("sample.tar.xz")
```

**Optionally setup the display variables which alters the speed display**

```go
b.SetSpeedInfo(300, "MB/s")
```

**Update the progress**
```go
b.Update(34)
```

[Doumentation](https://godoc.org/gihtub.com/aki237/pbar)