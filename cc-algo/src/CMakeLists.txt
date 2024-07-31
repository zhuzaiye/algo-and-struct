# Find all source files in the src directory and its subdirectories
file(GLOB_RECURSE SRC_FILES ${CMAKE_CURRENT_SOURCE_DIR}/*.cpp)

# 包含当前目录的头文件
include_directories(${CMAKE_CURRENT_SOURCE_DIR})

# 设置源文件为父作用域
set(SRC_FILES ${SRC_FILES} PARENT_SCOPE)