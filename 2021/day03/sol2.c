#include <stdio.h>
#include <string.h>

#define BITS 12
#define LINES 1000

/* TO ROZWIAZANIE NIE DZIALA :( */

int to_dec(char *s) {
   int res = 0;
   for (int idx = strlen(s) - 1, pow = 1; idx >= 0; idx--, pow *= 2) {
      res = res + (pow * (s[idx] - '0'));
   }
   return res;
}

int handleOnes(char arr[LINES][BITS + 1], int indices[LINES], int idx, int *first, int indsize) {
   int tmp = indsize;
   if (*first) {
      for (int i = 0; i < LINES; i++)
         if (arr[i][idx] == '1') {
            indices[i] = i;
            tmp++;
         }
      *first = 0;
   }
   else {
      // 10001 100011
      for (int i = 0; i < LINES; i++)
         if (indices[i] != -1)
            if (arr[indices[i]][idx] != '1'){
               indices[i] = -1;
               tmp--;
            }
   }
   return tmp;
}

int handleZeros(char arr[LINES][BITS + 1], int indices[LINES], int idx, int *first, int indsize) {
   int tmp = indsize;
   if (*first) {
      for (int i = 0; i < LINES; i++)
         if (arr[i][idx] == '0') {
            indices[i] = i;
            tmp++;
         }
      *first = 0;
   }
   else {
      // 10001 100011
      for (int i = 0; i < LINES; i++)
         if (indices[i] != -1)
            if (arr[indices[i]][idx] != '0'){
               indices[i] = -1;
               tmp--;
            }
   }
   return tmp;
}
int main(int argc, char *argv[]) {
   if (argc != 2) return 1;

   FILE *fp = fopen(argv[1], "r");

   char arr[LINES][BITS + 1];
   for (int i = 0; fscanf(fp, "%s", arr[i]) > 0; i++)
      ;

   int indices[LINES];
   int indsize = 0;
   for (int i = 0; i < LINES; i++) indices[i] = -1;
   int first = 1;
   for (int i = 0; i < BITS; i++) {
      int cnt = 0;
      for (int j = 0; j < LINES; j++) {
         if (first)
            cnt = (arr[j][i] == '1') ? cnt + 1 : cnt - 1;
         else if (indices[j] != -1)
            cnt = (arr[indices[j]][i] == '1') ? cnt + 1 : cnt - 1;
            
      }
      if (cnt >= 0) 
         handleOnes(arr, indices, i, &first, indsize);
      else handleZeros(arr, indices, i, &first, indsize);
   }

   char res1[BITS + 1];
   for (int i = 0; i < LINES; i++){
      if (indices[i] != -1)
         strcpy(res1, arr[indices[i]]);
   }

   for (int i = 0; i < LINES; i++) indices[i] = -1;
   first = 1;
   indsize = 0;
   for (int i = 0; i < BITS; i++) {
      int cnt = 0;
      for (int j = 0; j < LINES; j++) {
         if (first)
            cnt = (arr[j][i] == '1') ? cnt + 1 : cnt - 1;
         else if (indices[j] != -1)
            cnt = (arr[indices[j]][i] == '1') ? cnt + 1 : cnt - 1;
            
      }
         if (!first && indsize == abs(cnt))
            continue;
      
      if (cnt < 0) { 
         if ((indsize = handleOnes(arr, indices, i, &first, indsize)) == 1) 
            break;
      } else { 
         if ((indsize = handleZeros(arr, indices, i, &first, indsize)) == 1) 
            break;
      }
   }

   char res2[BITS + 1];
   for (int i = 0; i < LINES; i++) {
      if(indices[i] != -1);
         strcpy(res2, arr[indices[i]]);
   }

   printf("%s %s\n", res1, res2);
   //printf("%d\n", to_dec(res1) * to_dec(res2));
}
