# VTT2SRT
A simple Go script to convert VTT (WebVTT) subtitles into SRT (SubRip Text) format.

## MakeFile

Build the application
```bash
make build
```

Run the application
```bash
make run
```

Clean up binary from the last build:
```bash
make clean
```

## Usage
convert single vtt file:
```
go run main.go --path subtitle.vtt
```
or 
```
binaryname --path subtitle.vtt
```

convert whole diretory:
```
go run main.go --dir /dir/to/vttfiles
```
or
```
binaryname --dir /dir/to/vttfiles
```












