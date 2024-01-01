# Node Package
This is node package for s, the repo for this project can be found
at: [S](https://github.com/harryvince/s)

## Example usage
```typescript
import { s } from '@harryvince_/s';

s(() => {
    serve();
});
// OR
s(() => {
    serve();
}, { env: 'test' });
```
Injects all s secrets into the environment, and an environment can be specified
if required.

If an env is not specified `dev` is used by default.
