#include <stdio.h>

int get_numlines(char *filename) {
   FILE *fp = fopen(filename, "r");
   int lines = 0;
   int c;
   while (!feof(fp)){
      c = fgetc(fp);
      if (c == '\n') lines++;
   }
   fclose(fp);
   return lines;
}

void read_input(int *arr, int size, char *filename) {
   FILE *fp = fopen(filename, "r");
   for (int i = 0; i < size && !feof(fp); i++)
      fscanf(fp, "%d", &arr[i]);

   fclose(fp);
}

int main(int argc, char *argv[]) {
   if (argc != 2) {
      fprintf(stderr, "Usage: %s [filename]\n", argv[0]);
      return -1;
   }

   char *filename = argv[1];
   int size, cnt1 = 0, cnt2 = 0;
   size = get_numlines(filename);
   int arr[size];

   read_input(arr, size, filename);
   
   // part 1
   for (int i = 0; i < size - 1; i++)
      if (arr[i] < arr[i + 1]) cnt1++;

   // part 2
   for (int i = 0; i < size - 3; i++)
      if (arr[i] < arr[i + 3]) cnt2++;

   printf("Part #1: %d\nPart #2: %d\n", cnt1, cnt2);
}
