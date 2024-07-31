// Copyright (c) 2024/7/13 by hzzhu92
// LRUCache test of Doctest

#define DOCTEST_CONFIG_IMPLEMENT
#include "doctest.h"

#include "hash_map/lru_cache.cpp"

TEST_CASE("LRUCache Test" * doctest::test_suite("lru_cache_test")) {
    SUBCASE("Basic Operations") {
        LRUCache cache(2); // capacity of 2

        cache.put(1, 1);
        cache.put(2, 2);

        CHECK(cache.get(1) == 1);  // return 1
        CHECK(cache.get(2) == 2);  // return 2

        cache.put(3, 3);

        CHECK(cache.get(1) == -1);  // return -1 (not found)
        CHECK(cache.get(2) == 2);   // return 2
        CHECK(cache.get(3) == 3);   // return 3

        cache.put(4, 4);

        CHECK(cache.get(2) == -1); // returns -1 (not found)
        CHECK(cache.get(3) == 3);  // returns 3
        CHECK(cache.get(4) == 4);  // returns 4
    }

    SUBCASE("Edge Cases") {
        LRUCache cache(1); // Capacity of 1

        cache.put(1, 1);

        CHECK(cache.get(1) == 1); // returns 1

        cache.put(2, 2); // evicts key 1

        CHECK(cache.get(1) == -1); // returns -1 (not found)
        CHECK(cache.get(2) == 2);  // returns 2

        cache.put(3, 3); // evicts key 2

        CHECK(cache.get(2) == -1); // returns -1 (not found)
        CHECK(cache.get(3) == 3);  // returns 3
    }
}
