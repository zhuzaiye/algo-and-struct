# 包含 doctest 头文件目录
include_directories(${CMAKE_CURRENT_SOURCE_DIR}/lib)
include_directories(${CMAKE_SOURCE_DIR}/src)

# 定义 hash_map 和 link_list 目录中的源文件
file(GLOB_RECURSE HASH_MAP_SRC_FILES ${CMAKE_SOURCE_DIR}/src/hash_map/*.cpp)
file(GLOB_RECURSE LINK_LIST_SRC_FILES ${CMAKE_SOURCE_DIR}/src/link_list/*.cpp)

# 定义每个测试文件及其所需的源文件
add_executable(lru_cache_test hash_map/lru_cache_test.cpp ${SRC_HASH_MAP})
target_compile_definitions(lru_cache_test PRIVATE DOCTEST_CONFIG_IMPLEMENT_WITH_MAIN)

add_executable(two_sum_test hash_map/two_sum_test.cpp ${SRC_HASH_MAP})
target_compile_definitions(two_sum_test PRIVATE DOCTEST_CONFIG_IMPLEMENT_WITH_MAIN)

# 继续为其他测试文件创建可执行文件，并添加所需的源文件