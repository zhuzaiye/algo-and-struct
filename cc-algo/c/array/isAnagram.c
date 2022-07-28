#include <stdbool.h>
#include <stdio.h>
#include <string.h>

/*
    自定义：
        通过宏定义

        #define bool int
        #define true 1
        #define false 0
    引入:
        C99以后存在
        stdbool.h
 */

bool isAnagram1(char* s, char* t) {
    int slen = strlen(s);
    int tlen = strlen(t);

    // 如果长度不一致，直接返回false
    if (slen != tlen) {
        return false;
    }

    // 设定26个字母(小写)
    // 定义数组并初始化
    int stat[26] = {0};

    for (int i = 0; i < slen; i++) {
        stat[s[i] - 'a']++;
        stat[t[i] - 'a']--;
    }

    for (int i = 0; i < 26; i++) {
        if (stat[i] != 0) {
            return false;
        }
    }
    return true;
}

bool isAnagram2(char* s, char* t) {
    int hashTable[128] = {0};
    int count = 0;

    while (*t) {
        count++;
        hashTable[*t++]++;
    }

    while (*s) {
        if (hashTable[*s] > 0) {
            hashTable[*s]--;
            count--;
        } else {
            return false;
        }
        s++;
    }

    return (count ? false : true);
}

void main() {
    char* s1 = "anagram";
    char* s2 = "nagara";

    bool rn = isAnagram2(s1, s2);
    char* rst = "false";
    if (rn == true) {
        rst = "true";
        printf("resut: %s\n", rst);
    } else {
        printf("result: %s\n", rst);
    }
}

/*
    c-字符串：
        一个字符串是由一个或者多个字符组成，即末尾'\0'的字符数组。
    定义：
        char s[] = "lmj";
        char s[4] = {'l', 'm', 'j', '\0'};  // 这个也是常量
        char* s = "lmj";
    区别：
        char s[] = "lmj"; // 定义的是字符串变量，可以进行指定指针下的字符修改
        char* s = "lmj"; // 定义的是字符串常量，指针指向一块字符串常量，所以内部不能修改
    相关函数：
        // 传入指针变量或者数组名都可以
        size_t  strlen(const char*);
        char* strcpy(char*, const char*); // 字符串拷贝函数
        char* strcat(char*, const char*); // 字符串拼接函数
        int  strcmp(const char*, const char*); // 字符串比较函数
 */