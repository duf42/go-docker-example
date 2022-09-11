
cmake_minimum_required(VERSION 3.22)

project(
  GoIntegrationC
  VERSION 1.0
  LANGUAGES C)


SET(LIB_OUT_PATH ${GoIntegrationC_SOURCE_DIR}/lib)
SET(LIB_BIN_OUT_PATH ${GoIntegrationC_SOURCE_DIR}/test/bin)

include_directories(sources)
link_directories(sources)

# Library
add_subdirectory(sources)

# Tests
enable_testing()
add_subdirectory(tests)