## Hardcoding Factors
* In `vendor/github.com/ethereum/ethash/src/libethash/data_sizes.h`
* dag_sizes
  ```go
  static const uint64_t dag_sizes[2048] = {
    1073739904U, 1082130304U, 1090514816U, 1098906752U, 1107293056U,
    1115684224U, 1124070016U, 1132461952U, 1140849536U, 1149232768U,
    ...,
    18228444544U, 18236833408U, 18245220736U
  };
  ```
* cache_sizes
  ```go
  const uint64_t cache_sizes[2048] = {
    16776896U, 16907456U, 17039296U, 17170112U, 17301056U, 17432512U, 17563072U,
    17693888U, 17824192U, 17955904U, 18087488U, 18218176U, 18349504U, 18481088U,
    284950208U, 285081536U
  };
  ```
* We regard those two factors are fixed.
