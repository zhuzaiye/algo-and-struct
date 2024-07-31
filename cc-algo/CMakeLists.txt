cmake_minimum_required(VERSION 3.23)
project(cc_algo)

set(CMAKE_CXX_STANDARD 11)

# Include doctest header
#include_directories(test/lib)

# 包含 src 和 test 目录
add_subdirectory(src)
add_subdirectory(test)