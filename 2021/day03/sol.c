#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define BITS 12 
#define LINES 1000

void strrev(char *s) {
   if (!s || *s == '\0') return;

   char temp, *start, *end;
   start = s;
   end = start + strlen(s) - 1;
   while (start < end) {
      temp = *start;
      *start++ = *end;
      *end-- = temp;
   }
   
}

void gamma_to_ep(char *gamma, char *epsilon) {
   while (*gamma) {
      *epsilon = (*gamma == '1') ? '0' : '1';
      gamma++;
      epsilon++;
   }
   *epsilon = '\0';
}

int to_dec(char *s) {
   int res = 0;
   for (int pow = 1, idx = strlen(s) - 1; idx >= 0; idx--, pow *= 2) {
      res = res + (pow * (s[idx] - '0'));
   }
   return res;
}

int main(int argc, char *argv[]) {
   if (argc != 2) {
      fprintf(stderr, "Usage: %s [filename]\n", argv[0]);
      return 1;
   }

   FILE *fp = fopen(argv[1], "r");
   char arr[LINES][BITS];

   for (int i = 0; fscanf(fp, "%s", arr[i]) > 0; i++)
      ;
   
   char gamma[BITS + 1], epsilon[BITS + 1];
   int k, i;

   for (i = BITS - 1, k = 0; i >= 0; i--, k++) {
      int cnt = 0;
      for (int j = 0; j < LINES; j++) {
         cnt = (arr[j][i] == '1') ? cnt + 1 : cnt - 1;
      }
      gamma[k] = (cnt > 0) ? '1' : '0';
   }

   gamma[k] = '\0';   
   strrev(gamma);
   gamma_to_ep(gamma, epsilon);

   printf("Part #1: %d\n", to_dec(gamma) * to_dec(epsilon));
}
