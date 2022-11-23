#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(int argc, char *argv[]) {
   if (argc != 2) {
      fprintf(stderr, "Usage: %s [filname]\n", argv[0]);
      exit(1);
   }
   FILE *fp;
   if ((fp = fopen(argv[1], "r")) == NULL){
      fprintf(stderr, "Error while opening the file.\n");
       exit(2);
   }

   char direction[32];
   int val, x, y, y2, aim;
   x = y = y2 = aim = 0;

   while (fscanf(fp, "%s %d", direction, &val) > 0) {
      if (strcmp(direction, "forward") == 0) {
         x += val;
         y2 = y2 + (aim * val);
      }
      else if (strcmp(direction, "up") == 0) {
         y -= val;
         aim -= val;
      } 
      else if (strcmp(direction, "down") == 0) {
         y += val;
         aim += val;
      }
   }

   fclose(fp);
   printf("Part #1: %d\n", x * y);
   printf("Part #2: %d\n", x * y2);

}
