#ifndef _HASHUTIL_H
#define _HASHUTIL_H

unsigned int hashMysql(const char *key, unsigned int len);
unsigned int hashMysqlNRCaseUp(const char *key, unsigned int len);
unsigned long hashPhp(char *arKey, unsigned int nKeyLength);
unsigned long hashOpenSSL(char *str);
unsigned int hash(char *str);

#endif