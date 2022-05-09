# Example swagger with Fiber

To generate swagger docs:

```
swag init \ 
    --parseInternal=true \ 
    --generatedTime=true \ 
    --dir=cmd/server \ 
    --parseDependency=true \ 
    --output=docs
```